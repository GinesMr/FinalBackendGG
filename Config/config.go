package Config

import (
	"github.com/joho/godotenv"
	"log"
)

func CreateConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
