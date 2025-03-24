package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/service"
	"github.com/tjaszai/go-ms-gateway/internal/util"
	"log"
)

type SecurityController struct {
	Repository      *repository.UserRepository
	Validator       *service.Validator
	SecurityService *service.SecurityService
}

func NewSecurityController(r *repository.UserRepository, v *service.Validator, s *service.SecurityService) *SecurityController {
	return &SecurityController{Repository: r, Validator: v, SecurityService: s}
}

// Login func is used to authenticate the user
// @Description    User authentication
// @Tags           Security
// @Accept         json
// @Produce        json
// @Param          user body dto.LoginInputDto true "LoginInputDto dto object"
// @Success        201 {object} dto.DataRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Security/Login [post]
func (sc *SecurityController) Login(c *fiber.Ctx) error {
	inputDto := new(dto.LoginInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := sc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	u, _ := sc.Repository.FindByEmail(inputDto.Email)
	if u.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("User not found.", nil))
	}
	if !util.CompareUserPassword(u.Password, inputDto.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.NewErrRespDto("Invalid credentials.", nil))
	}
	token, err := sc.SecurityService.GenerateToken(u)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to generate token.", nil))
	}
	outputDto := map[string]string{"token": *token}
	return c.Status(fiber.StatusCreated).JSON(dto.NewRespDto[map[string]string]("Login successfully.", &outputDto))
}
