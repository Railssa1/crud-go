package errors_api

import "net/http"

type ApiErrors struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewApiErrors(message, err string, code int, causes []Causes) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func (e *ApiErrors) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     "Bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     "Bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     "Internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     "Not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewUnauthorizationError(message string) *ApiErrors {
	return &ApiErrors{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
