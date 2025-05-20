package initializers

import (
	"os"
)

func GetAuthSecret() string {
	LoadEnvVariables()

	return os.Getenv("AUTH_SECRET")
}

func GetRefreshSecret() string {
	LoadEnvVariables()

	return os.Getenv("REFRESH_SECRET")
}
