package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	RedisHost        string
	RedisPort        string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
	}

	return config, nil
}
