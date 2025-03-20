package service

import ssov1 "github.com/vladislavprovich/protobufContract/gen/go/sso"

type ConverterFromSSO struct{}

func NewConverterFromSSO() *ConverterFromSSO {
	return &ConverterFromSSO{}
}

func (s *ConverterFromSSO) ConvectorToSSOLoginResponse(req *ssov1.LoginResponse) *LoginUserResponse {
	return &LoginUserResponse{
		Token: req.Token,
	}
}

func (s *ConverterFromSSO) ConvectorToSSORegisterResponse(res *ssov1.RegisterResponse) *RegisterUserResponse {
	return &RegisterUserResponse{
		UserID: res.UsedId,
	}
}
