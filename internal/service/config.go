package service

type Config struct {
	AppID int32 `json:"app_id" env:"SERVICE_APP_ID" envDefault:"0"`
}
