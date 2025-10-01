package common

import "strings"

func ExtractPackage(fieldType string) string {
	if pkg, _, ok := strings.Cut(fieldType, "."); ok {
		return pkg
	}
	return ""
}

func KeyPath(values ...string) string {
	return strings.Join(values, ".")
}
