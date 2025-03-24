package handler

import "time"

type Config struct {
	SecretToValidationJWT string        `json:"secret_to_validation_jwt" env:"HANDLER_SECRET_TO_VALIDATION_JWT"`
	TimeOut               time.Duration `json:"time_out" env:"HANDLER_TIMEOUT" envDefault:"30s"`
	Address               string        `json:"address" env:"HANDLER_ADDRESS" envDefault:":8090"`
	ReadHandlerTimeout    time.Duration `json:"read_handler_timeout" env:"HANDLER_READ_TIMEOUT" envDefault:"30s"`
}
