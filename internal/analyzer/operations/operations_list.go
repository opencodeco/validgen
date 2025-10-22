package operations

import "github.com/opencodeco/validgen/internal/common"

var operationsList = map[string]Operation{
	"eq": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<INT>", "<FLOAT>", "<BOOL>"},
	},
	"required": {
		CountValues:      common.ZeroValue,
		IsFieldOperation: false,
		ValidTypes: []string{
			"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
			"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
			"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
	},
	"gt": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"gte": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"lte": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"lt": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<INT>", "<FLOAT>"},
	},
	"min": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes: []string{
			"<STRING>", "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
			"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
		},
	},
	"max": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes: []string{
			"<STRING>", "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
			"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
		},
	},
	"eq_ignore_case": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"len": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes: []string{
			"<STRING>", "[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
			"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
		},
	},
	"neq": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>", "<BOOL>", "<INT>", "<FLOAT>"},
	},
	"neq_ignore_case": {
		CountValues:      common.OneValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"in": {
		CountValues:      common.ManyValues,
		IsFieldOperation: false,
		ValidTypes: []string{
			"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
			"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
			"[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>",
			"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
		},
	},
	"nin": {
		CountValues:      common.ManyValues,
		IsFieldOperation: false,
		ValidTypes: []string{
			"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
			"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
			"[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>",
			"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
		},
	},
	"email": {
		CountValues:      common.ZeroValue,
		IsFieldOperation: false,
		ValidTypes:       []string{"<STRING>"},
	},
	"eqfield": {
		CountValues:      common.OneValue,
		IsFieldOperation: true,
		ValidTypes:       []string{"<STRING>", "<INT>", "<BOOL>"},
	},
	"neqfield": {
		CountValues:      common.OneValue,
		IsFieldOperation: true,
		ValidTypes:       []string{"<STRING>", "<INT>", "<BOOL>"},
	},
	"gtefield": {
		CountValues:      common.OneValue,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"gtfield": {
		CountValues:      common.OneValue,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"ltefield": {
		CountValues:      common.OneValue,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
	"ltfield": {
		CountValues:      common.OneValue,
		IsFieldOperation: true,
		ValidTypes:       []string{"<INT>"},
	},
}
