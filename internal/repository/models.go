package repository

type SaveTokenRequest struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type GetTokenRequest struct {
	UserID string `json:"user_id"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}
