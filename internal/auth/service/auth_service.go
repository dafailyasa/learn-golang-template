package service

import "github.com/dafailyasa/learn-golang-template/internal/auth/model"

type AuthService interface {
	Create(body *model.AuthRegisterRequest) error
	Login(body *model.AuthLoginRequest) (*model.LoginUserResponse, error)
}
