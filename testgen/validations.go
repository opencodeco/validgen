package main

import "github.com/opencodeco/validgen/internal/common"

type excludeIf uint32

const (
	cmpBenchTests excludeIf = 1 << iota
	noPointer
)

type typeValidation struct {
	typeClass    string
	validation   string
	validCase    string
	invalidCase  string
	errorMessage string
	excludeIf    excludeIf
}

var typesValidation = []struct {
	tag               string
	validatorTag      string
	isFieldValidation bool
	argsCount         common.CountValues
	testCases         []typeValidation
}{
	// email operations
	{
		tag:               "email",
		validatorTag:      `email`,
		isFieldValidation: false,
		argsCount:         common.ZeroValue,
		testCases: []typeValidation{
			{
				// email: "<STRING>"
				typeClass:    `<STRING>`,
				validation:   ``,
				validCase:    `"abcde@example.com"`,
				invalidCase:  `"abcde@example"`,
				errorMessage: `{{.FieldName}} must be a valid email`,
			},
		},
	},

	// required operations
	{
		tag:               "required",
		validatorTag:      `required`,
		isFieldValidation: false,
		argsCount:         common.ZeroValue,
		testCases: []typeValidation{
			// required: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				typeClass:    `<STRING>`,
				validation:   ``,
				validCase:    `"abcde"`,
				invalidCase:  `""`,
				errorMessage: `{{.FieldName}} is required`,
			},
			{
				typeClass:    `<INT>`,
				validation:   ``,
				validCase:    `32`,
				invalidCase:  `0`,
				errorMessage: `{{.FieldName}} is required`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   ``,
				validCase:    `12.34`,
				invalidCase:  `0`,
				errorMessage: `{{.FieldName}} is required`,
			},
			{
				typeClass:    `<BOOL>`,
				validation:   ``,
				validCase:    `true`,
				invalidCase:  `false`,
				errorMessage: `{{.FieldName}} is required`,
			},

			// required: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				typeClass:    `[]<STRING>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{"abcde"}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
			{
				typeClass:    `[]<INT>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{32}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
			{
				typeClass:    `[]<FLOAT>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{12.34}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
			{
				typeClass:    `[]<BOOL>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{true}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},

			// required: "[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>"
			{
				typeClass:    `[N]<STRING>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{"abcde"}`,
				invalidCase:  `--`,
				errorMessage: `{{.FieldName}} must not be empty`,
				excludeIf:    noPointer,
			},
			{
				typeClass:    `[N]<INT>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{32}`,
				invalidCase:  `--`,
				errorMessage: `{{.FieldName}} must not be empty`,
				excludeIf:    noPointer,
			},
			{
				typeClass:    `[N]<FLOAT>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{12.34}`,
				invalidCase:  `--`,
				errorMessage: `{{.FieldName}} must not be empty`,
				excludeIf:    noPointer,
			},
			{
				typeClass:    `[N]<BOOL>`,
				validation:   ``,
				validCase:    `{{.BasicType}}{true}`,
				invalidCase:  `--`,
				errorMessage: `{{.FieldName}} must not be empty`,
				excludeIf:    noPointer,
			},

			// required: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				typeClass:    `map[<STRING>]`,
				validation:   ``,
				validCase:    `{{.BasicType}}{"abcde":"value"}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
			{
				typeClass:    `map[<INT>]`,
				validation:   ``,
				validCase:    `{{.BasicType}}{32:64}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
			{
				typeClass:    `map[<FLOAT>]`,
				validation:   ``,
				validCase:    `{{.BasicType}}{12.34:56.78}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
			{
				typeClass:    `map[<BOOL>]`,
				validation:   ``,
				validCase:    `{{.BasicType}}{true:true}`,
				invalidCase:  `{{.BasicType}}{}`,
				errorMessage: `{{.FieldName}} must not be empty`,
			},
		},
	},

	// eq operations
	{
		tag:               "eq",
		validatorTag:      `eq`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// eq: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				typeClass:    `<STRING>`,
				validation:   `abcde`,
				validCase:    `"abcde"`,
				invalidCase:  `"fghij"`,
				errorMessage: `{{.FieldName}} must be equal to '{{.Target}}'`,
			},
			{
				typeClass:    `<INT>`,
				validation:   `32`,
				validCase:    `32`,
				invalidCase:  `64`,
				errorMessage: `{{.FieldName}} must be equal to {{.Target}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `12.34`,
				validCase:    `12.34`,
				invalidCase:  `34.56`,
				errorMessage: `{{.FieldName}} must be equal to {{.Target}}`,
			},
			{
				typeClass:    `<BOOL>`,
				validation:   `true`,
				validCase:    `true`,
				invalidCase:  `false`,
				errorMessage: `{{.FieldName}} must be equal to {{.Target}}`,
			},
		},
	},

	// neq operations
	{
		tag:               "neq",
		validatorTag:      `ne`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// neq: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				typeClass:    `<STRING>`,
				validation:   `abcde`,
				validCase:    `"fghij"`,
				invalidCase:  `"abcde"`,
				errorMessage: `{{.FieldName}} must not be equal to '{{.Target}}'`,
			},
			{
				typeClass:    `<INT>`,
				validation:   `32`,
				validCase:    `64`,
				invalidCase:  `32`,
				errorMessage: `{{.FieldName}} must not be equal to {{.Target}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `12.34`,
				validCase:    `34.56`,
				invalidCase:  `12.34`,
				errorMessage: `{{.FieldName}} must not be equal to {{.Target}}`,
			},
			{
				typeClass:    `<BOOL>`,
				validation:   `true`,
				validCase:    `false`,
				invalidCase:  `true`,
				errorMessage: `{{.FieldName}} must not be equal to {{.Target}}`,
			},
		},
	},

	// gt operations
	{
		tag:               "gt",
		validatorTag:      `gt`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// gt: "<INT>", "<FLOAT>"
			{
				typeClass:    `<INT>`,
				validation:   `32`,
				validCase:    `33`,
				invalidCase:  `31`,
				errorMessage: `{{.FieldName}} must be > {{.Target}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `12.34`,
				validCase:    `12.35`,
				invalidCase:  `12.34`,
				errorMessage: `{{.FieldName}} must be > {{.Target}}`,
			},
		},
	},

	// gte operations
	{
		tag:               "gte",
		validatorTag:      `gte`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// gte: "<INT>", "<FLOAT>"
			{
				typeClass:    `<INT>`,
				validation:   `32`,
				validCase:    `32`,
				invalidCase:  `31`,
				errorMessage: `{{.FieldName}} must be >= {{.Target}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `12.34`,
				validCase:    `12.34`,
				invalidCase:  `12.33`,
				errorMessage: `{{.FieldName}} must be >= {{.Target}}`,
			},
		},
	},

	// lt operations
	{
		tag:               "lt",
		validatorTag:      `lt`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// lt: "<INT>", "<FLOAT>"
			{
				typeClass:    `<INT>`,
				validation:   `32`,
				validCase:    `31`,
				invalidCase:  `33`,
				errorMessage: `{{.FieldName}} must be < {{.Target}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `12.34`,
				validCase:    `12.33`,
				invalidCase:  `12.35`,
				errorMessage: `{{.FieldName}} must be < {{.Target}}`,
			},
		},
	},

	// lte operations
	{
		tag:               "lte",
		validatorTag:      `lte`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// lte: "<INT>", "<FLOAT>"
			{
				typeClass:    `<INT>`,
				validation:   `32`,
				validCase:    `32`,
				invalidCase:  `33`,
				errorMessage: `{{.FieldName}} must be <= {{.Target}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `12.34`,
				validCase:    `12.34`,
				invalidCase:  `12.35`,
				errorMessage: `{{.FieldName}} must be <= {{.Target}}`,
			},
		},
	},

	// min operations
	{
		tag:               "min",
		validatorTag:      `min`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// min: "<STRING>"
			{
				typeClass:    `<STRING>`,
				validation:   `5`,
				validCase:    `"abcde"`,
				invalidCase:  `"abc"`,
				errorMessage: `{{.FieldName}} length must be >= {{.Target}}`,
			},

			// min: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				typeClass:    `[]<STRING>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{"abc", "def"}`,
				invalidCase:  `{{.BasicType}}{"abc"}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
			{
				typeClass:    `[]<INT>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{65, 67}`,
				invalidCase:  `{{.BasicType}}{65}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
			{
				typeClass:    `[]<FLOAT>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{65.65, 67.67}`,
				invalidCase:  `{{.BasicType}}{65.65}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
			{
				typeClass:    `[]<BOOL>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{true, false}`,
				invalidCase:  `{{.BasicType}}{true}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},

			// min: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				typeClass:    `map[<STRING>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{"a": "1", "b": "2"}`,
				invalidCase:  `{{.BasicType}}{"a": "1"}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
			{
				typeClass:    `map[<INT>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{1: 65, 2: 67}`,
				invalidCase:  `{{.BasicType}}{1: 65}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
			{
				typeClass:    `map[<FLOAT>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{1: 65.65, 2: 67.67}`,
				invalidCase:  `{{.BasicType}}{1: 65.65}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
			{
				typeClass:    `map[<BOOL>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{true: true, false: false}`,
				invalidCase:  `{{.BasicType}}{true: true}`,
				errorMessage: `{{.FieldName}} must have at least {{.Target}} elements`,
			},
		},
	},

	// max operations
	{
		tag:               "max",
		validatorTag:      `max`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// max: "<STRING>"
			{
				typeClass:    `<STRING>`,
				validation:   `3`,
				validCase:    `"abc"`,
				invalidCase:  `"abcde"`,
				errorMessage: `{{.FieldName}} length must be <= 3`,
			},

			// max: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				typeClass:    `[]<STRING>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{"abc", "def"}`,
				invalidCase:  `{{.BasicType}}{"abc", "def", "ghi"}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
			{
				typeClass:    `[]<INT>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{65, 67}`,
				invalidCase:  `{{.BasicType}}{65, 66, 67}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
			{
				typeClass:    `[]<FLOAT>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{65.65, 67.67}`,
				invalidCase:  `{{.BasicType}}{65.65, 66.66, 67.67}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
			{
				typeClass:    `[]<BOOL>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{true, false}`,
				invalidCase:  `{{.BasicType}}{true, false, true}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},

			// max: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				typeClass:    `map[<STRING>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{"a": "1", "b": "2"}`,
				invalidCase:  `{{.BasicType}}{"a": "1", "b": "2", "c": "3"}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
			{
				typeClass:    `map[<INT>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{1: 65, 2: 67}`,
				invalidCase:  `{{.BasicType}}{1: 65, 2: 67, 3: 68}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
			{
				typeClass:    `map[<FLOAT>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{1: 65.65, 2: 67.67}`,
				invalidCase:  `{{.BasicType}}{1: 65.65, 2: 66.66, 3: 67.67}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
			{
				typeClass:    `map[<BOOL>]`,
				validation:   `1`,
				validCase:    `{{.BasicType}}{true: true}`,
				invalidCase:  `{{.BasicType}}{true: true, false: false}`,
				errorMessage: `{{.FieldName}} must have at most {{.Target}} elements`,
			},
		},
	},

	// eq_ignore_case operations
	{
		tag:               "eq_ignore_case",
		validatorTag:      `eq_ignore_case`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// eq_ignore_case: "<STRING>"
			{
				typeClass:    `<STRING>`,
				validation:   `abcde`,
				validCase:    `"AbCdE"`,
				invalidCase:  `"a1b2c3"`,
				errorMessage: `{{.FieldName}} must be equal to '{{.Target}}'`,
			},
		},
	},

	// neq_ignore_case operations
	{
		tag:               "neq_ignore_case",
		validatorTag:      `ne_ignore_case`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// neq_ignore_case: "<STRING>"
			{
				typeClass:    `<STRING>`,
				validation:   `abcde`,
				validCase:    `"a1b2c3"`,
				invalidCase:  `"AbCdE"`,
				errorMessage: `{{.FieldName}} must not be equal to '{{.Target}}'`,
			},
		},
	},

	// len operations
	{
		tag:               "len",
		validatorTag:      `len`,
		isFieldValidation: false,
		argsCount:         common.OneValue,
		testCases: []typeValidation{
			// len: "<STRING>"
			{
				typeClass:    `<STRING>`,
				validation:   `2`,
				validCase:    `"ab"`,
				invalidCase:  `"abcde"`,
				errorMessage: `{{.FieldName}} length must be {{.Target}}`,
			},

			// len: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				typeClass:    `[]<STRING>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{"abc", "def"}`,
				invalidCase:  `{{.BasicType}}{"abc", "def", "ghi"}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
			{
				typeClass:    `[]<INT>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{65, 67}`,
				invalidCase:  `{{.BasicType}}{65, 66, 67}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
			{
				typeClass:    `[]<FLOAT>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{65.65, 67.67}`,
				invalidCase:  `{{.BasicType}}{65.65, 66.66, 67.67}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
			{
				typeClass:    `[]<BOOL>`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{true, false}`,
				invalidCase:  `{{.BasicType}}{true, false, true}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},

			// len: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				typeClass:    `map[<STRING>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{"a": "1", "b": "2"}`,
				invalidCase:  `{{.BasicType}}{"a": "1", "b": "2", "c": "3"}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
			{
				typeClass:    `map[<INT>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{1: 65, 2: 67}`,
				invalidCase:  `{{.BasicType}}{1: 65, 2: 67, 3: 68}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
			{
				typeClass:    `map[<FLOAT>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{1: 65.65, 2: 67.67}`,
				invalidCase:  `{{.BasicType}}{1: 65.65, 2: 66.66, 3: 67.67}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
			{
				typeClass:    `map[<BOOL>]`,
				validation:   `2`,
				validCase:    `{{.BasicType}}{true: true, false: false}`,
				invalidCase:  `{{.BasicType}}{true: true}`,
				errorMessage: `{{.FieldName}} must have exactly {{.Target}} elements`,
			},
		},
	},

	// in operations
	{
		tag:               "in",
		validatorTag:      `oneof`,
		isFieldValidation: false,
		argsCount:         common.ManyValues,
		testCases: []typeValidation{
			// in: "<STRING>", "<INT>", "<FLOAT>", "<BOOL>"
			{
				typeClass:    `<STRING>`,
				validation:   `ab cd ef`,
				validCase:    `"cd"`,
				invalidCase:  `"fg"`,
				errorMessage: `{{.FieldName}} must be one of {{.Targets}}`,
			},
			{
				typeClass:    `<INT>`,
				validation:   `12 34 56`,
				validCase:    `34`,
				invalidCase:  `78`,
				errorMessage: `{{.FieldName}} must be one of {{.Targets}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `22.22`,
				invalidCase:  `44.44`,
				errorMessage: `{{.FieldName}} must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `<BOOL>`,
				validation:   `true`,
				validCase:    `true`,
				invalidCase:  `false`,
				errorMessage: `{{.FieldName}} must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},

			// in: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				typeClass:    `[]<STRING>`,
				validation:   `ab cd ef`,
				validCase:    `{{.BasicType}}{"ab", "ef"}`,
				invalidCase:  `{{.BasicType}}{"ab", "gh", "ef"}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `[]<INT>`,
				validation:   `12 34 56`,
				validCase:    `{{.BasicType}}{12, 56}`,
				invalidCase:  `{{.BasicType}}{12, 78, 56}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `[]<FLOAT>`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `{{.BasicType}}{11.11, 22.22}`,
				invalidCase:  `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `[]<BOOL>`,
				validation:   `true`,
				validCase:    `{{.BasicType}}{true, true}`,
				invalidCase:  `{{.BasicType}}{true, false, true}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},

			// in: "[<N>]<STRING>", "[<N>]<INT>", "[<N>]<FLOAT>", "[<N>]<BOOL>"
			{
				typeClass:    `[N]<STRING>`,
				validation:   `ab cd ef`,
				validCase:    `{{.BasicType}}{"ab", "ef", "ab"}`,
				invalidCase:  `{{.BasicType}}{"ab", "gh", "ef"}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `[N]<INT>`,
				validation:   `12 34 56`,
				validCase:    `{{.BasicType}}{12, 56, 12}`,
				invalidCase:  `{{.BasicType}}{12, 78, 56}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `[N]<FLOAT>`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `{{.BasicType}}{11.11, 22.22, 11.11}`,
				invalidCase:  `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `[N]<BOOL>`,
				validation:   `true`,
				validCase:    `{{.BasicType}}{true, true, true}`,
				invalidCase:  `{{.BasicType}}{true, false, true}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},

			// in: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				typeClass:    `map[<STRING>]`,
				validation:   `a b c`,
				validCase:    `{{.BasicType}}{"a": "1", "b": "2", "c": "3"}`,
				invalidCase:  `{{.BasicType}}{"a": "1", "d": "9", "c": "3"}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `map[<INT>]`,
				validation:   `1 2 3`,
				validCase:    `{{.BasicType}}{1: 65, 2: 67, 3: 68}`,
				invalidCase:  `{{.BasicType}}{1: 65, 4: 69, 3: 68}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `map[<FLOAT>]`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `{{.BasicType}}{11.11: 11.11, 22.22: 22.22, 33.33: 33.33}`,
				invalidCase:  `{{.BasicType}}{11.11: 11.11, 44.44: 44.44, 33.33: 33.33}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
			{
				typeClass:    `map[<BOOL>]`,
				validation:   `false`,
				validCase:    `{{.BasicType}}{false: false}`,
				invalidCase:  `{{.BasicType}}{true: true, false: false}`,
				errorMessage: `{{.FieldName}} elements must be one of {{.Targets}}`,
				excludeIf:    cmpBenchTests,
			},
		},
	},

	// nin operations
	{
		tag:               "nin",
		validatorTag:      ``,
		isFieldValidation: false,
		argsCount:         common.ManyValues,
		testCases: []typeValidation{
			// nin: "[<STRING>]", "[<INT>]", "[<FLOAT>]", "[<BOOL>]"
			{
				typeClass:    `<STRING>`,
				validation:   `ab cd ef`,
				validCase:    `"fg"`,
				invalidCase:  `"cd"`,
				errorMessage: `{{.FieldName}} must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `<INT>`,
				validation:   `12 34 56`,
				validCase:    `78`,
				invalidCase:  `34`,
				errorMessage: `{{.FieldName}} must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `<FLOAT>`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `44.44`,
				invalidCase:  `22.22`,
				errorMessage: `{{.FieldName}} must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `<BOOL>`,
				validation:   `true`,
				validCase:    `false`,
				invalidCase:  `true`,
				errorMessage: `{{.FieldName}} must not be one of {{.Targets}}`,
			},

			// nin: "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>"
			{
				typeClass:    `[]<STRING>`,
				validation:   `ab cd ef`,
				validCase:    `{{.BasicType}}{"gh", "ij", "kl"}`,
				invalidCase:  `{{.BasicType}}{"ab", "ef"}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `[]<INT>`,
				validation:   `12 34 56`,
				validCase:    `{{.BasicType}}{78, 91}`,
				invalidCase:  `{{.BasicType}}{12, 78, 56}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `[]<FLOAT>`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `{{.BasicType}}{44.44, 55.55, 66.66}`,
				invalidCase:  `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `[]<BOOL>`,
				validation:   `true`,
				validCase:    `{{.BasicType}}{false, false, false}`,
				invalidCase:  `{{.BasicType}}{true, false, true}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},

			// nin: "[<N>]<STRING>", "[<N>]<INT>", "[<N>]<FLOAT>", "[<N>]<BOOL>"
			{
				typeClass:    `[N]<STRING>`,
				validation:   `ab cd ef`,
				validCase:    `{{.BasicType}}{"gh", "ij", "kl"}`,
				invalidCase:  `{{.BasicType}}{"ab", "gh", "ef"}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `[N]<INT>`,
				validation:   `12 34 56`,
				validCase:    `{{.BasicType}}{78, 91, 23}`,
				invalidCase:  `{{.BasicType}}{12, 78, 56}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `[N]<FLOAT>`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `{{.BasicType}}{44.44, 55.55, 66.66}`,
				invalidCase:  `{{.BasicType}}{11.11, 44.44, 33.33}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `[N]<BOOL>`,
				validation:   `true`,
				validCase:    `{{.BasicType}}{false, false, false}`,
				invalidCase:  `{{.BasicType}}{true, false, true}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},

			// nin: "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"
			{
				typeClass:    `map[<STRING>]`,
				validation:   `a b c`,
				validCase:    `{{.BasicType}}{"d": "1", "e": "2", "f": "3"}`,
				invalidCase:  `{{.BasicType}}{"a": "1", "d": "9", "c": "3"}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `map[<INT>]`,
				validation:   `1 2 3`,
				validCase:    `{{.BasicType}}{5: 55, 6: 66, 7: 77}`,
				invalidCase:  `{{.BasicType}}{1: 11, 4: 44, 3: 33}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `map[<FLOAT>]`,
				validation:   `11.11 22.22 33.33`,
				validCase:    `{{.BasicType}}{44.44: 44.44, 55.55: 55.55, 66.66: 66.66}`,
				invalidCase:  `{{.BasicType}}{11.11: 11.11, 44.44: 44.44, 33.33: 33.33}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
			{
				typeClass:    `map[<BOOL>]`,
				validation:   `false`,
				validCase:    `{{.BasicType}}{true: true}`,
				invalidCase:  `{{.BasicType}}{true: true, false: false}`,
				errorMessage: `{{.FieldName}} elements must not be one of {{.Targets}}`,
			},
		},
	},
}
