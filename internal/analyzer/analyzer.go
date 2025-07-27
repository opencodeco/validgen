package analyzer

import (
	"strconv"
	"strings"

	"github.com/opencodeco/validgen/internal/parser"
)

const validTag = "valid"

func AnalyzeStructs(structs []*parser.Struct) ([]*Struct, error) {
	result := []*Struct{}

	for _, st := range structs {
		analyzedStruct := &Struct{
			Struct: *st,
		}
		for _, fd := range st.Fields {
			fieldValidations, hasValidTag := parseFieldValidations(fd.Tag)
			if hasValidTag {
				analyzedStruct.HasValidTag = true
			}

			analyzedStruct.FieldsValidations = append(analyzedStruct.FieldsValidations, FieldValidations{fieldValidations})
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
