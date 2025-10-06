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

func HelperFromNormalizedToFieldTypes(t string) ([]FieldType, error) {
	switch t {
	case "<STRING>":
		return []FieldType{{BaseType: "string"}}, nil
	case "<BOOL>":
		return []FieldType{{BaseType: "bool"}}, nil
	case "<INT>":
		return []FieldType{
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
		}, nil
	case "map[<STRING>]":
		return []FieldType{{BaseType: "string", ComposedType: "map"}}, nil
	case "map[<BOOL>]":
		return []FieldType{{BaseType: "bool", ComposedType: "map"}}, nil
	case "map[<INT>]":
		return []FieldType{
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
		}, nil
	case "[]<STRING>":
		return []FieldType{{BaseType: "string", ComposedType: "[]"}}, nil
	case "[]<BOOL>":
		return []FieldType{{BaseType: "bool", ComposedType: "[]"}}, nil
	case "[]<INT>":
		return []FieldType{
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
		}, nil
	}

	// Try to remove [N] (array size)
	if len(t) > 0 && t[0] == '[' {
		closeBracketIndex := strings.Index(t, "]")
		if closeBracketIndex == -1 {
			return nil, types.NewValidationError("invalid array size %s", t)
		}

		size := t[1:closeBracketIndex]
		basicType := t[closeBracketIndex+1:]
		switch basicType {
		case "<STRING>":
			return []FieldType{{BaseType: "string", ComposedType: "[N]", Size: size}}, nil
		case "<BOOL>":
			return []FieldType{{BaseType: "bool", ComposedType: "[N]", Size: size}}, nil
		case "<INT>":
			return []FieldType{
				{BaseType: "int", ComposedType: "[N]", Size: size},
				{BaseType: "int8", ComposedType: "[N]", Size: size},
				{BaseType: "int16", ComposedType: "[N]", Size: size},
				{BaseType: "int32", ComposedType: "[N]", Size: size},
				{BaseType: "int64", ComposedType: "[N]", Size: size},
				{BaseType: "uint", ComposedType: "[N]", Size: size},
				{BaseType: "uint8", ComposedType: "[N]", Size: size},
				{BaseType: "uint16", ComposedType: "[N]", Size: size},
				{BaseType: "uint32", ComposedType: "[N]", Size: size},
				{BaseType: "uint64", ComposedType: "[N]", Size: size},
			}, nil
		default:
			return nil, types.NewValidationError("invalid array base type %s", basicType)
		}
	}

	return nil, types.NewValidationError("unknown normalized type %s", t)
}
