package handler

import (
	"log"
	"log/slog"

	"github.com/vladislavprovich/goldrush-integration/internal/middleware"
	"github.com/vladislavprovich/goldrush-integration/pkg/slog_writer"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func InitRouter(grHandler *GoldRushHandler, logger *slog.Logger, cfg *Config) *chi.Mux {
	r := chi.NewRouter()

	wrappedLogger := log.New(&slog_writer.SlogWriter{Logger: logger}, "", 0)

	// Middleware stack
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Logger)                                           // Default logger.
	r.Use(chiMiddleware.RequestLogger(&chiMiddleware.DefaultLogFormatter{ // Informative logger.
		Logger:  wrappedLogger,
		NoColor: false,
	}))
	r.Use(chiMiddleware.Timeout(cfg.TimeOut))
	r.Use(middleware.CORS)

	// API Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/register", grHandler.UserRegister)
		r.Post("/login", grHandler.UserLogin)

		r.Group(func(r chi.Router) {
			r.Use(middleware.RequireAuth(cfg.SecretToValidationJWT))
			r.Get("/transaction", grHandler.GetTransactionInfo)
			r.Get("/balances", grHandler.GetTokenBalances)
			r.Get("/recent-transactions", grHandler.GetRecentAddressTransaction)
			r.Get("/historical-transactions", grHandler.GetHistoricalPortfolioTransaction)
		})
	})
	logger.Info("Router initialized successfully")

	return r
}
