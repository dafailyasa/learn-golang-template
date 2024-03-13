package routes

import (
	"github.com/dafailyasa/learn-golang-template/internal/auth/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type RouteConfig struct {
	App         *fiber.App
	AuthHandler handler.AuthHandler
}

func (r *RouteConfig) Setup() {
	r.PublicRoute()
	r.PrivateRoute()
}

func (r *RouteConfig) PublicRoute() {
	r.App.Get("/", func(ctx *fiber.Ctx) error { return ctx.Status(fiber.StatusOK).SendString("Golang Template") })
	r.App.Get("/metrics", monitor.New(monitor.Config{Title: "Golang Template Metrics"}))
}

func (r *RouteConfig) PrivateRoute() {
	v1Prefix := r.App.Group("/api/v1")

	v1Prefix.Post("/auth/register", r.AuthHandler.RegisterUser)
	v1Prefix.Post("/auth/login", r.AuthHandler.LoginUser)
}
