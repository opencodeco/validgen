package common

type FieldType struct {
	ComposedType string // array ([N]), map (map) or slice ([])
	BaseType     string // base type (e.g. string, int, etc.)
	Size         string // for arrays
	// IsGoType     bool   // true if is a Go built-in type
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
