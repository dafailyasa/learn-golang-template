package model

import (
	"time"
)

type LoginUserResponse struct {
	//SessionUID            uuid.UUID `json:"session_id"`
	AccessTokenExpiredAt  time.Time `json:"access_token_expired_at"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
	RefreshToken          string    `json:"refresh_token"`
	AccessToken           string    `json:"access_token"`
}
