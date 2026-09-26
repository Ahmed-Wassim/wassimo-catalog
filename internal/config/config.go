package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT         string
	DATABASE_URL string
	ENV          string
}

func Load() (*Config, error) {
	if os.Getenv("ENV") != "prod" {
		_ = godotenv.Load()
	}

	cfg := &Config{
		PORT:         getEnv("PORT", "8080"),
		DATABASE_URL: getEnv("DATABASE_URL", ""),
		ENV:          getEnv("ENV", "development"),
	}

	if cfg.DATABASE_URL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
