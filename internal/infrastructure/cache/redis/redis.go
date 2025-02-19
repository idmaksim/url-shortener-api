package redis

import (
	"context"
	"time"

	"github.com/idmaksim/url-shortener-api/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedisCache(cfg *config.Config) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Addr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
			Username: cfg.Redis.User,
		}),
	}
}

func (c *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}
