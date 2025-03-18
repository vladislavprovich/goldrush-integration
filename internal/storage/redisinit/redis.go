package redisinit

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"goldrush-integration/internal/storage"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg *storage.Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return &RedisClient{
		Client: client,
	}, nil
}
