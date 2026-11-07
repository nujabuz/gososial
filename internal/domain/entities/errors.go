package entities

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrUnexpected     = errors.New("an unexpected error")
	ErrValidation     = errors.New("validation error")
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
)
