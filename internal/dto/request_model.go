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

func (d *MsReqDto) MsReqToModel(m *model.Microservice) *model.Microservice {
	if m == nil {
		m = &model.Microservice{}
	}
	m.Name = d.Name
	m.Description = util.ToNullString(d.Description)
	return m
}
