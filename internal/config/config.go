package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string
	//TODO other configs
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Port:     getEnv("PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "INFO"),
	}, nil
}

func getEnv(key, defValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defValue
}
