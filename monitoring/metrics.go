package monitoring

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	Name   string
	Labels []string
	Type   interface{} // TODO @shipperizer use type to create proper metric
}

func GaugesV1() map[string][]string {
	return map[string][]string{
		"http_api_time":  {"service", "route"},
		"http_pai_count": {"service", "route"},
	}
}

// NewGaugeVec is simply a wrapper around prometheus library call
// useful when specifying custom metrics
func NewGaugeVec(name string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: fmt.Sprintf("metric: %s", name),
		},
		labels,
	)
}

// GaugeVecs returns a map of prometheus gauges, keyed by metric name
func GaugeVecs(metrics map[string][]string) map[string]*prometheus.GaugeVec {
	measurements := make(map[string]*prometheus.GaugeVec)

	for metricName, labels := range metrics {
		measurements[metricName] = NewGaugeVec(metricName, labels)
	}

	return measurements
}
