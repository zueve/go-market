package rest

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	StatusCode int                 `json:"-"`
	Code       string              `json:"code"`
	Message    string              `json:"message"`
	Details    []map[string]string `json:"details"`
}

func (s *HTTPError) GetStatusCode() int {
	return s.StatusCode
}

func (s *HTTPError) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

func NewAuthErr(err error) HTTPError {
	return HTTPError{
		StatusCode: http.StatusUnauthorized,
		Code:       "InvalidLogin",
		Message:    err.Error(),
		Details:    nil,
	}
}

func NewLoginExistsErr(err error) HTTPError {
	return HTTPError{
		StatusCode: http.StatusConflict,
		Code:       "LoginExists",
		Message:    err.Error(),
		Details:    nil,
	}
}

func NewValidationError(message string, details []map[string]string) HTTPError {
	return HTTPError{
		StatusCode: http.StatusBadRequest,
		Code:       "ValidationError",
		Message:    message,
		Details:    details,
	}
}

func NewInvoiceError(message string) HTTPError {
	return HTTPError{
		StatusCode: http.StatusUnprocessableEntity,
		Code:       "ValidationError",
		Message:    message,
		Details:    nil,
	}
}


func NewBadRequest(message string, code string, details []map[string]string) HTTPError {
	return HTTPError{
		StatusCode: http.StatusBadRequest,
		Code:       code,
		Message:    message,
		Details:    details,
	}
}

func NewInternalError() HTTPError {
	return HTTPError{
		StatusCode: http.StatusInternalServerError,
		Code:       "InternalServerError",
		Message:    "Internal Server Error",
		Details:    make([]map[string]string, 0),
	}
}

func NewUnsupportedMediaType() HTTPError {
	return HTTPError{
		StatusCode: http.StatusUnsupportedMediaType,
		Code:       "UnsupportedMediaType",
		Message:    "Unsupported Media Type",
		Details:    nil,
	}
}

func NewNotFound(message string) HTTPError {
	return HTTPError{
		StatusCode: http.StatusNotFound,
		Code:       "Not Found",
		Message:    message,
		Details:    nil,
	}
}

func NewOutOfMoney(err error) HTTPError {
	return HTTPError{
		StatusCode: http.StatusPaymentRequired,
		Code:       "OutOfMoney",
		Message:    err.Error(),
		Details:    nil,
	}
}
