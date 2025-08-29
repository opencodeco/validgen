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
	ValidTypes       map[string]struct{}
}

var operations = map[string]Operation{
	"eq": {
		Name:             "eq",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
			"uint8":  {},
		},
	},
	"required": {
		Name:             "required",
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
			"uint8":  {},
		},
	},
	"gte": {
		Name:             "gte",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"lte": {
		Name:             "lte",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"min": {
		Name:             "min",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"max": {
		Name:             "max",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"eq_ignore_case": {
		Name:             "eq_ignore_case",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"len": {
		Name:             "len",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"neq": {
		Name:             "neq",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"neq_ignore_case": {
		Name:             "neq_ignore_case",
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"in": {
		Name:             "in",
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"nin": {
		Name:             "nin",
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"email": {
		Name:             "email",
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes: map[string]struct{}{
			"string": {},
		},
	},
	"eqfield": {
		Name:             "eqfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]struct{}{
			"string": {},
			"uint8":  {},
		},
	},
	"neqfield": {
		Name:             "neqfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]struct{}{
			"string": {},
			"uint8":  {},
		},
	},
	"gtefield": {
		Name:             "gtefield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"gtfield": {
		Name:             "gtfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"ltefield": {
		Name:             "ltefield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
	"ltfield": {
		Name:             "ltfield",
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes: map[string]struct{}{
			"uint8": {},
		},
	},
}

func countValuesOperation(op string) CountValues {
	cop, ok := operations[op]
	if !ok {
		return UNDEFINED
	}

	return cop.CountValues
}

func isFieldOperation(op string) bool {
	cop, ok := operations[op]
	if !ok {
		return false
	}

	return cop.IsFieldOperation
}

func isValidFieldOperationByType(op, fieldType string) bool {
	cop, ok := operations[op]
	if !ok {
		return false
	}

	_, ok = cop.ValidTypes[fieldType]

	return ok
}
