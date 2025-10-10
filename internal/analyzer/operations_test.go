package analyzer

import (
	"fmt"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
)

func TestValidFieldOperationsByType(t *testing.T) {
	tests := []struct {
		fieldType string
		op        string
	}{
		// eq operations
		{
			fieldType: "<STRING>",
			op:        "eq=abc",
		},
		{
			fieldType: "<INT>",
			op:        "eq=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "eq=123.45",
		},
		{
			fieldType: "<BOOL>",
			op:        "eq=true",
		},

		// required operations
		{
			fieldType: "<STRING>",
			op:        "required",
		},
		{
			fieldType: "<INT>",
			op:        "required",
		},
		{
			fieldType: "<FLOAT>",
			op:        "required",
		},
		{
			fieldType: "[]<STRING>",
			op:        "required",
		},
		{
			fieldType: "[]<INT>",
			op:        "required",
		},
		{
			fieldType: "map[<STRING>]",
			op:        "required",
		},
		{
			fieldType: "map[<INT>]",
			op:        "required",
		},

		// gt operations
		{
			fieldType: "<INT>",
			op:        "gt=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "gt=123.45",
		},

		// gte operations
		{
			fieldType: "<INT>",
			op:        "gte=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "gte=123.45",
		},

		// lt operations
		{
			fieldType: "<INT>",
			op:        "lt=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "lt=123.45",
		},

		// lte operations
		{
			fieldType: "<INT>",
			op:        "lte=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "lte=123.45",
		},

		// min operations
		{
			fieldType: "<STRING>",
			op:        "min=3",
		},
		{
			fieldType: "[]<STRING>",
			op:        "min=3",
		},
		{
			fieldType: "map[<STRING>]",
			op:        "min=3",
		},
		{
			fieldType: "[]<INT>",
			op:        "min=3",
		},
		{
			fieldType: "map[<INT>]",
			op:        "min=3",
		},

		// max operations
		{
			fieldType: "<STRING>",
			op:        "max=3",
		},
		{
			fieldType: "[]<STRING>",
			op:        "max=3",
		},
		{
			fieldType: "map[<STRING>]",
			op:        "max=3",
		},
		{
			fieldType: "[]<INT>",
			op:        "max=3",
		},
		{
			fieldType: "map[<INT>]",
			op:        "max=3",
		},

		// eq_ignore_case operations
		{
			fieldType: "<STRING>",
			op:        "eq_ignore_case=abc",
		},

		// len operations
		{
			fieldType: "<STRING>",
			op:        "len=3",
		},
		{
			fieldType: "[]<STRING>",
			op:        "len=3",
		},
		{
			fieldType: "map[<STRING>]",
			op:        "len=3",
		},
		{
			fieldType: "[]<INT>",
			op:        "len=3",
		},
		{
			fieldType: "map[<INT>]",
			op:        "len=3",
		},

		// neq operations
		{
			fieldType: "<STRING>",
			op:        "neq=abc",
		},
		{
			fieldType: "<INT>",
			op:        "neq=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "neq=123.45",
		},
		{
			fieldType: "<BOOL>",
			op:        "neq=true",
		},

		// neq_ignore_case operations
		{
			fieldType: "<STRING>",
			op:        "neq_ignore_case=abc",
		},

		// in operations
		{
			fieldType: "<STRING>",
			op:        "in=abc",
		},
		{
			fieldType: "<INT>",
			op:        "in=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "in=123.45",
		},
		{
			fieldType: "[]<STRING>",
			op:        "in=abc",
		},
		{
			fieldType: "[5]<STRING>",
			op:        "in=abc",
		},
		{
			fieldType: "map[<STRING>]",
			op:        "in=abc",
		},
		{
			fieldType: "[]<INT>",
			op:        "in=123",
		},
		{
			fieldType: "[5]<INT>",
			op:        "in=123",
		},
		{
			fieldType: "map[<INT>]",
			op:        "in=123",
		},

		// nin operations
		{
			fieldType: "<STRING>",
			op:        "nin=abc",
		},
		{
			fieldType: "<INT>",
			op:        "nin=123",
		},
		{
			fieldType: "<FLOAT>",
			op:        "nin=123.45",
		},
		{
			fieldType: "[]<STRING>",
			op:        "nin=abc",
		},
		{
			fieldType: "[5]<STRING>",
			op:        "nin=abc",
		},
		{
			fieldType: "map[<STRING>]",
			op:        "nin=abc",
		},
		{
			fieldType: "[]<INT>",
			op:        "nin=123",
		},
		{
			fieldType: "[5]<INT>",
			op:        "nin=123",
		},
		{
			fieldType: "map[<INT>]",
			op:        "nin=123",
		},

		// email operations
		{
			fieldType: "<STRING>",
			op:        "email",
		},
	}

	for _, tt := range tests {
		fieldTypes, _ := common.HelperFromNormalizedToFieldTypes(tt.fieldType)
		for _, fieldType := range fieldTypes {
			testName := fmt.Sprintf("valid %s %s", fieldType.ToType(), tt.op)
			t.Run(testName, func(t *testing.T) {
				arg := []*parser.Struct{
					{
						Fields: []parser.Field{
							{
								FieldName: "Field",
								Type:      fieldType,
								Tag:       fmt.Sprintf(`valid:"%s"`, tt.op),
							},
						},
					},
				}

				_, err := AnalyzeStructs(arg)
				if err != nil {
					t.Errorf("AnalyzeStructs() error = %v, wantErr %v", err, nil)
					return
				}
			})
		}
	}
}
