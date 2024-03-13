package routes

import "github.com/gofiber/fiber/v2"

func (r *RouteConfig) AuthRoutes(prefix fiber.Router) {
	authPrefix := prefix.Group("/auth")
	authPrefix.Post("/register", r.AuthHandler.RegisterUser)
	authPrefix.Post("/login", r.AuthHandler.LoginUser)
}
