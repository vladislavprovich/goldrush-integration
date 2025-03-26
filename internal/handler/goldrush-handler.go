package handler

import (
	"log/slog"

	"github.com/vladislavprovich/goldrush-integration/internal/service"
)

type GoldRushHandler struct {
	service service.Service
	logger  *slog.Logger
	cfg     *Config
}

func NewGoldRushHandler(srv *service.Service, log *slog.Logger, cfg *Config) *GoldRushHandler {
	return &GoldRushHandler{
		service: *srv,
		logger:  log,
		cfg:     cfg,
	}
}
