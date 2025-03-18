package service

import "time"

type Config struct {
	AppID      int32
	Expiration time.Duration `json:"expiration" default:"24h"`
}
