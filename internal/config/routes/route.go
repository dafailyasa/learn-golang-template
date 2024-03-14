package routes

import (
	"github.com/dafailyasa/learn-golang-template/internal/auth/handler"
	"github.com/dafailyasa/learn-golang-template/internal/config/routes/middlewares"
	"github.com/dafailyasa/learn-golang-template/pkg/token"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type RouteConfig struct {
	App         *fiber.App
	Maker       token.Maker
	AuthHandler handler.AuthHandler
}

func (r *RouteConfig) Setup() {
	v1Prefix := r.App.Group("/api/v1")

	r.PublicRoute(v1Prefix)

	authRoute := v1Prefix.Use(middlewares.NewAuth(r.Maker))
	r.PrivateRoute(authRoute)
}

func (r *RouteConfig) PublicRoute(v1Prefix fiber.Router) {
	r.App.Get("/", func(ctx *fiber.Ctx) error { return ctx.Status(fiber.StatusOK).SendString("Golang Template") })
	r.App.Get("/metrics", monitor.New(monitor.Config{Title: "Golang Template Metrics"}))

	// auth route public
	r.AuthRoutes(v1Prefix)
}

func (r *RouteConfig) PrivateRoute(authRoute fiber.Router) {
	authRoute.Get("test-authorization", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": ctx.Locals("auth")})
	})

}
