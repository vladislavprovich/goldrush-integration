package redis

import "time"

type Config struct {
	Addr         string        // Redis server address.
	Password     string        // Password (if required).
	DB           int           // Database number.
	PoolSize     int           // Maximum number of connections in the pool.
	MinIdleConns int           // Minimum number of idle connections to keep open.
	DialTimeout  time.Duration // Timeout for connecting to Redis.
	ReadTimeout  time.Duration // Timeout for reading from Redis.
	WriteTimeout time.Duration // Timeout for writing to Redis.
	PoolTimeout  time.Duration // Timeout for waiting for a connection from the pool.
}
