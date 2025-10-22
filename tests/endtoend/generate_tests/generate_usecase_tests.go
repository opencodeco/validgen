package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/opencodeco/validgen/internal/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type testCase struct {
	validation     string
	normalizedType string
	validCase      string
	invalidCase    string
	errorMessage   string
	generateOnly   string
}

var testCases = []struct {
	operation string
	testCases []testCase
}{
	// email operations
	{
		operation: "email",
		testCases: []testCase{
			{
				// email: "<STRING>"
				validation:     `email`,
				normalizedType: `<STRING>`,
				validCase:      `"abcde@example.com"`,
				invalidCase:    `"abcde@example"`,
				errorMessage:   `{{.FieldName}} must be a valid email`,
			},
		},
	},

	// required operations
	{
		operation: "required",
		testCases: []testCase{
			// required: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				validation:     `required`,
				normalizedType: `<STRING>`,
				validCase:      `"abcde"`,
				invalidCase:    `""`,
				errorMessage:   `{{.FieldName}} is required`,
			},
			{
				validation:     `required`,
				normalizedType: `<INT>`,
				validCase:      `32`,
				invalidCase:    `0`,
				errorMessage:   `{{.FieldName}} is required`,
			},
			{
				validation:     `required`,
				normalizedType: `<FLOAT>`,
				validCase:      `12.34`,
				invalidCase:    `0`,
				errorMessage:   `{{.FieldName}} is required`,
			},
			{
				validation:     `required`,
				normalizedType: `<BOOL>`,
				validCase:      `true`,
				invalidCase:    `false`,
				errorMessage:   `{{.FieldName}} is required`,
			},

			// required: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				validation:     `required`,
				normalizedType: `[]<STRING>`,
				validCase:      `{{.BasicType}}{"abcde"}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
			{
				validation:     `required`,
				normalizedType: `[]<INT>`,
				validCase:      `{{.BasicType}}{32}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
			{
				validation:     `required`,
				normalizedType: `[]<FLOAT>`,
				validCase:      `{{.BasicType}}{12.34}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
			{
				validation:     `required`,
				normalizedType: `[]<BOOL>`,
				validCase:      `{{.BasicType}}{true}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},

			// required: "[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>"
			{
				validation:     `required`,
				normalizedType: `[N]<STRING>`,
				validCase:      `{{.BasicType}}{"abcde"}`,
				invalidCase:    `--`,
				errorMessage:   `{{.FieldName}} must not be empty`,
				generateOnly:   "pointer",
			},
			{
				validation:     `required`,
				normalizedType: `[N]<INT>`,
				validCase:      `{{.BasicType}}{32}`,
				invalidCase:    `--`,
				errorMessage:   `{{.FieldName}} must not be empty`,
				generateOnly:   "pointer",
			},
			{
				validation:     `required`,
				normalizedType: `[N]<FLOAT>`,
				validCase:      `{{.BasicType}}{12.34}`,
				invalidCase:    `--`,
				errorMessage:   `{{.FieldName}} must not be empty`,
				generateOnly:   "pointer",
			},
			{
				validation:     `required`,
				normalizedType: `[N]<BOOL>`,
				validCase:      `{{.BasicType}}{true}`,
				invalidCase:    `--`,
				errorMessage:   `{{.FieldName}} must not be empty`,
				generateOnly:   "pointer",
			},

			// required: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				validation:     `required`,
				normalizedType: `map[<STRING>]`,
				validCase:      `{{.BasicType}}{"abcde":"value"}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
			{
				validation:     `required`,
				normalizedType: `map[<INT>]`,
				validCase:      `{{.BasicType}}{32:64}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
			{
				validation:     `required`,
				normalizedType: `map[<FLOAT>]`,
				validCase:      `{{.BasicType}}{12.34:56.78}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
			{
				validation:     `required`,
				normalizedType: `map[<BOOL>]`,
				validCase:      `{{.BasicType}}{true:true}`,
				invalidCase:    `{{.BasicType}}{}`,
				errorMessage:   `{{.FieldName}} must not be empty`,
			},
		},
	},

	// eq operations
	{
		operation: "eq",
		testCases: []testCase{
			// eq: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				validation:     `eq=abcde`,
				normalizedType: `<STRING>`,
				validCase:      `"abcde"`,
				invalidCase:    `"fghij"`,
				errorMessage:   `{{.FieldName}} must be equal to 'abcde'`,
			},
			{
				validation:     `eq=32`,
				normalizedType: `<INT>`,
				validCase:      `32`,
				invalidCase:    `64`,
				errorMessage:   `{{.FieldName}} must be equal to 32`,
			},
			{
				validation:     `eq=12.34`,
				normalizedType: `<FLOAT>`,
				validCase:      `12.34`,
				invalidCase:    `34.56`,
				errorMessage:   `{{.FieldName}} must be equal to 12.34`,
			},
			{
				validation:     `eq=true`,
				normalizedType: `<BOOL>`,
				validCase:      `true`,
				invalidCase:    `false`,
				errorMessage:   `{{.FieldName}} must be equal to true`,
			},
		},
	},

	// neq operations
	{
		operation: "neq",
		testCases: []testCase{
			// neq: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				validation:     `neq=abcde`,
				normalizedType: `<STRING>`,
				validCase:      `"fghij"`,
				invalidCase:    `"abcde"`,
				errorMessage:   `{{.FieldName}} must not be equal to 'abcde'`,
			},
			{
				validation:     `neq=32`,
				normalizedType: `<INT>`,
				validCase:      `64`,
				invalidCase:    `32`,
				errorMessage:   `{{.FieldName}} must not be equal to 32`,
			},
			{
				validation:     `neq=12.34`,
				normalizedType: `<FLOAT>`,
				validCase:      `34.56`,
				invalidCase:    `12.34`,
				errorMessage:   `{{.FieldName}} must not be equal to 12.34`,
			},
			{
				validation:     `neq=true`,
				normalizedType: `<BOOL>`,
				validCase:      `false`,
				invalidCase:    `true`,
				errorMessage:   `{{.FieldName}} must not be equal to true`,
			},
		},
	},

	// gt operations
	{
		operation: "gt",
		testCases: []testCase{
			// gt: "<INT>", "<FLOAT>"
			{
				validation:     `gt=32`,
				normalizedType: `<INT>`,
				validCase:      `33`,
				invalidCase:    `31`,
				errorMessage:   `{{.FieldName}} must be > 32`,
			},
			{
				validation:     `gt=12.34`,
				normalizedType: `<FLOAT>`,
				validCase:      `12.35`,
				invalidCase:    `12.34`,
				errorMessage:   `{{.FieldName}} must be > 12.34`,
			},
		},
	},

	// gte operations
	{
		operation: "gte",
		testCases: []testCase{
			// gte: "<INT>", "<FLOAT>"
			{
				validation:     `gte=32`,
				normalizedType: `<INT>`,
				validCase:      `32`,
				invalidCase:    `31`,
				errorMessage:   `{{.FieldName}} must be >= 32`,
			},
			{
				validation:     `gte=12.34`,
				normalizedType: `<FLOAT>`,
				validCase:      `12.34`,
				invalidCase:    `12.33`,
				errorMessage:   `{{.FieldName}} must be >= 12.34`,
			},
		},
	},

	// lt operations
	{
		operation: "lt",
		testCases: []testCase{
			// lt: "<INT>", "<FLOAT>"
			{
				validation:     `lt=32`,
				normalizedType: `<INT>`,
				validCase:      `31`,
				invalidCase:    `33`,
				errorMessage:   `{{.FieldName}} must be < 32`,
			},
			{
				validation:     `lt=12.34`,
				normalizedType: `<FLOAT>`,
				validCase:      `12.33`,
				invalidCase:    `12.35`,
				errorMessage:   `{{.FieldName}} must be < 12.34`,
			},
		},
	},

	// lte operations
	{
		operation: "lte",
		testCases: []testCase{
			// lte: "<INT>", "<FLOAT>"
			{
				validation:     `lte=32`,
				normalizedType: `<INT>`,
				validCase:      `32`,
				invalidCase:    `33`,
				errorMessage:   `{{.FieldName}} must be <= 32`,
			},
			{
				validation:     `lte=12.34`,
				normalizedType: `<FLOAT>`,
				validCase:      `12.34`,
				invalidCase:    `12.35`,
				errorMessage:   `{{.FieldName}} must be <= 12.34`,
			},
		},
	},

	// min operations
	{
		operation: "min",
		testCases: []testCase{
			// min: "<STRING>"
			{
				validation:     `min=5`,
				normalizedType: `<STRING>`,
				validCase:      `"abcde"`,
				invalidCase:    `"abc"`,
				errorMessage:   `{{.FieldName}} length must be >= 5`,
			},

			// min: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				validation:     `min=2`,
				normalizedType: `[]<STRING>`,
				validCase:      `{{.BasicType}}{"abc", "def"}`,
				invalidCase:    `{{.BasicType}}{"abc"}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
			{
				validation:     `min=2`,
				normalizedType: `[]<INT>`,
				validCase:      `{{.BasicType}}{65, 67}`,
				invalidCase:    `{{.BasicType}}{65}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
			{
				validation:     `min=2`,
				normalizedType: `[]<FLOAT>`,
				validCase:      `{{.BasicType}}{65.65, 67.67}`,
				invalidCase:    `{{.BasicType}}{65.65}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
			{
				validation:     `min=2`,
				normalizedType: `[]<BOOL>`,
				validCase:      `{{.BasicType}}{true, false}`,
				invalidCase:    `{{.BasicType}}{true}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},

			// min: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				validation:     `min=2`,
				normalizedType: `map[<STRING>]`,
				validCase:      `{{.BasicType}}{"a": "1", "b": "2"}`,
				invalidCase:    `{{.BasicType}}{"a": "1"}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
			{
				validation:     `min=2`,
				normalizedType: `map[<INT>]`,
				validCase:      `{{.BasicType}}{1: 65, 2: 67}`,
				invalidCase:    `{{.BasicType}}{1: 65}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
			{
				validation:     `min=2`,
				normalizedType: `map[<FLOAT>]`,
				validCase:      `{{.BasicType}}{1: 65.65, 2: 67.67}`,
				invalidCase:    `{{.BasicType}}{1: 65.65}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
			{
				validation:     `min=2`,
				normalizedType: `map[<BOOL>]`,
				validCase:      `{{.BasicType}}{true: true, false: false}`,
				invalidCase:    `{{.BasicType}}{true: true}`,
				errorMessage:   `{{.FieldName}} must have at least 2 elements`,
			},
		},
	},

	// max operations
	{
		operation: "max",
		testCases: []testCase{
			// max: "<STRING>"
			{
				validation:     `max=3`,
				normalizedType: `<STRING>`,
				validCase:      `"abc"`,
				invalidCase:    `"abcde"`,
				errorMessage:   `{{.FieldName}} length must be <= 3`,
			},

			// max: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				validation:     `max=2`,
				normalizedType: `[]<STRING>`,
				validCase:      `{{.BasicType}}{"abc", "def"}`,
				invalidCase:    `{{.BasicType}}{"abc", "def", "ghi"}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},
			{
				validation:     `max=2`,
				normalizedType: `[]<INT>`,
				validCase:      `{{.BasicType}}{65, 67}`,
				invalidCase:    `{{.BasicType}}{65, 66, 67}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},
			{
				validation:     `max=2`,
				normalizedType: `[]<FLOAT>`,
				validCase:      `{{.BasicType}}{65.65, 67.67}`,
				invalidCase:    `{{.BasicType}}{65.65, 66.66, 67.67}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},
			{
				validation:     `max=2`,
				normalizedType: `[]<BOOL>`,
				validCase:      `{{.BasicType}}{true, false}`,
				invalidCase:    `{{.BasicType}}{true, false, true}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},

			// max: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				validation:     `max=2`,
				normalizedType: `map[<STRING>]`,
				validCase:      `{{.BasicType}}{"a": "1", "b": "2"}`,
				invalidCase:    `{{.BasicType}}{"a": "1", "b": "2", "c": "3"}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},
			{
				validation:     `max=2`,
				normalizedType: `map[<INT>]`,
				validCase:      `{{.BasicType}}{1: 65, 2: 67}`,
				invalidCase:    `{{.BasicType}}{1: 65, 2: 67, 3: 68}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},
			{
				validation:     `max=2`,
				normalizedType: `map[<FLOAT>]`,
				validCase:      `{{.BasicType}}{1: 65.65, 2: 67.67}`,
				invalidCase:    `{{.BasicType}}{1: 65.65, 2: 66.66, 3: 67.67}`,
				errorMessage:   `{{.FieldName}} must have at most 2 elements`,
			},
			{
				validation:     `max=1`,
				normalizedType: `map[<BOOL>]`,
				validCase:      `{{.BasicType}}{true: true}`,
				invalidCase:    `{{.BasicType}}{true: true, false: false}`,
				errorMessage:   `{{.FieldName}} must have at most 1 elements`,
			},
		},
	},

	// eq_ignore_case operations
	{
		operation: "eq_ignore_case",
		testCases: []testCase{
			// eq_ignore_case: "<STRING>"
			{
				validation:     `eq_ignore_case=abcde`,
				normalizedType: `<STRING>`,
				validCase:      `"AbCdE"`,
				invalidCase:    `"a1b2c3"`,
				errorMessage:   `{{.FieldName}} must be equal to 'abcde'`,
			},
		},
	},

	// neq_ignore_case operations
	{
		operation: "neq_ignore_case",
		testCases: []testCase{
			// neq_ignore_case: "<STRING>"
			{
				validation:     `neq_ignore_case=abcde`,
				normalizedType: `<STRING>`,
				validCase:      `"a1b2c3"`,
				invalidCase:    `"AbCdE"`,
				errorMessage:   `{{.FieldName}} must not be equal to 'abcde'`,
			},
		},
	},

	// len operations
	{
		operation: "len",
		testCases: []testCase{
			// len: "<STRING>"
			{
				validation:     `len=2`,
				normalizedType: `<STRING>`,
				validCase:      `"ab"`,
				invalidCase:    `"abcde"`,
				errorMessage:   `{{.FieldName}} length must be 2`,
			},

			// len: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				validation:     `len=2`,
				normalizedType: `[]<STRING>`,
				validCase:      `{{.BasicType}}{"abc", "def"}`,
				invalidCase:    `{{.BasicType}}{"abc", "def", "ghi"}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
			{
				validation:     `len=2`,
				normalizedType: `[]<INT>`,
				validCase:      `{{.BasicType}}{65, 67}`,
				invalidCase:    `{{.BasicType}}{65, 66, 67}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
			{
				validation:     `len=2`,
				normalizedType: `[]<FLOAT>`,
				validCase:      `{{.BasicType}}{65.65, 67.67}`,
				invalidCase:    `{{.BasicType}}{65.65, 66.66, 67.67}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
			{
				validation:     `len=2`,
				normalizedType: `[]<BOOL>`,
				validCase:      `{{.BasicType}}{true, false}`,
				invalidCase:    `{{.BasicType}}{true, false, true}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},

			// len: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				validation:     `len=2`,
				normalizedType: `map[<STRING>]`,
				validCase:      `{{.BasicType}}{"a": "1", "b": "2"}`,
				invalidCase:    `{{.BasicType}}{"a": "1", "b": "2", "c": "3"}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
			{
				validation:     `len=2`,
				normalizedType: `map[<INT>]`,
				validCase:      `{{.BasicType}}{1: 65, 2: 67}`,
				invalidCase:    `{{.BasicType}}{1: 65, 2: 67, 3: 68}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
			{
				validation:     `len=2`,
				normalizedType: `map[<FLOAT>]`,
				validCase:      `{{.BasicType}}{1: 65.65, 2: 67.67}`,
				invalidCase:    `{{.BasicType}}{1: 65.65, 2: 66.66, 3: 67.67}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
			{
				validation:     `len=2`,
				normalizedType: `map[<BOOL>]`,
				validCase:      `{{.BasicType}}{true: true, false: false}`,
				invalidCase:    `{{.BasicType}}{true: true}`,
				errorMessage:   `{{.FieldName}} must have exactly 2 elements`,
			},
		},
	},

	// in operations
	{
		operation: "in",
		testCases: []testCase{
			// in: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				validation:     `in=ab cd ef`,
				normalizedType: `<STRING>`,
				validCase:      `"cd"`,
				invalidCase:    `"fg"`,
				errorMessage:   `{{.FieldName}} must be one of 'ab' 'cd' 'ef'`,
			},
			{
				validation:     `in=12 34 56`,
				normalizedType: `<INT>`,
				validCase:      `34`,
				invalidCase:    `78`,
				errorMessage:   `{{.FieldName}} must be one of '12' '34' '56'`,
			},
			{
				validation:     `in=11.11 22.22 33.33`,
				normalizedType: `<FLOAT>`,
				validCase:      `22.22`,
				invalidCase:    `44.44`,
				errorMessage:   `{{.FieldName}} must be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `in=true`,
				normalizedType: `<BOOL>`,
				validCase:      `true`,
				invalidCase:    `false`,
				errorMessage:   `{{.FieldName}} must be one of 'true'`,
			},

			// in: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				validation:     `in=ab cd ef`,
				normalizedType: `[]<STRING>`,
				validCase:      `{{.BasicType}}{"ab", "ef"}`,
				invalidCase:    `{{.BasicType}}{"ab", "gh", "ef"}`,
				errorMessage:   `{{.FieldName}} elements must be one of 'ab' 'cd' 'ef'`,
			},
			{
				validation:     `in=12 34 56`,
				normalizedType: `[]<INT>`,
				validCase:      `{{.BasicType}}{12, 56}`,
				invalidCase:    `{{.BasicType}}{12, 78, 56}`,
				errorMessage:   `{{.FieldName}} elements must be one of '12' '34' '56'`,
			},
			{
				validation:     `in=11.11 22.22 33.33`,
				normalizedType: `[]<FLOAT>`,
				validCase:      `{{.BasicType}}{11.11, 22.22}`,
				invalidCase:    `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage:   `{{.FieldName}} elements must be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `in=true`,
				normalizedType: `[]<BOOL>`,
				validCase:      `{{.BasicType}}{true, true}`,
				invalidCase:    `{{.BasicType}}{true, false, true}`,
				errorMessage:   `{{.FieldName}} elements must be one of 'true'`,
			},

			// in: "[<N>]<STRING>", "[<N>]<INT>", "[<N>]<FLOAT>", "[<N>]<BOOL>"
			{
				validation:     `in=ab cd ef`,
				normalizedType: `[N]<STRING>`,
				validCase:      `{{.BasicType}}{"ab", "ef", "ab"}`,
				invalidCase:    `{{.BasicType}}{"ab", "gh", "ef"}`,
				errorMessage:   `{{.FieldName}} elements must be one of 'ab' 'cd' 'ef'`,
			},
			{
				validation:     `in=12 34 56`,
				normalizedType: `[N]<INT>`,
				validCase:      `{{.BasicType}}{12, 56, 12}`,
				invalidCase:    `{{.BasicType}}{12, 78, 56}`,
				errorMessage:   `{{.FieldName}} elements must be one of '12' '34' '56'`,
			},
			{
				validation:     `in=11.11 22.22 33.33`,
				normalizedType: `[N]<FLOAT>`,
				validCase:      `{{.BasicType}}{11.11, 22.22, 11.11}`,
				invalidCase:    `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage:   `{{.FieldName}} elements must be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `in=true`,
				normalizedType: `[N]<BOOL>`,
				validCase:      `{{.BasicType}}{true, true, true}`,
				invalidCase:    `{{.BasicType}}{true, false, true}`,
				errorMessage:   `{{.FieldName}} elements must be one of 'true'`,
			},

			// in: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				validation:     `in=a b c`,
				normalizedType: `map[<STRING>]`,
				validCase:      `{{.BasicType}}{"a": "1", "b": "2", "c": "3"}`,
				invalidCase:    `{{.BasicType}}{"a": "1", "d": "9", "c": "3"}`,
				errorMessage:   `{{.FieldName}} elements must be one of 'a' 'b' 'c'`,
			},
			{
				validation:     `in=1 2 3`,
				normalizedType: `map[<INT>]`,
				validCase:      `{{.BasicType}}{1: 65, 2: 67, 3: 68}`,
				invalidCase:    `{{.BasicType}}{1: 65, 4: 69, 3: 68}`,
				errorMessage:   `{{.FieldName}} elements must be one of '1' '2' '3'`,
			},
			{
				validation:     `in=11.11 22.22 33.33`,
				normalizedType: `map[<FLOAT>]`,
				validCase:      `{{.BasicType}}{11.11: 11.11, 22.22: 22.22, 33.33: 33.33}`,
				invalidCase:    `{{.BasicType}}{11.11: 11.11, 44.44: 44.44, 33.33: 33.33}`,
				errorMessage:   `{{.FieldName}} elements must be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `in=false`,
				normalizedType: `map[<BOOL>]`,
				validCase:      `{{.BasicType}}{false: false}`,
				invalidCase:    `{{.BasicType}}{true: true, false: false}`,
				errorMessage:   `{{.FieldName}} elements must be one of 'false'`,
			},
		},
	},

	// nin operations
	{
		operation: "nin",
		testCases: []testCase{
			// nin: "[<STRING>]", "[<INT>]", "[<FLOAT>]", "[<BOOL>]"
			{
				validation:     `nin=ab cd ef`,
				normalizedType: `<STRING>`,
				validCase:      `"fg"`,
				invalidCase:    `"cd"`,
				errorMessage:   `{{.FieldName}} must not be one of 'ab' 'cd' 'ef'`,
			},
			{
				validation:     `nin=12 34 56`,
				normalizedType: `<INT>`,
				validCase:      `78`,
				invalidCase:    `34`,
				errorMessage:   `{{.FieldName}} must not be one of '12' '34' '56'`,
			},
			{
				validation:     `nin=11.11 22.22 33.33`,
				normalizedType: `<FLOAT>`,
				validCase:      `44.44`,
				invalidCase:    `22.22`,
				errorMessage:   `{{.FieldName}} must not be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `nin=true`,
				normalizedType: `<BOOL>`,
				validCase:      `false`,
				invalidCase:    `true`,
				errorMessage:   `{{.FieldName}} must not be one of 'true'`,
			},

			// nin: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				validation:     `nin=ab cd ef`,
				normalizedType: `[]<STRING>`,
				validCase:      `{{.BasicType}}{"gh", "ij", "kl"}`,
				invalidCase:    `{{.BasicType}}{"ab", "ef"}`,
				errorMessage:   `{{.FieldName}} elements must not be one of 'ab' 'cd' 'ef'`,
			},
			{
				validation:     `nin=12 34 56`,
				normalizedType: `[]<INT>`,
				validCase:      `{{.BasicType}}{78, 91}`,
				invalidCase:    `{{.BasicType}}{12, 78, 56}`,
				errorMessage:   `{{.FieldName}} elements must not be one of '12' '34' '56'`,
			},
			{
				validation:     `nin=11.11 22.22 33.33`,
				normalizedType: `[]<FLOAT>`,
				validCase:      `{{.BasicType}}{44.44, 55.55, 66.66}`,
				invalidCase:    `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage:   `{{.FieldName}} elements must not be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `nin=true`,
				normalizedType: `[]<BOOL>`,
				validCase:      `{{.BasicType}}{false, false, false}`,
				invalidCase:    `{{.BasicType}}{true, false, true}`,
				errorMessage:   `{{.FieldName}} elements must not be one of 'true'`,
			},

			// nin: "[<N>]<STRING>", "[<N>]<INT>", "[<N>]<FLOAT>", "[<N>]<BOOL>"
			{
				validation:     `nin=ab cd ef`,
				normalizedType: `[N]<STRING>`,
				validCase:      `{{.BasicType}}{"gh", "ij", "kl"}`,
				invalidCase:    `{{.BasicType}}{"ab", "gh", "ef"}`,
				errorMessage:   `{{.FieldName}} elements must not be one of 'ab' 'cd' 'ef'`,
			},
			{
				validation:     `nin=12 34 56`,
				normalizedType: `[N]<INT>`,
				validCase:      `{{.BasicType}}{78, 91, 23}`,
				invalidCase:    `{{.BasicType}}{12, 78, 56}`,
				errorMessage:   `{{.FieldName}} elements must not be one of '12' '34' '56'`,
			},
			{
				validation:     `nin=11.11 22.22 33.33`,
				normalizedType: `[N]<FLOAT>`,
				validCase:      `{{.BasicType}}{44.44, 55.55, 66.66}`,
				invalidCase:    `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage:   `{{.FieldName}} elements must not be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `nin=true`,
				normalizedType: `[N]<BOOL>`,
				validCase:      `{{.BasicType}}{false, false, false}`,
				invalidCase:    `{{.BasicType}}{true, false, true}`,
				errorMessage:   `{{.FieldName}} elements must not be one of 'true'`,
			},

			// nin: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				validation:     `nin=a b c`,
				normalizedType: `map[<STRING>]`,
				validCase:      `{{.BasicType}}{"d": "1", "e": "2", "f": "3"}`,
				invalidCase:    `{{.BasicType}}{"a": "1", "d": "9", "c": "3"}`,
				errorMessage:   `{{.FieldName}} elements must not be one of 'a' 'b' 'c'`,
			},
			{
				validation:     `nin=1 2 3`,
				normalizedType: `map[<INT>]`,
				validCase:      `{{.BasicType}}{5: 55, 6: 66, 7: 77}`,
				invalidCase:    `{{.BasicType}}{1: 11, 4: 44, 3: 33}`,
				errorMessage:   `{{.FieldName}} elements must not be one of '1' '2' '3'`,
			},
			{
				validation:     `nin=11.11 22.22 33.33`,
				normalizedType: `map[<FLOAT>]`,
				validCase:      `{{.BasicType}}{44.44: 44.44, 55.55: 55.55, 66.66: 66.66}`,
				invalidCase:    `{{.BasicType}}{11.11: 11.11, 44.44: 44.44, 33.33: 33.33}`,
				errorMessage:   `{{.FieldName}} elements must not be one of '11.11' '22.22' '33.33'`,
			},
			{
				validation:     `nin=false`,
				normalizedType: `map[<BOOL>]`,
				validCase:      `{{.BasicType}}{true: true}`,
				invalidCase:    `{{.BasicType}}{true: true, false: false}`,
				errorMessage:   `{{.FieldName}} elements must not be one of 'false'`,
			},
		},
	},
}

type AllTestCasesToGenerate struct {
	TestCases []TestCaseToGenerate
}

type TestCaseToGenerate struct {
	StructName string
	Tests      []TestCase
}

type TestCase struct {
	FieldName    string
	Validation   string
	FieldType    string
	BasicType    string
	ValidCase    string
	InvalidCase  string
	ErrorMessage string
}

func generateTestCases() {
	generateTestCasesFile("no_pointer_tests.tpl", "generated_no_pointer_tests.go", false)
	generateTestCasesFile("pointer_tests.tpl", "generated_pointer_tests.go", true)
}

func generateTestCasesFile(tpl, dest string, pointer bool) {
	log.Printf("Generating test cases file: tpl[%s] dest[%s] pointer[%v]\n", tpl, dest, pointer)

	allTestsToGenerate := AllTestCasesToGenerate{}

	for _, testCase := range testCases {
		structName := testCase.operation + "StructFields"
		if pointer {
			structName += "Pointer"
		}
		allTestsToGenerate.TestCases = append(allTestsToGenerate.TestCases, TestCaseToGenerate{
			StructName: structName,
		})
		for _, toGenerate := range testCase.testCases {
			// Default ("") gen no pointer and pointer test.
			if toGenerate.generateOnly != "" {
				if toGenerate.generateOnly == "pointer" && !pointer {
					continue
				}
				if toGenerate.generateOnly == "nopointer" && pointer {
					continue
				}
			}
			normalizedType := toGenerate.normalizedType
			if pointer {
				normalizedType = "*" + normalizedType
			}
			fTypes := common.HelperFromNormalizedToBasicTypes(normalizedType)
			sNames := common.HelperFromNormalizedToStringNames(normalizedType)
			for i := range fTypes {
				op, _, _ := strings.Cut(toGenerate.validation, "=")
				fieldName := "Field" + cases.Title(language.Und).String(op) + sNames[i]
				basicType, _ := strings.CutPrefix(fTypes[i], "*")
				allTestsToGenerate.TestCases[len(allTestsToGenerate.TestCases)-1].Tests = append(allTestsToGenerate.TestCases[len(allTestsToGenerate.TestCases)-1].Tests, TestCase{
					FieldName:    fieldName,
					Validation:   toGenerate.validation,
					FieldType:    fTypes[i],
					BasicType:    basicType,
					ValidCase:    strings.ReplaceAll(toGenerate.validCase, "{{.BasicType}}", basicType),
					InvalidCase:  strings.ReplaceAll(toGenerate.invalidCase, "{{.BasicType}}", basicType),
					ErrorMessage: strings.ReplaceAll(toGenerate.errorMessage, "{{.FieldName}}", fieldName),
				})
			}
		}
	}

	if err := allTestsToGenerate.GenerateFile(tpl, dest); err != nil {
		log.Fatalf("error generation usecases file %s", err)

	}

	log.Printf("Generating %s done\n", dest)
}

func (at *AllTestCasesToGenerate) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New("UsecaseTests").Parse(string(tpl))
	if err != nil {
		return err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, at); err != nil {
		return err
	}

	formattedCode, err := format.Source(code.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(output, formattedCode, 0644); err != nil {
		return err
	}

	return nil
}
