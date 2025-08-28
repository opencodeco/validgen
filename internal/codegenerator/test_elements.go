package codegenerator

import (
	"slices"
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/types"
)

const (
	EqIgnoreCaseOp  = "eq_ignore_case"
	NeqIgnoreCaseOp = "neq_ignore_case"
	InOp            = "in"
	NotInOp         = "nin"
)

type TestElements struct {
	leftOperand    string
	operator       string
	rightOperands  []string
	concatOperator string
	errorMessage   string
}

func DefineTestElements(fieldName, fieldType string, fieldValidation *analyzer.Validation) (TestElements, error) {

	type ConditionTable struct {
		loperand     string
		operator     string
		roperand     string
		errorMessage string
	}

	conditionTable := map[string]ConditionTable{
		"eq,string":              {"{{.Name}}", "==", `"{{.Target}}"`, "{{.Name}} must be equal to '{{.Target}}'"},
		"required,string":        {"{{.Name}}", "!=", `""`, "{{.Name}} is required"},
		"required,uint8":         {"{{.Name}}", "!=", `0`, "{{.Name}} is required"},
		"gte,uint8":              {"{{.Name}}", ">=", `{{.Target}}`, "{{.Name}} must be >= {{.Target}}"},
		"lte,uint8":              {"{{.Name}}", "<=", `{{.Target}}`, "{{.Name}} must be <= {{.Target}}"},
		"min,string":             {"len({{.Name}})", ">=", `{{.Target}}`, "{{.Name}} length must be >= {{.Target}}"},
		"max,string":             {"len({{.Name}})", "<=", `{{.Target}}`, "{{.Name}} length must be <= {{.Target}}"},
		"eq_ignore_case,string":  {"types.ToLower({{.Name}})", "==", `"{{.Target}}"`, "{{.Name}} must be equal to '{{.Target}}'"},
		"len,string":             {"len({{.Name}})", "==", `{{.Target}}`, "{{.Name}} length must be {{.Target}}"},
		"neq,string":             {"{{.Name}}", "!=", `"{{.Target}}"`, "{{.Name}} must not be equal to '{{.Target}}'"},
		"neq_ignore_case,string": {"types.ToLower({{.Name}})", "!=", `"{{.Target}}"`, "{{.Name}} must not be equal to '{{.Target}}'"},
		"in,string":              {"{{.Name}}", "==", `"{{.Target}}"`, "{{.Name}} must be one of {{.Targets}}"},
		"nin,string":             {"{{.Name}}", "!=", `"{{.Target}}"`, "{{.Name}} must not be one of {{.Targets}}"},
		"email,string":           {"types.IsValidEmail({{.Name}})", "==", `true`, "{{.Name}} must be a valid email"},
		"required,[]string":      {"len({{.Name}})", "!=", `0`, "{{.Name}} must not be empty"},
		"min,[]string":           {"len({{.Name}})", ">=", `{{.Target}}`, "{{.Name}} must have at least {{.Target}} elements"},
		"max,[]string":           {"len({{.Name}})", "<=", `{{.Target}}`, "{{.Name}} must have at most {{.Target}} elements"},
		"len,[]string":           {"len({{.Name}})", "==", `{{.Target}}`, "{{.Name}} must have exactly {{.Target}} elements"},
		"in,[]string":            {"", "", `types.SlicesContains({{.Name}}, "{{.Target}}")`, "{{.Name}} elements must be one of {{.Targets}}"},
		"nin,[]string":           {"", "", `!types.SlicesContains({{.Name}}, "{{.Target}}")`, "{{.Name}} elements must not be one of {{.Targets}}"},
		"eqfield,string":         {"{{.Name}}", "==", `obj.{{.Target}}`, "{{.Name}} must be equal to {{.Target}}"},
		"neqfield,string":        {"{{.Name}}", "!=", `obj.{{.Target}}`, "{{.Name}} must not be equal to {{.Target}}"},
	}

	condition, ok := conditionTable[fieldValidation.Operation+","+fieldType]
	if !ok {
		return TestElements{}, types.NewValidationError("unsupported operation %s type %s", fieldValidation.Operation, fieldType)
	}

	normalizedValues := slices.Clone(fieldValidation.Values)
	if fieldValidation.Operation == EqIgnoreCaseOp || fieldValidation.Operation == NeqIgnoreCaseOp {
		for i := range normalizedValues {
			normalizedValues[i] = strings.ToLower(normalizedValues[i])
		}
	}

	roperands := []string{}
	targetValue := ""
	targetValues := ""

	switch fieldValidation.ExpectedValues {
	case analyzer.ZERO_VALUE:
		roperands = append(roperands, replaceNameAndTargetWithPrefix(condition.roperand, fieldName, condition.roperand))
		targetValue = condition.roperand
		targetValues = "'" + condition.roperand + "' "
	case analyzer.ONE_VALUE, analyzer.MANY_VALUES:
		for _, value := range normalizedValues {
			roperands = append(roperands, replaceNameAndTargetWithPrefix(condition.roperand, fieldName, value))
			targetValue = value
			targetValues += "'" + value + "' "
		}
	}

	var concatOperator string
	switch fieldValidation.Operation {
	case InOp:
		concatOperator = "||"
	case NotInOp:
		concatOperator = "&&"
	default:
		concatOperator = ""
	}

	if len(roperands) > 1 && concatOperator == "" {
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
		concatOperator: concatOperator,
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
