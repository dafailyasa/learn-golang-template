package customErr

import "errors"

var (
	// errInvalidToken is returned when access token was invalid
	ErrInvalidToken = errors.New("invalid token")
	// errTokenExpired is returned when access token was expired
	ErrTokenExpired = errors.New("token has expired")
	// ErrExpiredToken is returned when access token was expired
	ErrEmailDuplicate = errors.New("email has been taken, please use another email")
	// ErrIncorrectPassword is returned when password inccorect
	ErrIncorrectPassword = errors.New("incorrect password")
	//ErrFailedCreateAccessToken is returned when access token failed to create
	ErrFailedCreateAccessToken = errors.New("failed to create access token")
	//ErrFailedCreateRefreshToken is returned when refresh token failed to create
	ErrFailedCreateRefreshToken = errors.New("failed to create refresh token")
	//ErrUserNotFound is returned when user was not found
	ErrUserNotFound = errors.New("user was not found")
)
