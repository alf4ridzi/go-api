package utils

import (
	"api/initializers"
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
