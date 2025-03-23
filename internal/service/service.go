package service

import (
	"context"
	"log/slog"

	"github.com/vladislavprovich/goldrush-integration/internal/repository"
	"github.com/vladislavprovich/goldrush-integration/pkg/client"

	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const otelName = "goldrush-integration.service"

type UserService interface {
	RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
	TransactionInfo(ctx context.Context, req *TransactionInfoRequest) (*TransactionInfoResponse, error)
	RecentAddressTransaction(ctx context.Context, req *RecentAddressTransactionRequest) (*RecentAddressTransactionResponse, error)
	TokenBalances(ctx context.Context, req *TokenBalancesRequest) (*TokenBalancesResponse, error)
	HistoricalPortfolioValue(ctx context.Context, req *HistoricalPortfolioValueRequest) (*HistoricalPortfolioValueResponse, error)
}

type Service struct {
	tracer                trace.Tracer
	log                   *slog.Logger
	cfg                   *Config
	ssoService            ssov1.AuthClient
	storage               repository.TokenRepository
	convectorToSOO        *ConverterToSSO
	convectorFromSOO      *ConverterFromSSO
	convectorToClient     *ConvectorToClient
	convectorFromClient   *ConvectorFromClient
	convectorToRepository *ConvectorToRepository
	client                client.Client
}

type Params struct {
	Log        *slog.Logger
	Cfg        *Config
	SSOService ssov1.AuthClient
	Storage    repository.TokenRepository
	Client     client.Client
}

func NewService(p Params) *Service {
	return &Service{
		tracer:                otel.Tracer(otelName),
		log:                   p.Log,
		cfg:                   p.Cfg,
		ssoService:            p.SSOService,
		storage:               p.Storage,
		client:                p.Client,
		convectorToSOO:        NewConvectorToSSO(),
		convectorFromSOO:      NewConverterFromSSO(),
		convectorToClient:     NewConvectorToClient(),
		convectorFromClient:   NewConvectorFromClient(),
		convectorToRepository: NewConvectorToRepository(),
	}
}
