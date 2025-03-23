package jwt_parsing

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

const claimsUserID = "user_id"

func ExtractUserID(tokenString string) (int, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	log.Printf("Token claims: %+v\n", claims)

	if userIDFloat, exists := claims[claimsUserID].(float64); exists {
		return int(userIDFloat), nil
	}

	return 0, fmt.Errorf("user_id not found in token")
}
