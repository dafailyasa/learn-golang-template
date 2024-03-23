package config

import (
	"fmt"

	accountHdl "github.com/dafailyasa/learn-golang-template/internal/account/handler"
	accountRepo "github.com/dafailyasa/learn-golang-template/internal/account/repository"
	accountService "github.com/dafailyasa/learn-golang-template/internal/account/service"
	authHdl "github.com/dafailyasa/learn-golang-template/internal/auth/handler"
	authRepo "github.com/dafailyasa/learn-golang-template/internal/auth/repository"
	authService "github.com/dafailyasa/learn-golang-template/internal/auth/service"
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
	authRepo := authRepo.NewAuthRepository(config.DB)
	accountRepo := accountRepo.NewAccountRepository(config.DB)

	// service
	authService := authService.NewAuthService(authRepo, config.DB, tokenMaker, config.Config)
	accountService := accountService.NewAccountService(authRepo, accountRepo, config.DB)

	// handler
	authHandler := authHdl.NewAuthHandler(authService, config.Validate)
	accountHandler := accountHdl.NewAccountHandler(accountService, *config.Validate)

	routeConfig := routes.RouteConfig{
		App:            config.App,
		Maker:          tokenMaker,
		AuthHandler:    authHandler,
		AccountHandler: accountHandler,
	}

	routeConfig.Setup()
}
