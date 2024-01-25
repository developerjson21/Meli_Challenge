package handler

import (
	"net/http"

)

type ErrorApi struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
}

func (e *ErrorApi) Error() string {
	return e.Message
}

func NewNotFoundApiError(message string) error {
	return &ErrorApi {http.StatusNotFound, message}
}

func NewBadRequestApiError(message string) error {
	return &ErrorApi {http.StatusBadRequest, message}
}

func NewInternalServerApiError(message string) error {
	return &ErrorApi {http.StatusInternalServerError, message}
}

func NewQueryStatementError(message string) error {
	return &ErrorApi {http.StatusInternalServerError, message}
}

func DataNotFoundError(message string) error {
	return &ErrorApi {http.StatusOK, message}
}

func JSONSerializeError(message string) error {
	return &ErrorApi {http.StatusInternalServerError, message}
}

func AuthFailedError(message string) error {
	return &ErrorApi {http.StatusForbidden, message}
}

func AuthUnauthorizedError(message string) error {
	return &ErrorApi {http.StatusUnauthorized, message}
}