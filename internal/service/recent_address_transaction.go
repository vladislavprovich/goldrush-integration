package service

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/codes"
)

func (s *Service) RecentAddressTransaction(
	ctx context.Context,
	req *RecentAddressTransactionRequest,
) (*RecentAddressTransactionResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.RecentAddressTransaction")
	defer span.End()
	s.log.InfoContext(ctx, "RecentAddressTransaction call")

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
