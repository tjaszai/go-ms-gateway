package dto

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
)

// TODO: ms version dto...

type MsRespDto RespDto[MsDto]
type MsListRespDto RespDto[[]MsDto]

type MsDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
}

func NewMsRespDtoFromModel(m *model.Microservice) *MsDto {
	return &MsDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: util.FromNullString(m.Description),
	}
}

func NewMsRespListDtoFromModels(models []model.Microservice) []MsDto {
	var mrd []MsDto
	for _, m := range models {
		d := NewMsRespDtoFromModel(&m)
		mrd = append(mrd, *d)
	}
	return mrd
}
