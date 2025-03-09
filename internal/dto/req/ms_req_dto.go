package req

import "github.com/tjaszai/go-ms-gateway/internal/entity"

type MSReqDto struct {
	Name        string  `json:"name" validate:"required,gte=1,lte=255,regex_pattern=^[a-z-]+$"`
	Description *string `json:"description"`
}

func (reqDto MSReqDto) ToEntity() *entity.Microservice {
	return &entity.Microservice{
		Name:        reqDto.Name,
		Description: reqDto.Description,
	}
}
