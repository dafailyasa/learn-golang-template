package middlewares

import (
	"strings"

	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"
	"github.com/dafailyasa/learn-golang-template/pkg/token"
	util "github.com/dafailyasa/learn-golang-template/utils"
	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeader     string = "Authorization"
	authorizationType       string = "Bearer"
	authorizationPayloadKey string = "auth"
)

func NewAuth(maker token.Maker) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authReq := ctx.Get(authorizationHeader, "")
		if len(authReq) == 0 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(util.ApiResponse{Errors: customErr.ErrAuthorizationNotFound.Error()})
		}

		field := strings.Split(authReq, " ")
		if len(field) < 2 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(util.ApiResponse{Errors: customErr.ErrInvalidHeaderFormat.Error()})
		}

		authType := strings.TrimSpace(field[0])
		if authType != authorizationType {
			return ctx.Status(fiber.StatusUnauthorized).JSON(util.ApiResponse{Errors: customErr.ErrUnsupportAuthType.Error()})
		}

		token := field[1]
		payload, err := maker.VerifyToken(token)
		if err != nil {
			if err := ctx.Status(fiber.StatusUnauthorized).JSON(util.ApiResponse{Errors: err.Error()}); err != nil {
				return err
			}
			return nil
		}

		ctx.Locals(authorizationPayloadKey, payload)
		return ctx.Next()
	}
}
