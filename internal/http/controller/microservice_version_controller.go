package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/dto"
	"github.com/tjaszai/go-ms-gateway/internal/repository"
	"github.com/tjaszai/go-ms-gateway/internal/service"
	"log"
)

type MicroserviceVersionController struct {
	MsRepository        *repository.MicroserviceRepository
	MsVersionRepository *repository.MicroserviceVersionRepository
	Validator           *service.Validator
}

func NewMicroserviceVersionController(
	r *repository.MicroserviceRepository,
	vr *repository.MicroserviceVersionRepository,
	v *service.Validator) *MicroserviceVersionController {
	return &MicroserviceVersionController{MsRepository: r, MsVersionRepository: vr, Validator: v}
}

// Create func create a microservice version
// @Description    Create a microservice version
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Param          MicroserviceVersion body dto.MsVersionInputDto true "MicroserviceVersion dto object"
// @Success        201 {object} dto.MsVersionRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id}/Versions [post]
func (mvc *MicroserviceVersionController) Create(c *fiber.Ctx) error {
	id := c.Params("id")
	ms, _ := mvc.MsRepository.Find(id, false)
	if ms.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	inputDto := new(dto.MsVersionInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := mvc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	m, err := mvc.MsVersionRepository.CreateFrom(id, inputDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to create microservice.", nil))
	}
	outputDto := dto.NewMsVersionOutputDtoFromModel(m)
	return c.Status(fiber.StatusCreated).JSON(dto.NewRespDto[*dto.MsVersionOutputDto]("Microservice Created.", &outputDto))
}

// GetOne func get one microservice version by ID
// @Description    Get one microservice version by ID
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id  path string true "Microservice ID"
// @Param          vid path string true "Microservice version ID"
// @Success        200 {object} dto.MsVersionRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id}/Versions/{vid} [get]
func (mvc *MicroserviceVersionController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	ms, _ := mvc.MsRepository.Find(id, false)
	if ms.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	vID := c.Params("vid")
	m, _ := mvc.MsVersionRepository.Find(vID)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice version not found.", nil))
	}
	outputDto := dto.NewMsVersionOutputDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.MsVersionOutputDto]("Microservice version Found.", &outputDto))
}

// Update func update a microservice by ID
// @Description    Update a microservice by ID
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id  path string true "Microservice ID"
// @Param          vid path string true "Microservice version ID"
// @Param          MicroserviceVersion body dto.MsVersionInputDto true "MicroserviceVersion dto object"
// @Success        200 {object} dto.MsVersionRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        422 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id}/Versions/{vid} [put]
func (mvc *MicroserviceVersionController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	ms, _ := mvc.MsRepository.Find(id, false)
	if ms.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	vID := c.Params("vid")
	m, _ := mvc.MsVersionRepository.Find(vID)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice version not found.", nil))
	}
	inputDto := new(dto.MsVersionInputDto)
	if err := c.BodyParser(inputDto); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", nil))
	}
	if err := mvc.Validator.ValidateObject(inputDto); err != nil {
		log.Println(err)
		errList := map[string]any{"error": err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.NewErrRespDto("Invalid request body", &errList))
	}
	var err error
	m, err = mvc.MsVersionRepository.UpdateFrom(m, inputDto)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to update Microservice.", nil))
	}
	outputDto := dto.NewMsVersionOutputDtoFromModel(m)
	return c.JSON(dto.NewRespDto[*dto.MsVersionOutputDto]("Microservice Updated.", &outputDto))
}

// Delete func delete a microservice version by ID
// @Description    Delete a microservice version by ID
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id  path string true "Microservice ID"
// @Param          vid path string true "Microservice version ID"
// @Success        200 {object} dto.MessageRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        403 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id}/Versions/{vid} [delete]
func (mvc *MicroserviceVersionController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	ms, _ := mvc.MsRepository.Find(id, false)
	if ms.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	vID := c.Params("vid")
	m, _ := mvc.MsVersionRepository.Find(vID)
	if m.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice version not found.", nil))
	}
	err := mvc.MsVersionRepository.Delete(id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Failed to delete microservice version.", nil))
	}
	return c.JSON(dto.NewRespDto[*string]("Microservice version Deleted.", nil))
}

// GetAll func gets all existing microservice versions
// @Description    Get all existing microservice versions
// @Security       BearerAuth
// @Tags           Microservices
// @Accept         json
// @Produce        json
// @Param          id path string true "Microservice ID"
// @Success        200 {object} dto.MsVersionListRespDto
// @Failure        401 {object} dto.ErrRespDto
// @Failure        404 {object} dto.ErrRespDto
// @Failure        500 {object} dto.ErrRespDto
// @Router         /api/Microservices/{id}/Versions [get]
func (mvc *MicroserviceVersionController) GetAll(c *fiber.Ctx) error {
	id := c.Params("id")
	ms, _ := mvc.MsRepository.Find(id, false)
	if ms.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.NewErrRespDto("Microservice not found.", nil))
	}
	m, err := mvc.MsVersionRepository.FindAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrRespDto("Unexpected error.", nil))
	}
	outputListDto := dto.NewMsVersionOutputListDtoFromModels(m)
	return c.JSON(dto.NewRespDto[[]dto.MsVersionOutputDto]("Microservices Found.", &outputListDto))
}
