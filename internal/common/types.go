package common

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
		"string": {},
		"bool":   {},
		"int":    {},
		"int8":   {},
		"int16":  {},
		"int32":  {},
		"int64":  {},
		"uint":   {},
		"uint8":  {},
		"uint16": {},
		"uint32": {},
		"uint64": {},
	}

	_, ok := goTypes[ft.BaseType]

	return ok
}

func (ft FieldType) NormalizeBaseType() NormalizedBaseType {
	// Base type grouping by type (e.g. string, bool, int and float)

	normalizedBaseType := map[string]NormalizedBaseType{
		"string": StringType,
		"bool":   BoolType,
		"int":    IntType,
		"int8":   IntType,
		"int16":  IntType,
		"int32":  IntType,
		"int64":  IntType,
		"uint":   IntType,
		"uint8":  IntType,
		"uint16": IntType,
		"uint32": IntType,
		"uint64": IntType,
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

// TODO: precisa funcao abaixo?
