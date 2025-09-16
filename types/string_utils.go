package types

import (
	"regexp"
	"slices"
	"strings"
)

// emailRegex is a pre-compiled regex for email validation
// This avoids recompiling the regex on every validation call
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// IsValidEmail validates if a string is a valid email format
// Returns true for valid email format, false otherwise
// Empty string returns true (for optional email fields)
func IsValidEmail(email string) bool {
	if email == "" {
		return true // Empty email is valid for optional fields
	}

	// Use pre-compiled regex for better performance
	return emailRegex.MatchString(email)
}

func SliceOnlyContains[S ~[]E, V ~[]E, E comparable](s S, v V) bool {
	for _, item := range s {
		if !slices.Contains(v, item) {
			return false
		}
	}

	return true
}

func SliceNotContains[S ~[]E, V ~[]E, E comparable](s S, v V) bool {
	for _, item := range s {
		if slices.Contains(v, item) {
			return false
		}
	}

	return true
}
