package router

import (
	"github.com/gofiber/fiber/v2"
	defaultController "github.com/tjaszai/go-ms-gateway/internal/controller/default_controller"
	microserviceController "github.com/tjaszai/go-ms-gateway/internal/controller/microservice_controller"
)

func SetupRoutes(app *fiber.App) {
	SetupDefaultRoutes(app)

	apiRouter := app.Group("/api")
	SetupMSRoutes(apiRouter)
}

func SetupDefaultRoutes(app *fiber.App) {
	app.Get("/", defaultController.Index)
	// TODO: app.Get("/docs/*", ...)
	app.Get("/HealthCheck", defaultController.HealthCheck)
}

func SetupMSRoutes(router fiber.Router) {
	msRouter := router.Group("/microservice")
	msRouter.Get("/", microserviceController.GetAll)
	msRouter.Post("/", microserviceController.Create)
}
