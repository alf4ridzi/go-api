package utils

import (
	"api/initializers"
	"os"
)

func Env(key string) string {
	initializers.LoadEnvVariables()

	return os.Getenv(key)
}
