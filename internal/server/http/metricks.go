package internalhttp

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	RequestsTotal    *prometheus.CounterVec
	RequestsDuration *prometheus.SummaryVec
}

func GetMetricksServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return mux
}

func CreateMetrics() Metrics {
	metricks := Metrics{}
	metricks.RequestsDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "cbrwsdltojson",
		Subsystem:  "http",
		Name:       "app_request_duration",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"status", "handler"})

	metricks.RequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "cbrwsdltojson",
		Subsystem: "http",
		Name:      "app_request_counter_total",
	}, []string{"status", "handler"})

	return metricks
}

func (s *Server) observeRequestDuration(status string, handler string, d time.Duration) {
	s.metrics.RequestsDuration.WithLabelValues(status, handler).Observe(d.Seconds())
}

func (s *Server) observeRequestCounterTotal(status string, handler string) {
	s.metrics.RequestsTotal.With(prometheus.Labels{"status": status, "handler": handler}).Add(1)
}
