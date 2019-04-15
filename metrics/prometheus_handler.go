package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HandlerTimer = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "go_pinger",
			Subsystem: "metrics",
			Name:      "request_latency_seconds",
			Help:      "Bucketed histogram of handler timings",

			// 1ms to 5min
			Buckets: prometheus.ExponentialBuckets(.001, 2, 13),
		},
		[]string{"handler"},
	)

	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "go_pinger",
			Subsystem: "metrics",
			Name:      "requests_total",
			Help:      "number of total requests to this server",
		},
		[]string{"handler"},
	)
)

func init() {
	prometheus.MustRegister(HandlerTimer)
	prometheus.MustRegister(RequestCounter)
}
