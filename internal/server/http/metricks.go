package internalhttp

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetMetricksServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return mux
	//return http.ListenAndServe(address, mux)
}

var requestMetrics = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:  "cbrwsdltojson",
	Subsystem:  "http",
	Name:       "request",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
}, []string{"status"})

func observeRequest(status string, d time.Duration) {
	requestMetrics.WithLabelValues(status).Observe(d.Seconds())
}
