package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
)

type DefaultController struct{}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (dc *DefaultController) Index(c *fiber.Ctx) error {
	return c.JSON(dto.NewRespDto[*string]("Hello world!", nil))
}
