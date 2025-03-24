package redis

import (
	"context"
	defaultLog "log"

	"github.com/redis/go-redis/v9"
)

type ClientRedis struct {
	Client *redis.Client
}

func NewRedisClient(cfg Config) *ClientRedis {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
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

	defaultLog.Println("Connected to Redis:", cfg.Addr)
	return &ClientRedis{Client: client}
}

func (r *ClientRedis) Close() error {
	return r.Client.Close()
}
