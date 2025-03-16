package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
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

func (r *Redis) SaveToken(ctx context.Context, email string, token string, expiration time.Duration) error {
	key := fmt.Sprintf("user_token:%s", email)
	return r.client.Set(ctx, key, token, expiration).Err()
}

func (r *Redis) GetToken(ctx context.Context, email string) (string, error) {
	key := fmt.Sprintf("user_token:%s", email)
	token, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("token not found for email: %s", email)
	} else if err != nil {
		return "", err
	}
	return token, nil
}

func (r *Redis) HasToken(ctx context.Context, email string) (bool, error) {
	key := fmt.Sprintf("user_token:%s", email)
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *Redis) DeleteToken(ctx context.Context, email string) error {
	key := fmt.Sprintf("user_token:%s", email)
	return r.client.Del(ctx, key).Err()
}
