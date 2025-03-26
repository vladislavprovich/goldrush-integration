package middleware

import (
	"context"
	"encoding/json"
	defaultLog "log"
	"net/http"

	"github.com/vladislavprovich/goldrush-integration/pkg/tokenjwtparsing"
)

// UserIDKey is the context key for the user ID.
type contextKey string

const UserIDKey contextKey = "user_id"

const claimsUserIDKey = "user_id"

// JWTAuthMiddleware is a middleware function for JWT verification.
func JWTAuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := tokenjwtparsing.ExtractTokenFromHeader(r)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims, err := tokenjwtparsing.VerifyJWT(tokenString, secret)
		if err != nil {
			defaultLog.Printf("JWT verification failed: %v", err)
			writeJSONError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		userIDFloat, ok := claims[claimsUserIDKey].(float64)
		if !ok {
			writeJSONError(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, int(userIDFloat))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireAuth is a middleware that ensures the request has a valid JWT token.
func RequireAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return JWTAuthMiddleware(next, secret)
	}
}

// writeJSONError sends a JSON error response.
func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
