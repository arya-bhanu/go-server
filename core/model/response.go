package model

import "net/http"

const (
	Success      = iota
	SystemFailed = iota
)

type ResponseType = int
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func GenerateResponse(resType ResponseType, message string, data any) *Response {
	switch resType {
	case Success:
		return &Response{
			Status:  http.StatusOK,
			Message: message,
			Data:    data,
		}
	case SystemFailed:
		return &Response{
			Status:  http.StatusInternalServerError,
			Message: message,
			Data:    data,
		}
	default:
		return nil
	}
}
