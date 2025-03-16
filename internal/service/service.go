package service

import (
	"context"
	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"goldrush-integration/internal/storage"
	"log/slog"
	"time"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
}

type Service struct {
	tracer         trace.Tracer
	log            *slog.Logger
	cfg            *Config
	ssoService     ssov1.AuthClient
	storage        storage.Redis
	convectorToSOO *SSOConverter
}

func NewService(log *slog.Logger, cfg *Config, ssoService ssov1.AuthClient, storage storage.Redis) *Service {
	return &Service{
		tracer:         otel.Tracer("service"),
		log:            log,
		cfg:            cfg,
		ssoService:     ssoService,
		storage:        storage,
		convectorToSOO: NewConvectorToSSO(),
	}
}

func (s *Service) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.RegisterUser")
	defer span.End()

	registerReq := s.convectorToSOO.ConvectorToSSORegisterReq(req)

	registerSSORes, err := s.ssoService.Register(ctx, registerReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "RegisterUser call error", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToSOO.ConvectorToSSORegisterRes(registerSSORes)

	return res, nil
}

func (s *Service) LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.LoginUser")
	defer span.End()

	loginReq := s.convectorToSOO.ConvectorToSSOLoginReq(req, s.cfg.AppID)

	loginSSORes, err := s.ssoService.Login(ctx, loginReq)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "LoginUser call error", slog.Any("error", err))
		return nil, err
	}

	res := s.convectorToSOO.ConvectorToSSOLoginRes(loginSSORes)

	err = s.storage.SaveToken(ctx, req.Email, res.Token, time.Hour)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return res, nil
}
