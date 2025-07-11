package validgen

import (
	"fmt"
	"strings"
)

const (
	EqIgnoreCaseTag  = "eq_ignore_case"
	NeqIgnoreCaseTag = "neq_ignore_case"
)

type TestElements struct {
	loperand     string
	operator     string
	roperand     string
	errorMessage string
}

func GetTestElements(fieldName, fieldValidation, fieldType string) (TestElements, error) {
	ifCode := map[string]TestElements{
		"eq,string":              {"{{.Name}}", "==", `"{{.Target}}"`, "{{.Name}} must be equal to '{{.Target}}'"},
		"required,string":        {"{{.Name}}", "!=", `""`, "{{.Name}} required"},
		"required,uint8":         {"{{.Name}}", "!=", `0`, "{{.Name}} required"},
		"gte,uint8":              {"{{.Name}}", ">=", `{{.Target}}`, "{{.Name}} must be >= {{.Target}}"},
		"lte,uint8":              {"{{.Name}}", "<=", `{{.Target}}`, "{{.Name}} must be <= {{.Target}}"},
		"min,string":             {"len({{.Name}})", ">=", `{{.Target}}`, "{{.Name}} length must be >= {{.Target}}"},
		"max,string":             {"len({{.Name}})", "<=", `{{.Target}}`, "{{.Name}} length must be <= {{.Target}}"},
		"eq_ignore_case,string":  {"types.ToLower({{.Name}})", "==", `"{{.Target}}"`, "{{.Name}} must be equal to '{{.Target}}'"},
		"len,string":             {"len({{.Name}})", "==", `{{.Target}}`, "{{.Name}} length must be {{.Target}}"},
		"neq,string":             {"{{.Name}}", "!=", `"{{.Target}}"`, "{{.Name}} must be not equal to '{{.Target}}'"},
		"neq_ignore_case,string": {"types.ToLower({{.Name}})", "!=", `"{{.Target}}"`, "{{.Name}} must be not equal to '{{.Target}}'"},
	}

	splitField := strings.Split(fieldValidation, "=")
	if len(splitField) > 2 {
		return TestElements{}, fmt.Errorf("malformed validation %s type %s", fieldValidation, fieldType)
	}

	validation := splitField[0]
	target := ""
	if len(splitField) > 1 {
		target = splitField[1]
	}

	ifData, ok := ifCode[validation+","+fieldType]
	if !ok {
		return TestElements{}, fmt.Errorf("unsupported validation %s type %s", fieldValidation, fieldType)
	}

	if validation == EqIgnoreCaseTag || validation == NeqIgnoreCaseTag {
		target = strings.ToLower(target)
	}

	ifData.loperand = replaceNameAndTargetWithPrefix(ifData.loperand, fieldName, target)
	ifData.roperand = replaceNameAndTargetWithPrefix(ifData.roperand, fieldName, target)
	ifData.errorMessage = replaceNameAndTargetWithoutPrefix(ifData.errorMessage, fieldName, target)

	return ifData, nil
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
