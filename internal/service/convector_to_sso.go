package service

import (
	ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"
)

type SSOConverter struct{}

func NewConvectorToSSO() *SSOConverter {
	return &SSOConverter{}
}

func (s *SSOConverter) ConvectorToSSORegisterRequest(req *RegisterUserRequest) *ssov1.RegisterRequest {
	return &ssov1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (s *SSOConverter) ConvectorToSSORegisterResponse(res *ssov1.RegisterResponse) *RegisterUserResponse {
	return &RegisterUserResponse{
		UserID: res.UsedId,
	}
}

func (s *SSOConverter) ConvectorToSSOLoginRequest(req *LoginUserRequest, appID int32) *ssov1.LoginRequest {
	return &ssov1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		AppId:    appID,
	}
}

func (s *SSOConverter) ConvectorToSSOLoginResponse(req *ssov1.LoginResponse) *LoginUserResponse {
	return &LoginUserResponse{
		Token: req.Token,
	}
}
