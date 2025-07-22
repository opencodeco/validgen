package benchtests

import (
	"errors"
	"fmt"
)

var ErrValidation = errors.New("validation error")

func ManualCodingValidate(obj *StructManualCoding) []error {
	var errs []error

	if obj.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName is required", ErrValidation))
	}

	if obj.LastName == "" {
		errs = append(errs, fmt.Errorf("%w: LastName is required", ErrValidation))
	}

	if obj.Age < 18 {
		errs = append(errs, fmt.Errorf("%w: Age must be >= 18", ErrValidation))
	}

	if obj.Age > 130 {
		errs = append(errs, fmt.Errorf("%w: Age must be <= 130", ErrValidation))
	}

	if len(obj.UserName) < 5 {
		errs = append(errs, fmt.Errorf("%w: UserName length must be >= 5", ErrValidation))
	}

	if len(obj.UserName) > 10 {
		errs = append(errs, fmt.Errorf("%w: UserName length must be <= 10", ErrValidation))
	}

	return errs
}
