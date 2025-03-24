package repository

import (
	"context"
	"errors"
	"fmt"

	redisintegration "github.com/vladislavprovich/goldrush-integration/internal/repository/redis"

	"github.com/redis/go-redis/v9"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, req *SaveTokenRequest) error
	GetToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error)
}

type RedisTokenRepository struct {
	client *redisintegration.ClientRedis
	cfg    Config
}

func NewRedisTokenRepository(client *redisintegration.ClientRedis, cfg Config) *RedisTokenRepository {
	return &RedisTokenRepository{
		client: client,
		cfg:    cfg,
	}
}

func (r *RedisTokenRepository) SaveToken(ctx context.Context, req *SaveTokenRequest) error {
	key := fmt.Sprintf("user_token:%d", req.UserID)
	return r.client.Client.Set(ctx, key, req.Token, r.cfg.Expiration).Err()
}

func (r *RedisTokenRepository) GetToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error) {
	key := fmt.Sprintf("user_token:%d", req.UserID)
	token, err := r.client.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("token not found for email: %d", req.UserID)
	} else if err != nil {
		return nil, err
	}
	return &GetTokenResponse{Token: token}, nil
}
