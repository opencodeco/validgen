package analyzer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/types"
)

const validTag = "valid"

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
	fieldValidations := []string{}
	hasValidTag := false
	prefixToSearch := validTag + ":"

	if strings.HasPrefix(fieldTag, prefixToSearch) {
		hasValidTag = true
		tagWithoutPrefix, _ := strings.CutPrefix(fieldTag, prefixToSearch)
		tagWithoutQuotes, _ := strconv.Unquote(tagWithoutPrefix)
		fieldValidations = strings.Split(tagWithoutQuotes, ",")
	}

	return fieldValidations, hasValidTag
}

func checkForInvalidOperations(structs []*Struct) error {

	for _, st := range structs {
		for i, fd := range st.Fields {
			for _, val := range st.FieldsValidations[i].Validations {
				// Check if is a valid operation.
				op := val.Operation
				if operations[op].Name == "" {
					return types.NewValidationError("unsupported operation %s", op)
				}

				// Check if is a valid operation for this type.
				fdType := fd.Type
				if fdType.IsGoType() && !operations[op].ValidTypes[fdType.ToString()] {
					return types.NewValidationError("operation %s: invalid %s type", op, fdType.BaseType)
				}
			}
		}
	}

	return nil
}

func analyzeFieldOperations(structs []*Struct) error {

	// Map all fields and their types.
	fieldsType := map[string]common.FieldType{}
	for _, st := range structs {
		for _, fd := range st.Fields {
			fieldsType[common.KeyPath(st.PackageName, st.StructName, fd.FieldName)] = fd.Type
		}
	}

	for _, st := range structs {
		for i, fd := range st.Fields {
			for _, val := range st.FieldsValidations[i].Validations {
				// Check if is a field operation.
				op := val.Operation
				if !operations[op].IsFieldOperation {
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
