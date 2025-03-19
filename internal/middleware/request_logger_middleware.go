package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// RequestLogger creates a middleware that logs HTTP requests using slog.
func RequestLogger(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Process request
			next.ServeHTTP(w, r)

			// Calculate request duration
			duration := time.Since(start)

			// Log request details
			log.InfoContext(r.Context(),
				"HTTP Request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.Int("status", getStatusCode(w)),
				slog.Duration("duration", duration),
			)
		})
	}
}

// Helper function to get status code from ResponseWriter.
func getStatusCode(w http.ResponseWriter) int {
	if rw, ok := w.(interface{ Status() int }); ok {
		return rw.Status()
	}
	return http.StatusOK
}
