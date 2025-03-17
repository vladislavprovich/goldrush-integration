package service

import (
	"context"
	"errors"
	"fmt"
)

// ErrInvalidToken is returned when the token is invalid or not found
var ErrInvalidToken = errors.New("invalid or missing token")

// VerifyToken checks if the token is valid for the given email
func (s *Service) VerifyToken(ctx context.Context, email, token string) error {
	if email == "" || token == "" {
		return ErrInvalidToken
	}

	// Get the token from storage
	storedToken, err := s.storage.GetToken(ctx, email)
	if err != nil {
		s.log.ErrorContext(ctx, "error getting token", "error", err, "email", email)
		return fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	// Compare the tokens
	if storedToken != token {
		s.log.WarnContext(ctx, "token mismatch", "email", email)
		return ErrInvalidToken
	}

	return nil
}
