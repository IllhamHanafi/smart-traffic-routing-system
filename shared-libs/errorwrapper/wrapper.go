package errorwrapper

import (
	"fmt"
	"net/http"
)

var (
	ErrInternalServerError = ErrorWrapper{
		httpStatusCode: http.StatusInternalServerError,
		message:        "internal server error",
		code:           "GEN_001",
	}
	ErrBadRequest = ErrorWrapper{
		httpStatusCode: http.StatusBadRequest,
		message:        "bad request",
		code:           "GEN_002",
	}
	ErrNotFound = ErrorWrapper{
		httpStatusCode: http.StatusNotFound,
		message:        "resource not found",
		code:           "GEN_003",
	}
	ErrUnauthorized = ErrorWrapper{
		httpStatusCode: http.StatusUnauthorized,
		message:        "unauthorized",
		code:           "GEN_004",
	}
	ErrForbidden = ErrorWrapper{
		httpStatusCode: http.StatusForbidden,
		message:        "forbidden",
		code:           "GEN_005",
	}
)

type ErrorWrapper struct {
	httpStatusCode int
	message        string
	code           string
	err            error
}

func (e ErrorWrapper) Error() string {
	if e.err == nil {
		return fmt.Sprintf("%s: %s", e.code, e.message)
	}
	return e.err.Error()
}

func (e ErrorWrapper) Unwrap() error {
	return e.err
}

func (e ErrorWrapper) GetHttpStatusCode() int {
	if e.httpStatusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.httpStatusCode
}

func (e ErrorWrapper) GetMessage() string {
	if e.message == "" {
		return "internal server error"
	}
	return e.message
}

func (e ErrorWrapper) GetCode() string {
	if e.code == "" {
		return "GEN_001"
	}
	return e.code
}
