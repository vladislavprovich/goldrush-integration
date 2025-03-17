package service

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
	"time"
)

func (s *Service) LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.LoginUser")
	defer span.End()

	loginReq := s.convectorToSOO.ConvectorToSSOLoginRequest(req, s.cfg.AppID)

	loginSSORes, err := s.ssoService.Login(ctx, loginReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "LoginUser call error", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToSOO.ConvectorToSSOLoginResponse(loginSSORes)

	err = s.storage.SaveToken(ctx, req.Email, res.Token, time.Hour)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "save token error", slog.Any("error", err))
		return nil, err
	}

	return res, nil
}
