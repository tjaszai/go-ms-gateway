package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
)

type DefaultController struct {
	DatabaseManager *db.DatabaseManager
}

func NewDefaultController(m *db.DatabaseManager) *DefaultController {
	return &DefaultController{DatabaseManager: m}
}

func (dc *DefaultController) Index(c *fiber.Ctx) error {
	return c.JSON(dto.NewRespDto[*string]("Hello world!", nil))
}

// HealthCheck func check the status of the application
// @Description    Check the status of the application
// @Tags           Default
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.MessageRespDto
// @Failure        502 {object} dto.ErrRespDto
// @Router         /HealthCheck [get]
func (dc *DefaultController) HealthCheck(c *fiber.Ctx) error {
	if dc.DatabaseManager.CheckConnection() != nil {
		return c.Status(fiber.StatusBadGateway).JSON(dto.NewErrRespDto("Service is down", nil))
	}
	return c.JSON(dto.NewRespDto[*string]("Service is up", nil))
}
