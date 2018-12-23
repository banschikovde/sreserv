package metrics

import (
	"github.com/vlamug/elementary-cacher/metrics"
	"net/http"
)

func MetricsMiddleware(handler string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// increase count of set requests
		metrics.RequestsTotal.WithLabelValues(handler).Inc()
		next.ServeHTTP(w, r)
	})
}
