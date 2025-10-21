package common

import (
	"strings"

	"github.com/opencodeco/validgen/types"
)

func HelperFromNormalizedToBasicTypes(t string) []string {
	fieldTypes, err := HelperFromNormalizedToFieldTypes(t)
	if err != nil {
		return []string{"invalid"}
	}

	result := []string{}
	for _, ft := range fieldTypes {
		result = append(result, ft.ToType())
	}

	return result
}

func HelperFromNormalizedToStringNames(t string) []string {
	fieldTypes, err := HelperFromNormalizedToFieldTypes(t)
	if err != nil {
		return []string{"invalid"}
	}

	result := []string{}
	for _, ft := range fieldTypes {
		result = append(result, ft.ToStringName())
	}

	return result
}

func HelperFromNormalizedToFieldTypes(t string) ([]FieldType, error) {
	fieldTypes := []FieldType{}
	t, isPointer := strings.CutPrefix(t, "*")
	switch t {
	case "<STRING>":
		fieldTypes = []FieldType{{BaseType: "string"}}
	case "<BOOL>":
		fieldTypes = []FieldType{{BaseType: "bool"}}
	case "<INT>":
		fieldTypes = []FieldType{
			{BaseType: "int"},
			{BaseType: "int8"},
			{BaseType: "int16"},
			{BaseType: "int32"},
			{BaseType: "int64"},
			{BaseType: "uint"},
			{BaseType: "uint8"},
			{BaseType: "uint16"},
			{BaseType: "uint32"},
			{BaseType: "uint64"},
		}
	case "<FLOAT>":
		fieldTypes = []FieldType{
			{BaseType: "float32"},
			{BaseType: "float64"},
		}
	case "map[<STRING>]":
		fieldTypes = []FieldType{{BaseType: "string", ComposedType: "map"}}
	case "map[<BOOL>]":
		fieldTypes = []FieldType{{BaseType: "bool", ComposedType: "map"}}
	case "map[<INT>]":
		fieldTypes = []FieldType{
			{BaseType: "int", ComposedType: "map"},
			{BaseType: "int8", ComposedType: "map"},
			{BaseType: "int16", ComposedType: "map"},
			{BaseType: "int32", ComposedType: "map"},
			{BaseType: "int64", ComposedType: "map"},
			{BaseType: "uint", ComposedType: "map"},
			{BaseType: "uint8", ComposedType: "map"},
			{BaseType: "uint16", ComposedType: "map"},
			{BaseType: "uint32", ComposedType: "map"},
			{BaseType: "uint64", ComposedType: "map"},
		}
	case "map[<FLOAT>]":
		fieldTypes = []FieldType{
			{BaseType: "float32", ComposedType: "map"},
			{BaseType: "float64", ComposedType: "map"},
		}
	case "[]<STRING>":
		fieldTypes = []FieldType{{BaseType: "string", ComposedType: "[]"}}
	case "[]<BOOL>":
		fieldTypes = []FieldType{{BaseType: "bool", ComposedType: "[]"}}
	case "[]<INT>":
		fieldTypes = []FieldType{
			{BaseType: "int", ComposedType: "[]"},
			{BaseType: "int8", ComposedType: "[]"},
			{BaseType: "int16", ComposedType: "[]"},
			{BaseType: "int32", ComposedType: "[]"},
			{BaseType: "int64", ComposedType: "[]"},
			{BaseType: "uint", ComposedType: "[]"},
			{BaseType: "uint8", ComposedType: "[]"},
			{BaseType: "uint16", ComposedType: "[]"},
			{BaseType: "uint32", ComposedType: "[]"},
			{BaseType: "uint64", ComposedType: "[]"},
		}
	case "[]<FLOAT>":
		fieldTypes = []FieldType{
			{BaseType: "float32", ComposedType: "[]"},
			{BaseType: "float64", ComposedType: "[]"},
		}
	case "[N]<STRING>":
		fieldTypes = []FieldType{{BaseType: "string", ComposedType: "[N]", Size: "3"}}
	case "[N]<BOOL>":
		fieldTypes = []FieldType{{BaseType: "bool", ComposedType: "[N]", Size: "3"}}
	case "[N]<INT>":
		fieldTypes = []FieldType{
			{BaseType: "int", ComposedType: "[N]", Size: "3"},
			{BaseType: "int8", ComposedType: "[N]", Size: "3"},
			{BaseType: "int16", ComposedType: "[N]", Size: "3"},
			{BaseType: "int32", ComposedType: "[N]", Size: "3"},
			{BaseType: "int64", ComposedType: "[N]", Size: "3"},
			{BaseType: "uint", ComposedType: "[N]", Size: "3"},
			{BaseType: "uint8", ComposedType: "[N]", Size: "3"},
			{BaseType: "uint16", ComposedType: "[N]", Size: "3"},
			{BaseType: "uint32", ComposedType: "[N]", Size: "3"},
			{BaseType: "uint64", ComposedType: "[N]", Size: "3"},
		}
	case "[N]<FLOAT>":
		fieldTypes = []FieldType{
			{BaseType: "float32", ComposedType: "[N]", Size: "3"},
			{BaseType: "float64", ComposedType: "[N]", Size: "3"},
		}
	}

	if len(fieldTypes) > 0 {
		if isPointer {
			for i := range fieldTypes {
				fieldTypes[i].ComposedType = "*" + fieldTypes[i].ComposedType
			}
		}
		return fieldTypes, nil
	}

	return nil, types.NewValidationError("unknown normalized type %s", t)
}
