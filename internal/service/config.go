package service

import "time"

type Config struct {
	AppID             int32         `json:"app_id" env:"SERVICE_APP_ID" envDefault:"0"`
	EndPoint          string        `json:"end_point" env:"SERVICE_ENDPOINT" envDefault:"http://localhost:44043"`
	MinConnectTimeout time.Duration `json:"min_connect_time_out" env:"SERVICE_MIN_CONNECT_TIMEOUT" envDefault:"5s"`
}
