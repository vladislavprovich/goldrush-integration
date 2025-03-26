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

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/sdk/trace"

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
		defaultLog.Fatalf("Config is nil! %v", cfg)
	}
	ctx := context.Background()

	log := setupLogger(ctx, cfg)

	err := initTelemetry(ctx, log, cfg)
	if err != nil {
		defaultLog.Fatalf("failed to initialize telemetry", slog.Any("error", err))
	}

	tracerProvider := initTraceProvider(ctx, cfg, log)
	defer initTraceProviderShutdown(ctx, tracerProvider)

	redisClient := initRedisClient(ctx, cfg)
	defer initRedisClientShutdown(ctx, redisClient)

	tokenRepo := initTokenRepository(ctx, redisClient, cfg)

	conn := initConnToSSO(ctx, cfg)
	defer initConnToSSOShutdown(ctx, conn)

	// Initialize SSO client with configuration.
	initProtoSSO, httpClient := initProtobufSSO(ctx, conn, cfg)

	integrationClient := initGoldRushClient(ctx, httpClient, tracerProvider, log, cfg)

	params := service.Params{
		Log:        log,
		Cfg:        &cfg.Service,
		SSOService: initProtoSSO,
		Storage:    tokenRepo,
		Client:     *integrationClient,
	}

	grService := initGoldRushService(ctx, params)

	handlerGoldRush := initGoldRushHandler(ctx, grService, log, cfg)

	r := initRouter(ctx, handlerGoldRush, log, cfg)

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
	initGracefulShutdown(ctx, server, log)

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

func initTraceProvider(ctx context.Context, cfg *config.Config, log *slog.Logger) *trace.TracerProvider {
	tracerProvider, err := telemetry.InitTracing(ctx, cfg.Metrics.Endpoint, log)
	if err != nil {
		defaultLog.Fatalf("failed to init tracing: %v", err)
	}

	return tracerProvider
}

func initTraceProviderShutdown(_ context.Context, tracerProvider *trace.TracerProvider) {
	if err := tracerProvider.Shutdown(context.Background()); err != nil {
		defaultLog.Fatalf("failed to shutdown tracer: %v", err)
	}
}

func initRedisClient(_ context.Context, cfg *config.Config) *redisintegration.ClientRedis {
	redisClient := redisintegration.NewRedisClient(cfg.Redis)
	if redisClient == nil {
		defaultLog.Fatal("failed to init redis client")
	}

	return redisClient
}

func initRedisClientShutdown(_ context.Context, redisClient *redisintegration.ClientRedis) {
	err := redisClient.Close()
	if err != nil {
		defaultLog.Fatalf("failed to close redis client: %v", err)
	}
}

func initTokenRepository(
	_ context.Context,
	redisClient *redisintegration.ClientRedis,
	cfg *config.Config,
) *repository.RedisTokenRepository {
	tokenRepo := repository.NewRedisTokenRepository(redisClient, cfg.Repository)
	return tokenRepo
}

func initConnToSSO(_ context.Context, cfg *config.Config) *grpc.ClientConn {
	conn, err := grpc.NewClient(cfg.Service.EndPoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: cfg.Service.MinConnectTimeout,
		}),
	)
	if err != nil {
		defaultLog.Fatalf("failed to connect to SSO service: %v", err)
	}

	return conn
}

func initConnToSSOShutdown(_ context.Context, conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		defaultLog.Fatalf("failed to close SSO connection: %v", err)
	}
}

func initProtobufSSO(_ context.Context, conn *grpc.ClientConn, cfg *config.Config) (ssov1.AuthClient, *http.Client) {
	initProtoSSO := ssov1.NewAuthClient(conn)

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

	return initProtoSSO, httpClient
}

func initGoldRushClient(
	_ context.Context,
	httpClient *http.Client,
	tracerProvider *trace.TracerProvider,
	log *slog.Logger,
	cfg *config.Config,
) *client.Client {
	integrationClient := client.NewClient(httpClient, tracerProvider, log, &cfg.Client)
	return integrationClient
}

func initGoldRushService(_ context.Context, params service.Params) *service.Service {
	grService := service.NewService(params)
	return grService
}

func initGoldRushHandler(
	_ context.Context,
	grService *service.Service,
	log *slog.Logger,
	cfg *config.Config,
) *handler.GoldRushHandler {
	handlerGoldRush := handler.NewGoldRushHandler(grService, log, &cfg.Handler)
	return handlerGoldRush
}

func initRouter(
	_ context.Context,
	handlerGoldRush *handler.GoldRushHandler,
	log *slog.Logger,
	cfg *config.Config,
) *chi.Mux {
	router := handler.InitRouter(handlerGoldRush, log, &cfg.Handler)
	return router
}

func initTelemetry(ctx context.Context, log *slog.Logger, cfg *config.Config) error {
	err := telemetry.EnsureLogDir(cfg.Logging.LogDir)
	if err != nil {
		log.ErrorContext(ctx, "failed to ensure log dir",
			slog.String("dir", cfg.Logging.LogDir),
			slog.String("error ", err.Error()))
		return err
	}

	// Init metrics.
	_, err = telemetry.InitMetrics(ctx, log, &cfg.Metrics)
	if err != nil {
		defaultLog.Fatalf("failed to init metrics: %v", err)
		return err
	}

	return nil
}

func initGracefulShutdown(ctx context.Context, server *http.Server, log *slog.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.ErrorContext(ctx, "failed to shutdown server gracefully", slog.String("error", err.Error()))
	}
}
