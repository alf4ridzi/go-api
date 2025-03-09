package utils

import (
	"api/initializers"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(initializers.GetJwtSecret())

func CreateTokenJwt(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(3 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func VerifyTokenJwt(tokenJwt string) error {
	token, err := jwt.Parse(tokenJwt, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func DecodeJwtToken(tokenJwt string) (map[string]any, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenJwt, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func GetUsernameFromJwt(tokenJwt string) (string, error) {
	dec, err := DecodeJwtToken(tokenJwt)
	if err != nil {
		return "", err
	}

	username, ok := dec["username"].(string)
	if !ok {
		return "", errors.New("username not found")
	}

	return username, nil
}
