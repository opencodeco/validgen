package codegenerator

import (
	"slices"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/types"
)

type Operation struct {
	ConditionByTypes []ConditionByType
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

type ConditionByType struct {
	AcceptedTypes []string
	ConditionTable
}

var conditionTable = map[string]Operation{
	"eq": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} == "{{.Target}}"`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
				},
			},
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>", "<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} == {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} == "{{.Target}}"`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>", "*<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} == {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to {{.Target}}",
				},
			},
		},
	},
	"required": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != ""`,
					concatOperator: "",
					errorMessage:   "{{.Name}} is required",
				},
			},
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != 0`,
					concatOperator: "",
					errorMessage:   "{{.Name}} is required",
				},
			},
			{
				AcceptedTypes: []string{"<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != false`,
					concatOperator: "",
					errorMessage:   "{{.Name}} is required",
				},
			},
			{
				AcceptedTypes: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) != 0`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be empty",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} != ""`,
					concatOperator: "",
					errorMessage:   "{{.Name}} is required",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} != 0`,
					concatOperator: "",
					errorMessage:   "{{.Name}} is required",
				},
			},
			{
				AcceptedTypes: []string{"*<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} != false`,
					concatOperator: "",
					errorMessage:   "{{.Name}} is required",
				},
			},
			{
				AcceptedTypes: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) != 0`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be empty",
				},
			},
			{
				AcceptedTypes: []string{"*[N]<STRING>", "*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be empty",
				},
			},
		},
	},
	"gte": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} >= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be >= {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} >= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be >= {{.Target}}",
				},
			},
		},
	},
	"gt": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} > {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be > {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} > {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be > {{.Target}}",
				},
			},
		},
	},
	"lte": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} <= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be <= {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} <= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be <= {{.Target}}",
				},
			},
		},
	},
	"lt": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} < {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be < {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} < {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be < {{.Target}}",
				},
			},
		},
	},
	"min": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} length must be >= {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) >= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) >= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} length must be >= {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) >= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must have at least {{.Target}} elements",
				},
			},
		},
	},
	"max": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} length must be <= {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) <= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} length must be <= {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) <= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
				},
			},
			{
				AcceptedTypes: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) <= {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must have at most {{.Target}} elements",
				},
			},
		},
	},
	"eq_ignore_case": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `types.EqualFold(obj.{{.Name}}, "{{.Target}}")`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.EqualFold(*obj.{{.Name}}, "{{.Target}}")`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to '{{.Target}}'",
				},
			},
		},
	},
	"len": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) == {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} length must be {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `len(obj.{{.Name}}) == {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) == {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} length must be {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && len(*obj.{{.Name}}) == {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must have exactly {{.Target}} elements",
				},
			},
		},
	},
	"neq": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != "{{.Target}}"`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
				},
			},
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>", "<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} != "{{.Target}}"`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>", "*<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && *obj.{{.Name}} != {{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
				},
			},
		},
	},
	"neq_ignore_case": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `!types.EqualFold(obj.{{.Name}}, "{{.Target}}")`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && !types.EqualFold(*obj.{{.Name}}, "{{.Target}}")`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to '{{.Target}}'",
				},
			},
		},
	},
	"in": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} == "{{.Target}}"`,
					concatOperator: "||",
					errorMessage:   "{{.Name}} must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>", "<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} == {{.Target}}`,
					concatOperator: "||",
					errorMessage:   "{{.Name}} must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceOnlyContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<INT>", "[]<FLOAT>", "[]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceOnlyContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[N]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceOnlyContains(obj.{{.Name}}[:], {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceOnlyContains(obj.{{.Name}}[:], {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"map[<STRING>]"},
				ConditionTable: ConditionTable{
					operation:      `types.MapOnlyContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `types.MapOnlyContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `(obj.{{.Name}} != nil && *obj.{{.Name}} == "{{.Target}}")`,
					concatOperator: "||",
					errorMessage:   "{{.Name}} must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>", "*<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `(obj.{{.Name}} != nil && *obj.{{.Name}} == {{.Target}})`,
					concatOperator: "||",
					errorMessage:   "{{.Name}} must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceOnlyContains(*obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceOnlyContains(*obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[N]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceOnlyContains(obj.{{.Name}}[:], {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceOnlyContains(obj.{{.Name}}[:], {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*map[<STRING>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.MapOnlyContains(*obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.MapOnlyContains(*obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must be one of {{.Targets}}",
				},
			},
		},
	},
	"nin": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != "{{.Target}}"`,
					concatOperator: "&&",
					errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>", "<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != {{.Target}}`,
					concatOperator: "&&",
					errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceNotContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[]<INT>", "[]<FLOAT>", "[]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceNotContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[N]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceNotContains(obj.{{.Name}}[:], {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `types.SliceNotContains(obj.{{.Name}}[:], {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"map[<STRING>]"},
				ConditionTable: ConditionTable{
					operation:      `types.MapNotContains(obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `types.MapNotContains(obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `(obj.{{.Name}} != nil && *obj.{{.Name}} != "{{.Target}}")`,
					concatOperator: "&&",
					errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*<INT>", "*<FLOAT>", "*<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `(obj.{{.Name}} != nil && *obj.{{.Name}} != {{.Target}})`,
					concatOperator: "&&",
					errorMessage:   "{{.Name}} must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceNotContains(*obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceNotContains(*obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[N]<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceNotContains(obj.{{.Name}}[:], {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.SliceNotContains(obj.{{.Name}}[:], {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*map[<STRING>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.MapNotContains(*obj.{{.Name}}, {{.TargetsAsStringSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
			{
				AcceptedTypes: []string{"*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.MapNotContains(*obj.{{.Name}}, {{.TargetsAsNumericSlice}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} elements must not be one of {{.Targets}}",
				},
			},
		},
	},
	"email": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `types.IsValidEmail(obj.{{.Name}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be a valid email",
				},
			},
			{
				AcceptedTypes: []string{"*<STRING>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != nil && types.IsValidEmail(*obj.{{.Name}})`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be a valid email",
				},
			},
		},
	},
	"eqfield": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>", "<INT>", "<FLOAT>", "<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} == obj.{{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be equal to {{.Target}}",
				},
			},
		},
	},
	"neqfield": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<STRING>", "<INT>", "<FLOAT>", "<BOOL>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} != obj.{{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must not be equal to {{.Target}}",
				},
			},
		},
	},
	"gtefield": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} >= obj.{{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be >= {{.Target}}",
				},
			},
		},
	},
	"gtfield": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} > obj.{{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be > {{.Target}}",
				},
			},
		},
	},
	"ltefield": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} <= obj.{{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be <= {{.Target}}",
				},
			},
		},
	},
	"ltfield": {
		ConditionByTypes: []ConditionByType{
			{
				AcceptedTypes: []string{"<INT>", "<FLOAT>"},
				ConditionTable: ConditionTable{
					operation:      `obj.{{.Name}} < obj.{{.Target}}`,
					concatOperator: "",
					errorMessage:   "{{.Name}} must be < {{.Target}}",
				},
			},
		},
	},
}

func GetConditionTable(operation string, fieldType common.FieldType) (ConditionTable, error) {
	op, ok := conditionTable[operation]
	if !ok {
		return ConditionTable{}, types.NewValidationError("INTERNAL ERROR: unsupported operation %s", operation)
	}

	normalizedType := fieldType.ToNormalizedString()
	for _, conditionByType := range op.ConditionByTypes {
		if slices.Contains(conditionByType.AcceptedTypes, normalizedType) {
			return conditionByType.ConditionTable, nil
		}
	}

	return ConditionTable{}, types.NewValidationError("INTERNAL ERROR: unsupported operation %s type %s (%s)", operation, normalizedType, fieldType.BaseType)
}
