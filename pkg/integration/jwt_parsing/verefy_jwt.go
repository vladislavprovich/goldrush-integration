package jwt_parsing

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// ExtractTokenFromHeader extracts JWT token from the Authorization header.
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}

	return tokenParts[1], nil
}

// VerifyJWT parses and verifies the JWT token using the provided secret key.
func VerifyJWT(tokenString string, secret []byte) (jwt.MapClaims, error) {
	if len(secret) == 0 {
		return nil, errors.New("JWT secret key is missing")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
