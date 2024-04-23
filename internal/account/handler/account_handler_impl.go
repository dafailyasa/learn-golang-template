package handler

import (
	"github.com/dafailyasa/learn-golang-template/internal/account/model"
	"github.com/dafailyasa/learn-golang-template/internal/account/service"
	"github.com/dafailyasa/learn-golang-template/pkg/token"
	"github.com/dafailyasa/learn-golang-template/pkg/validator"
	util "github.com/dafailyasa/learn-golang-template/utils"
	"github.com/gofiber/fiber/v2"
)

type accountHandler struct {
	AccountService service.AccountService
	Validator      validator.Validator
}

func NewAccountHandler(accountService service.AccountService, validator validator.Validator) *accountHandler {
	return &accountHandler{
		AccountService: accountService,
		Validator:      validator,
	}
}

func (h *accountHandler) CreateAccount(ctx *fiber.Ctx) error {
	body := new(model.CreateAccountRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	if err := h.Validator.ValidateStruct(body); len(err) > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(util.ApiResponse{Errors: err})
	}

	if err := util.IsSupportCurrency(body.Currency); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	authLocals := ctx.Locals("auth")
	auth := authLocals.(*token.CustomClaim)

	account, err := h.AccountService.CreateAccount(body, auth.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(util.ApiResponse{Data: account})
}

func (h *accountHandler) UserAccounts(ctx *fiber.Ctx) error {
	authLocals := ctx.Locals("auth")
	auth := authLocals.(*token.CustomClaim)

	accounts, err := h.AccountService.FindAccounts(auth.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(util.ApiResponse{Data: accounts})
}

func (h *accountHandler) AccountDetail(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	authLocals := ctx.Locals("auth")
	auth := authLocals.(*token.CustomClaim)

	account, err := h.AccountService.FindAccountDetail(id, auth.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(util.ApiResponse{Data: account})
}
