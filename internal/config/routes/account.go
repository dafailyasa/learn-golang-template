package routes

import "github.com/gofiber/fiber/v2"

func (r *RouteConfig) accountRoutes(prefix fiber.Router) {
	accountsGroup := prefix.Group("/accounts")

	accountsGroup.Post("", r.AccountHandler.CreateAccount)
	accountsGroup.Get("", r.AccountHandler.UserAccounts)
	accountsGroup.Get("/:id", r.AccountHandler.AccountDetail)
}
