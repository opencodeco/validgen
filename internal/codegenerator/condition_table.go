package codegenerator

type Operation struct {
	Name            string
	ConditionByType map[string]ConditionTable
}

// ConditionTable defines the template for generating a condition check for a specific type.
type ConditionTable struct {
	// operation is the complete operation expression (e.g., "{{.Name}} == {{.Target}}").
	operation string
	// concatOperator is an optional operator used to concatenate multiple conditions (e.g., "&&", "||").
	concatOperator string
	// errorMessage is the message to display when the condition fails.
	errorMessage string
}

var operationTable = map[string]Operation{
	"eq": {
		Name: "eq",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} == "{{.Target}}"`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
			"<BOOL>": {
				operation:      `obj.{{.Name}} == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
		},
	},
	"required": {
		Name: "required",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} != ""`,
				concatOperator: "",
				errorMessage:   "{{.Name}} is required",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} != 0`,
				concatOperator: "",
				errorMessage:   "{{.Name}} is required",
			},
			"[]<STRING>": {
				operation:      `len(obj.{{.Name}}) != 0`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be empty",
			},
			"[]<INT>": {
				operation:      `len(obj.{{.Name}}) != 0`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be empty",
			},
			"map[<STRING>]": {
				operation:      `len(obj.{{.Name}}) >= 1`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be empty",
			},
			"map[<INT>]": {
				operation:      `len(obj.{{.Name}}) >= 1`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be empty",
			},
		},
	},
	"gte": {
		Name: "gte",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} >= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be >= {{.Target}}",
			},
		},
	},
	"gt": {
		Name: "gt",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} > {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be > {{.Target}}",
			},
		},
	},
	"lt": {
		Name: "lt",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} < {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be < {{.Target}}",
			},
		},
	},
	"lte": {
		Name: "lte",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} <= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be <= {{.Target}}",
			},
		},
	},
	"min": {
		Name: "min",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} length must be >= {{.Target}}",
			},
			"[]<STRING>": {
				operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
			},
			"[]<INT>": {
				operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
			},
			"map[<STRING>]": {
				operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
			},
			"map[<INT>]": {
				operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
			},
		},
	},
	"max": {
		Name: "max",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} length must be <= {{.Target}}",
			},
			"[]<STRING>": {
				operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
			},
			"[]<INT>": {
				operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
			},
			"map[<STRING>]": {
				operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
			},
			"map[<INT>]": {
				operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
			},
		},
	},
	"eq_ignore_case": {
		Name: "eq_ignore_case",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `types.EqualFold(obj.{{.Name}}, "{{.Target}}")`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
			},
		},
	},
	"len": {
		Name: "len",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `len(obj.{{.Name}}) == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} length must be {{.Target}}",
			},
			"[]<STRING>": {
				operation:      `len(obj.{{.Name}}) == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
			},
			"[]<INT>": {
				operation:      `len(obj.{{.Name}}) == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
			},
			"map[<STRING>]": {
				operation:      `len(obj.{{.Name}}) == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
			},
			"map[<INT>]": {
				operation:      `len(obj.{{.Name}}) == {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
			},
		},
	},
	"neq": {
		Name: "neq",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} != "{{.Target}}"`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
			},
			"<BOOL>": {
				operation:      `obj.{{.Name}} != {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} != {{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
		},
	},
	"neq_ignore_case": {
		Name: "neq_ignore_case",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `!types.EqualFold(obj.{{.Name}}, "{{.Target}}")`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
			},
		},
	},
	"in": {
		Name: "in",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} == "{{.Target}}"`,
				concatOperator: "||",
				errorMessage:   "{{.Name}} must be one of {{.Targets}}",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} == {{.Target}}`,
				concatOperator: "||",
				errorMessage:   "{{.Name}} must be one of {{.Targets}}",
			},
			"[]<STRING>": {
				operation:      `types.SliceOnlyContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
			"[]<INT>": {
				operation:      `types.SliceOnlyContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
			"[N]<STRING>": {
				operation:      `types.SliceOnlyContains(obj.{{.Name}}[:], {{.TargetsAsStringSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
			"[N]<INT>": {
				operation:      `types.SliceOnlyContains(obj.{{.Name}}[:], {{.TargetsAsNumericSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
			"map[<STRING>]": {
				operation:      `types.MapOnlyContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
			"map[<INT>]": {
				operation:      `types.MapOnlyContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
			},
		},
	},
	"nin": {
		Name: "nin",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} != "{{.Target}}"`,
				concatOperator: "&&",
				errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} != {{.Target}}`,
				concatOperator: "&&",
				errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
			},

			"[]<STRING>": {
				operation:      `types.SliceNotContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
			"[]<INT>": {
				operation:      `types.SliceNotContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
			"[N]<STRING>": {
				operation:      `types.SliceNotContains(obj.{{.Name}}[:], {{.TargetsAsStringSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
			"[N]<INT>": {
				operation:      `types.SliceNotContains(obj.{{.Name}}[:], {{.TargetsAsNumericSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
			"map[<STRING>]": {
				operation:      `types.MapNotContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
			"map[<INT>]": {
				operation:      `types.MapNotContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
			},
		},
	},
	"email": {
		Name: "email",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `types.IsValidEmail(obj.{{.Name}})`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be a valid email",
			},
		},
	},
	"eqfield": {
		Name: "eqfield",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} == obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} == obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
			"<BOOL>": {
				operation:      `obj.{{.Name}} == obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be equal to {{.Target}}",
			},
		},
	},
	"neqfield": {
		Name: "neqfield",
		ConditionByType: map[string]ConditionTable{
			"<STRING>": {
				operation:      `obj.{{.Name}} != obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
			"<INT>": {
				operation:      `obj.{{.Name}} != obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
			"<BOOL>": {
				operation:      `obj.{{.Name}} != obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
			},
		},
	},
	"gtefield": {
		Name: "gtefield",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} >= obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be >= {{.Target}}",
			},
		},
	},
	"gtfield": {
		Name: "gtfield",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} > obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be > {{.Target}}",
			},
		},
	},
	"ltefield": {
		Name: "ltefield",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} <= obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be <= {{.Target}}",
			},
		},
	},
	"ltfield": {
		Name: "ltfield",
		ConditionByType: map[string]ConditionTable{
			"<INT>": {
				operation:      `obj.{{.Name}} < obj.{{.Target}}`,
				concatOperator: "",
				errorMessage:   "{{.Name}} must be < {{.Target}}",
			},
		},
	},
}
