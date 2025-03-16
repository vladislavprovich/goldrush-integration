package client

import (
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	tracer     trace.Tracer
	log        *slog.Logger
	cfg        *ConfigClient
}

func NewClient(httpClient *http.Client, tracer trace.TracerProvider, log *slog.Logger, cfg *ConfigClient) *Client {
	return &Client{
		httpClient: httpClient,
		tracer:     tracer.Tracer("client"),
		log:        log,
		cfg:        cfg,
	}
}
