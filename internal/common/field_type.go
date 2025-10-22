package common

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FieldType struct {
	ComposedType string // array ([N]), map (map) or slice ([])
	BaseType     string // base type (e.g. string, int, etc.)
	Size         string // for arrays
}

func (ft FieldType) IsGoType() bool {
	if ft.ComposedType == "map" || ft.ComposedType == "*map" {
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
	}

	return normalizedBaseType[ft.BaseType]
}

func (ft FieldType) ToGenericType() string {
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

func (ft FieldType) ToType() string {
	switch ft.ComposedType {
	case "[N]":
		return fmt.Sprintf("[%s]%s", ft.Size, ft.BaseType)
	case "[]":
		return "[]" + ft.BaseType
	case "map":
		return "map[" + ft.BaseType + "]" + ft.BaseType
	case "*":
		return "*" + ft.BaseType
	case "*[N]":
		return fmt.Sprintf("*[%s]%s", ft.Size, ft.BaseType)
	case "*[]":
		return "*[]" + ft.BaseType
	case "*map":
		return "*map[" + ft.BaseType + "]" + ft.BaseType
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
	case "*[N]":
		return "*[N]" + ft.NormalizeBaseType().String()
	case "*[]":
		return "*[]" + ft.NormalizeBaseType().String()
	case "*map":
		return "*map[" + ft.NormalizeBaseType().String() + "]"
	case "*":
		return "*" + ft.NormalizeBaseType().String()
	}

	return ft.NormalizeBaseType().String()
}

func (ft FieldType) ToStringName() string {
	result := ""

	switch ft.ComposedType {
	case "[N]":
		result = "Array"
	case "[]":
		result = "Slice"
	case "map":
		result = "Map"
	case "*[N]":
		result = "ArrayPointer"
	case "*[]":
		result = "SlicePointer"
	case "*map":
		result = "MapPointer"
	case "*":
		result = "Pointer"
	}

	return cases.Title(language.Und).String(ft.BaseType) + result
}
