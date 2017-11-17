package irc

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	namespace string
	metrics   map[string]prometheus.Metric
}

func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		namespace: namespace,
		metrics:   make(map[string]prometheus.Metric),
	}
}

func (m *Metrics) NewCounter(subsystem, name, help string) prometheus.Counter {
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: m.namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
	)

	key := fmt.Sprintf("%s_%s", subsystem, name)
	m.metrics[key] = counter
	prometheus.MustRegister(counter)

	return counter
}

func (m *Metrics) NewCounterFunc(subsystem, name, help string, f func() float64) prometheus.CounterFunc {
	counter := prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Namespace: m.namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		f,
	)

	key := fmt.Sprintf("%s_%s", subsystem, name)
	m.metrics[key] = counter
	prometheus.MustRegister(counter)

	return counter
}

func (m *Metrics) NewGauge(subsystem, name, help string) prometheus.Gauge {
	guage := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: m.namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
	)

	key := fmt.Sprintf("%s_%s", subsystem, name)
	m.metrics[key] = guage
	prometheus.MustRegister(guage)

	return guage
}

func (m *Metrics) NewGaugeFunc(subsystem, name, help string, f func() float64) prometheus.GaugeFunc {
	guage := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: m.namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		f,
	)

	key := fmt.Sprintf("%s_%s", subsystem, name)
	m.metrics[key] = guage
	prometheus.MustRegister(guage)

	return guage
}

func (m *Metrics) Counter(subsystem, name string) prometheus.Counter {
	key := fmt.Sprintf("%s_%s", subsystem, name)
	return m.metrics[key].(prometheus.Counter)
}

func (m *Metrics) Gauge(subsystem, name string) prometheus.Gauge {
	key := fmt.Sprintf("%s_%s", subsystem, name)
	return m.metrics[key].(prometheus.Gauge)
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

func (m *Metrics) Run(addr string) {
	http.Handle("/", m.Handler())
	log.Infof("metrics endpoint listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
