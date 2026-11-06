package entity

import "errors"

var (
	// common errors.

	ErrValidation = errors.New("failed validation")
	ErrNotFound   = errors.New("not found")
)
