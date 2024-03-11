package token

import (
	"time"

	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewCustomClaim(email string, duration time.Duration) (*CustomClaim, error) {
	claims := CustomClaim{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return &claims, nil
}

func (payload *CustomClaim) Valid() error {
	expiredAt, _ := payload.GetExpirationTime()

	if expiredAt.Before(time.Now()) {
		return customErr.ErrTokenExpired
	}
	return nil
}
