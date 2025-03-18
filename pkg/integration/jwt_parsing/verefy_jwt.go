package jwt_parsing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// add .env file.
var secretKey []byte

func init() {

	_ = godotenv.Load()

	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY is not set")
	}
	secretKey = []byte(secret)
}

// ExtractTokenFromHeader extracts JWT token from the Authorization header.
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid token format")
	}

	return tokenParts[1], nil
}

// VerifyJWT parses and verifies the JWT token.
func VerifyJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetUserIDFromClaims extracts user_id from JWT claims.
func GetUserIDFromClaims(claims *jwt.MapClaims) (string, error) {
	userID, ok := (*claims)["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token payload: user_id not found or not a string")
	}
	return userID, nil
}

// VerifyAndGetUserID combines token verification and user ID extraction.
func VerifyAndGetUserID(tokenString string) (string, error) {
	claims, err := VerifyJWT(tokenString)
	if err != nil {
		return "", err
	}
	return GetUserIDFromClaims(claims)
}

// CreateMiddlewareFunc creates a middleware function for JWT verification.
func CreateMiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := ExtractTokenFromHeader(r)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims, err := VerifyJWT(tokenString)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		userID, err := GetUserIDFromClaims(claims)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// WriteJSONError sends a JSON error response.
func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
