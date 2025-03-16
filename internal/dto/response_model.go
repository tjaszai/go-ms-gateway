package dto

import (
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
)

// TODO: ms version dto...

type MsRespDto RespDto[*MsDto]
type MsListRespDto RespDto[[]MsDto]

type MsDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
}

func NewMsDtoFromModel(m *model.Microservice) *MsDto {
	return &MsDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: util.FromNullString(m.Description),
	}
}

func NewMsListDtoFromModels(models []model.Microservice) []MsDto {
	var mld []MsDto
	for _, m := range models {
		d := NewMsDtoFromModel(&m)
		mld = append(mld, *d)
	}
	return mld
}

type UserRespDto RespDto[*UserDto]
type UserListRespDto RespDto[[]UserDto]

type UserDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func NewUserDtoFromModel(m *model.User) *UserDto {
	return &UserDto{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}

func NewUserListDtoFromModels(models []model.User) []UserDto {
	var uld []UserDto
	for _, m := range models {
		d := NewUserDtoFromModel(&m)
		uld = append(uld, *d)
	}
	return uld
}
