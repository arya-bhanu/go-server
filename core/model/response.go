package model

type ResponseType = int
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func GenerateResponse(status uint, message string, data any) (*Response, int) {
	return &Response{
		Status:  int(status),
		Message: message,
		Data:    data,
	}, int(status)
}
