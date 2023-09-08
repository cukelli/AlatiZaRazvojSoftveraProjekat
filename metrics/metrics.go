package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	apiHits = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_hits_total",
			Help: "Total number of hits to API endpoints",
		},
		[]string{"endpoint"},
	)
)

func Count(handler http.HandlerFunc, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiHits.WithLabelValues(endpoint).Inc()
		handler(w, r)
	}
}
