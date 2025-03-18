package service

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
)

func (s *Service) RecentAddressTransaction(ctx context.Context, req *RecentAddressTransactionRequest, token string) (*RecentAddressTransactionResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.RecentAddressTransaction")
	defer span.End()
	s.log.InfoContext(ctx, "RecentAddressTransaction call")

	// Verify token before proceeding.
	if err := s.VerifyToken(ctx, token); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token verification failed", slog.Any("error", err))
		return nil, err
	}

	integrationReq := s.convectorToClient.ConvectorToRecentAddressTransactionRequest(req)

	clientRes, err := s.client.GetRecentTransactionForAddress(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "GetRecentTransactionForAddress", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorFromClient.ConvectorToRecentAddressTransactionResponse(clientRes)

	return res, nil
}
