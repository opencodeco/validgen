package types

import (
	"fmt"
)

type ValidationError struct {
	Msg string
}

func NewValidationError(msg string) ValidationError {
	return ValidationError{
		Msg: msg,
	}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}
