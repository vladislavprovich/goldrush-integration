package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recovery creates a middleware that recovers from panics and logs them using slog.
func Recovery(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Get stack trace.
					stackTrace := debug.Stack()

					// Log the error with stack trace.
					log.ErrorContext(r.Context(),
						"panic recovered",
						slog.Any("error", err),
						slog.String("stack_trace", string(stackTrace)),
						slog.String("method", r.Method),
						slog.String("url", r.URL.String()),
					)

					// Return error response.
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					response := map[string]string{
						"error": "Internal Server Error",
					}
					_ = json.NewEncoder(w).Encode(response)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
