package codegenerator

type Operation struct {
	Name            string
	ConditionByType map[string]ConditionTable
}

// ConditionTable defines the template for generating a condition check for a specific type.
type ConditionTable struct {
	// loperand is the left-hand side operand in the condition expression (e.g., the variable name).
	loperand string
	// operator is the comparison operator used in the condition (e.g., "==", "!=", ">=").
	operator string
	// roperand is the right-hand side operand in the condition expression (e.g., the target value).
	roperand string
	// concatOperator is an optional operator used to concatenate multiple conditions (e.g., "&&", "||").
	concatOperator string
	// errorMessage is the message to display when the condition fails.
	errorMessage string
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
				errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `{{.Target}}`,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} is required",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `0`,
				concatOperator: "",
				errorMessage:   "{{.Name}} is required",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       "!=",
				roperand:       `0`,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} length must be >= {{.Target}}",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       ">=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} length must be <= {{.Target}}",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       "<=",
				roperand:       `{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
			},
		},
	},
	"eq_ignore_case": {
		Name: "eq_ignore_case",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       `types.EqualFold({{.Name}},"{{.Target}}")`,
				operator:       "",
				roperand:       ``,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} length must be {{.Target}}",
			},
			"[]string": {
				loperand:       "len({{.Name}})",
				operator:       "==",
				roperand:       `{{.Target}}`,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
			},
		},
	},
	"neq_ignore_case": {
		Name: "neq_ignore_case",
		ConditionByType: map[string]ConditionTable{
			"string": {
				loperand:       `!types.EqualFold({{.Name}}, "{{.Target}}")`,
				operator:       "",
				roperand:       "",
				concatOperator: "",
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
				errorMessage:   "{{.Name}} must be one of {{.Targets}}",
			},
			"[]string": {
				loperand:       "",
				operator:       "",
				roperand:       `types.SlicesContains({{.Name}}, "{{.Target}}")`,
				concatOperator: "||",
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
				errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
			},
			"[]string": {
				loperand:       "",
				operator:       "",
				roperand:       `!types.SlicesContains({{.Name}}, "{{.Target}}")`,
				concatOperator: "&&",
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
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "==",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
			"uint8": {
				loperand:       "{{.Name}}",
				operator:       "!=",
				roperand:       `obj.{{.Target}}`,
				concatOperator: "",
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
				errorMessage:   "{{.Name}} must be < {{.Target}}",
			},
		},
	},
}
