package validgen

import (
	"fmt"
)

func BuildValidationCode(fieldName, fieldType string, fieldValidations []string) (string, error) {

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

	if len(testElements.rightOperands) > 1 && testElements.concatOperator == "" {
		return "", fmt.Errorf("missed concat operator")
	}

	booleanCondition := ""
	for _, roperand := range testElements.rightOperands {
		if booleanCondition != "" {
			booleanCondition += " " + testElements.concatOperator + " "
		}

		booleanCondition += fmt.Sprintf("%s %s %s", testElements.leftOperand, testElements.operator, roperand)
	}

	return fmt.Sprintf(
		`
	if !(%s) {
		errs = append(errs, types.NewValidationError("%s"))
	}
`, booleanCondition, testElements.errorMessage), nil
}
