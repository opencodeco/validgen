package common

import (
	"strings"

	"github.com/opencodeco/validgen/types"
)

func HelperFromNormalizedToBasicTypes(t string) []string {
	switch t {
	case "<STRING>":
		return []string{"string"}
	case "<BOOL>":
		return []string{"bool"}
	case "<INT>":
		return []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"}
	case "map[<STRING>]":
		return []string{"map[string]"}
	case "map[<BOOL>]":
		return []string{"map[bool]"}
	case "map[<INT>]":
		return []string{"map[int]", "map[int8]", "map[int16]", "map[int32]", "map[int64]", "map[uint]", "map[uint8]", "map[uint16]", "map[uint32]", "map[uint64]"}
	case "[]<STRING>":
		return []string{"[]string"}
	case "[]<BOOL>":
		return []string{"[]bool"}
	case "[]<INT>":
		return []string{"[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64"}
	}

	// Try to remove [N]
	if len(t) > 0 && t[0] == '[' {
		closeBracketIndex := strings.Index(t, "]")
		if closeBracketIndex == -1 {
			return []string{"invalid"}
		}

		size := t[1:closeBracketIndex]
		basicType := t[closeBracketIndex+1:]
		sizeInType := "[" + size + "]"
		result := []string{}
		switch basicType {
		case "<STRING>":
			result = []string{"[N]string"}
		case "<BOOL>":
			result = []string{"[N]bool"}
		case "<INT>":
			result = []string{"[N]int", "[N]int8", "[N]int16", "[N]int32", "[N]int64", "[N]uint", "[N]uint8", "[N]uint16", "[N]uint32", "[N]uint64"}
		}

		for i := range result {
			result[i] = strings.ReplaceAll(result[i], "[N]", sizeInType)
		}

		return result
	}

	return []string{"invalid"}
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
