package defaultController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tjaszai/go-ms-gateway/internal/database"
)

// TODO: swagger annotations...

func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true, "message": "Hello World!"})
}

func HealthCheck(c *fiber.Ctx) error {
	if database.Manager.CheckConnection() != nil {
		return c.JSON(fiber.Map{"status": "down", "message": "Service is down"})
	}

	return c.JSON(fiber.Map{"status": "up", "message": "Service is up"})
}
