package config

import "github.com/caarlos0/env/v11"

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

func LoadRedisConfig() (*RedisConfig, error) {
	var config RedisConfig

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
