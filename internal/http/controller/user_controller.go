package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/service"
	"log"
)

type UserController struct {
	Repository *repository.UserRepository
	Validator  *service.Validator
}

func NewUserController(r *repository.UserRepository, v *service.Validator) *UserController {
	return &UserController{Repository: r, Validator: v}
}

// Create func create a user
// @Description    Create a user
// @Security       BearerAuth
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          user body dto.UserInputDto true "User dto object"
// @Success        201 {object} dto.UserRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Users [post]
func (uc *UserController) Create(c *fiber.Ctx) error {
	inputDto := new(dto.UserInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := uc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	m, err := uc.Repository.CreateFrom(inputDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to create user.", nil))
	}
	outputDto := dto.NewUserOutputDtoFromModel(m)
	return c.Status(fiber.StatusCreated).JSON(dto.NewRespDto[*dto.UserOutputDto]("User Created.", &outputDto))
}

// GetOne func get one user by ID
// @Description    Get one user by ID
// @Security       BearerAuth
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          id path string true "User ID"
// @Success        200 {object} dto.UserRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Router         /api/Users/{id} [get]
func (uc *UserController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := uc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("User not found.", nil))
	}
	outputDto := dto.NewUserOutputDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.UserOutputDto]("User Found.", &outputDto))
}

// Update func update a user by ID
// @Description    Update a user by ID
// @Security       BearerAuth
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          id path string true "User ID"
// @Param          user body dto.UserInputDto true "User dto object"
// @Success        200 {object} dto.UserRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Users/{id} [put]
func (uc *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := uc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("User not found.", nil))
	}
	inputDto := new(dto.UserInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := uc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	var err error
	m, err = uc.Repository.UpdateFrom(m, inputDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to update User.", nil))
	}
	outputDto := dto.NewUserOutputDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.UserOutputDto]("User Updated.", &outputDto))
}

// Delete func delete a user by ID
// @Description    Delete a user by ID
// @Security       BearerAuth
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          id path string true "User ID"
// @Success        200 {object} dto.MessageRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Users/{id} [delete]
func (uc *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := uc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("User not found.", nil))
	}
	err := uc.Repository.Delete(id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to delete user.", nil))
	}
	return c.JSON(dto.NewRespDto[*string]("User Deleted.", nil))
}

// GetAll func gets all existing users
// @Description    Get all existing users
// @Security       BearerAuth
// @Tags           Users
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.UserListRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Users [get]
func (uc *UserController) GetAll(c *fiber.Ctx) error {
	m, err := uc.Repository.FindAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Unexpected error.", nil))
	}
	outputListDto := dto.NewUserOutputListDtoFromModels(m)
	return c.JSON(dto.NewRespDto[[]dto.UserOutputDto]("Users Found.", &outputListDto))
}
