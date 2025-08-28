package analyzer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/types"
)

const validTag = "valid"

func AnalyzeStructs(structs []*parser.Struct) ([]*Struct, error) {
	result, err := analyzeFieldValidations(structs)
	if err != nil {
		return nil, err
	}

	err = analyzeFieldOperations(result)
	if err != nil {
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

func analyzeFieldOperations(structs []*Struct) error {

	fieldOperations := map[string]struct{}{
		"eqfield":  {},
		"neqfield": {},
		"gtefield": {},
		"gtfield":  {},
		"ltefield": {},
		"ltfield":  {},
	}

	// For string type, eqfield and nefield.
	// For uint8 type, eqfield, gtefield, gtfield, ltefield, ltfield and nefield.
	fieldOperationsByType := map[string]struct{}{
		"string,eqfield":  {},
		"string,neqfield": {},
	}

	for _, st := range structs {
		fieldsType := map[string]string{}
		for _, fd := range st.Fields {
			fieldsType[fd.FieldName] = fd.Type
		}

		for i, fd := range st.Fields {
			for _, val := range st.FieldsValidations[i].Validations {
				// Check if is a field operation.
				op := val.Operation
				if _, ok := fieldOperations[op]; !ok {
					continue
				}

				// Check if is a valid operation for a type.
				fd1Type := fd.Type
				if _, ok := fieldOperationsByType[fd1Type+","+op]; !ok {
					return types.NewValidationError("invalid operation %s to %s type", op, fd1Type)
				}

				fd1Name := fd.FieldName
				fd2Name := val.Values[0]

				// Check if field exits.
				f2Type, ok := fieldsType[fd2Name]
				if !ok {
					return types.NewValidationError("operation %s: undefined field %s", op, fd2Name)
				}

				// Check if fields has the same type.
				if fd.Type != f2Type {
					return types.NewValidationError("operation %s: mismatched types between %s and %s", op, fd1Name, fd2Name)
				}
			}
		}
	}

	return nil
}
