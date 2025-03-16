package dto

type DataRespDto RespDto[map[string]string]
type MessageRespDto RespDto[*string]

type RespDto[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

func NewRespDto[T any](message string, data *T) *RespDto[T] {
	return &RespDto[T]{Success: true, Message: message, Data: data}
}

type ErrRespDto struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

func NewErrRespDto(message string, errors []string) *ErrRespDto {
	return &ErrRespDto{Success: false, Message: message, Errors: errors}
}
