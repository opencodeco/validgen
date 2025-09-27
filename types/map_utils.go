package types

import "slices"

func MapOnlyContains[K comparable, V any](m map[K]V, e []K) bool {
	for k := range m {
		if !slices.Contains(e, k) {
			return false
		}
	}

	return true
}

func MapNotContains[K comparable, V any](m map[K]V, e []K) bool {
	for k := range m {
		if slices.Contains(e, k) {
			return false
		}
	}

	return true
}
