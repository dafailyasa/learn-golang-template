package routes

import "github.com/gofiber/fiber/v2"

func (r *RouteConfig) ProductRoutes(prefix fiber.Router) {
	authPrefix := prefix.Group("/products")
	authPrefix.Post("", r.ProductHandler.CreateProduct)
	authPrefix.Get("/search", r.ProductHandler.Search)
}
