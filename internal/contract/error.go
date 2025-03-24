package contract

type Error struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Details *map[string]any `json:"details,omitempty"`
}

func NewError(code int, message string, details *map[string]any) *Error {
	return &Error{Code: code, Message: message, Details: details}
}

func (e *Error) AddDetail(key string, value any) {
	if e.Details == nil {
		details := make(map[string]any)
		e.Details = &details
	}
	(*e.Details)[key] = value
}
