package token

import (
	"fmt"
	"time"

	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"
	"github.com/golang-jwt/jwt/v5"
)

type JwtMaker struct {
	secretKey string
}

var minLen = 32
var errKey = fmt.Errorf("invalid key size: must be at least %d characters", minLen)

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minLen {
		return nil, errKey
	}
	return &JwtMaker{secretKey}, nil
}

func (maker *JwtMaker) CreateToken(email string, duration time.Duration) (string, *CustomClaim, error) {
	payload, err := NewCustomClaim(email, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err

}

func (maker *JwtMaker) VerifyToken(token string) (*CustomClaim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaim{}, keyFunc)
	if err != nil {
		return nil, customErr.ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*CustomClaim)
	if !ok {
		return nil, customErr.ErrInvalidToken
	}

	return claims, nil
}
