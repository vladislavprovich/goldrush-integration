package service

import (
	"context"
	"errors"
	"log/slog"

	"go.opentelemetry.io/otel/codes"
)

func (s *Service) HistoricalPortfolioValue(
	ctx context.Context,
	req *HistoricalPortfolioValueRequest,
) (*HistoricalPortfolioValueResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.HistoricalPortfolioValue")
	defer span.End()
	s.log.InfoContext(ctx, "HistoricalPortfolioValue call")

	if req.WalletAddress == "" || req.ChainName == "" {
		err := errors.New("wallet_address and chain_name are required")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "validation error", slog.Any("error", err))
		return nil, err
	}

	integrationReq := s.convectorToClient.ConvectorToHistoricalPortfolioValueRequest(req)

	clientRes, err := s.client.GetHistoricalPortfolioValueOverTime(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "GetHistoricalPortfolioValue", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorFromClient.ConvectorToHistoricalPortfolioValueResponse(clientRes)

	return res, nil
}
