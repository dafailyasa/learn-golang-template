package repository

import "github.com/dafailyasa/learn-golang-template/internal/account/entity"

type AccountRepository interface {
	Create(account *entity.Account) (*entity.Account, error)
	FindAccountsByUser(userId string) (*[]entity.Account, error)
	FindAccountDetail(id string, userId string) (*entity.Account, error)
}
