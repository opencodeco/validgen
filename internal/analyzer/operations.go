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
			"string": true,
			"uint8":  true,
		},
	},
	"required": {
		Name:             "required",
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string":   true,
			"uint8":    true,
			"[]string": true,
		},
	},
	"gte": {
		Name:             "gte",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"uint8": true,
		},
	},
	"lte": {
		Name:             "lte",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"uint8": true,
		},
	},
	"min": {
		Name:             "min",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string":   true,
			"[]string": true,
		},
	},
	"max": {
		Name:             "max",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string":   true,
			"[]string": true,
		},
	},
	"eq_ignore_case": {
		Name:             "eq_ignore_case",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string": true,
		},
	},
	"len": {
		Name:             "len",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string":   true,
			"[]string": true,
		},
	},
	"neq": {
		Name:             "neq",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string": true,
		},
	},
	"neq_ignore_case": {
		Name:             "neq_ignore_case",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string": true,
		},
	},
	"in": {
		Name:             "in",
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string":    true,
			"[]string":  true,
			"[N]string": true,
		},
	},
	"nin": {
		Name:             "nin",
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string":    true,
			"[]string":  true,
			"[N]string": true,
		},
	},
	"email": {
		Name:             "email",
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]bool{
			"string": true,
		},
	},
	"eqfield": {
		Name:             "eqfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"string": true,
			"uint8":  true,
		},
	},
	"neqfield": {
		Name:             "neqfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"string": true,
			"uint8":  true,
		},
	},
	"gtefield": {
		Name:             "gtefield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"uint8": true,
		},
	},
	"gtfield": {
		Name:             "gtfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"uint8": true,
		},
	},
	"ltefield": {
		Name:             "ltefield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"uint8": true,
		},
	},
	"ltfield": {
		Name:             "ltfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]bool{
			"uint8": true,
		},
	},
}
