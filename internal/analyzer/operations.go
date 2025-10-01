package analyzer

type CountValues int

const (
	UNDEFINED CountValues = iota
	ZERO_VALUE
	ONE_VALUE
	MANY_VALUES
)

type Operation struct {
	Name             string
	CountValues      CountValues
	IsFieldOperation bool
	ValidTypes       map[string]bool
}

var operations = map[string]Operation{
	"eq": {
		Name:             "eq",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>": true,
			"<INT>":    true,
			"<BOOL>":   true,
		},
	},
	"required": {
		Name:             "required",
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>":      true,
			"<INT>":         true,
			"[]<STRING>":    true,
			"map[<STRING>]": true,
			"map[<INT>]":    true,
		},
	},
	"gt": {
		Name:             "gt",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"gte": {
		Name:             "gte",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"lte": {
		Name:             "lte",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"lt": {
		Name:             "lt",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"min": {
		Name:             "min",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>":      true,
			"[]<STRING>":    true,
			"map[<STRING>]": true,
			"map[<INT>]":    true,
		},
	},
	"max": {
		Name:             "max",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>":      true,
			"[]<STRING>":    true,
			"map[<STRING>]": true,
			"map[<INT>]":    true,
		},
	},
	"eq_ignore_case": {
		Name:             "eq_ignore_case",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>": true,
		},
	},
	"len": {
		Name:             "len",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>":      true,
			"[]<STRING>":    true,
			"map[<STRING>]": true,
			"map[<INT>]":    true,
		},
	},
	"neq": {
		Name:             "neq",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>": true,
			"<BOOL>":   true,
			"<INT>":    true,
		},
	},
	"neq_ignore_case": {
		Name:             "neq_ignore_case",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>": true,
		},
	},
	"in": {
		Name:             "in",
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>":      true,
			"<INT>":         true,
			"[]<STRING>":    true,
			"[N]<STRING>":   true,
			"map[<STRING>]": true,
			"map[<INT>]":    true,
		},
	},
	"nin": {
		Name:             "nin",
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>":      true,
			"<INT>":         true,
			"[]<STRING>":    true,
			"[N]<STRING>":   true,
			"map[<STRING>]": true,
			"map[<INT>]":    true,
		},
	},
	"email": {
		Name:             "email",
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"<STRING>": true,
		},
	},
	"eqfield": {
		Name:             "eqfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"<STRING>": true,
			"<INT>":    true,
			"<BOOL>":   true,
		},
	},
	"neqfield": {
		Name:             "neqfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"<STRING>": true,
			"<INT>":    true,
			"<BOOL>":   true,
		},
	},
	"gtefield": {
		Name:             "gtefield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"gtfield": {
		Name:             "gtfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"ltefield": {
		Name:             "ltefield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
	"ltfield": {
		Name:             "ltfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"<INT>": true,
		},
	},
}
