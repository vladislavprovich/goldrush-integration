package repository

import "time"

type Config struct {
	Expiration time.Duration `json:"expiration" default:"24h"`
}
