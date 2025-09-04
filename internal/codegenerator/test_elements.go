package codegenerator

import (
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/types"
)

type TestElements struct {
	leftOperand    string
	operator       string
	rightOperands  []string
	concatOperator string
	errorMessage   string
}

func DefineTestElements(fieldName, fieldType string, fieldValidation *analyzer.Validation) (TestElements, error) {

	op, ok := operationTable[fieldValidation.Operation]
	if !ok {
		return TestElements{}, types.NewValidationError("INTERNAL ERROR: unsupported operation %s", fieldValidation.Operation)
	}

	condition, ok := op.ConditionByType[fieldType]
	if !ok {
		return TestElements{}, types.NewValidationError("INTERNAL ERROR: unsupported operation %s type %s", fieldValidation.Operation, fieldType)
	}

	values := fieldValidation.Values
	roperands := []string{}
	targetValue := ""
	targetValues := ""

	switch fieldValidation.ExpectedValues {
	case analyzer.ZERO_VALUE:
		roperands = append(roperands, replaceNameAndTargetWithPrefix(condition.roperand, fieldName, condition.roperand))
		targetValue = condition.roperand
		targetValues = "'" + condition.roperand + "' "
	case analyzer.ONE_VALUE, analyzer.MANY_VALUES:
		for _, value := range values {
			roperands = append(roperands, replaceNameAndTargetWithPrefix(condition.roperand, fieldName, value))
			targetValue = value
			targetValues += "'" + value + "' "
		}
	}

	if len(roperands) > 1 && condition.concatOperator == "" {
		return TestElements{}, types.NewValidationError("missed concat operator")
	}

	targetValues = strings.TrimSpace(targetValues)
	errorMsg := condition.errorMessage
	errorMsg = replaceNameAndTargetWithoutPrefix(errorMsg, fieldName, targetValue)
	errorMsg = replaceTargetInErrors(errorMsg, targetValue, targetValues)

	return TestElements{
		leftOperand:    replaceNameAndTargetWithPrefix(condition.loperand, fieldName, targetValue),
		operator:       condition.operator,
		rightOperands:  roperands,
		concatOperator: condition.concatOperator,
		errorMessage:   errorMsg,
	}, nil
}

func replaceNameAndTargetWithPrefix(text, name, target string) string {
	text = strings.ReplaceAll(text, "{{.Name}}", "obj."+name)
	text = strings.ReplaceAll(text, "{{.Target}}", target)

	return text
}

func replaceNameAndTargetWithoutPrefix(text, name, target string) string {
	text = strings.ReplaceAll(text, "{{.Name}}", name)
	text = strings.ReplaceAll(text, "{{.Target}}", target)

	return text
}

func replaceTargetInErrors(text, target, targets string) string {
	text = strings.ReplaceAll(text, "{{.Target}}", target)
	text = strings.ReplaceAll(text, "{{.Targets}}", targets)

	return text
}
