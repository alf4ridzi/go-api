package utils

import (
	"api/initializers"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// create auth
func CreateAuthToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(initializers.GetAuthSecret()))
}

func CreateRefreshToken() (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(claims)
}

func VerifyTokenJwt(jwtSecret []byte, tokenJwt string) error {
	token, err := jwt.Parse(tokenJwt, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return errors.New("token is expired")
		}

		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func VerifyJwtAuth(tokenJwt string) error {
	return VerifyTokenJwt([]byte(initializers.GetAuthSecret()), tokenJwt)
}

func VerifyJwtRefresh(tokenJwt string) error {
	return VerifyTokenJwt([]byte(initializers.GetRefreshSecret()), tokenJwt)
}

func DecodeJwtToken(tokenJwt string) (map[string]any, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenJwt, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(initializers.GetAuthSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func GetUsernameFromJwtAuth(tokenJwt string) (string, error) {
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
