package service

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
)

func (s *Service) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.RegisterUser")
	defer span.End()

	registerReq := s.convectorToSOO.ConvectorToSSORegisterRequest(req)

	registerSSORes, err := s.ssoService.Register(ctx, registerReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "RegisterUser call error", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToSOO.ConvectorToSSORegisterResponse(registerSSORes)

	return res, nil
}
