package types

import (
	"fmt"
)

type ValidationError struct {
	Msg string
}

func NewValidationError(format string, a ...any) error {
	return ValidationError{
		Msg: fmt.Sprintf(format, a...),
	}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}
