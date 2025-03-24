package dto

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
)

type MsRespDto RespDto[*MsOutputDto]
type MsListRespDto RespDto[[]MsOutputDto]

type MsOutputDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
}

func NewMsOutputDtoFromModel(m *model.Microservice) *MsOutputDto {
	return &MsOutputDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: util.FromNullString(m.Description),
	}
}

func NewMsOutputListDtoFromModels(models []model.Microservice) []MsOutputDto {
	var mld []MsOutputDto
	for _, m := range models {
		d := NewMsOutputDtoFromModel(&m)
		mld = append(mld, *d)
	}
	return mld
}

type UserRespDto RespDto[*UserOutputDto]
type UserListRespDto RespDto[[]UserOutputDto]

type UserOutputDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func NewUserOutputDtoFromModel(m *model.User) *UserOutputDto {
	return &UserOutputDto{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}

func NewUserOutputListDtoFromModels(models []model.User) []UserOutputDto {
	var uld []UserOutputDto
	for _, m := range models {
		d := NewUserOutputDtoFromModel(&m)
		uld = append(uld, *d)
	}
	return uld
}
