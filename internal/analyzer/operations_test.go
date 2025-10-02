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
			fieldType: "string",
			op:        "eq=abc",
		},
		{
			fieldType: "uint8",
			op:        "eq=123",
		},
		{
			fieldType: "bool",
			op:        "eq=true",
		},

		// required operations
		{
			fieldType: "string",
			op:        "required",
		},
		{
			fieldType: "uint8",
			op:        "required",
		},
		{
			fieldType: "[]string",
			op:        "required",
		},
		{
			fieldType: "map[string]",
			op:        "required",
		},
		{
			fieldType: "map[uint8]",
			op:        "required",
		},

		// gte operations
		{
			fieldType: "uint8",
			op:        "gte=123",
		},

		// lte operations
		{
			fieldType: "uint8",
			op:        "lte=123",
		},

		// min operations
		{
			fieldType: "string",
			op:        "min=3",
		},
		{
			fieldType: "[]string",
			op:        "min=3",
		},
		{
			fieldType: "map[string]",
			op:        "min=3",
		},
		{
			fieldType: "map[uint8]",
			op:        "min=3",
		},

		// max operations
		{
			fieldType: "string",
			op:        "max=3",
		},
		{
			fieldType: "[]string",
			op:        "max=3",
		},
		{
			fieldType: "map[string]",
			op:        "max=3",
		},
		{
			fieldType: "map[uint8]",
			op:        "max=3",
		},

		// eq_ignore_case operations
		{
			fieldType: "string",
			op:        "eq_ignore_case=abc",
		},

		// len operations
		{
			fieldType: "string",
			op:        "len=3",
		},
		{
			fieldType: "[]string",
			op:        "len=3",
		},
		{
			fieldType: "map[string]",
			op:        "len=3",
		},
		{
			fieldType: "map[uint8]",
			op:        "len=3",
		},

		// neq operations
		{
			fieldType: "string",
			op:        "neq=abc",
		},
		{
			fieldType: "bool",
			op:        "neq=true",
		},

		// neq_ignore_case operations
		{
			fieldType: "string",
			op:        "neq_ignore_case=abc",
		},

		// in operations
		{
			fieldType: "string",
			op:        "in=abc",
		},
		{
			fieldType: "[]string",
			op:        "in=abc",
		},
		{
			fieldType: "[5]string",
			op:        "in=abc",
		},
		{
			fieldType: "map[string]",
			op:        "in=abc",
		},
		{
			fieldType: "map[uint8]",
			op:        "in=123",
		},

		// nin operations
		{
			fieldType: "string",
			op:        "nin=abc",
		},
		{
			fieldType: "[]string",
			op:        "nin=abc",
		},
		{
			fieldType: "[5]string",
			op:        "nin=abc",
		},
		{
			fieldType: "map[string]",
			op:        "nin=abc",
		},
		{
			fieldType: "map[uint8]",
			op:        "nin=123",
		},

		// email operations
		{
			fieldType: "string",
			op:        "email",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("valid %s %s", tt.fieldType, tt.op)
		t.Run(testName, func(t *testing.T) {
			arg := []*parser.Struct{
				{
					Fields: []parser.Field{
						{
							FieldName: "Field",
							Type:      common.FieldType{BaseType: tt.fieldType},
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
