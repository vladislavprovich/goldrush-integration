package service

import (
	"goldrush-integration/internal/storage"
	"time"
)

type ConvectorToRepository struct{}

func NewConvectorToRepository() *ConvectorToRepository {
	return &ConvectorToRepository{}
}

func (c *ConvectorToRepository) ConvectorToSaveToken(userID, token string, expiration time.Duration) *storage.SaveTokenRequest {
	return &storage.SaveTokenRequest{
		UserID:     userID,
		Token:      token,
		Expiration: expiration,
	}
}

func (c *ConvectorToRepository) ConvectorToGetToken(userID string) *storage.GetTokenRequest {
	return &storage.GetTokenRequest{
		UserID: userID,
	}
}
