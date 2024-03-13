package service

import (
	"errors"

	"github.com/dafailyasa/learn-golang-template/internal/auth/entity"
	"github.com/dafailyasa/learn-golang-template/internal/auth/model"
	"github.com/dafailyasa/learn-golang-template/internal/auth/repository"
	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"
	"github.com/dafailyasa/learn-golang-template/pkg/token"
	util "github.com/dafailyasa/learn-golang-template/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type authService struct {
	DB         *gorm.DB
	AuthRepo   repository.AuthRepository
	Config     *viper.Viper
	TokenMaker token.Maker
}

func NewAuthService(authRepo repository.AuthRepository, db *gorm.DB, token token.Maker, config *viper.Viper) *authService {
	return &authService{
		AuthRepo:   authRepo,
		DB:         db,
		Config:     config,
		TokenMaker: token,
	}
}

func (s *authService) Create(body *model.AuthRegisterRequest) error {
	hashedPass, err := util.HashPassword(body.Password)
	if err != nil {
		return err
	}

	user := new(entity.User)
	user.Email = body.Email
	user.Name = body.Name
	user.Password = hashedPass

	if err := s.AuthRepo.Create(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return customErr.ErrEmailDuplicate
		}
	}

	return nil
}

func (s *authService) Login(body *model.AuthLoginRequest) (*model.LoginUserResponse, error) {
	user, err := s.AuthRepo.FindOneByEmail(body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErr.ErrUserNotFound
		}

		return nil, err
	}

	if err = util.CheckPassword(body.Password, user.Password); err != nil {
		return nil, customErr.ErrIncorrectPassword
	}

	accessTokenDuration := s.Config.GetDuration("jwt.accessTokenExpired")
	accessToken, accessTokenPayload, err := s.TokenMaker.CreateToken(user.Email, accessTokenDuration)
	if err != nil {
		return nil, customErr.ErrFailedCreateAccessToken
	}

	refreshTokenDuration := s.Config.GetDuration("jwt.refreshTokenExpired")
	refreshToken, refreshTokenPayload, err := s.TokenMaker.CreateToken(user.Email, refreshTokenDuration)
	if err != nil {
		return nil, customErr.ErrFailedCreateRefreshToken
	}

	return &model.LoginUserResponse{
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  accessTokenPayload.ExpiresAt.Time,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshTokenPayload.ExpiresAt.Time,
	}, nil

}
