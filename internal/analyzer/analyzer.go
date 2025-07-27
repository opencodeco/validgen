package analyzer

import (
	"strconv"
	"strings"

	"github.com/opencodeco/validgen/internal/parser"
)

const validTag = "valid"

func AnalyzeStructs(structs []*parser.Struct) ([]*parser.Struct, error) {
	for _, st := range structs {
		for fdIndex, fd := range st.Fields {
			fieldValidations, hasValidTag := parseFieldValidations(fd.Tag)
			if hasValidTag {
				st.HasValidTag = true
			}

			st.Fields[fdIndex].FieldAnalyzerInfo.Validations = fieldValidations
		}
	}

	return structs, nil
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
