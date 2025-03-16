package service

type RegisterUserRequest struct {
	Email    string
	Password string
}

type RegisterUserResponse struct {
	UserID int64
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	Token string
}
