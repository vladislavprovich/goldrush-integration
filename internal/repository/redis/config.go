package redis

import "time"

type Config struct {
	// Redis server address.
	Addr string `json:"addr" env:"REDIS_ADDR" envDefault:"localhost:6379"`
	// Password (if required).
	Password string `json:"password" env:"REDIS_PASSWORD" envDefault:""`
	// Database number.
	DB int `json:"db" env:"REDIS_DB" envDefault:"0"`
	// Maximum number of connections in the pool.
	PoolSize int `json:"pool_size" env:"REDIS_POOL_SIZE" envDefault:"50"`
	// Minimum number of idle connections to keep open.
	MinIdleConns int `json:"min_idle_conns" env:"REDIS_MIN_IDLE_CONNS" envDefault:"10"`
	// Timeout for connecting to Redis.
	DialTimeout time.Duration `json:"dial_timeout" env:"REDIS_DIAL_TIMEOUT" envDefault:"5s"`
	// Timeout for reading from Redis.
	ReadTimeout time.Duration `json:"read_timeout" env:"REDIS_READ_TIMEOUT" envDefault:"3s"`
	// Timeout for writing to Redis.
	WriteTimeout time.Duration `json:"write_timeout" env:"REDIS_WRITE_TIMEOUT" envDefault:"3s"`
	// Timeout for waiting for a connection from the pool.
	PoolTimeout time.Duration `json:"pool_timeout" env:"REDIS_POOL_TIMEOUT" envDefault:"5s"`
}
