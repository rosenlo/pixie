package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"pixielabs.ai/pixielabs/src/shared/services"
	"pixielabs.ai/pixielabs/src/shared/services/election"
	"pixielabs.ai/pixielabs/src/shared/services/healthz"
	"pixielabs.ai/pixielabs/src/shared/services/httpmiddleware"
	"pixielabs.ai/pixielabs/src/shared/version"
	"pixielabs.ai/pixielabs/src/vizier/services/metadata/controllers"
	"pixielabs.ai/pixielabs/src/vizier/services/metadata/controllers/kvstore"
	"pixielabs.ai/pixielabs/src/vizier/services/metadata/metadataenv"
	"pixielabs.ai/pixielabs/src/vizier/services/metadata/metadatapb"
)

const (
	cacheFlushPeriod = 1 * time.Second
	cacheClearPeriod = 2 * time.Second
)

func init() {
	pflag.String("md_etcd_server", "https://etcd.pl.svc:2379", "The address to metadata etcd server.")
	pflag.String("cluster_id", "", "The Cluster ID to use for Pixie Cloud")
	pflag.Duration("max_expected_clock_skew", 750, "Duration in ms of expected maximum clock skew in a cluster")
	pflag.Duration("renew_period", 500, "Duration in ms of the time to wait to renew lease")
	pflag.String("pod_namespace", "pl", "The namespace this pod runs in. Used for leader elections")
}

func etcdTLSConfig() (*tls.Config, error) {
	tlsCert := viper.GetString("client_tls_cert")
	tlsKey := viper.GetString("client_tls_key")
	tlsCACert := viper.GetString("tls_ca_cert")

	tlsInfo := transport.TLSInfo{
		CertFile:      tlsCert,
		KeyFile:       tlsKey,
		TrustedCAFile: tlsCACert,
	}

	return tlsInfo.ClientConfig()
}

func main() {
	log.WithField("service", "metadata").
		WithField("version", version.GetVersion().ToString()).
		Info("Starting service")

	services.SetupService("metadata", 50400)
	services.SetupSSLClientFlags()
	services.PostFlagSetupAndParse()
	services.CheckServiceFlags()
	services.CheckSSLClientFlags()
	services.SetupServiceLogging()

	flush := services.InitDefaultSentry(viper.GetString("cluster_id"))
	defer flush()

	var tlsConfig *tls.Config
	if !viper.GetBool("disable_ssl") {
		var err error
		tlsConfig, err = etcdTLSConfig()
		if err != nil {
			log.WithError(err).Fatal("Failed to load SSL for ETCD")
		}
	}

	// Connect to etcd.
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{viper.GetString("md_etcd_server")},
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
	})
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to etcd at " + viper.GetString("md_etcd_server"))
	}
	defer etcdClient.Close()

	// Set up leader election.
	isLeader := false
	leaderMgr, err := election.NewK8sLeaderElectionMgr(
		viper.GetString("pod_namespace"),
		viper.GetDuration("max_expected_clock_skew"),
		viper.GetDuration("renew_period"),
		"metadata-election",
	)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to leader election manager.")
	}

	// Cancel callback causes leader to resign.
	leaderCtx, cancel := context.WithCancel(context.Background())
	go func() {
		// Campaign in background. Metadata replicas that are not the leader should
		// do everything that the leader does, except write to the metadata store
		err = leaderMgr.Campaign(leaderCtx)
		if err != nil {
			log.WithError(err).Fatal("Failed to become leader")
		}
		log.Info("Gained leadership")
		isLeader = true
	}()

	// Resign leadership after the server stops.
	defer func() {
		log.Info("Resigning leadership")
		cancel()
	}()

	etcdDS := kvstore.NewEtcdStore(etcdClient)
	cache := kvstore.NewCache(etcdDS)
	mds, err := controllers.NewKVMetadataStore(cache)
	if err != nil {
		log.WithError(err).Fatal("Could not create metadata store")
	}

	agtMgr := controllers.NewAgentManager(mds)
	keepAlive := true
	go func() {
		for keepAlive {
			if isLeader {
				agtMgr.UpdateAgentState()
			}
			time.Sleep(10 * time.Second)
		}
	}()
	go func() {
		for keepAlive {
			if isLeader {
				cache.FlushToDatastore()
				time.Sleep(cacheFlushPeriod)
			} else {
				cache.Clear()
				time.Sleep(cacheClearPeriod)
			}

		}
	}()

	defer func() {
		keepAlive = false
	}()

	// Listen for K8s metadata updates.
	mdHandler, err := controllers.NewMetadataHandler(mds, &isLeader)
	if err != nil {
		log.WithError(err).Fatal("Failed to create metadata handler")
	}
	mdHandler.AddSubscriber(agtMgr)

	mdHandler.ProcessSubscriberUpdates()

	mc, err := controllers.NewMessageBusController("pl-nats", "update_agent", agtMgr, mds, mdHandler, &isLeader)

	if err != nil {
		log.WithError(err).Fatal("Failed to connect to message bus")
	}
	defer mc.Close()

	k8sMc, err := controllers.NewK8sMetadataController(mdHandler)
	defer k8sMc.Stop()

	// Set up server.
	env, err := metadataenv.New()
	if err != nil {
		log.WithError(err).Fatal("Failed to create api environment")
	}
	mux := http.NewServeMux()
	healthz.RegisterDefaultChecks(mux)

	server, err := controllers.NewServer(env, agtMgr, mds)
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize GRPC server funcs")
	}

	log.Info("Metadata Server: " + version.GetVersion().ToString())

	s := services.NewPLServer(env,
		httpmiddleware.WithBearerAuthMiddleware(env, mux))
	metadatapb.RegisterMetadataServiceServer(s.GRPCServer(), server)
	s.Start()
	s.StopOnInterrupt()
}
