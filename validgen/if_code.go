package validgen

import (
	"fmt"
)

func Condition(fieldName, fieldType string, fieldValidations []string) (string, error) {

	tests := ""
	for _, fieldValidation := range fieldValidations {
		testCode, err := IfCode(fieldName, fieldValidation, fieldType)
		if err != nil {
			return "", err
		}

		tests += testCode
	}

	return tests, nil
}

func IfCode(fieldName, fieldValidation, fieldType string) (string, error) {
	testElements, err := GetTestElements(fieldName, fieldValidation, fieldType)
	if err != nil {
		return "", fmt.Errorf("field %s: %w", fieldName, err)
	}

	return fmt.Sprintf(
		`
	if !(%s %s %s) {
		errs = append(errs, types.NewValidationError("%s"))
	}
`, testElements.leftOperand, testElements.operator, testElements.rightOperand, testElements.errorMessage), nil
}
