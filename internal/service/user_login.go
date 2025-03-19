package service

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"goldrush-integration/pkg/integration/jwt_parsing"
	"log/slog"
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

	res := s.convectorFromSOO.ConvectorToSSOLoginResponse(loginSSORes)

	// Get userID from jwt token.
	userID, err := jwt_parsing.ExtractUserID(res.Token)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "ExtractUserID call error", slog.Any("error", err))
		return nil, err
	}

	reqToRepo := s.convectorToRepository.ConvectorToSaveToken(userID, res.Token)

	err = s.storage.SaveToken(ctx, reqToRepo)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "save token error", slog.Any("error", err))
		return nil, err
	}

	return res, nil
}
