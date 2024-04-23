package repository

import "github.com/dafailyasa/learn-golang-template/internal/auth/entity"

type authRepository interface {
	FindOneByEmail(email string) (*entity.User, error)
	Create(user *entity.User) error
}
