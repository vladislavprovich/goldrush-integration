package service

import (
	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
)

type ConverterToSSO struct{}

func NewConvectorToSSO() *ConverterToSSO {
	return &ConverterToSSO{}
}

func (s *ConverterToSSO) ConvectorToSSORegisterRequest(req *RegisterUserRequest) *ssov1.RegisterRequest {
	return &ssov1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (s *ConverterToSSO) ConvectorToSSOLoginRequest(req *LoginUserRequest, appID int32) *ssov1.LoginRequest {
	return &ssov1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		AppId:    appID,
	}
}
