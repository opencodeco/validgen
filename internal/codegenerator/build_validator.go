package codegenerator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/opencodeco/validgen/internal/common"
)

var funcValidatorTpl = `
func {{.StructName}}Validate(obj *{{.StructName}}) []error {
	var errs []error
{{range .Fields}}{{buildValidationCode .FieldName .Type .Validations}}{{end}}
	return errs
}
`

type structTpl struct {
	StructName string
	Fields     []fieldTpl
}

type fieldTpl struct {
	FieldName   string
	Type        string
	Validations []string
}

func (gv *genValidations) BuildFuncValidatorCode() (string, error) {

	stTpl := StructToTpl(gv.Struct)

	funcMap := template.FuncMap{
		"buildValidationCode": gv.buildValidationCode,
	}

	tmpl, err := template.New("FuncValidator").Funcs(funcMap).Parse(funcValidatorTpl)
	if err != nil {
		return "", err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, stTpl); err != nil {
		return "", err
	}

	return code.String(), nil
}

func (gv *genValidations) buildValidationCode(fieldName, fieldType string, fieldValidations []string) (string, error) {

	tests := ""
	for _, fieldValidation := range fieldValidations {
		var testCode = ""
		var err error

		if common.IsGoType(fieldType) {
			testCode, err = gv.buildIfCode(fieldName, fieldType, fieldValidation)
			if err != nil {
				return "", err
			}
		} else {
			testCode, err = gv.buildIfNestedCode(fieldName, fieldType)
			if err != nil {
				return "", err
			}
		}

		tests += testCode
	}

	return tests, nil
}

func (gv *genValidations) buildIfCode(fieldName, fieldType, fieldValidation string) (string, error) {
	testElements, err := DefineTestElements(fieldName, fieldType, fieldValidation)
	if err != nil {
		return "", fmt.Errorf("field %s: %w", fieldName, err)
	}

	booleanCondition := ""
	for _, roperand := range testElements.rightOperands {
		if booleanCondition != "" {
			booleanCondition += " " + testElements.concatOperator + " "
		}

		booleanCondition += strings.TrimSpace(fmt.Sprintf("%s %s %s", testElements.leftOperand, testElements.operator, roperand))
	}

	return fmt.Sprintf(
		`
	if !(%s) {
		errs = append(errs, types.NewValidationError("%s"))
	}
`, booleanCondition, testElements.errorMessage), nil
}

func (gv *genValidations) buildIfNestedCode(fieldName, fieldType string) (string, error) {
	_, ok := gv.StructsWithValidation[fieldType]
	if !ok {
		return "", fmt.Errorf("no validator found for struct type %s", fieldType)
	}

	pkg := common.ExtractPackage(fieldType)
	if pkg == gv.Struct.PackageName {
		fieldType = strings.TrimPrefix(fieldType, pkg+".")
	}

	funcName := fieldType + "Validate"
	fieldParam := "&obj." + fieldName

	return fmt.Sprintf(
		`
	errs = append(errs, %s(%s)...)
`, funcName, fieldParam), nil
}
