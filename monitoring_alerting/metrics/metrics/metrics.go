package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "cacher"

var (
	// total requests
	RequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "requests_total",
		Help:      "Total requests",
	}, []string{"handler"})
	// request duration
	RequestDurationSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds",
	}, []string{"handler"})
)

func init() {
	prometheus.MustRegister(
		RequestsTotal,
		RequestDurationSeconds,
	)
}
