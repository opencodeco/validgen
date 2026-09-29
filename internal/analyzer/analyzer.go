package analyzer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer/operations"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/types"
)

func AnalyzeStructs(structs []*parser.Struct) ([]*Struct, error) {
	result, err := analyzeFieldValidations(structs)
	if err != nil {
		return nil, err
	}

	if err := checkForInvalidOperations(result); err != nil {
		return nil, err
	}

	if err := analyzeFieldOperations(result); err != nil {
		return nil, err
	}

	return result, nil
}

func analyzeFieldValidations(structs []*parser.Struct) ([]*Struct, error) {

	result := []*Struct{}

	for _, st := range structs {
		analyzedStruct := &Struct{
			Struct: *st,
		}
		for i, fd := range st.Fields {
			fieldValidations, hasValidTag := parseFieldValidations(fd.Tag)
			if hasValidTag {
				analyzedStruct.HasValidTag = true
			}

			analyzedStruct.FieldsValidations = append(analyzedStruct.FieldsValidations, FieldValidations{})

			for _, validation := range fieldValidations {
				val, err := ParserValidation(validation)
				if err != nil {
					return nil, types.NewValidationError("%s", fmt.Errorf("parser validation %s: %w", validation, err))
				}

				analyzedStruct.FieldsValidations[i].Validations = append(analyzedStruct.FieldsValidations[i].Validations, val)
			}
		}

		result = append(result, analyzedStruct)
	}

	return result, nil
}

func parseFieldValidations(fieldTag string) ([]string, bool) {
	if fieldTag == "" {
		return nil, false
	}

	tag := reflect.StructTag(fieldTag)
	value, ok := tag.Lookup("valid")
	if !ok {
		value, ok = tag.Lookup("validate")
	}
	if !ok {
		return nil, false
	}

	return strings.Split(value, ","), true
}

func checkForInvalidOperations(structs []*Struct) error {

	structsWithValidation := map[string]bool{}

	for _, st := range structs {
		structsWithValidation[common.KeyPath(st.PackageName, st.StructName)] = true
	}

	ops := operations.New()

	for _, st := range structs {
		for i, fd := range st.Fields {
			err := checkValidations(ops, fd.FieldName, fd.Type, st.FieldsValidations[i].Validations, structsWithValidation, false, 0)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkValidations(ops *operations.Operations, fieldName string, current common.FieldType, validations []*Validation, structs map[string]bool, dived bool, keysDepth int) error {
	for i, val := range validations {
		op := val.Operation
		if !ops.IsValid(op) {
			return types.NewValidationError("unsupported operation %s", op)
		}

		switch op {
		case "keys":
			return types.NewValidationError("operation keys: field %s: keys must immediately follow dive", fieldName)
		case "endkeys":
			return types.NewValidationError("operation endkeys: field %s: endkeys without keys", fieldName)
		case "dive":
			if i+1 < len(validations) && validations[i+1].Operation == "keys" {
				if keysDepth > 0 {
					return types.NewValidationError("operation keys: field %s: nested keys are not supported", fieldName)
				}
				return checkMapKeys(ops, fieldName, current, validations[i+1:], structs)
			}

			next, err := current.DiveInto()
			if err != nil {
				return types.NewValidationError("operation dive: field %s: %s", fieldName, err.Error())
			}
			current = next
			dived = true
			continue
		}

		if err := validateOperation(ops, op, current, structs, dived); err != nil {
			return err
		}
	}

	return nil
}

func checkMapKeys(ops *operations.Operations, fieldName string, current common.FieldType, validations []*Validation, structs map[string]bool) error {
	keyType, err := current.MapKey()
	if err != nil {
		return types.NewValidationError("operation keys: field %s: %s", fieldName, err.Error())
	}

	valueType, err := current.DiveInto()
	if err != nil {
		return types.NewValidationError("operation dive: field %s: %s", fieldName, err.Error())
	}

	keyVals, valueVals, err := SplitKeysBlock(fieldName, validations)
	if err != nil {
		return err
	}

	if err := checkValidations(ops, fieldName, keyType, keyVals, structs, true, 1); err != nil {
		return err
	}

	return checkValidations(ops, fieldName, valueType, valueVals, structs, true, 1)
}

// SplitKeysBlock splits the validations that follow dive, starting at keys.
// Tags between keys and endkeys apply to the map key. Tags after endkeys apply to the map value.
func SplitKeysBlock(fieldName string, rest []*Validation) ([]*Validation, []*Validation, error) {
	if len(rest) == 0 || rest[0].Operation != "keys" {
		return nil, nil, types.NewValidationError("operation keys: field %s: keys must immediately follow dive", fieldName)
	}

	end := -1
	for i := 1; i < len(rest); i++ {
		switch rest[i].Operation {
		case "endkeys":
			end = i
		case "dive":
			return nil, nil, types.NewValidationError("operation dive: field %s: nested dive inside keys is not supported", fieldName)
		case "keys":
			return nil, nil, types.NewValidationError("operation keys: field %s: nested keys are not supported", fieldName)
		}
		if end != -1 {
			break
		}
	}
	if end == -1 {
		return nil, nil, types.NewValidationError("operation keys: field %s: missing endkeys", fieldName)
	}

	for _, val := range rest[end+1:] {
		switch val.Operation {
		case "keys":
			return nil, nil, types.NewValidationError("operation keys: field %s: nested keys are not supported", fieldName)
		case "endkeys":
			return nil, nil, types.NewValidationError("operation endkeys: field %s: extra endkeys", fieldName)
		}
	}

	return rest[1:end], rest[end+1:], nil
}

func validateOperation(ops *operations.Operations, op string, ft common.FieldType, structs map[string]bool, dived bool) error {
	if dived && ops.IsFieldOperation(op) {
		return types.NewValidationError("operation %s: field comparisons are not supported after dive", op)
	}
	if ft.ElemPointer && !common.IsLenOperation(op) {
		return types.NewValidationError("operation %s: cannot apply to a slice or array of pointers", op)
	}

	if ft.IsNestedStruct() {
		if !structs[ft.BaseType] {
			return types.NewValidationError("unsupported operation %s with unknown go type %s", op, ft.BaseType)
		}
		if dived && op != "required" {
			return types.NewValidationError("operation %s: cannot apply to struct %s", op, ft.BaseType)
		}
		return nil
	}

	accept := func(candidate common.FieldType) bool {
		return ops.IsValidByType(op, candidate.ToNormalizedString())
	}
	if _, ok := ft.OperationType(op, accept); ok {
		return nil
	}

	lookup := ft.ForCatalog()
	if !lookup.IsGoType() {
		return types.NewValidationError("unsupported operation %s with unknown go type %s", op, ft.BaseType)
	}

	return types.NewValidationError("operation %s: invalid %s(%s) type", op, lookup.BaseType, lookup.ToNormalizedString())
}

func analyzeFieldOperations(structs []*Struct) error {

	// Map all fields and their types.
	fieldsType := map[string]common.FieldType{}
	for _, st := range structs {
		for _, fd := range st.Fields {
			fieldsType[common.KeyPath(st.PackageName, st.StructName, fd.FieldName)] = fd.Type
		}
	}

	ops := operations.New()

	for _, st := range structs {
		for i, fd := range st.Fields {
			for _, val := range st.FieldsValidations[i].Validations {
				// Check if is a field operation.
				op := val.Operation
				if !ops.IsFieldOperation(op) {
					continue
				}

				fd1Name := fd.FieldName
				fd2Name := val.Values[0]

				// Check if field exists.
				fd2NameToSearch := ""
				qualifiedField, qualifiedNestedField, ok := strings.Cut(fd2Name, ".")
				if !ok {
					// If the operation is with an inner field, assume it's in the same struct.
					fd2NameToSearch = common.KeyPath(st.PackageName, st.StructName, fd2Name)
				} else {
					// If the field is qualified with a nested field, use its type.
					qFieldType, ok := fieldsType[common.KeyPath(st.PackageName, st.StructName, qualifiedField)]
					if !ok {
						return types.NewValidationError("operation %s: undefined nested field %s", op, qualifiedField)
					}
					fd2NameToSearch = common.KeyPath(qFieldType.BaseType, qualifiedNestedField)
				}

				f2Type, ok := fieldsType[fd2NameToSearch]
				if !ok {
					return types.NewValidationError("operation %s: undefined field %s", op, fd2Name)
				}

				// Check if fields have the same type.
				if fd.Type != f2Type {
					return types.NewValidationError("operation %s: mismatched types between %s and %s", op, fd1Name, fd2Name)
				}
			}
		}
	}

	return nil
}
