package core_errors

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrConflict            = errors.New("conflict")
	ErrInvalidArgument     = errors.New("invalid argument")
)
