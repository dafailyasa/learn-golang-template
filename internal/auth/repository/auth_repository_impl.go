package repository

import (
	"github.com/dafailyasa/learn-golang-template/internal/auth/entity"
	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		DB: db,
	}
}

var _ AuthRepository = (*authRepository)(nil)

func (r *authRepository) FindOneByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Model(&entity.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *authRepository) Create(user *entity.User) error {
	if err := r.DB.Model(&entity.User{}).Create(user).Error; err != nil {
		return err
	}

	return nil
}
