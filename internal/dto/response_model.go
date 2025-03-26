package dto

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tjaszai/go-ms-gateway/internal/contract"
	"github.com/tjaszai/go-ms-gateway/internal/model"
	"github.com/tjaszai/go-ms-gateway/internal/util"
)

type MsRespDto RespDto[*MsOutputDto]
type MsListRespDto RespDto[[]MsOutputDto]

type MsOutputDto struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Links       []contract.Link `json:"links,omitempty"`
}

func NewMsOutputDtoFromModel(m *model.Microservice) *MsOutputDto {
	return &MsOutputDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: util.FromNullString(m.Description),
		Links: []contract.Link{
			{
				Rel:  "self",
				Href: fmt.Sprintf("/api/Microservices/%s", m.ID),
				Type: "get",
			},
			{
				Rel:  "versions",
				Href: fmt.Sprintf("/api/Microservices/%s/Versions", m.ID),
				Type: "get",
			},
		},
	}
}

func NewMsOutputListDtoFromModels(models []model.Microservice) []MsOutputDto {
	var ld []MsOutputDto
	for _, m := range models {
		d := NewMsOutputDtoFromModel(&m)
		ld = append(ld, *d)
	}
	return ld
}

type MsVersionRespDto RespDto[*MsVersionOutputDto]

type MsVersionListRespDto RespDto[[]MsVersionOutputDto]

type MsVersionOutputDto struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Url         string          `json:"url"`
	OpenAPIUrl  string          `json:"openapi_url"`
	Links       []contract.Link `json:"links,omitempty"`
}

func NewMsVersionOutputDtoFromModel(m *model.MicroserviceVersion) *MsVersionOutputDto {
	return &MsVersionOutputDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: util.FromNullString(m.Description),
		Url:         m.Url,
		OpenAPIUrl:  m.OpenAPIUrl,
		Links: []contract.Link{
			{
				Rel:  "self",
				Href: fmt.Sprintf("/api/Microservices/%s/Versions/%s", m.MicroserviceID, m.ID),
				Type: "get",
			},
			{
				Rel:  "microservice",
				Href: fmt.Sprintf("/api/Microservices/%s", m.MicroserviceID),
				Type: "get",
			},
		},
	}
}

func NewMsVersionOutputListDtoFromModels(models []model.MicroserviceVersion) []MsVersionOutputDto {
	var ld []MsVersionOutputDto
	for _, m := range models {
		d := NewMsVersionOutputDtoFromModel(&m)
		ld = append(ld, *d)
	}
	return ld
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
