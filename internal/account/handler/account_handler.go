package handler

import "github.com/gofiber/fiber/v2"

type AccountHandler interface {
	CreateAccount(ctx *fiber.Ctx) error
	UserAccounts(ctx *fiber.Ctx) error
	AccountDetail(ctx *fiber.Ctx) error
}
