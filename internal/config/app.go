package config

import (
	"fmt"

	"github.com/dafailyasa/learn-golang-template/internal/auth/handler"
	"github.com/dafailyasa/learn-golang-template/internal/auth/repository"
	"github.com/dafailyasa/learn-golang-template/internal/auth/service"
	"github.com/dafailyasa/learn-golang-template/internal/config/routes"
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
	authRepo := repository.NewAuthRepository(config.DB)

	// service
	authService := service.NewAuthService(authRepo, config.DB, tokenMaker, config.Config)

	// handler
	authHandler := handler.NewAuthHandler(authService, config.Validate)

	routeConfig := routes.RouteConfig{
		App:         config.App,
		AuthHandler: &authHandler,
	}

	routeConfig.Setup()
}
