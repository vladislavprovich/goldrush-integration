package repository

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Expiration time.Duration `json:"expiration" env:"REPOSITORY_EXPIRATION" envDefault:"8h"`
}

// NewConfig creates a new Config with values from environment variables.
func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
