package service

import (
	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
)

type ConverterToSSO struct{}

func NewConvectorToSSO() *ConverterToSSO {
	return &ConverterToSSO{}
}

// convector to sso
func (s *ConverterToSSO) ConvectorToSSORegisterRequest(req *RegisterUserRequest) *ssov1.RegisterRequest {
	return &ssov1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

// convector to sso
func (s *ConverterToSSO) ConvectorToSSOLoginRequest(req *LoginUserRequest, appID int32) *ssov1.LoginRequest {
	return &ssov1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		AppId:    appID,
	}
}

type ConverterFromSSO struct{}

func NewConverterFromSSO() *ConverterFromSSO {
	return &ConverterFromSSO{}
}

// convector from sso
func (s *ConverterFromSSO) ConvectorToSSOLoginResponse(req *ssov1.LoginResponse) *LoginUserResponse {
	return &LoginUserResponse{
		Token: req.Token,
	}
}

// convector from sso
func (s *ConverterFromSSO) ConvectorToSSORegisterResponse(res *ssov1.RegisterResponse) *RegisterUserResponse {
	return &RegisterUserResponse{
		UserID: res.UsedId,
	}
}
