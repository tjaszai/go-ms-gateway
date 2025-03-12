package dto

import "github.com/tjaszai/go-ms-gateway/internal/model"

type MsReqDto struct {
	// TODO: unique name validator...
	Name        string `json:"name" validate:"required,gte=1,lte=255,regex_pattern=^[a-z-]+$"`
	Description string `json:"description,omitempty"`
}

func (d *MsReqDto) MsReqToModel(m *model.Microservice) *model.Microservice {
	if m == nil {
		m = &model.Microservice{}
	}
	m.Name = d.Name
	m.Description = d.Description
	return m
}
