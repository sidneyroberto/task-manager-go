package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI(variableName string, envFilePath string) string {
	var err error

	if envFilePath != "" {
		err = godotenv.Load(envFilePath)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(variableName)
}
