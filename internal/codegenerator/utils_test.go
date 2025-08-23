package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
)

func AssertParserValidation(t *testing.T, validation string) *analyzer.Validation {
	t.Helper()

	val, err := analyzer.ParserValidation(validation)
	if err != nil {
		t.Fatalf("failed to parse validation %q: %v", validation, err)
	}

	return val
}
