package service

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/codes"
)

func (s *Service) TokenBalances(ctx context.Context, req *TokenBalancesRequest) (*TokenBalancesResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.TokenBalances")
	defer span.End()
	s.log.InfoContext(ctx, "TokenBalances call")

	integrationReq := s.convectorToClient.ConvectorToTokenBalancesRequest(req)

	clientRes, err := s.client.GetTokenBalancesForAddress(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "GetTokenBalancesForAddress", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorFromClient.ConvectorToTokenBalancesResponse(clientRes)

	return res, nil
}
