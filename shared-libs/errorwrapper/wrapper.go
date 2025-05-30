package errorwrapper

import (
	"fmt"
	"net/http"
)

var (
	ErrInternalServerError = ErrorWrapper{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "internal server error",
		Code:           "GEN_001",
	}
	ErrBadRequest = ErrorWrapper{
		HttpStatusCode: http.StatusBadRequest,
		Message:        "bad request",
		Code:           "GEN_002",
	}
	ErrNotFound = ErrorWrapper{
		HttpStatusCode: http.StatusNotFound,
		Message:        "resource not found",
		Code:           "GEN_003",
	}
	ErrUnauthorized = ErrorWrapper{
		HttpStatusCode: http.StatusUnauthorized,
		Message:        "unauthorized",
		Code:           "GEN_004",
	}
	ErrForbidden = ErrorWrapper{
		HttpStatusCode: http.StatusForbidden,
		Message:        "forbidden",
		Code:           "GEN_005",
	}
	ErrValidationNotPassed = ErrorWrapper{
		HttpStatusCode: http.StatusBadRequest,
		Message:        "validation not passed",
		Code:           "GEN_006",
	}
)

type ErrorWrapper struct {
	HttpStatusCode int
	Message        string
	Code           string
	Err            error
	Detail         map[string]any
}

func (e ErrorWrapper) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}
	return e.Err.Error()
}

func (e ErrorWrapper) Unwrap() error {
	return e.Err
}

func (e ErrorWrapper) GetHttpStatusCode() int {
	if e.HttpStatusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.HttpStatusCode
}

func (e ErrorWrapper) GetMessage() string {
	if e.Message == "" {
		return "internal server error"
	}
	return e.Message
}

func (e ErrorWrapper) GetCode() string {
	if e.Code == "" {
		return "GEN_001"
	}
	return e.Code
}

func (e ErrorWrapper) GetDetail() map[string]any {
	if e.Detail == nil {
		return map[string]any{}
	}
	return e.Detail
}

func (e ErrorWrapper) WithDetail(detail map[string]any) ErrorWrapper {
	e.Detail = detail
	return e
}

func (e ErrorWrapper) WithMessage(message string) ErrorWrapper {
	e.Message = message
	return e
}
func (e ErrorWrapper) WithCode(code string) ErrorWrapper {
	e.Code = code
	return e
}

func (e ErrorWrapper) WithError(err error) ErrorWrapper {
	e.Err = err
	return e
}

func (e ErrorWrapper) IsError() bool {
	return e.Code != ""
}
