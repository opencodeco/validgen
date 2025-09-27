package common

import "strings"

func IsGoType(fieldType string) bool {
	if strings.HasPrefix(fieldType, "map[") {
		// REFACTOR!
		return true
	}

	fieldType = strings.TrimPrefix(fieldType, "[]")
	fieldType = strings.TrimPrefix(fieldType, "[N]")

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
		"map": {},
	}

	_, ok := goTypes[fieldType]

	return ok
}

func ExtractPackage(fieldType string) string {
	if pkg, _, ok := strings.Cut(fieldType, "."); ok {
		return pkg
	}
	return ""
}

func KeyPath(values ...string) string {
	return strings.Join(values, ".")
}
