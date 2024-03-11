package repository

import (
	"github.com/dafailyasa/learn-golang-template/internal/auth/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) FindOneByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.DB.Model(&entity.User{}).Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) Create(user *entity.User) error {
	user.ID = uuid.NewString()
	if err := r.DB.Model(&entity.User{}).Create(user).Error; err != nil {
		return err
	}

	return nil
}
