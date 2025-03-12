package dto

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/model"
)

type MsRespDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
}

func NewMsRespDtoFromModel(m *model.Microservice) *MsRespDto {
	return &MsRespDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
	}
}

func NewMsRespListDtoFromModels(models []model.Microservice) []MsRespDto {
	var mrd []MsRespDto
	for _, m := range models {
		d := NewMsRespDtoFromModel(&m)
		mrd = append(mrd, *d)
	}
	return mrd
}
