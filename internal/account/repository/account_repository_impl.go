package repository

import (
	"github.com/dafailyasa/learn-golang-template/internal/account/entity"
	repository "github.com/dafailyasa/learn-golang-template/pkg/base-repository"
	"gorm.io/gorm"
)

type accountRepository struct {
	repository.Repository[entity.Account]
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

func (r *accountRepository) FindAccountsByUser(userId string) (*[]entity.Account, error) {
	var accounts *[]entity.Account

	if err := r.DB.Model(&entity.Account{}).Where("user_id = ?", userId).Find(&accounts).Error; err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (r *accountRepository) FindAccountDetail(id string, userId string) (*entity.Account, error) {
	var account entity.Account

	if err := r.DB.Model(&entity.Account{}).Where("user_id = ?", userId).Where("id = ?", id).First(&account).Error; err != nil {
		return &account, err
	}

	return &account, nil
}
