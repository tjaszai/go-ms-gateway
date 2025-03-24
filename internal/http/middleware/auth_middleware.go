package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/service"
)

type AuthMiddleware struct {
	SecurityService *service.SecurityService
}

func NewAuthMiddleware(s *service.SecurityService) *AuthMiddleware {
	return &AuthMiddleware{
		SecurityService: s,
	}
}

func (m *AuthMiddleware) Check(c *fiber.Ctx) error {
	if err := m.SecurityService.Auth(c); err != nil {
		return c.Status(err.Code).JSON(dto.NewErrRespDto(err.Message, err.Details))
	}
	return c.Next()
}
