package service

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
)

func (s *Service) TransactionInfo(ctx context.Context, req *TransactionInfoRequest) (*TransactionInfoResponse, error) {
	ctx, span := s.tracer.Start(context.Background(), "service.TransactionInfo")
	defer span.End()
	s.log.InfoContext(ctx, "transaction info call")

	integrationReq := s.convectorToClient.ConvectorToTransactionInfoRequest(req)

	clientRes, err := s.client.GetTransaction(ctx, integrationReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "error getting transaction info", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorFromClient.ConvectorToTransactionInfoResponse(clientRes)

	return res, nil
}
