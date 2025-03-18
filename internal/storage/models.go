package storage

import "time"

type SaveTokenRequest struct {
	UserID     string `json:"user_id"`
	Token      string `json:"token"`
	Expiration time.Duration
}

type GetTokenRequest struct {
	UserID string `json:"user_id"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}
