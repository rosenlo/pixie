#include <functional>
#include <memory>
#include <string>
#include <vector>

#include "src/stirling/bcc_connector.h"
#include "src/stirling/bpftrace_connector.h"
#include "src/stirling/cgroup_stats_connector.h"
#include "src/stirling/proc_stat_connector.h"
#include "src/stirling/proto/collector_config.pb.h"
#include "src/stirling/seq_gen_connector.h"
#include "src/stirling/socket_trace_connector.h"
#include "src/stirling/source_registry.h"

namespace pl {
namespace stirling {

void RegisterAllSources(SourceRegistry* registry) {
  CHECK(registry != nullptr);
  registry->RegisterOrDie<SeqGenConnector>("sequences");
  registry->RegisterOrDie<FakeProcStatConnector>("fake_proc_stat");
  registry->RegisterOrDie<ProcStatConnector>("proc_stat");
  registry->RegisterOrDie<PIDCPUUseBCCConnector>("bcc_cpu_stat");
  registry->RegisterOrDie<CPUStatBPFTraceConnector>("bpftrace_cpu_stats");
  registry->RegisterOrDie<PIDCPUUseBPFTraceConnector>("bpftrace_pid_cpu_usage");
  registry->RegisterOrDie<SocketTraceConnector>("socket_tracer");
  registry->RegisterOrDie<CGroupStatsConnector>("cgroup_stats");
}

}  // namespace stirling
}  // namespace pl
