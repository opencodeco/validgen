package main

import (
	"strings"
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/analyzer/operations"
	"github.com/opencodeco/validgen/internal/codegenerator"
	"github.com/opencodeco/validgen/internal/common"
)

// generatedOperations is the non-field set make testgen emits.
var generatedOperations = []string{
	"email",
	"required",
	"eq",
	"neq",
	"gt",
	"gte",
	"lt",
	"lte",
	"min",
	"max",
	"eq_ignore_case",
	"neq_ignore_case",
	"len",
	"in",
	"nin",
}

// fieldOperations stay in hand-written tests. Integer field operations are issue #78.
var fieldOperations = []string{
	"eqfield",
	"neqfield",
	"gtfield",
	"gtefield",
	"ltfield",
	"ltefield",
}

// normalizedTypeClasses is every class HelperFromNormalizedToFieldTypes accepts.
var normalizedTypeClasses = []string{
	"<STRING>", "<INT>", "<FLOAT>", "<BOOL>", "<COMPLEX>",
	"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "[]<COMPLEX>",
	"[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>", "[N]<COMPLEX>",
	"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]", "map[<COMPLEX>]",
}

func TestTypesValidationListsNonFieldOperations(t *testing.T) {
	ops := operations.New()
	byTag := map[string]int{}
	for _, entry := range typesValidation {
		byTag[entry.tag]++
		if entry.isFieldValidation {
			t.Errorf("%s is marked as a field validation", entry.tag)
		}
		if !ops.IsValid(entry.tag) {
			t.Errorf("%s is missing from the operations list", entry.tag)
			continue
		}
		if ops.IsFieldOperation(entry.tag) {
			t.Errorf("%s is a field operation", entry.tag)
		}
		if got := ops.ArgsCount(entry.tag); got != entry.argsCount {
			t.Errorf("%s argument count is %v, catalog has %v", entry.tag, got, entry.argsCount)
		}

		classes := map[string]struct{}{}
		for _, tc := range entry.testCases {
			classes[tc.typeClass] = struct{}{}
			if tc.validCase == "" || tc.errorMessage == "" {
				t.Errorf("%s %s is missing a valid case or error message", entry.tag, tc.typeClass)
			}
			if tc.invalidCase == "" {
				t.Errorf("%s %s is missing an invalid case", entry.tag, tc.typeClass)
			}
			if tc.invalidCase == "--" && tc.excludeIf&noPointer == 0 {
				t.Errorf("%s %s uses -- without noPointer", entry.tag, tc.typeClass)
			}
		}
		for _, class := range normalizedTypeClasses {
			if !ops.IsValidByType(entry.tag, class) {
				continue
			}
			if _, ok := classes[class]; !ok {
				t.Errorf("%s accepts %s and typesValidation has no case", entry.tag, class)
			}
		}
	}

	inCatalog := map[string]struct{}{}
	for tag := range byTag {
		inCatalog[tag] = struct{}{}
	}

	for _, op := range generatedOperations {
		if byTag[op] != 1 {
			t.Errorf("%s appears %d times in typesValidation, want 1", op, byTag[op])
		}
		delete(byTag, op)
	}
	for tag := range byTag {
		t.Errorf("%s is in typesValidation and is outside the generated set", tag)
	}

	for _, op := range fieldOperations {
		if !ops.IsValid(op) || !ops.IsFieldOperation(op) {
			t.Errorf("%s should stay a field operation in the operations list", op)
		}
		if _, ok := inCatalog[op]; ok {
			t.Errorf("%s is a field operation and is in typesValidation", op)
		}
	}
}

func TestTypesValidationCasesBuildValidationCode(t *testing.T) {
	gv := codegenerator.GenValidations{}

	for _, entry := range typesValidation {
		for _, tc := range entry.testCases {
			for _, pointer := range []bool{false, true} {
				if tc.excludeIf&noPointer != 0 && !pointer {
					continue
				}

				normalizedType := tc.typeClass
				if pointer {
					normalizedType = "*" + normalizedType
				}

				fieldTypes, err := common.HelperFromNormalizedToFieldTypes(normalizedType)
				if err != nil {
					t.Errorf("%s %s pointer %v: %v", entry.tag, tc.typeClass, pointer, err)
					continue
				}

				validation := entry.tag
				if entry.argsCount != common.ZeroValue {
					validation += "=" + tc.validation
				}
				parsed, err := analyzer.ParserValidation(validation)
				if err != nil {
					t.Errorf("parse %s: %v", validation, err)
					continue
				}

				for _, fieldType := range fieldTypes {
					const fieldName = "Field"
					got, err := gv.BuildValidationCode(fieldName, fieldType, []*analyzer.Validation{parsed})
					if err != nil {
						t.Errorf("%s %s %s pointer %v: %v", entry.tag, tc.typeClass, fieldType.ToType(), pointer, err)
						continue
					}
					if !errorTextMatches(got, fieldName, tc.errorMessage) {
						t.Errorf("%s %s %s pointer %v: generated error text missing %q\n%s",
							entry.tag, tc.typeClass, fieldType.ToType(), pointer, tc.errorMessage, got)
					}
				}
			}
		}
	}
}

func errorTextMatches(generated, fieldName, catalogMessage string) bool {
	message := strings.ReplaceAll(catalogMessage, "{{.FieldName}}", fieldName)
	message = strings.ReplaceAll(message, "{{.Target}}", "\x00")
	message = strings.ReplaceAll(message, "{{.Targets}}", "\x00")

	from := 0
	for _, part := range strings.Split(message, "\x00") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		idx := strings.Index(generated[from:], part)
		if idx < 0 {
			return false
		}
		from += idx + len(part)
	}
	return from > 0
}
