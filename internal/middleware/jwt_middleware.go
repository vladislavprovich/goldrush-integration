package middleware

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"goldrush-integration/internal/service"
	"goldrush-integration/pkg/integration/jwt_parsing"
	"net/http"
)

// UserIDKey is the context key for the user ID.
type contextKey string

const UserIDKey contextKey = "user_id"

// JWTAuthMiddleware is a middleware function for JWT verification.
func JWTAuthMiddleware(next http.Handler, secret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := jwt_parsing.ExtractTokenFromHeader(r)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims, err := jwt_parsing.VerifyJWT(tokenString, secret)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		userID, ok := claims[string(UserIDKey)].(string)
		if !ok || userID == "" {
			writeJSONError(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireAuth is a middleware that ensures the request has a valid JWT token.
func RequireAuth(next http.Handler, cfg service.Config) http.Handler {
	secretKey := int32ToBytes(cfg.AppID)
	return JWTAuthMiddleware(next, secretKey)
}

// writeJSONError sends a JSON error response.
func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func int32ToBytes(n int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(n))
	return buf
}
