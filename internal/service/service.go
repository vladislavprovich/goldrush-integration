package service

import (
	"context"
	"goldrush-integration/internal/storage"
	"goldrush-integration/pkg/integration/client"
	"log/slog"

	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
	TransactionInfo(ctx context.Context, req *TransactionInfoRequest, email, token string) (*TransactionInfoResponse, error)
	RecentAddressTransaction(ctx context.Context, req *RecentAddressTransactionRequest, email, token string) (*RecentAddressTransactionResponse, error)
	TokenBalances(ctx context.Context, req *TokenBalancesRequest, email, token string) (*TokenBalancesResponse, error)
	HistoricalPortfolioValue(ctx context.Context, req *HistoricalPortfolioValueRequest, email, token string) (*HistoricalPortfolioValueResponse, error)
}

type Service struct {
	tracer                trace.Tracer
	log                   *slog.Logger
	cfg                   *Config
	ssoService            ssov1.AuthClient
	storage               storage.TokenRepository
	convectorToSOO        *ConverterToSSO
	convectorFromSOO      *ConverterFromSSO
	convectorToClient     *ConvectorToClient
	convectorFromClient   *ConvectorFromClient
	convectorToRepository *ConvectorToRepository
	client                client.Client
}

func NewService(log *slog.Logger, cfg *Config, ssoService ssov1.AuthClient, storage storage.TokenRepository) *Service {
	return &Service{
		tracer:                otel.Tracer("service"),
		log:                   log,
		cfg:                   cfg,
		ssoService:            ssoService,
		storage:               storage,
		convectorToSOO:        NewConvectorToSSO(),
		convectorFromSOO:      NewConverterFromSSO(),
		convectorToClient:     NewConvectorToClient(),
		convectorFromClient:   NewConvectorFromClient(),
		convectorToRepository: NewConvectorToRepository(),
	}
}
