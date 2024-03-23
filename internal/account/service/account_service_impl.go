package service

import (
	"errors"

	"github.com/dafailyasa/learn-golang-template/internal/account/entity"
	"github.com/dafailyasa/learn-golang-template/internal/account/model"
	accountRepo "github.com/dafailyasa/learn-golang-template/internal/account/repository"
	authRepo "github.com/dafailyasa/learn-golang-template/internal/auth/repository"
	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"
	"gorm.io/gorm"
)

type accountService struct {
	AccountRepository accountRepo.AccountRepository
	AuthRepository    authRepo.AuthRepository
	DB                *gorm.DB
}

var _ AccountService = (*accountService)(nil)

func NewAccountService(authRepository authRepo.AuthRepository, accountRepository accountRepo.AccountRepository, db *gorm.DB) *accountService {
	return &accountService{
		AccountRepository: accountRepository,
		AuthRepository:    authRepository,
		DB:                db,
	}
}

func (s *accountService) CreateAccount(body *model.CreateAccountRequest, email string) (*model.CreateAccountResponse, error) {
	user, err := s.AuthRepository.FindOneByEmail(email)
	if err != nil {
		if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
			return nil, customErr.ErrUserNotFound
		}
		return nil, err
	}

	payload := &entity.Account{
		UserID:   user.ID.String(),
		Balance:  0,
		Currency: body.Currency,
	}

	account, err := s.AccountRepository.Create(payload)
	if err != nil {
		if ok := errors.Is(err, gorm.ErrDuplicatedKey); ok {
			return nil, customErr.ErrAccountCurrencyAlreadCreated
		}
		return nil, err
	}

	return &model.CreateAccountResponse{
		ID:        account.ID,
		Balance:   account.Balance,
		Currency:  account.Currency,
		CreatedAt: account.CreatedAt,
	}, nil
}
