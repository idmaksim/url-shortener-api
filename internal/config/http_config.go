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

	return &config, nil

}
