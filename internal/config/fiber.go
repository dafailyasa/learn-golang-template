package config

import (
	"fmt"
	"time"

	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/spf13/viper"
)

type Fiber struct {
	Fiber  *fiber.App
	Config *viper.Viper
}

func NewFiber(config *viper.Viper) *Fiber {
	app := fiber.New(fiber.Config{
		AppName:           config.GetString("app.name"),
		ErrorHandler:      NewErrorHandler(),
		Prefork:           config.GetBool("web.prefork"),
		EnablePrintRoutes: config.GetBool("web.printRoute"),
	})

	app.Use(cors.New())

	// app limiter
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        5,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"errors": customErr.ErrToManyRequest.Error(),
			})
		},
	}))

	return &Fiber{
		Fiber:  app,
		Config: config,
	}
}

func (f *Fiber) Run() error {
	port := f.Config.GetInt("app.port")
	return f.Fiber.Listen(fmt.Sprintf(":%d", port))
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
