package repository

import (
	"github.com/dafailyasa/learn-golang-template/internal/account/entity"
	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

var _ AccountRepository = (*accountRepository)(nil)

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{
		DB: db,
	}
}

func (r *accountRepository) Create(account *entity.Account) (*entity.Account, error) {
	res := r.DB.Model(&entity.Account{}).Create(account)
	if res.Error != nil {
		return nil, res.Error
	}

	return account, nil
}
