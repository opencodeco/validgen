package codegenerator

import (
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer"
)

func StructToTpl(st *analyzer.Struct) *structTpl {
	stTpl := &structTpl{
		StructName:  st.StructName,
		PackageName: st.PackageName,
		Imports:     st.Imports,
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

func IsGoType(fieldType string) bool {
	goTypes := map[string]struct{}{
		"string": {},
		"uint8":  {},
	}

	_, ok := goTypes[fieldType]

	return ok
}

func ExtractPackage(fieldType string) string {
	dotIdx := strings.IndexByte(fieldType, '.')
	if dotIdx == -1 {
		return ""
	}

	pkg := fieldType[:dotIdx]

	return pkg
}
