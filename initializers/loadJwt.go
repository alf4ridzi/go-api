package initializers

import (
	"os"
)

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}
