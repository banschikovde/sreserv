package middleware

import (
	"net/http"
	"time"

	"github.com/vlamug/elementary-cacher/metrics"
)

func MetricsMiddleware(handler string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// increase count of requests
		metrics.RequestsTotal.WithLabelValues(handler).Inc()

		// start handler timing
		started := time.Now()

		// serve request
		next.ServeHTTP(w, r)

		// observer request duration
		sinceSeconds := float64(time.Since(started)) / float64(time.Second)
		metrics.RequestDurationSeconds.WithLabelValues(handler).Observe(sinceSeconds)
	})
}
