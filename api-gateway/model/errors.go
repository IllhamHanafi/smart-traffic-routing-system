package model

import (
	"fmt"
	"net/http"

	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/errorwrapper"
)

var (
	ErrInvalidName     = fmt.Errorf("invalid name")
	ErrInvalidEmail    = fmt.Errorf("invalid email")
	ErrInvalidPassword = fmt.Errorf("invalid password")
	ErrInvalidRole     = fmt.Errorf("invalid role")
)

var (
	ErrInvalidInput = errorwrapper.ErrorWrapper{
		HttpStatusCode: http.StatusBadRequest,
		Message:        "invalid input",
		Code:           "API_REG_001",
		Err:            fmt.Errorf("invalid input"),
	}
	ErrCreateUser = errorwrapper.ErrorWrapper{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "failed to create user",
		Code:           "API_REG_002",
		Err:            fmt.Errorf("failed to create user"),
	}

	ErrInvalidCredentials = errorwrapper.ErrorWrapper{
		HttpStatusCode: http.StatusUnauthorized,
		Message:        "invalid credentials",
		Code:           "AUTH_001",
		Err:            fmt.Errorf("invalid email and password"),
	}
)
