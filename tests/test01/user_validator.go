package main

import (
	"fmt"
)

func UserValidate(u *User) []error {
	var errs []error

	if u.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}

	if u.LastName == "" {
		errs = append(errs, fmt.Errorf("%w: LastName required", ErrValidation))
	}

	if u.Age == 0 {
		errs = append(errs, fmt.Errorf("%w: Age required", ErrValidation))
	}

	return errs
}
