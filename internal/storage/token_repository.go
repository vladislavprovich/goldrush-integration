package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"goldrush-integration/internal/storage/redisinit"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, req *SaveTokenRequest) error
	GetToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error)
	Close() error
}

type RedisTokenRepository struct {
	client *redisinit.RedisClient
}

func NewRedisTokenRepository(client *redisinit.RedisClient) *RedisTokenRepository {
	return &RedisTokenRepository{client: client}
}

func (r *RedisTokenRepository) SaveToken(ctx context.Context, req *SaveTokenRequest) error {
	key := fmt.Sprintf("user_token:%s", req.UserID)
	return r.client.Client.Set(ctx, key, req.Token, req.Expiration).Err()
}

func (r *RedisTokenRepository) GetToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error) {
	key := fmt.Sprintf("user_token:%s", req.UserID)
	token, err := r.client.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("token not found for email: %s", req.UserID)
	} else if err != nil {
		return nil, err
	}
	return &GetTokenResponse{Token: token}, nil
}

func (r *RedisTokenRepository) Close() error {
	return r.client.Client.Close()
}
