package handler

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	CreateProduct(ctx *fiber.Ctx) error
	Search(ctx *fiber.Ctx) error
}
