package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Kevinrestrepoh/event-simulator/logger"
	"github.com/Kevinrestrepoh/event-simulator/metrics"
	"github.com/google/uuid"
)

const requestIDKey string = "request_id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}

func Logging(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			log.Info("request",
				map[string]any{
					"method":     r.Method,
					"path":       r.URL.Path,
					"duration":   time.Since(start).String(),
					"request_id": GetRequestID(r.Context()),
				},
			)
		})
	}
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		metrics.RequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()
		metrics.RequestDuration.WithLabelValues(r.URL.Path).
			Observe(time.Since(start).Seconds())
	})
}
