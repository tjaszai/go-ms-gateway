package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tjaszai/go-ms-gateway/internal/config"
	"github.com/tjaszai/go-ms-gateway/internal/service/router"
	"log"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Use(logger.New())

	port := config.Config("APP_PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
