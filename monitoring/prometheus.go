package monitoring

import (
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shipperizer/miniature-monkey/monitoring"
	"go.uber.org/zap"
)

// PrometheusConfig object
type PrometheusConfig struct {
	Metrics map[string]*prometheus.GaugeVec
	Service string
	Logger  *zap.SugaredLogger
}

// PrometheusMonitor object
type PrometheusMonitor struct {
	Metrics map[string]*prometheus.GaugeVec
	service string
	regex   *regexp.Regexp

	logger *zap.SugaredLogger
}

func (m *PrometheusMonitor) init() {
	m.logger.Debugf("Registering metrics: %v", m.Metrics)
	for _, metric := range m.Metrics {
		err := prometheus.Register(metric)
		if err != nil {
			m.logger.Debugf("Metric already registered with the name %v", metric)
		}
	}
}

func (m *PrometheusMonitor) GetService() string {
	return m.service
}

func (m *PrometheusMonitor) GetMetric(metric string) *prometheus.GaugeVec {
	measure, ok := m.Metrics[metric]
	if !ok {
		m.logger.Debugf("No metric found with the name %s", metric)
		return nil
	}
	return measure
}

// Incr is a wrapper for the prometheus Gauge.Inc function,
func (m PrometheusMonitor) Incr(metric string, tags map[string]string) {
	// TODO @shipperizer add description to function signature
	measure := m.GetMetric(metric)
	if measure == nil {
		return
	}

	measure.With(tags).Inc()
}

// NewPrometheusMonitor creates a new monitor object
func NewPrometheusMonitor(cfg PrometheusConfig) monitoring.MonitorInterface {
	m := new(PrometheusMonitor)
	m.Metrics = cfg.Metrics
	if m.Metrics == nil {
		m.Metrics = GaugeVecs(GaugesV1())
	}

	m.regex = regexp.MustCompile(monitoring.IDPathRegex)
	m.service = cfg.Service
	m.logger = cfg.Logger

	if m.logger == nil {
		m.logger = zap.NewNop().Sugar()
	}

	m.init()

	return m
}
