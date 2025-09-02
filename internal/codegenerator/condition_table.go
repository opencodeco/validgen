package codegenerator

import "strings"

type Operation struct {
	Name            string
	ConditionByType map[string]ConditionTable
}

type ConditionTable struct {
	loperand       string
	operator       string
	roperand       string
	concatOperator string
	normalizeFunc  func(string) string
	errorMessage   string
}

var operationTable = map[string]Operation{
	"eq": {
		Name: "eq",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `"{{.Target}}"`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
		},
	},
	"required": {
		Name: "required",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `""`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} is required",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `0`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} is required",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       "!=",
				roperand:       `0`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must not be empty",
			},
		},
	},
	"gte": {
		Name: "gte",
		ConditionByType: map[string]ConditionTable{
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       ">=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be >= {{.Target}}",
			},
		},
	},
	"lte": {
		Name: "lte",
		ConditionByType: map[string]ConditionTable{
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "<=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be <= {{.Target}}",
			},
		},
	},
	"min": {
		Name: "min",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "len({{.Name}})",
				operator:       ">=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} length must be >= {{.Target}}",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       ">=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
			},
		},
	},
	"max": {
		Name: "max",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "len({{.Name}})",
				operator:       "<=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} length must be <= {{.Target}}",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       "<=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
			},
		},
	},
	"eq_ignore_case": {
		Name: "eq_ignore_case",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "types.ToLower({{.Name}})",
				operator:       "==",
				roperand:       `"{{.Target}}"`,
				concatOperator: "",
				normalizeFunc:  strings.ToLower,
				errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
			},
		},
	},
	"len": {
		Name: "len",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "len({{.Name}})",
				operator:       "==",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} length must be {{.Target}}",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       "==",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
			},
		},
	},
	"neq": {
		Name: "neq",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `"{{.Target}}"`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
			},
		},
	},
	"neq_ignore_case": {
		Name: "neq_ignore_case",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "types.ToLower({{.Name}})",
				operator:       "!=",
				roperand:       `"{{.Target}}"`,
				concatOperator: "",
				normalizeFunc:  strings.ToLower,
				errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
			},
		},
	},
	"in": {
		Name: "in",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `"{{.Target}}"`,
				concatOperator: "||",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be one of {{.Targets}}",
			},
			"[]string": {
				loperand:       "",
				operator:       "",
				roperand:       `types.SlicesContains({{.Name}}, "{{.Target}}")`,
				concatOperator: "||",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
		},
	},
	"nin": {
		Name: "nin",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `"{{.Target}}"`,
				concatOperator: "&&",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
			},
			"[]string": {
				loperand:       "",
				operator:       "",
				roperand:       `!types.SlicesContains({{.Name}}, "{{.Target}}")`,
				concatOperator: "&&",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
		},
	},
	"email": {
		Name: "email",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "types.IsValidEmail({{.Name}})",
				operator:       "==",
				roperand:       `true`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be a valid email",
			},
		},
	},
	"eqfield": {
		Name: "eqfield",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
		},
	},
	"neqfield": {
		Name: "neqfield",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
		},
	},
	"gtefield": {
		Name: "gtefield",
		ConditionByType: map[string]ConditionTable{
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       ">=",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be >= {{.Target}}",
			},
		},
	},
	"gtfield": {
		Name: "gtfield",
		ConditionByType: map[string]ConditionTable{
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       ">",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be > {{.Target}}",
			},
		},
	},
	"ltefield": {
		Name: "ltefield",
		ConditionByType: map[string]ConditionTable{
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "<=",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be <= {{.Target}}",
			},
		},
	},
	"ltfield": {
		Name: "ltfield",
		ConditionByType: map[string]ConditionTable{
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "<",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
				normalizeFunc:  nil,
				errorMessage:   "{{.Name}} must be < {{.Target}}",
			},
		},
	},
}
