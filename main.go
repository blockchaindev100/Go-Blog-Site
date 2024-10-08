package main

import (
	"log"

	"github.com/blockchaindev100/Go-Blog-Site/config"
	_ "github.com/blockchaindev100/Go-Blog-Site/docs"
	logger "github.com/blockchaindev100/Go-Blog-Site/logger"
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	routers "github.com/blockchaindev100/Go-Blog-Site/router"
	"github.com/blockchaindev100/Go-Blog-Site/service"
	"github.com/gofiber/fiber/v2"
)

// @title Swagger For BlogSite API
// @version 1.0
// @description This is a Blog site server. Here you can able to post blogs and read them.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	service.RedisInit()
	defer service.RedisConnectionClose()
	db := config.InitDB()
	app := fiber.New()
	logger := logger.Logging()
	routers.InitRouter(app, repository.AquireDatabase(db, logger), logger)
	err := app.Listen(":8080")
	if err != nil {
		logger.Error(err)
		log.Fatal("Failed to start the server\n", err)
	}
}
