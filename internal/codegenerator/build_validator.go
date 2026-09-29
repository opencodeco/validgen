package codegenerator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/types"
)

var funcValidatorTpl = `func {{.StructName}}Validate(obj *{{.StructName}}) []error {
var errs []error
{{range .Fields}}{{buildValidationCode .FieldName .Type .Validations}}{{end}}return errs
}
`

type structTpl struct {
	StructName string
	Fields     []fieldTpl
}

type fieldTpl struct {
	FieldName   string
	Type        common.FieldType
	Validations []*analyzer.Validation
}

func (gv *GenValidations) BuildFuncValidatorCode() (string, error) {

	stTpl := StructToTpl(gv.Struct)

	funcMap := template.FuncMap{
		"buildValidationCode": gv.BuildValidationCode,
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

func (gv *GenValidations) BuildUnmarshalJSONCode() string {
	return fmt.Sprintf(
		`func (obj *%s) UnmarshalJSON(b []byte) error {
	type alias %s
	if err := json.Unmarshal(b, (*alias)(obj)); err != nil {
		return err
	}
	if errs := %sValidate(obj); len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
`,
		gv.Struct.StructName, gv.Struct.StructName, gv.Struct.StructName,
	)
}

func (gv *GenValidations) BuildValidationCode(fieldName string, fieldType common.FieldType, fieldValidations []*analyzer.Validation) (string, error) {
	return gv.emitValidations("obj."+fieldName, fieldName, fieldType, fieldValidations, false, 0, 0)
}

func (gv *GenValidations) emitValidations(expr, fieldName string, fieldType common.FieldType, fieldValidations []*analyzer.Validation, dived bool, depth, keysDepth int) (string, error) {
	tests := ""
	for i, fieldValidation := range fieldValidations {
		switch fieldValidation.Operation {
		case "dive":
			if i+1 < len(fieldValidations) && fieldValidations[i+1].Operation == "keys" {
				if keysDepth > 0 {
					return "", types.NewValidationError("operation keys: field %s: nested keys are not supported", fieldName)
				}
				loop, err := gv.emitMapKeys(expr, fieldName, fieldType, fieldValidations[i+1:], depth)
				if err != nil {
					return "", err
				}
				return tests + loop, nil
			}

			loop, err := gv.emitDive(expr, fieldName, fieldType, fieldValidations[i+1:], depth, keysDepth)
			if err != nil {
				return "", err
			}
			return tests + loop, nil
		case "keys":
			return "", types.NewValidationError("operation keys: field %s: keys must immediately follow dive", fieldName)
		case "endkeys":
			return "", types.NewValidationError("operation endkeys: field %s: endkeys without keys", fieldName)
		}

		testCode, err := gv.emitOne(expr, fieldName, fieldType, fieldValidation, dived)
		if err != nil {
			return "", err
		}
		tests += testCode
	}

	if dived && fieldType.IsNestedStruct() {
		nested, err := gv.emitStructCall(expr, fieldType)
		if err != nil {
			return "", err
		}
		tests += nested
	}

	return tests, nil
}

func (gv *GenValidations) emitDive(expr, fieldName string, fieldType common.FieldType, rest []*analyzer.Validation, depth, keysDepth int) (string, error) {
	elemType, err := fieldType.DiveInto()
	if err != nil {
		return "", fmt.Errorf("field %s: %w", fieldName, err)
	}

	depth++
	elemExpr := fmt.Sprintf("elem%d", depth)
	body, err := gv.emitValidations(elemExpr, fieldName, elemType, rest, true, depth, keysDepth)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(body) == "" {
		return "", nil
	}

	rangeExpr := expr
	prefix := ""
	suffix := ""
	if fieldType.PointerToContainer() {
		rangeExpr = "*" + expr
		prefix = fmt.Sprintf("if %s != nil {\n", expr)
		suffix = "}\n"
	}

	return prefix + fmt.Sprintf("for _, %s := range %s {\n%s}\n", elemExpr, rangeExpr, body) + suffix, nil
}

func (gv *GenValidations) emitMapKeys(expr, fieldName string, fieldType common.FieldType, rest []*analyzer.Validation, depth int) (string, error) {
	keyType, err := fieldType.MapKey()
	if err != nil {
		return "", types.NewValidationError("operation keys: field %s: %s", fieldName, err.Error())
	}
	valueType, err := fieldType.DiveInto()
	if err != nil {
		return "", fmt.Errorf("field %s: %w", fieldName, err)
	}

	keyVals, valueVals, err := analyzer.SplitKeysBlock(fieldName, rest)
	if err != nil {
		return "", err
	}

	depth++
	keyExpr := fmt.Sprintf("key%d", depth)
	elemExpr := fmt.Sprintf("elem%d", depth)
	keyBody, err := gv.emitValidations(keyExpr, fieldName, keyType, keyVals, true, depth, 1)
	if err != nil {
		return "", err
	}
	valueBody, err := gv.emitValidations(elemExpr, fieldName, valueType, valueVals, true, depth, 1)
	if err != nil {
		return "", err
	}
	body := keyBody + valueBody
	if strings.TrimSpace(body) == "" {
		return "", nil
	}

	keyName := "_"
	if strings.TrimSpace(keyBody) != "" {
		keyName = keyExpr
	}
	elemName := "_"
	if strings.TrimSpace(valueBody) != "" {
		elemName = elemExpr
	}

	rangeExpr := expr
	prefix := ""
	suffix := ""
	if fieldType.PointerToContainer() {
		rangeExpr = "*" + expr
		prefix = fmt.Sprintf("if %s != nil {\n", expr)
		suffix = "}\n"
	}

	loop := fmt.Sprintf("for %s, %s := range %s {\n%s}\n", keyName, elemName, rangeExpr, body)
	if elemName == "_" {
		loop = fmt.Sprintf("for %s := range %s {\n%s}\n", keyName, rangeExpr, body)
	}

	return prefix + loop + suffix, nil
}

func (gv *GenValidations) emitOne(expr, fieldName string, fieldType common.FieldType, fieldValidation *analyzer.Validation, dived bool) (string, error) {
	if fieldType.IsNestedStruct() {
		if !dived {
			return gv.buildIfNestedCode(fieldName, fieldType)
		}
		if fieldValidation.Operation == "required" && fieldType.ComposedType == "*" {
			return fmt.Sprintf("if !(%s != nil) {\nerrs = append(errs, types.NewValidationError(\"%s is required\"))\n}\n", expr, fieldName), nil
		}
		if fieldValidation.Operation == "required" {
			return "", nil
		}
		return "", fmt.Errorf("operation %s: cannot apply to struct %s", fieldValidation.Operation, fieldType.BaseType)
	}

	accept := func(candidate common.FieldType) bool {
		_, err := GetConditionTable(fieldValidation.Operation, candidate)
		return err == nil
	}
	target, ok := fieldType.OperationType(fieldValidation.Operation, accept)
	if !ok {
		return "", fmt.Errorf("field %s: operation %s is not supported for %s", fieldName, fieldValidation.Operation, fieldType.ToType())
	}

	return gv.buildIfCode(expr, fieldName, target, fieldValidation)
}

func (gv *GenValidations) emitStructCall(expr string, fieldType common.FieldType) (string, error) {
	_, ok := gv.StructsWithValidation[fieldType.BaseType]
	if !ok {
		return "", fmt.Errorf("no validator found for struct type %s", fieldType.ToType())
	}

	funcName := nestedFuncName(gv.Struct.PackageName, fieldType.BaseType)
	if fieldType.ComposedType == "*" {
		return fmt.Sprintf("if %s != nil {\nerrs = append(errs, %s(%s)...)\n}\n", expr, funcName, expr), nil
	}
	if fieldType.ComposedType != "" {
		return "", fmt.Errorf("no validator found for struct type %s", fieldType.ToType())
	}

	return fmt.Sprintf("errs = append(errs, %s(&%s)...)\n", funcName, expr), nil
}

func nestedFuncName(packageName, baseType string) string {
	pkg := common.ExtractPackage(baseType)
	if pkg == packageName {
		baseType = strings.TrimPrefix(baseType, pkg+".")
	}

	return baseType + "Validate"
}

func (gv *GenValidations) buildIfCode(expr, fieldName string, fieldType common.FieldType, fieldValidation *analyzer.Validation) (string, error) {
	testElements, err := defineTestElements(expr, fieldName, fieldType, fieldValidation)
	if err != nil {
		return "", fmt.Errorf("field %s: %w", fieldName, err)
	}

	booleanCondition := ""
	for _, condition := range testElements.conditions {
		if booleanCondition != "" {
			booleanCondition += " " + testElements.concatOperator + " "
		}

		booleanCondition += condition
	}

	return fmt.Sprintf(
		`if !(%s) {
errs = append(errs, types.NewValidationError("%s"))
}
`, booleanCondition, testElements.errorMessage), nil
}

func (gv *GenValidations) buildIfNestedCode(fieldName string, fieldType common.FieldType) (string, error) {
	_, ok := gv.StructsWithValidation[fieldType.BaseType]
	if !ok {
		return "", fmt.Errorf("no validator found for struct type %s", fieldType.ToType())
	}

	pkg := common.ExtractPackage(fieldType.BaseType)
	if pkg == gv.Struct.PackageName {
		fieldType.BaseType = strings.TrimPrefix(fieldType.BaseType, pkg+".")
	}

	funcName := fieldType.BaseType + "Validate"
	fieldParam := "&obj." + fieldName

	return fmt.Sprintf("errs = append(errs, %s(%s)...)\n", funcName, fieldParam), nil
}
