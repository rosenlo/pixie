package autocomplete

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"

	"pixielabs.ai/pixielabs/src/cloud/cloudapipb"
)

// ElasticSuggester provides suggestions based on the given index in Elastic.
type ElasticSuggester struct {
	client          *elastic.Client
	mdIndexName     string
	scriptIndexName string
}

// EsMDEntity is information about metadata entities, stored in Elastic.
// This is copied from the indexer. When the indexer is checked in, we should consider
// putting this in a shared directory.
type EsMDEntity struct {
	OrgID string `json:"orgID"`
	UID   string `json:"uid"`
	Name  string `json:"name"`
	NS    string `json:"ns"`
	Kind  string `json:"kind"`

	TimeStartedNS int64 `json:"timeStartedNS"`
	TimeStoppedNS int64 `json:"timeStoppedNS"`

	RelatedEntityNames []string `json:"relatedEntityNames"`

	ResourceVersion string `json:"resourceVersion"`
	Test            string `json:"test"`
}

var protoToElasticLabelMap = map[cloudapipb.AutocompleteEntityKind]string{
	cloudapipb.AEK_SVC:       "service",
	cloudapipb.AEK_POD:       "pod",
	cloudapipb.AEK_SCRIPT:    "script",
	cloudapipb.AEK_NAMESPACE: "script",
}

var elasticLabelToProtoMap = map[string]cloudapipb.AutocompleteEntityKind{
	"service":   cloudapipb.AEK_SVC,
	"pod":       cloudapipb.AEK_POD,
	"script":    cloudapipb.AEK_SCRIPT,
	"namespace": cloudapipb.AEK_NAMESPACE,
}

// NewElasticSuggester creates a suggester based on an elastic index.
func NewElasticSuggester(client *elastic.Client, mdIndex string, scriptIndex string) *ElasticSuggester {
	return &ElasticSuggester{client, mdIndex, scriptIndex}
}

// GetSuggestions get suggestions for the given input using Elastic.
func (e *ElasticSuggester) GetSuggestions(orgID uuid.UUID, input string, allowedKinds []cloudapipb.AutocompleteEntityKind, allowedArgs []cloudapipb.AutocompleteEntityKind) ([]*Suggestion, bool, error) {
	q := elastic.NewBoolQuery()

	q.Should(e.getMDEntityQuery(orgID, input, allowedKinds))
	// TODO(michelle): Add script query here once that is ready: q.Should(e.getScriptQuery(orgID, input, allowedArgs))
	resp, err := e.client.Search().
		Query(q).
		Do(context.TODO())

	if err != nil {
		return nil, false, err
	}

	results := make([]*Suggestion, 0)

	// Convert elastic entity into a suggestion object.
	for _, h := range resp.Hits.Hits {
		res := &EsMDEntity{}
		err = json.Unmarshal(h.Source, res)
		if err != nil {
			return nil, false, err
		}
		results = append(results, &Suggestion{
			Name:  res.NS + "/" + res.Name,
			Score: float64(*h.Score),
			Kind:  elasticLabelToProtoMap[res.Kind],
		})
	}

	exactMatch := len(results) > 0 && results[0].Name == input
	return results, exactMatch, nil
}

func (e *ElasticSuggester) getMDEntityQuery(orgID uuid.UUID, input string, allowedKinds []cloudapipb.AutocompleteEntityKind) *elastic.BoolQuery {
	entityQuery := elastic.NewBoolQuery()
	entityQuery.Must(elastic.NewTermQuery("_index", e.mdIndexName))

	// Search by name + namespace.
	splitInput := strings.Split(input, "/") // If contains "/", then everything preceding "/" is a namespace.
	name := input
	if len(splitInput) > 1 {
		entityQuery.Must(elastic.NewTermQuery("ns", splitInput[0]))
		name = splitInput[1]
	}
	entityQuery.Must(elastic.NewMatchQuery("name", name))

	// Only search for entities in org.
	entityQuery.Must(elastic.NewTermQuery("orgID", orgID.String()))

	// Only search for allowed kinds.
	kindsQuery := elastic.NewBoolQuery()
	for _, k := range allowedKinds {
		kindsQuery.Should(elastic.NewTermQuery("kind", protoToElasticLabelMap[k]))
	}
	entityQuery.Must(kindsQuery)

	return entityQuery
}

func (e *ElasticSuggester) getScriptQuery(orgID uuid.UUID, input string, allowedArgs []cloudapipb.AutocompleteEntityKind) *elastic.BoolQuery {
	// TODO(michelle): Handle scripts once we get a better idea of what the index looks like.
	scriptQuery := elastic.NewBoolQuery()
	scriptQuery.Must(elastic.NewTermQuery("_index", "scripts"))

	return scriptQuery
}
