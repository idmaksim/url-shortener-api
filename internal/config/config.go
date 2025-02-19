package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	Database *DatabaseConfig
	Http     *HTTPConfig
	Redis    *RedisConfig
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	database, err := LoadDatabaseConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading database config: %w", err)
	}

	http, err := LoadHTTPConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading HTTP config: %w", err)
	}

	redis, err := LoadRedisConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading Redis config: %w", err)
	}

	return &Config{
		Database: database,
		Http:     http,
		Redis:    redis,
	}, nil
}
