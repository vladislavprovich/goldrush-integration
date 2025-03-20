package jwt_parsing

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

const claimsUserID = "user_id"

func ExtractUserID(tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, exists := claims[claimsUserID].(string)
	if !exists {
		return "", fmt.Errorf("user_id not found in token")
	}

	return userID, nil
}
