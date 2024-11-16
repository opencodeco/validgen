package main

import (
	"fmt"
	"strings"
)

var testFieldTpl = `	if %[1]s {
		errs = append(errs, fmt.Errorf("%[2]s: %[3]s", ErrValidation))
	}
`

type FieldInfo struct {
	Name string
	Type string
	Tag  string
}

// TODO: NewFieldValidation to validate params and build the object.

func (f *FieldInfo) GenerateTestField() (string, error) {
	tag := f.Tag
	tag, _ = strings.CutPrefix(tag, "validate:")
	testCondition := ""
	errorMessage := ""

	// TODO: refactor
	if tag == `"required"` {
		if f.Type == "string" {
			testCondition = "u." + f.Name + " == \"\""
			errorMessage = f.Name + " required"
		} else if f.Type == "uint8" {
			testCondition = "u." + f.Name + " == 0"
			errorMessage = f.Name + " required"
		} else {
			return "", fmt.Errorf("unsupported type %s", f.Type)
		}
	}

	return fmt.Sprintf(testFieldTpl, testCondition, "%w", errorMessage), nil
}
