package model

import "fmt"

var (
	ErrInvalidName     = fmt.Errorf("invalid name")
	ErrInvalidEmail    = fmt.Errorf("invalid email")
	ErrInvalidPassword = fmt.Errorf("invalid password")
	ErrInvalidRole     = fmt.Errorf("invalid role")
)
