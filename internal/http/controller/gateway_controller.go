package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tjaszai/go-ms-gateway/internal/db"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
)

type GatewayController struct {
	DatabaseManager *db.DatabaseManager
}

func NewGatewayController(m *db.DatabaseManager) *GatewayController {
	return &GatewayController{DatabaseManager: m}
}

// HealthCheck func check the status of the application
// @Description    Check the status of the application
// @Tags           Gateway
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.MessageRespDto
// @Failure        502 {object} dto.ErrRespDto
// @Router         /api/HealthCheck [get]
func (gc *GatewayController) HealthCheck(c *fiber.Ctx) error {
	if gc.DatabaseManager.CheckConnection() != nil {
		return c.Status(fiber.StatusBadGateway).JSON(dto.NewErrRespDto("Service is down", nil))
	}
	return c.JSON(dto.NewRespDto[*string]("Service is up", nil))
}

// CallMs func provides an interface for calling registered microservices
// @Description    It provides an interface for calling registered microservices
// @Tags           Gateway
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.MessageRespDto
// @Router         /api/CallMs [post]
func (gc *GatewayController) CallMs(c *fiber.Ctx) error {
	// TODO: implement method
	return c.JSON(dto.NewRespDto[*string]("Todo.", nil))
}
