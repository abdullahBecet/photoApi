package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error .env")
	}
	mongoURI := os.Getenv("Mongo_URI")
	return mongoURI
}
