package form

import (
	"errors"
)

var (
	ErrNotPointerToStruct = errors.New("the value is not a pointer to struct")
	ErrInvalidFormat = errors.New("invalid form format")
	ErrFieldNotFound = errors.New("field not found")
)
