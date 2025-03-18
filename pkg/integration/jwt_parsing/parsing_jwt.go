package jwt_parsing

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractUserID(tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, exists := claims["user_id"].(string)
	if !exists {
		return "", fmt.Errorf("user_id not found in token")
	}

	return userID, nil
}
