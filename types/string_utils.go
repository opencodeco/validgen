package types

import (
	"regexp"
	"strings"
)

// emailRegex is a pre-compiled regex for email validation
// This avoids recompiling the regex on every validation call
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ToLower(str string) string {
	return strings.ToLower(str)
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
