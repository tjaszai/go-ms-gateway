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
	Repository     *repository.UserRepository
	ModelValidator *service.ModelValidator
}

func NewUserController(r *repository.UserRepository, v *service.ModelValidator) *UserController {
	return &UserController{Repository: r, ModelValidator: v}
}

// Create func create a user
// @Description    Create a user
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          user body dto.CreateUserReqDto true "User dto object"
// @Success        201 {object} dto.UserRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Users [post]
func (uc *UserController) Create(c *fiber.Ctx) error {
	reqDto := new(dto.CreateUserReqDto)
	if err := c.BodyParser(reqDto); err != nil {
		log.Println(err)
		return c.JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := uc.ModelValidator.Validate(reqDto); err != nil {
		log.Println(err)
		errList := []string{err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", errList))
	}
	m, err := uc.Repository.CreateFromReqDto(reqDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to create user.", nil))
	}
	uDto := dto.NewUserDtoFromModel(m)
	return c.Status(fiber.StatusCreated).JSON(dto.NewRespDto[*dto.UserDto]("User Created.", &uDto))
}

// GetOne func get one user by ID
// @Description    Get one user by ID
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          id path string true "User ID"
// @Success        200 {object} dto.UserRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Router         /api/Users/{id} [get]
func (uc *UserController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := uc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("User not found.", nil))
	}
	uDto := dto.NewUserDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.UserDto]("User Found.", &uDto))
}

// Update func update a user by ID
// @Description    Update a user by ID
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          id path string true "User ID"
// @Param          user body dto.UpdateUserReqDto true "User dto object"
// @Success        200 {object} dto.UserRespDto
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
	reqDto := new(dto.UpdateUserReqDto)
	if err := c.BodyParser(reqDto); err != nil {
		log.Println(err)
		return c.JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := uc.ModelValidator.Validate(reqDto); err != nil {
		log.Println(err)
		errList := []string{err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", errList))
	}
	m = reqDto.UserReqDtoToModel(m)
	if err := uc.Repository.Update(reqDto, m); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to update User.", nil))
	}
	uDto := dto.NewUserDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.UserDto]("User Updated.", &uDto))
}

// Delete func delete a user by ID
// @Description    Delete a user by ID
// @Tags           Users
// @Accept         json
// @Produce        json
// @Param          id path string true "User ID"
// @Success        200 {object} dto.MessageRespDto
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
// @Tags           Users
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.UserListRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Users [get]
func (uc *UserController) GetAll(c *fiber.Ctx) error {
	m, err := uc.Repository.FindAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Unexpected error.", nil))
	}
	uListDto := dto.NewUserListDtoFromModels(m)
	return c.JSON(dto.NewRespDto[[]dto.UserDto]("Users Found.", &uListDto))
}
