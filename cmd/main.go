package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tjaszai/go-ms-gateway/internal/config"
	"github.com/tjaszai/go-ms-gateway/internal/database"
	"github.com/tjaszai/go-ms-gateway/internal/service/router"
	"log"
	"time"
)

func main() {
	app := fiber.New()
	// Database
	database.InitConnection(5, 5*time.Second)
	// Routes
	router.SetupRoutes(app)
	// Middlewares
	app.Use(logger.New())

	port := config.Config("APP_PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
