package dto

type RespDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // TODO: generic field...
}

func NewRespDto(message string, data any) *RespDto {
	return &RespDto{Success: true, Message: message, Data: data}
}

type ErrRespDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"` // TODO: generic field...
}

func NewErrRespDto(message string, errors any) *ErrRespDto {
	return &ErrRespDto{Success: false, Message: message, Errors: errors}
}
