package dto

import (
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
	"strings"
)

type MsInputDto struct {
	Name        string  `json:"name" validate:"required,gte=1,lte=100,regex_pattern=^[a-z-]+$"`
	Description *string `json:"description,omitempty"`
}

func (d *MsInputDto) ToModel(m *model.Microservice) *model.Microservice {
	if m == nil {
		m = &model.Microservice{}
	}
	m.Name = d.Name
	m.Description = util.ToNullString(d.Description)
	return m
}

type MsVersionInputDto struct {
	Name        string  `json:"name" validate:"required,gte=1,lte=30,regex_pattern=^[a-z-]+$"`
	Description *string `json:"description,omitempty"`
	Url         string  `json:"url" validate:"required,gte=1,lte=255,http_url"`
	OpenAPIUrl  string  `json:"openapi_url" validate:"required,gte=1,lte=255,http_url"`
}

func (d *MsVersionInputDto) ToModel(m *model.MicroserviceVersion) *model.MicroserviceVersion {
	if m == nil {
		m = &model.MicroserviceVersion{}
	}
	m.Name = d.Name
	m.Description = util.ToNullString(d.Description)
	m.Url = d.Url
	m.OpenAPIUrl = d.OpenAPIUrl
	return m
}

type UserInputDto struct {
	Name     string `json:"name" validate:"required,gte=1,lte=255,regex_pattern=^[a-zA-Z-]+$"`
	Email    string `json:"email" validate:"required,gte=1,lte=255,email"`
	Password string `json:"password" validate:"required,gte=1,lte=255"`
}

func (d *UserInputDto) ToModel(m *model.User) *model.User {
	if m == nil {
		m = &model.User{}
	}
	m.Name = d.Name
	m.Email = strings.ToLower(d.Email)
	m.Password = d.Password
	return m
}

type LoginInputDto struct {
	Email    string `json:"email" validate:"required,gte=1,lte=255,email"`
	Password string `json:"password" validate:"required,gte=1,lte=255"`
}
