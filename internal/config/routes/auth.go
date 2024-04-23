package routes

import "github.com/gofiber/fiber/v2"

func (r *RouteConfig) authRoutes(prefix fiber.Router) {
	authGroup := prefix.Group("/auth")
	authGroup.Post("/register", r.AuthHandler.RegisterUser)
	authGroup.Post("/login", r.AuthHandler.LoginUser)
}
