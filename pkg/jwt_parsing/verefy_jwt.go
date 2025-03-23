package jwt_parsing

import (
	"errors"
	"fmt"
	defaultLog "log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
	//defaultLog.Println("AAAAAAAAAAAAAA")
	parser := jwt.NewParser(jwt.WithLeeway(24 * time.Hour))
	token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		//defaultLog.Println("BBBBBBBBBBBBBBBB", token.Header["alg"])
		//defaultLog.Println("BBBBBBBBBBBBBBBB2222", token)
		//defaultLog.Println("BBBBBBBBBBBBBBBB33333", tokenString)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//defaultLog.Println("CCCCCCCCCCCCCCCCCC")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		//defaultLog.Println("DDDDDDDDDDDDDDDDDDDDDDDDD")
		// Convert secret to bytes for HMAC signing
		secretBytes := []byte(secret)
		//defaultLog.Println("JJJJJJJJJJJJJJJJJJJJJ", secretBytes)
		//defaultLog.Println("JJJJJJJJJJJJJJJJJJJJJ1111111111", secret)
		if len(secretBytes) == 0 {
			return nil, fmt.Errorf("invalid secret key format")
		}
		return secretBytes, nil
	})

	//defaultLog.Println("FFFFFFFFFFFFFFFFFFFF", token.Claims)
	//defaultLog.Println(time.Unix(1742742649, 0))

	if err != nil {
		//defaultLog.Println("FFFFFFFFFFFFFFFFFFFFF1111", err)
		//defaultLog.Println("Token parsing error:", err)
		return nil, err
	}

	//defaultLog.Println("PPPPPPPPPPPPPPPPPPPPPPPPPP", err)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		//defaultLog.Println("GGGGGGGGGGGGGGGGG", claims)
		//defaultLog.Println("GGGGGGGGGGGGGGGGG111111111", ok)
		//defaultLog.Println("GGGGGGGGGGGGGGG2222222", token.Valid)
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
