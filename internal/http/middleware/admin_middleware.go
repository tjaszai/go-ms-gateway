package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/service"
)

type AdminGuardMiddleware struct {
	SecurityService *service.SecurityService
}

func NewAdminGuardMiddleware(s *service.SecurityService) *AdminGuardMiddleware {
	return &AdminGuardMiddleware{
		SecurityService: s,
	}
}

func (m *AdminGuardMiddleware) Check(c *fiber.Ctx) error {
	if err := m.SecurityService.AdminGuard(c); err != nil {
		return c.Status(err.Code).JSON(dto.NewErrRespDto(err.Message, err.Details))
	}
	return c.Next()
}
