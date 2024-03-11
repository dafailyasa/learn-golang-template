package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
		EnablePrintRoutes: true,
	})

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
