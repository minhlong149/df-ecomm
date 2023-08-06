package util

import (
	"github.com/golang-jwt/jwt/v5"

	"df-ecomm/pkg/model"
)

func CreateToken(claims map[string]interface{}, secretKey string) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for key, value := range claims {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string, secretKey string) (map[string]interface{}, bool) {
	token, _ := jwt.Parse(tokenString, verify(secretKey))
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return token.Claims.(jwt.MapClaims), true
	} else {
		return nil, false
	}
}

func verify(secretKey string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrInvalidToken
		}
		return []byte(secretKey), nil
	}
}
