package codegenerator

import (
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
)

type TestElements struct {
	conditions     []string
	concatOperator string
	errorMessage   string
}

func DefineTestElements(fieldName string, fieldType common.FieldType, fieldValidation *analyzer.Validation) (TestElements, error) {

	condition, err := GetConditionTable(fieldValidation.Operation, fieldType)
	if err != nil {
		return TestElements{}, err
	}

	values := fieldValidation.Values
	roperands := []string{}
	targetValue := ""
	targetValues := ""

	switch fieldValidation.ExpectedValues {
	case common.ZeroValue: // REFACTOR: codegenerator should inform how many values are expected
		roperands = append(roperands, replaceNameAndTarget(condition.operation, fieldName, ""))
		targetValue = condition.operation
		targetValues = "'" + condition.operation + "' "
	case common.OneValue, common.ManyValues:
		valuesAsNumericSlice, valuesAsStringSlice := normalizeSlicesAsCode(fieldType.BaseType, values)

		for _, value := range values {
			operation := replaceNameAndTarget(condition.operation, fieldName, value)
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
	errorMsg = replaceNameAndTarget(errorMsg, fieldName, targetValue)
	errorMsg = replaceTargetInErrors(errorMsg, targetValue, targetValues)

	return TestElements{
		conditions:     roperands,
		concatOperator: condition.concatOperator,
		errorMessage:   errorMsg,
	}, nil
}

func replaceNameAndTarget(text, name, target string) string {
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
	valuesAsStringSlice := "[]" + "string" + "{"

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
