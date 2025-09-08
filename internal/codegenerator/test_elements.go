package codegenerator

import (
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/types"
)

type TestElements struct {
	conditions     []string
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
		roperands = append(roperands, replaceNameAndTargetWithPrefix(condition.operation, fieldName, condition.operation))
		targetValue = condition.operation
		targetValues = "'" + condition.operation + "' "
	case analyzer.ONE_VALUE, analyzer.MANY_VALUES:
		basicType := strings.TrimPrefix(fieldType, "[N]")
		basicType = strings.TrimPrefix(basicType, "[]")
		valuesAsNumericSlice, valuesAsStringSlice := normalizeSlicesAsCode(basicType, values)

		for _, value := range values {
			operation := replaceNameAndTargetWithPrefix(condition.operation, fieldName, value)
			operation = replaceSlicesTargets(operation, valuesAsStringSlice, valuesAsNumericSlice)
			roperands = append(roperands, operation)
			targetValue = value
			targetValues += "'" + value + "' "
		}
	}

	if len(roperands) > 1 && condition.concatOperator == "" {
		// REFACTOR!
		roperands = roperands[:1]
	}

	targetValues = strings.TrimSpace(targetValues)
	errorMsg := condition.errorMessage
	errorMsg = replaceNameAndTargetWithoutPrefix(errorMsg, fieldName, targetValue)
	errorMsg = replaceTargetInErrors(errorMsg, targetValue, targetValues)

	return TestElements{
		conditions:     roperands,
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

func replaceSlicesTargets(text, stringTargets, numericTargets string) string {
	text = strings.ReplaceAll(text, "{{.TargetsAsStringSlice}}", stringTargets)
	text = strings.ReplaceAll(text, "{{.TargetsAsNumericSlice}}", numericTargets)

	return text
}

func replaceTargetInErrors(text, target, targets string) string {
	text = strings.ReplaceAll(text, "{{.Target}}", target)
	text = strings.ReplaceAll(text, "{{.Targets}}", targets)

	return text
}

func normalizeSlicesAsCode(basicType string, values []string) (string, string) {

	valuesAsNumericSlice := "[]" + basicType + "{"
	valuesAsStringSlice := "[]" + basicType + "{"

	for i, value := range values {
		if i != 0 {
			valuesAsNumericSlice += ", "
			valuesAsStringSlice += ", "
		}

		valuesAsNumericSlice += value
		valuesAsStringSlice += "\"" + value + "\""
	}

	valuesAsNumericSlice += "}"
	valuesAsStringSlice += "}"

	return valuesAsNumericSlice, valuesAsStringSlice
}
