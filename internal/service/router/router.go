package router

import (
	"github.com/gofiber/fiber/v2"
	defaultController "github.com/tjaszai/go-ms-gateway/internal/controller/default_controller"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", defaultController.IndexAction)
	// apiRouter := app.Group("/api")
}
