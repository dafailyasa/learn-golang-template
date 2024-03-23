package handler

import "github.com/gofiber/fiber/v2"

type AccountHandler interface {
	CreateAccount(ctx *fiber.Ctx) error
}
