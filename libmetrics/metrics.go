package libmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RegisterPrometheusMetrics is used to register the prometheus metric counters
// that are used by the program
func RegisterPrometheusMetrics() {
	prometheus.MustRegister(HTTPRequestsTotal, HTTPResponseSize)
}

var (
	// HTTPRequestsTotal keeps track of the different requests that we've
	// received throughout the API
	HTTPRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The number of API requests made to the server by endpoint",
	}, []string{"method", "endpoint", "code"})

	// HTTPResponseSize keeps track of the response sizes that we've sent through
	// the API
	HTTPResponseSize = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_response_size",
		Help:    "The size of HTTP responses from the server",
		Buckets: prometheus.LinearBuckets(200, 200, 10),
	})
)
