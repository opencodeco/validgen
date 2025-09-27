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
		// "int":        {},
		// "int8":       {},
		// "int16":      {},
		// "int32":      {},
		// "int64":      {},
		// "uint":       {},
		"uint8": {},
		// "uint16":     {},
		// "uint32":     {},
		// "uint64":     {},
		// "float32":    {},
		// "float64":    {},
		// "complex64":  {},
		// "complex128": {},
	}

	_, ok := goTypes[ft.BaseType]

	return ok
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
