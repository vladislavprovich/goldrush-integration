package service

import (
	"context"
	"goldrush-integration/internal/storage"
	"goldrush-integration/pkg/integration/client"
	"log/slog"

	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
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
	tracer                              trace.Tracer
	log                                 *slog.Logger
	cfg                                 *Config
	ssoService                          ssov1.AuthClient
	storage                             storage.Redis
	convectorToSOO                      *SSOConverter
	convectorToRecentAddressTransaction *RecentAddressTransactionConvector
	convectorToTransactionInfo          *TransactionInfoConvector
	convectorToTokenBalances            *TokenBalancesConvector
	convectorToHistoricalPortfolioValue *HistoricalPortfolioValueConvector
	client                              client.Client
}

func NewService(log *slog.Logger, cfg *Config, ssoService ssov1.AuthClient, storage storage.Redis) *Service {
	return &Service{
		tracer:                              otel.Tracer("service"),
		log:                                 log,
		cfg:                                 cfg,
		ssoService:                          ssoService,
		storage:                             storage,
		convectorToSOO:                      NewConvectorToSSO(),
		convectorToRecentAddressTransaction: NewRecentAddressTransactionConvector(),
		convectorToTransactionInfo:          NewTransactionInfoConvector(),
		convectorToTokenBalances:            NewTokenBalancesConvector(),
		convectorToHistoricalPortfolioValue: NewHistoricalPortfolioValueConvector(),
	}
}

func (s *Service) TransactionInfo(ctx context.Context, req *TransactionInfoRequest, email, token string) (*TransactionInfoResponse, error) {
	ctx, span := s.tracer.Start(context.Background(), "service.TransactionInfo")
	defer span.End()
	s.log.InfoContext(ctx, "transaction info call")

	// Verify token before proceeding
	if err := s.VerifyToken(ctx, email, token); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token verification failed", slog.Any("error", err))
		return nil, err
	}

	integrationReq := s.convectorToTransactionInfo.ConvectorToTransactionInfoRequest(req)

	clientRes, err := s.client.GetTransaction(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "error getting transaction info", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToTransactionInfo.ConvectorToTransactionInfoResponse(clientRes)

	return res, nil
}

func (s *Service) RecentAddressTransaction(ctx context.Context, req *RecentAddressTransactionRequest, email, token string) (*RecentAddressTransactionResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.RecentAddressTransaction")
	defer span.End()
	s.log.InfoContext(ctx, "RecentAddressTransaction call")

	// Verify token before proceeding
	if err := s.VerifyToken(ctx, email, token); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token verification failed", slog.Any("error", err))
		return nil, err
	}

	integrationReq := s.convectorToRecentAddressTransaction.ConvectorToRecentAddressTransactionRequest(req)

	clientRes, err := s.client.GetRecentTransactionForAddress(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "GetRecentTransactionForAddress", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToRecentAddressTransaction.ConvectorToRecentAddressTransactionResponse(clientRes)

	return res, nil
}

func (s *Service) TokenBalances(ctx context.Context, req *TokenBalancesRequest, email, token string) (*TokenBalancesResponse, error) {
	ctx, span := s.tracer.Start(context.Background(), "service.TokenBalances")
	defer span.End()
	s.log.InfoContext(ctx, "TokenBalances call")

	// Verify token before proceeding
	if err := s.VerifyToken(ctx, email, token); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token verification failed", slog.Any("error", err))
		return nil, err
	}

	integrationReq := s.convectorToTokenBalances.ConvectorToTokenBalancesRequest(req)

	clientRes, err := s.client.GetTokenBalancesForAddress(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "GetTokenBalancesForAddress", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToTokenBalances.ConvectorToTokenBalancesResponse(clientRes)

	return res, nil
}

func (s *Service) HistoricalPortfolioValue(ctx context.Context, req *HistoricalPortfolioValueRequest, email, token string) (*HistoricalPortfolioValueResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.HistoricalPortfolioValue")
	defer span.End()
	s.log.InfoContext(ctx, "HistoricalPortfolioValue call")

	// Verify token before proceeding
	if err := s.VerifyToken(ctx, email, token); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token verification failed", slog.Any("error", err))
		return nil, err
	}

	integrationReq := s.convectorToHistoricalPortfolioValue.ConvectorToHistoricalPortfolioValueRequest(req)

	clientRes, err := s.client.GetHistoricalPortfolioValueOverTime(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "GetHistoricalPortfolioValue", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToHistoricalPortfolioValue.ConvectorToHistoricalPortfolioValueResponse(clientRes)

	return res, nil
}
