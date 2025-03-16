package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	client *redis.Client
}

func New(cfg *Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Redis{
		client: client,
	}
}

func (r *Redis) Close() error {
	return r.client.Close()
}

func (r *Redis) SaveToken(ctx context.Context, email, token string, expiration time.Duration) error {
	key := fmt.Sprintf("user_token:%s", email)
	return r.client.Set(ctx, key, token, expiration).Err()
}
