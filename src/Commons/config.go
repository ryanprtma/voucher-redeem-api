package commons

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() map[string]string {
	envFile := ".env"
	if os.Getenv("NODE_ENV") == "test" {
		envFile = ".test.env"
	}

	// Load file env
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	config := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASS"),
		"database": os.Getenv("DB_NAME"),
		"app_port": os.Getenv("APP_PORT"),
	}

	return config
}
