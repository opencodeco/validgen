package operations

import (
	"fmt"
	"testing"
)

func TestIsValidTypeOperation(t *testing.T) {
	tests := []struct {
		fieldType string
		op        string
		valid     bool
	}{
		// eq operations
		{
			fieldType: "<STRING>",
			op:        "eq",
			valid:     true,
		},
		{
			fieldType: "<INT>",
			op:        "eq",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "eq",
			valid:     true,
		},
		{
			fieldType: "<BOOL>",
			op:        "eq",
			valid:     true,
		},

		// required operations
		{
			fieldType: "<STRING>",
			op:        "required",
			valid:     true,
		},
		{
			fieldType: "<INT>",
			op:        "required",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "required",
			valid:     true,
		},
		{
			fieldType: "[]<STRING>",
			op:        "required",
			valid:     true,
		},
		{
			fieldType: "[]<INT>",
			op:        "required",
			valid:     true,
		},
		{
			fieldType: "map[<STRING>]",
			op:        "required",
			valid:     true,
		},
		{
			fieldType: "map[<INT>]",
			op:        "required",
			valid:     true,
		},

		// gt operations
		{
			fieldType: "<INT>",
			op:        "gt",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "gt",
			valid:     true,
		},

		// gte operations
		{
			fieldType: "<INT>",
			op:        "gte",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "gte",
			valid:     true,
		},

		// lt operations
		{
			fieldType: "<INT>",
			op:        "lt",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "lt",
			valid:     true,
		},

		// lte operations
		{
			fieldType: "<INT>",
			op:        "lte",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "lte",
			valid:     true,
		},

		// min operations
		{
			fieldType: "<STRING>",
			op:        "min",
			valid:     true,
		},
		{
			fieldType: "[]<STRING>",
			op:        "min",
			valid:     true,
		},
		{
			fieldType: "map[<STRING>]",
			op:        "min",
			valid:     true,
		},
		{
			fieldType: "[]<INT>",
			op:        "min",
			valid:     true,
		},
		{
			fieldType: "map[<INT>]",
			op:        "min",
			valid:     true,
		},

		// max operations
		{
			fieldType: "<STRING>",
			op:        "max",
			valid:     true,
		},
		{
			fieldType: "[]<STRING>",
			op:        "max",
			valid:     true,
		},
		{
			fieldType: "map[<STRING>]",
			op:        "max",
			valid:     true,
		},
		{
			fieldType: "[]<INT>",
			op:        "max",
			valid:     true,
		},
		{
			fieldType: "map[<INT>]",
			op:        "max",
			valid:     true,
		},

		// eq_ignore_case operations
		{
			fieldType: "<STRING>",
			op:        "eq_ignore_case",
			valid:     true,
		},

		// len operations
		{
			fieldType: "<STRING>",
			op:        "len",
			valid:     true,
		},
		{
			fieldType: "[]<STRING>",
			op:        "len",
			valid:     true,
		},
		{
			fieldType: "map[<STRING>]",
			op:        "len",
			valid:     true,
		},
		{
			fieldType: "[]<INT>",
			op:        "len",
			valid:     true,
		},
		{
			fieldType: "map[<INT>]",
			op:        "len",
			valid:     true,
		},

		// neq operations
		{
			fieldType: "<STRING>",
			op:        "neq",
			valid:     true,
		},
		{
			fieldType: "<INT>",
			op:        "neq",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "neq",
			valid:     true,
		},
		{
			fieldType: "<BOOL>",
			op:        "neq",
			valid:     true,
		},

		// neq_ignore_case operations
		{
			fieldType: "<STRING>",
			op:        "neq_ignore_case",
			valid:     true,
		},

		// in operations
		{
			fieldType: "<STRING>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "<INT>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "[]<STRING>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "[N]<STRING>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "map[<STRING>]",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "[]<INT>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "[N]<INT>",
			op:        "in",
			valid:     true,
		},
		{
			fieldType: "map[<INT>]",
			op:        "in",
			valid:     true,
		},

		// nin operations
		{
			fieldType: "<STRING>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "<INT>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "<FLOAT>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "[]<STRING>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "[N]<STRING>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "map[<STRING>]",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "[]<INT>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "[N]<INT>",
			op:        "nin",
			valid:     true,
		},
		{
			fieldType: "map[<INT>]",
			op:        "nin",
			valid:     true,
		},

		// email operations
		{
			fieldType: "<STRING>",
			op:        "email",
			valid:     true,
		},

		// invalid cases
		{
			fieldType: "<BOOL>",
			op:        "email",
			valid:     false,
		},
		{
			fieldType: "<BOOL>",
			op:        "gt",
			valid:     false,
		},
		{
			fieldType: "<INT>",
			op:        "neq_ignore_case",
			valid:     false,
		},
	}

	ops := New()
	for _, tt := range tests {
		testType := ""
		if tt.valid {
			testType = "valid"
		} else {
			testType = "invalid"
		}

		testName := fmt.Sprintf("%s %s %s", testType, tt.fieldType, tt.op)

		t.Run(testName, func(t *testing.T) {
			valid := ops.IsValidType(tt.op, tt.fieldType)
			if valid != tt.valid {
				t.Errorf("IsValidTypeOperation() = %v, want %v", valid, tt.valid)
				return
			}
		})
	}
}
