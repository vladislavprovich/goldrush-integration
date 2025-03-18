package service

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
)

func (s *Service) HistoricalPortfolioValue(ctx context.Context, req *HistoricalPortfolioValueRequest, token string) (*HistoricalPortfolioValueResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.HistoricalPortfolioValue")
	defer span.End()
	s.log.InfoContext(ctx, "HistoricalPortfolioValue call")

	// Verify token before proceeding.
	if err := s.VerifyToken(ctx, token); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token verification failed", slog.Any("error", err))
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
