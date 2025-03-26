package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/service"
	"log"
)

type MicroserviceController struct {
	Repository *repository.MicroserviceRepository
	Validator  *service.Validator
}

func NewMicroserviceController(r *repository.MicroserviceRepository, v *service.Validator) *MicroserviceController {
	return &MicroserviceController{Repository: r, Validator: v}
}

// Create func create a microservice
// @Description    Create a microservice
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          Microservice body dto.MsInputDto true "Microservice dto object"
// @Success        201 {object} dto.MsRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices [post]
func (mc *MicroserviceController) Create(c *fiber.Ctx) error {
	inputDto := new(dto.MsInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := mc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	m, err := mc.Repository.CreateFrom(inputDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to create microservice.", nil))
	}
	outputDto := dto.NewMsOutputDtoFromModel(m)
	return c.Status(fiber.StatusCreated).JSON(dto.NewRespDto[*dto.MsOutputDto]("Microservice Created.", &outputDto))
}

// GetOne func get one microservice by ID
// @Description    Get one microservice by ID
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Success        200 {object} dto.MsRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id} [get]
func (mc *MicroserviceController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := mc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	outputDto := dto.NewMsOutputDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.MsOutputDto]("Microservice Found.", &outputDto))
}

// Update func update a microservice by ID
// @Description    Update a microservice by ID
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Param          Microservice body dto.MsInputDto true "Microservice dto object"
// @Success        200 {object} dto.MsRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id} [put]
func (mc *MicroserviceController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := mc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	inputDto := new(dto.MsInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := mc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	var err error
	m, err = mc.Repository.UpdateFrom(m, inputDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to update Microservice.", nil))
	}
	outputDto := dto.NewMsOutputDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.MsOutputDto]("Microservice Updated.", &outputDto))
}

// Delete func delete a microservice by ID
// @Description    Delete a microservice by ID
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Success        200 {object} dto.MessageRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id} [delete]
func (mc *MicroserviceController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := mc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	err := mc.Repository.Delete(id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to delete microservice.", nil))
	}
	return c.JSON(dto.NewRespDto[*string]("Microservice Deleted.", nil))
}

// GetAll func gets all existing microservices
// @Description    Get all existing microservices
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.MsListRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices [get]
func (mc *MicroserviceController) GetAll(c *fiber.Ctx) error {
	m, err := mc.Repository.FindAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Unexpected error.", nil))
	}
	outputListDto := dto.NewMsOutputListDtoFromModels(m)
	return c.JSON(dto.NewRespDto[[]dto.MsOutputDto]("Microservices Found.", &outputListDto))
}
