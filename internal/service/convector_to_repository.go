package service

import (
	"github.com/vladislavprovich/goldrush-integration/internal/repository"
)

type ConvectorToRepository struct{}

func NewConvectorToRepository() *ConvectorToRepository {
	return &ConvectorToRepository{}
}

func (c *ConvectorToRepository) ConvectorToSaveToken(userID int, token string) *repository.SaveTokenRequest {
	return &repository.SaveTokenRequest{
		UserID: userID,
		Token:  token,
	}
}

func (c *ConvectorToRepository) ConvectorToGetToken(userID int) *repository.GetTokenRequest {
	return &repository.GetTokenRequest{
		UserID: userID,
	}
}
