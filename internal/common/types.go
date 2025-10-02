package common

import "strings"

type FieldType struct {
	ComposedType string // array ([N]), map (map) or slice ([])
	BaseType     string // base type (e.g. string, int, etc.)
	Size         string // for arrays
}

func (ft FieldType) IsGoType() bool {
	if ft.ComposedType == "map" {
		return true
	}

	goTypes := map[string]struct{}{
		"string":  {},
		"bool":    {},
		"int":     {},
		"int8":    {},
		"int16":   {},
		"int32":   {},
		"int64":   {},
		"uint":    {},
		"uint8":   {},
		"uint16":  {},
		"uint32":  {},
		"uint64":  {},
		"float32": {},
		"float64": {},
		// "complex64":  {},
		// "complex128": {},
	}

	_, ok := goTypes[ft.BaseType]

	return ok
}

func (ft FieldType) NormalizeBaseType() NormalizedBaseType {
	// Base type grouping by type (e.g. string, bool, int and float)

	normalizedBaseType := map[string]NormalizedBaseType{
		"string":  StringType,
		"bool":    BoolType,
		"int":     IntType,
		"int8":    IntType,
		"int16":   IntType,
		"int32":   IntType,
		"int64":   IntType,
		"uint":    IntType,
		"uint8":   IntType,
		"uint16":  IntType,
		"uint32":  IntType,
		"uint64":  IntType,
		"float32": FloatType,
		"float64": FloatType,
		// "complex64":  ComplexType,
		// "complex128": ComplexType,
	}

	return normalizedBaseType[ft.BaseType]
}

func (ft FieldType) ToString() string {
	switch ft.ComposedType {
	case "[N]":
		return "[N]" + ft.BaseType
	case "[]":
		return "[]" + ft.BaseType
	case "map":
		return "map[" + ft.BaseType + "]"
	}

	return ft.BaseType
}

func (ft FieldType) ToNormalizedString() string {
	switch ft.ComposedType {
	case "[N]":
		return "[N]" + ft.NormalizeBaseType().String()
	case "[]":
		return "[]" + ft.NormalizeBaseType().String()
	case "map":
		return "map[" + ft.NormalizeBaseType().String() + "]"
	}

	return ft.NormalizeBaseType().String()
}

func FromNormalizedToBasicTypes(t string) []string {
	switch t {
	case "<STRING>":
		return []string{"string"}
	case "<BOOL>":
		return []string{"bool"}
	case "<INT>":
		return []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"}
	case "<FLOAT>":
		return []string{"float32", "float64"}
	case "map[<STRING>]":
		return []string{"map[string]"}
	case "map[<BOOL>]":
		return []string{"map[bool]"}
	case "map[<INT>]":
		return []string{"map[int]", "map[int8]", "map[int16]", "map[int32]", "map[int64]", "map[uint]", "map[uint8]", "map[uint16]", "map[uint32]", "map[uint64]"}
	case "map[<FLOAT>]":
		return []string{"map[float32]", "map[float64]"}
	case "[]<STRING>":
		return []string{"[]string"}
	case "[]<BOOL>":
		return []string{"[]bool"}
	case "[]<INT>":
		return []string{"[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64"}
	case "[]<FLOAT>":
		return []string{"[]float32", "[]float64"}
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
		case "<FLOAT>":
			result = []string{"[N]float32", "[N]float64"}
		}

		for i := range result {
			result[i] = strings.ReplaceAll(result[i], "[N]", sizeInType)
		}

		return result
	}

	return []string{"invalid"}
}
