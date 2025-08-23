package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
)

func StructToTpl(st *analyzer.Struct) *structTpl {
	stTpl := &structTpl{
		StructName: st.StructName,
	}

	for i, field := range st.Fields {
		fldTpl := fieldTpl{
			FieldName:   field.FieldName,
			Type:        field.Type,
			Validations: st.FieldsValidations[i].Validations,
		}

		stTpl.Fields = append(stTpl.Fields, fldTpl)
	}

	return stTpl
}

func AssertParserValidation(t *testing.T, validation string) *analyzer.Validation {
	t.Helper()

	val, err := analyzer.ParserValidation(validation)
	if err != nil {
		t.Fatalf("failed to parse validation %q: %v", validation, err)
	}

	return val
}
