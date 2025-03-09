package microserviceController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/database"
	"github.com/tjaszai/go-ms-gateway/internal/dto/req"
	"github.com/tjaszai/go-ms-gateway/internal/dto/res"
	"github.com/tjaszai/go-ms-gateway/internal/entity"
	"github.com/tjaszai/go-ms-gateway/internal/service/validator"
)

// GetAll func gets all existing microservices
// @Description Get all existing microservices
// @Tags Microservices
// @Accept json
// @Produce json
// @Success 200 {array} res.MSResDto
// @router /api/microservice [get]
func GetAll(c *fiber.Ctx) error {
	var entities []entity.Microservice

	dbManager := database.Manager
	dbManager.GetDB().Find(&entities)

	DTOs := res.FromEntities(&entities)

	return c.JSON(fiber.Map{"success": true, "message": "Microservices Found.", "data": DTOs})
}

// Create func create a microservice entity
// @Description Create a microservice entity
// @Tags Microservices
// @Accept json
// @Produce json
// @Param microservice body req.MSReqDto true "Microservice object"
// @Success 200 {object} res.MSResDto
// @router /api/microservice [post]
func Create(c *fiber.Ctx) error {
	dto := new(req.MSReqDto)
	dbManager := database.Manager

	err := c.BodyParser(dto)
	if err != nil {
		return c.JSON(fiber.Map{"success": false, "message": "Invalid data.", "data": nil})
	}

	checker := validator.Checker

	if err := checker.ValidateDto(dto); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"success": false, "message": "Invalid data.", "data": nil, "error": err.Error()})
	}

	msEntity := dto.ToEntity()
	msEntity.ID = uuid.New()

	err = dbManager.GetDB().Create(&msEntity).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Could not create the Microservice entity.", "data": nil})
	}

	resDto := res.FromEntity(msEntity)

	return c.JSON(fiber.Map{"success": true, "message": "Created Microservice.", "data": resDto})
}
