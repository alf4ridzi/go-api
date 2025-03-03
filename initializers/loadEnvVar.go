package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting root path : %v", err)
	}

	envPath := filepath.Join(rootPath, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
