package operations

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestOperationsIsValid(t *testing.T) {
	tests := []struct {
		op   string
		want bool
	}{
		{op: "eq", want: true},
		{op: "required", want: true},
		{op: "gt", want: true},
		{op: "gte", want: true},
		{op: "lte", want: true},
		{op: "lt", want: true},
		{op: "min", want: true},
		{op: "max", want: true},
		{op: "eq_ignore_case", want: true},
		{op: "len", want: true},
		{op: "neq", want: true},
		{op: "neq_ignore_case", want: true},
		{op: "in", want: true},
		{op: "nin", want: true},
		{op: "email", want: true},
		{op: "eqfield", want: true},
		{op: "neqfield", want: true},
		{op: "gtefield", want: true},
		{op: "gtfield", want: true},
		{op: "ltefield", want: true},
		{op: "ltfield", want: true},
		{op: "invalid_op", want: false},
	}

	ops := New()
	for _, tt := range tests {
		testName := fmt.Sprintf("%s operation", tt.op)
		t.Run(testName, func(t *testing.T) {
			if got := ops.IsValid(tt.op); got != tt.want {
				t.Errorf("Operations.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationsIsValidByType(t *testing.T) {
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
		{
			fieldType: "<XPTO>",
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
			valid := ops.IsValidByType(tt.op, tt.fieldType)
			if valid != tt.valid {
				t.Errorf("IsValidTypeOperation() = %v, want %v", valid, tt.valid)
				return
			}
		})
	}
}

func TestOperationsIsFieldOperation(t *testing.T) {
	tests := []struct {
		op   string
		want bool
	}{
		{op: "eq", want: false},
		{op: "required", want: false},
		{op: "gt", want: false},
		{op: "gte", want: false},
		{op: "lte", want: false},
		{op: "lt", want: false},
		{op: "min", want: false},
		{op: "max", want: false},
		{op: "eq_ignore_case", want: false},
		{op: "len", want: false},
		{op: "neq", want: false},
		{op: "neq_ignore_case", want: false},
		{op: "in", want: false},
		{op: "nin", want: false},
		{op: "email", want: false},
		{op: "eqfield", want: true},
		{op: "neqfield", want: true},
		{op: "gtefield", want: true},
		{op: "gtfield", want: true},
		{op: "ltefield", want: true},
		{op: "ltfield", want: true},
		{op: "invalid_op", want: false},
	}

	ops := New()
	for _, tt := range tests {
		testName := fmt.Sprintf("%s operation", tt.op)
		t.Run(testName, func(t *testing.T) {
			if got := ops.IsFieldOperation(tt.op); got != tt.want {
				t.Errorf("Operations.IsFieldOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationsArgsCount(t *testing.T) {
	tests := []struct {
		op   string
		want common.CountValues
	}{
		{op: "eq", want: common.OneValue},
		{op: "required", want: common.ZeroValue},
		{op: "gt", want: common.OneValue},
		{op: "gte", want: common.OneValue},
		{op: "lte", want: common.OneValue},
		{op: "lt", want: common.OneValue},
		{op: "min", want: common.OneValue},
		{op: "max", want: common.OneValue},
		{op: "eq_ignore_case", want: common.OneValue},
		{op: "len", want: common.OneValue},
		{op: "neq", want: common.OneValue},
		{op: "neq_ignore_case", want: common.OneValue},
		{op: "in", want: common.ManyValues},
		{op: "nin", want: common.ManyValues},
		{op: "email", want: common.ZeroValue},
		{op: "eqfield", want: common.OneValue},
		{op: "neqfield", want: common.OneValue},
		{op: "gtefield", want: common.OneValue},
		{op: "gtfield", want: common.OneValue},
		{op: "ltefield", want: common.OneValue},
		{op: "ltfield", want: common.OneValue},
		{op: "invalid_op", want: common.UndefinedValue},
	}

	ops := New()
	for _, tt := range tests {
		testName := fmt.Sprintf("%s operation", tt.op)
		t.Run(testName, func(t *testing.T) {
			if got := ops.ArgsCount(tt.op); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operations.ArgsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
