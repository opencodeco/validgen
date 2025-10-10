package analyzer

import "slices"

type CountValues int

const (
	UNDEFINED CountValues = iota
	ZERO_VALUE
	ONE_VALUE
	MANY_VALUES
)

type Operation struct {
	CountValues      CountValues
	IsFieldOperation bool
	ValidTypes       []string
}

var operations = map[string]Operation{
	"eq": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "<BOOL>"},
	},
	"required": {
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"gt": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"gte": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"lte": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"lt": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"min": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"max": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"eq_ignore_case": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"len": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"neq": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<BOOL>", "<INT>", "<FLOAT>"},
	},
	"neq_ignore_case": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"in": {
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "[]<STRING>", "[]<INT>", "[N]<STRING>", "[N]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"nin": {
		CountValues:      MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "[]<STRING>", "[]<INT>", "[N]<STRING>", "[N]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"email": {
		CountValues:      ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"eqfield": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<STRING>", "<INT>", "<BOOL>"},
	},
	"neqfield": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<STRING>", "<INT>", "<BOOL>"},
	},
	"gtefield": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"gtfield": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"ltefield": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"ltfield": {
		CountValues:      ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
}

func IsValidOperation(op string) bool {
	_, ok := operations[op]

	return ok
}

func IsValidTypeOperation(op, fieldType string) bool {
	operation, ok := operations[op]
	if !ok {
		return false
	}

	return slices.Contains(operation.ValidTypes, fieldType)
}

func IsFieldOperation(op string) bool {
	operation, ok := operations[op]
	if !ok {
		return false
	}

	return operation.IsFieldOperation
}

func CountValuesByOperation(op string) CountValues {
	operation, ok := operations[op]
	if !ok {
		return UNDEFINED
	}

	return operation.CountValues
}
