package tokenjwtparsing

import (
	"errors"
	defaultLog "log"

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
		defaultLog.Println("Claims: ", claims)
		return 0, errors.New("invalid token claims")
	}

	defaultLog.Printf("Token claims: %+v\n", claims)

	if userIDFloat, exists := claims[claimsUserID].(float64); exists {
		return int(userIDFloat), nil
	}

	return 0, errors.New("user_id not found in token")
}
