package telemetry

import "time"

type MetricsConfig struct {
	Endpoint          string        `env:"OTEL_ENDPOINT" envDefault:"http://otel-lgtm:4317"`
	MetricsPort       int           `env:"OTEL_METRICS_PORT" envDefault:"9090"`
	Adr               string        `env:"OTEL_ADR" envDefault:"localhost"`
	ReadTimeout       time.Duration `env:"OTEL_READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout      time.Duration `env:"OTEL_WRITE_TIMEOUT" envDefault:"5s"`
	ReadHeaderTimeout time.Duration `env:"OTEL_READ_HEADER_TIMEOUT" envDefault:"5s"`
}

type LoggingConfig struct {
	LevelLoki string `env:"LOG_LEVEL_LOKI" envDefault:"info"`
	Format    string `env:"LOG_FORMAT" envDefault:"json"`
	LokiURL   string `env:"LOG_LOKI_URL" envDefault:"http://loki:3100"`
	LogDir    string `env:"LOG_DIR" envDefault:"./logs"`
}

type TracingConfig struct {
	TempoURL  string `env:"TRACING_TEMPO_URL" envDefault:"http://tempo:14268"`
	NameSpase string `env:"TRACING_NAME_SPASE" envDefault:"goldrush"`
}
