package repository

import (
	"github.com/dafailyasa/learn-golang-template/internal/auth/entity"
	repository "github.com/dafailyasa/learn-golang-template/pkg/base-repository"
	"gorm.io/gorm"
)

type AuthRepository struct {
	repository.Repository[entity.User]
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

var _ authRepository = (*AuthRepository)(nil)

func (r *AuthRepository) FindOneByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Model(&entity.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *AuthRepository) Create(user *entity.User) error {
	if err := r.DB.Model(&entity.User{}).Create(user).Error; err != nil {
		return err
	}

	return nil
}
