package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/service"
	"log"
)

type SecurityController struct {
	Repository      *repository.UserRepository
	ModelValidator  *service.ModelValidator
	SecurityService *service.SecurityService
}

func NewSecurityController(r *repository.UserRepository, v *service.ModelValidator, s *service.SecurityService) *SecurityController {
	return &SecurityController{Repository: r, ModelValidator: v, SecurityService: s}
}

// Login func is used to authenticate the user
// @Description    User authentication
// @Tags           Security
// @Accept         json
// @Produce        json
// @Param          user body dto.LoginUserReqDto true "LoginUser dto object"
// @Success        201 {object} dto.DataRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Auth/Login [post]
func (sc *SecurityController) Login(c *fiber.Ctx) error {
	reqDto := new(dto.LoginUserReqDto)
	if err := c.BodyParser(reqDto); err != nil {
		log.Println(err)
		return c.JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := sc.ModelValidator.Validate(reqDto); err != nil {
		log.Println(err)
		errList := []string{err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", errList))
	}
	u, _ := sc.Repository.FindByEmail(reqDto.Email)
	if u.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("User not found.", nil))
	}
	if !u.CheckPassword(reqDto.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrRespDto("Invalid credentials.", nil))
	}
	token, err := sc.SecurityService.GenerateToken(u)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to generate token.", nil))
	}
	dDto := map[string]string{"token": *token}
	return c.JSON(dto.NewRespDto[map[string]string]("Login successfully.", &dDto))
}
