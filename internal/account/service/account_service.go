package service

import (
	"github.com/dafailyasa/learn-golang-template/internal/account/model"
)

type AccountService interface {
	CreateAccount(body *model.CreateAccountRequest, email string) (*model.CreateAccountResponse, error)
}
