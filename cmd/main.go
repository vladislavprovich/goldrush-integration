package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	defaultLog "log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/vladislavprovich/goldrush-integration/cmd/config"
	"github.com/vladislavprovich/goldrush-integration/internal/handler"
	"github.com/vladislavprovich/goldrush-integration/internal/repository"
	redisintegration "github.com/vladislavprovich/goldrush-integration/internal/repository/redis"
	"github.com/vladislavprovich/goldrush-integration/internal/service"
	"github.com/vladislavprovich/goldrush-integration/pkg/client"
	"github.com/vladislavprovich/goldrush-integration/pkg/telemetry"
	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	if cfg == nil {
		panic("Config is nil!")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log := setupLogger(ctx, cfg)

	err := telemetry.EnsureLogDir(cfg.Logging.LogDir)
	if err != nil {
		log.Error("failed to ensure log dir",
			slog.String("dir", cfg.Logging.LogDir),
			slog.String("error ", err.Error()))
	}

	_, err = telemetry.InitMetrics(ctx, log, &cfg.Metrics)
	if err != nil {
		defaultLog.Fatalf("failed to init metrics: %v", err)
	}

	tracerProvider, err := telemetry.InitTracing(ctx, cfg.Metrics.Endpoint, log)
	if err != nil {
		defaultLog.Fatalf("failed to init tracing: %v", err)
	}
	defer func() {
		if err = tracerProvider.Shutdown(context.Background()); err != nil {
			log.Error("failed to shutdown tracer", slog.String("error", err.Error()))
		}
	}()

	redisClient := redisintegration.NewRedisClient(cfg.Redis)
	if redisClient == nil {
		defaultLog.Fatal("failed to init redis client")
	}
	defer func() {
		err = redisClient.Close()
		if err != nil {
			defaultLog.Fatalf("failed to close redis client: %v", err)
		}
	}()

	tokenRepo := repository.NewRedisTokenRepository(redisClient, cfg.Repository)

	conn, err := grpc.NewClient(cfg.Service.EndPoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: cfg.Service.MinConnectTimeout,
		}),
	)
	if err != nil {
		defaultLog.Fatalf("failed to connect to SSO service: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			defaultLog.Fatalf("failed to close SSO connection: %v", err)
		}
	}()

	// Initialize SSO client with configuration.
	initProtoSSO := ssov1.NewAuthClient(conn)
	if err != nil {
		defaultLog.Fatalf("failed to init SSO client: %v", err)
	}

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:    rootCAs,
			MinVersion: tls.VersionTLS12,
		},
	}

	httpClient := &http.Client{
		Timeout:   cfg.Handler.TimeOut,
		Transport: tr,
	}
	integrationClient := client.NewClient(httpClient, tracerProvider, log, &cfg.Client)

	params := service.Params{
		Log:        log,
		Cfg:        &cfg.Service,
		SSOService: initProtoSSO,
		Storage:    tokenRepo,
		Client:     *integrationClient,
	}

	grService := service.NewService(params)

	handlerGoldRush := handler.NewGoldRushHandler(grService, log, &cfg.Handler)

	r := handler.InitRouter(handlerGoldRush, log, &cfg.Handler)

	// Start HTTP server
	server := &http.Server{
		Addr:              cfg.Handler.Address,
		ReadHeaderTimeout: cfg.Handler.ReadHandlerTimeout,
		Handler:           r,
	}

	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("failed to start server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	log.Info("Server started", slog.String("address", cfg.Handler.Address))

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second)
	defer shutdownCancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		log.Error("failed to shutdown server gracefully", slog.String("error", err.Error()))
	}

	log.Info("Server stopped gracefully")
}

func setupLogger(ctx context.Context, cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	logFilePath := filepath.Join(cfg.Logging.LogDir, "app.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// Use io.MultiWriter to write logs to both stdout and file.
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	switch cfg.Logger.Env {
	case envLocal:
		log = slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev, envProd:
		log = slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		log = slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	// Now that log is initialized, we can handle the error
	if err != nil {
		log.ErrorContext(ctx, "failed to open log file",
			slog.String("logFilePath", logFilePath),
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	return log
}
