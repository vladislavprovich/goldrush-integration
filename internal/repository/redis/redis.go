package redis

import (
	"context"
	"fmt"
	defaultLog "log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		//Username:     cfg.Username,
		//Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolTimeout:  cfg.PoolTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		defaultLog.Fatalf("redis ping failed: %v", err)
		return nil
	}

	fmt.Println("Connected to Redis:", cfg.Addr)
	return &RedisClient{Client: client}
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
