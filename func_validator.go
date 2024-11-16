package main

import (
	"fmt"
)

var funcValidatorTpl = `func %[1]sValidate(u *%[1]s) []error {
	var errs []error
%[2]s
	return errs
}
`

type FuncValidator struct {
	Name              string
	Path              string
	Package           string
	FieldsValidations []FieldValidation
	HasValidateTag    bool
}

func (s *FuncValidator) Generate() (string, error) {
	if !s.HasValidateTag {
		return "", nil
	}

	fieldValidations := ""
	for _, field := range s.FieldsValidations {
		validations, err := field.Generate()
		if err != nil {
			return "", err
		}
		fieldValidations += "\n" + validations
	}

	return fmt.Sprintf(funcValidatorTpl, s.Name, fieldValidations), nil
}
