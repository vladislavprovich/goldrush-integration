package service

import (
	"goldrush-integration/internal/repository"
)

type ConvectorToRepository struct{}

func NewConvectorToRepository() *ConvectorToRepository {
	return &ConvectorToRepository{}
}

func (c *ConvectorToRepository) ConvectorToSaveToken(userID, token string) *repository.SaveTokenRequest {
	return &repository.SaveTokenRequest{
		UserID: userID,
		Token:  token,
	}
}

func (c *ConvectorToRepository) ConvectorToGetToken(userID string) *repository.GetTokenRequest {
	return &repository.GetTokenRequest{
		UserID: userID,
	}
}
