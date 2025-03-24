package client

import (
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

type Client struct {
	httpClient *http.Client
	tracer     trace.Tracer
	log        *slog.Logger
	cfg        *Config
}

func NewClient(httpClient *http.Client, tracer trace.TracerProvider, log *slog.Logger, cfg *Config) *Client {
	return &Client{
		httpClient: httpClient,
		tracer:     tracer.Tracer("client"),
		log:        log,
		cfg:        cfg,
	}
}
