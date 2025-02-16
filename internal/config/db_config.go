package config

import "github.com/caarlos0/env/v11"

type DatabaseConfig struct {
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSL_MODE"`
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	var config DatabaseConfig

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
