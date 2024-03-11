package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	RegisterUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
}
