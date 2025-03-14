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
	Repository     *repository.MicroserviceRepository
	ModelValidator *service.ModelValidator
}

func NewMicroserviceController(r *repository.MicroserviceRepository, v *service.ModelValidator) *MicroserviceController {
	return &MicroserviceController{Repository: r, ModelValidator: v}
}

// Create func create a microservice
// @Description    Create a microservice
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          microservice body dto.MsReqDto true "Microservice dto object"
// @Success        201 {object} dto.MsRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices [post]
func (mc *MicroserviceController) Create(c *fiber.Ctx) error {
	reqDto := new(dto.MsReqDto)
	if err := c.BodyParser(reqDto); err != nil {
		log.Println(err)
		return c.JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := mc.ModelValidator.Validate(reqDto); err != nil {
		log.Println(err)
		errList := []string{err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", errList))
	}
	m, err := mc.Repository.CreateFromReqDto(reqDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to create microservice.", nil))
	}
	resDto := dto.NewMsRespDtoFromModel(m)
	return c.Status(fiber.StatusCreated).JSON(dto.NewRespDto[*dto.MsDto]("Microservice Created.", &resDto))
}

// GetOne func get one microservice by ID
// @Description    Get one microservice by ID
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Success        200 {object} dto.MsRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id} [get]
func (mc *MicroserviceController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	m, _ := mc.Repository.Find(id)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	resDto := dto.NewMsRespDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.MsDto]("Microservice Found.", &resDto))
}

// Update func update a microservice by ID
// @Description    Update a microservice by ID
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Param          microservice body dto.MsReqDto true "Microservice dto object"
// @Success        200 {object} dto.MsRespDto
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
	reqDto := new(dto.MsReqDto)
	if err := c.BodyParser(reqDto); err != nil {
		log.Println(err)
		return c.JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := mc.ModelValidator.Validate(reqDto); err != nil {
		log.Println(err)
		errList := []string{err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", errList))
	}
	m = reqDto.MsReqToModel(m)
	if err := mc.Repository.Update(m); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to update Microservice.", nil))
	}
	resDto := dto.NewMsRespDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.MsDto]("Microservice Updated.", &resDto))
}

// Delete func delete a microservice by ID
// @Description    Delete a microservice by ID
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Success        200 {object} dto.MessageRespDto
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
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Success        200 {object} dto.MsListRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices [get]
func (mc *MicroserviceController) GetAll(c *fiber.Ctx) error {
	m, err := mc.Repository.FindAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Unexpected error.", nil))
	}
	resDto := dto.NewMsRespListDtoFromModels(m)
	return c.JSON(dto.NewRespDto[[]dto.MsDto]("Microservices Found.", &resDto))
}
