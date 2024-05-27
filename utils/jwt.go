package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "DECRET_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webToken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return t, err
	}

	return t, nil

}
