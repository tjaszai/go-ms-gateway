package res

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/entity"
)

type MSResDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func FromEntity(service *entity.Microservice) *MSResDto {
	return &MSResDto{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
	}
}

func FromEntities(services *[]entity.Microservice) *[]MSResDto {
	var DTOs []MSResDto
	for _, ms := range *services {
		dto := FromEntity(&ms)
		DTOs = append(DTOs, *dto)
	}
	return &DTOs
}
