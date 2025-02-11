// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/storegateway/metadata_fetcher_metrics.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package storegateway

import (
	dskit_metrics "github.com/grafana/dskit/metrics"
	"github.com/prometheus/client_golang/prometheus"

	util_log "github.com/grafana/mimir/pkg/util/log"
)

// This struct aggregates metrics exported by Thanos MetaFetcher
// and re-exports those aggregates as Mimir metrics.
type MetadataFetcherMetrics struct {
	regs *dskit_metrics.TenantRegistries

	// Exported metrics, gathered from Thanos MetaFetcher
	syncs                *prometheus.Desc
	syncFailures         *prometheus.Desc
	syncDuration         *prometheus.Desc
	syncConsistencyDelay *prometheus.Desc
	synced               *prometheus.Desc

	// Ignored:
	// blocks_meta_modified
	// blocks_meta_base_syncs_total
}

func NewMetadataFetcherMetrics() *MetadataFetcherMetrics {
	return &MetadataFetcherMetrics{
		regs: dskit_metrics.NewTenantRegistries(util_log.Logger),

		// When mapping new metadata fetcher metrics from Thanos, please remember to add these metrics
		// to our internal fetcherMetrics implementation too.
		syncs: prometheus.NewDesc(
			"cortex_blocks_meta_syncs_total",
			"Total blocks metadata synchronization attempts",
			nil, nil),
		syncFailures: prometheus.NewDesc(
			"cortex_blocks_meta_sync_failures_total",
			"Total blocks metadata synchronization failures",
			nil, nil),
		syncDuration: prometheus.NewDesc(
			"cortex_blocks_meta_sync_duration_seconds",
			"Duration of the blocks metadata synchronization in seconds",
			nil, nil),
		syncConsistencyDelay: prometheus.NewDesc(
			"cortex_blocks_meta_sync_consistency_delay_seconds",
			"Configured consistency delay in seconds.",
			nil, nil),
		synced: prometheus.NewDesc(
			"cortex_blocks_meta_synced",
			"Reflects current state of synced blocks (over all tenants).",
			[]string{"state"}, nil),
	}
}

func (m *MetadataFetcherMetrics) AddUserRegistry(user string, reg *prometheus.Registry) {
	m.regs.AddTenantRegistry(user, reg)
}

func (m *MetadataFetcherMetrics) RemoveUserRegistry(user string) {
	m.regs.RemoveTenantRegistry(user, false)
}

func (m *MetadataFetcherMetrics) Describe(out chan<- *prometheus.Desc) {
	out <- m.syncs
	out <- m.syncFailures
	out <- m.syncDuration
	out <- m.syncConsistencyDelay
	out <- m.synced
}

func (m *MetadataFetcherMetrics) Collect(out chan<- prometheus.Metric) {
	data := m.regs.BuildMetricFamiliesPerTenant()

	data.SendSumOfCounters(out, m.syncs, "blocks_meta_syncs_total")
	data.SendSumOfCounters(out, m.syncFailures, "blocks_meta_sync_failures_total")
	data.SendSumOfHistograms(out, m.syncDuration, "blocks_meta_sync_duration_seconds")
	data.SendMaxOfGauges(out, m.syncConsistencyDelay, "consistency_delay_seconds")
	data.SendSumOfGaugesWithLabels(out, m.synced, "blocks_meta_synced", "state")
}
