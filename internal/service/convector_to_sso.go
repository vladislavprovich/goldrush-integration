package service

import ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"

type SSOConverter struct{}

func NewConvectorToSSO() *SSOConverter {
	return &SSOConverter{}
}

func (s *SSOConverter) ConvectorToSSORegisterReq(req *RegisterUserRequest) *ssov1.RegisterRequest {
	return &ssov1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (s *SSOConverter) ConvectorToSSORegisterRes(res *ssov1.RegisterResponse) *RegisterUserResponse {
	return &RegisterUserResponse{
		UserID: res.UsedId,
	}
}

func (s *SSOConverter) ConvectorToSSOLoginReq(req *LoginUserRequest, appID int32) *ssov1.LoginRequest {
	return &ssov1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		AppId:    appID,
	}
}

func (s *SSOConverter) ConvectorToSSOLoginRes(req *ssov1.LoginResponse) *LoginUserResponse {
	return &LoginUserResponse{
		Token: req.Token,
	}
}
