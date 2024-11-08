package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER             string
	DB_NAME             string
	DB_HOST             string
	DB_PASSWORD         string
	DB_PORT             string
	DB_REPLICA_USER     string
	DB_REPLICA_NAME     string
	DB_REPLICA_HOST     string
	DB_REPLICA_PASSWORD string
	DB_REPLICA_PORT     string
	REDIS_HOST          string
	REDIS_PORT          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		DB_USER:             os.Getenv("DB_USER"),
		DB_PASSWORD:         os.Getenv("DB_PASSWORD"),
		DB_NAME:             os.Getenv("DB_NAME"),
		DB_HOST:             os.Getenv("DB_HOST"),
		DB_PORT:             os.Getenv("DB_PORT"),
		DB_REPLICA_USER:     os.Getenv("DB_REPLICA_USER"),
		DB_REPLICA_PASSWORD: os.Getenv("DB_REPLICA_PASSWORD"),
		DB_REPLICA_NAME:     os.Getenv("DB_REPLICA_NAME"),
		DB_REPLICA_HOST:     os.Getenv("DB_REPLICA_HOST"),
		DB_REPLICA_PORT:     os.Getenv("DB_REPLICA_PORT"),
		REDIS_HOST:          os.Getenv("REDIS_HOST"),
		REDIS_PORT:          os.Getenv("REDIS_PORT"),
	}

	return config, nil
}
