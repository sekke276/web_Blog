package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MongoURI string
}

func EnvMongoURI() (*Config, error) {
	cfg := Config{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg.MongoURI = os.Getenv("MONGOURI")
	cfg.Port = os.Getenv("PORT")
	return &cfg, err
}
