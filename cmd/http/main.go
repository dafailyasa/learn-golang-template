package main

import (
	"fmt"

	"github.com/dafailyasa/learn-golang-template/internal/config"
	"github.com/dafailyasa/learn-golang-template/pkg/logger"
)

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	logger := logger.NewLogger(viperConfig)
	validate := config.NewValidator()
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app.Fiber,
		DB:       db,
		Logger:   logger,
		Validate: validate,
		Config:   viperConfig,
	})

	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
