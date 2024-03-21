package config

import (
	"fmt"

	authHdl "github.com/dafailyasa/learn-golang-template/internal/auth/handler"
	auth "github.com/dafailyasa/learn-golang-template/internal/auth/repository"
	authService "github.com/dafailyasa/learn-golang-template/internal/auth/service"
	"github.com/dafailyasa/learn-golang-template/internal/config/routes"
	productHdl "github.com/dafailyasa/learn-golang-template/internal/product/handler"
	product "github.com/dafailyasa/learn-golang-template/internal/product/repository"
	"github.com/dafailyasa/learn-golang-template/internal/product/service"
	"github.com/dafailyasa/learn-golang-template/pkg/logger"
	"github.com/dafailyasa/learn-golang-template/pkg/token"
	"github.com/dafailyasa/learn-golang-template/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Config   *viper.Viper
	Validate *validator.Validator
	Logger   *logger.Logger
}

func Bootstrap(config *BootstrapConfig) {
	tokenMaker, err := token.NewJwtMaker(config.Config.GetString("jwt.secretKey"))
	if err != nil {
		panic(fmt.Errorf("cannot create token maker: %w", err))
	}

	// repository
	authRepo := auth.NewAuthRepository(config.DB)
	productRepo := product.NewProductRepository(config.DB)

	// service
	authService := authService.NewAuthService(authRepo, config.DB, tokenMaker, config.Config)
	productService := service.NewProductService(productRepo, authRepo, config.DB)

	// handler
	authHandler := authHdl.NewAuthHandler(authService, config.Validate)
	productHandler := productHdl.NewProductHandler(productService, config.Validate)

	routeConfig := routes.RouteConfig{
		App:            config.App,
		Maker:          tokenMaker,
		AuthHandler:    &authHandler,
		ProductHandler: &productHandler,
	}

	routeConfig.Setup()
}
