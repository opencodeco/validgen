package codegenerator

import (
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
