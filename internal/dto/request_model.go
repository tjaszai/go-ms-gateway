package dto

import (
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
)

// TODO: unique name validator...
// TODO: ms version dto...

type MsReqDto struct {
	Name        string  `json:"name" validate:"required,gte=1,lte=255,regex_pattern=^[a-z-]+$"`
	Description *string `json:"description,omitempty"`
}

func (d *MsReqDto) MsReqDtoToModel(m *model.Microservice) *model.Microservice {
	if m == nil {
		m = &model.Microservice{}
	}
	m.Name = d.Name
	m.Description = util.ToNullString(d.Description)
	return m
}

type CreateUserReqDto struct {
	Name     string `json:"name" validate:"required,gte=1,lte=255"`
	Email    string `json:"email" validate:"required,gte=1,lte=255,email"`
	Password string `json:"password" validate:"required,gte=1,lte=255"`
}

func (d *CreateUserReqDto) UserReqDtoToModel() *model.User {
	m := &model.User{}
	m.Name = d.Name
	m.Email = d.Email
	m.Password = d.Password
	return m
}

type UpdateUserReqDto struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,gte=1,lte=255"`
	Email    *string `json:"email,omitempty" validate:"omitempty,gte=1,lte=255,email"`
	Password *string `json:"password,omitempty" validate:"omitempty,gte=1,lte=255"`
}

func (d *UpdateUserReqDto) UserReqDtoToModel(m *model.User) *model.User {
	if d.Name != nil {
		m.Name = *d.Name
	}
	if d.Email != nil {
		m.Email = *d.Email
	}
	if d.Password != nil {
		m.Password = *d.Password
	}
	return m
}

type LoginUserReqDto struct {
	Email    string `json:"email" validate:"required,gte=1,lte=255,email"`
	Password string `json:"password" validate:"required,gte=1,lte=255"`
}
