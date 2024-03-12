package handler

import (
	"github.com/dafailyasa/learn-golang-template/internal/auth/model"
	"github.com/dafailyasa/learn-golang-template/internal/auth/service"
	util "github.com/dafailyasa/learn-golang-template/utils"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) authHandler {
	return authHandler{
		AuthService: authService,
	}
}

func (h *authHandler) RegisterUser(ctx *fiber.Ctx) error {
	body := new(model.AuthRegisterRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	if err := h.AuthService.Create(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON("OK")
}

func (h *authHandler) LoginUser(ctx *fiber.Ctx) error {
	body := new(model.AuthLoginRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err})
	}

	data, err := h.AuthService.Login(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.ApiResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(util.ApiResponse{
		Data: data,
	})
}
