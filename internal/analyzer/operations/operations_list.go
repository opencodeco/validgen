package operations

import "github.com/opencodeco/validgen/internal/common"

var operationsList = map[string]Operation{
	"eq": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "<BOOL>"},
	},
	"required": {
		CountValues:      common.ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"gt": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"gte": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"lte": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"lt": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"min": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"max": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"eq_ignore_case": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"len": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "[]<STRING>", "[]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"neq": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<BOOL>", "<INT>", "<FLOAT>"},
	},
	"neq_ignore_case": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"in": {
		CountValues:      common.MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "[]<STRING>", "[]<INT>", "[N]<STRING>", "[N]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"nin": {
		CountValues:      common.MANY_VALUES,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "[]<STRING>", "[]<INT>", "[N]<STRING>", "[N]<INT>", "map[<STRING>]", "map[<INT>]"},
	},
	"email": {
		CountValues:      common.ZERO_VALUE,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"eqfield": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<STRING>", "<INT>", "<BOOL>"},
	},
	"neqfield": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<STRING>", "<INT>", "<BOOL>"},
	},
	"gtefield": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"gtfield": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"ltefield": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"ltfield": {
		CountValues:      common.ONE_VALUE,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
}
