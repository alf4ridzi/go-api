package initializers

import (
	"os"
)

func GetAuthSecret() string {
	return os.Getenv("AUTH_SECRET")
}

func GetRefreshSecret() string {
	return os.Getenv("REFRESH_SECRET")
}
