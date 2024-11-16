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

type StructInfo struct {
	Name           string
	Path           string
	PackageName    string
	FieldsInfo     []FieldInfo
	HasValidateTag bool
}

func (s *StructInfo) GenerateFuncValidator() (string, error) {
	if !s.HasValidateTag {
		return "", nil
	}

	fieldValidations := ""
	for _, field := range s.FieldsInfo {
		validations, err := field.GenerateTestField()
		if err != nil {
			return "", err
		}
		fieldValidations += "\n" + validations
	}

	return fmt.Sprintf(funcValidatorTpl, s.Name, fieldValidations), nil
}

func (s *StructInfo) PrintInfo() {
	fmt.Println("Struct:", s.Name)
	fmt.Println("\tHasValidateTag:", s.HasValidateTag)

	for _, f := range s.FieldsInfo {
		fmt.Println("\tField:", f.Name, f.Type, f.Tag)
	}
}
