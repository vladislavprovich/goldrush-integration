package tokenjwtparsing

import (
	"errors"
	defaultLog "log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const timeValidation = 9 // Hour.

// ExtractTokenFromHeader extracts JWT token from the Authorization header.
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Token" {
		return "", errors.New("invalid authorization format")
	}

	return tokenParts[1], nil
}

// VerifyJWT parses and verifies the JWT token using the provided secret key.
func VerifyJWT(tokenString string, secret string) (jwt.MapClaims, error) {
	if secret == "" {
		return nil, errors.New("JWT secret key is missing")
	}

	defaultLog.Println("secret key is", secret)

	parser := jwt.NewParser(jwt.WithLeeway(timeValidation * time.Hour))
	token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			defaultLog.Println("unexpected signing method: ", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		// Convert secret to bytes for HMAC signing
		secretBytes := []byte(secret)
		if len(secretBytes) == 0 {
			return nil, errors.New("invalid secret key format")
		}
		return secretBytes, nil
	})

	if err != nil {
		defaultLog.Println("Token parsing error:", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		defaultLog.Println("Claims: ", claims)
		defaultLog.Println("Res ok or !ok: ", ok)
		defaultLog.Println("Token validation: ", token.Valid)
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
