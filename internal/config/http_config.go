package config

import (
	"github.com/caarlos0/env/v11"
)

type HTTPConfig struct {
	Port uint   `env:"HTTP_PORT"`
	Host string `env:"HTTP_HOST"`
}

func LoadHTTPConfig() (*HTTPConfig, error) {
	var config HTTPConfig

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	if config.Port == 0 {
		config.Port = 8080
	}

	return &config, nil

}
