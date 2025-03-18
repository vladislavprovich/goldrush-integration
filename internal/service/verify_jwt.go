package service

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/codes"
	"goldrush-integration/pkg/integration/jwt_parsing"
	"log/slog"
)

// VerifyToken verifies the JWT token by extracting the user ID and checking if it exists in Redis
func (s *Service) VerifyToken(ctx context.Context, token string) error {
	ctx, span := s.tracer.Start(ctx, "service.VerifyToken")
	defer span.End()
	s.log.InfoContext(ctx, "Verifying token")

	// Extract user ID from token
	userID, err := jwt_parsing.VerifyAndGetUserID(token)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Failed to extract user ID from token", slog.Any("error", err))
		return fmt.Errorf("invalid token: %w", err)
	}

	reqToRepo := s.convectorToRepository.ConvectorToGetToken(userID)

	// Check if token exists in Redis
	res, err := s.storage.GetToken(ctx, reqToRepo)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Failed to get token from storage", slog.Any("error", err))
		return fmt.Errorf("token verification failed: %w", err)
	}

	// Verify that the stored token matches the provided token
	if res.Token != token {
		err = fmt.Errorf("token mismatch")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		s.log.ErrorContext(ctx, "Token mismatch", slog.String("user_id", userID))
		return err
	}

	return nil
}
