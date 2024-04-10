package helpers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadGodotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

}
