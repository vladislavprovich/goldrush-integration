package config

import (
	"flag"
	"os"

	"github.com/vladislavprovich/goldrush-integration/internal/handler"
	"github.com/vladislavprovich/goldrush-integration/internal/repository"
	"github.com/vladislavprovich/goldrush-integration/internal/repository/redis"
	"github.com/vladislavprovich/goldrush-integration/internal/service"
	"github.com/vladislavprovich/goldrush-integration/pkg/client"
	"github.com/vladislavprovich/goldrush-integration/pkg/telemetry"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Service    service.Config
	Logger     LoggerConfig
	Tracing    telemetry.TracingConfig
	Metrics    telemetry.MetricsConfig
	Logging    telemetry.LoggingConfig
	Client     client.Config
	Redis      redis.Config
	Repository repository.Config
	Handler    handler.Config
}

type LoggerConfig struct {
	Env string `env:"LOG_ENV" default:"local"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
