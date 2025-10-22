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
		op         string
		fieldTypes []string
		valid      bool
	}{
		// eq operations
		{
			op: "eq",
			fieldTypes: []string{
				"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
				"*<STRING>", "*<INT>", "*<FLOAT>", "*<BOOL>",
			},
			valid: true,
		},

		// required operations
		{
			op: "required",
			fieldTypes: []string{
				"<STRING>", "<INT>", "<FLOAT>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"*<STRING>", "*<INT>", "*<FLOAT>", "*<BOOL>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*[N]<STRING>", "*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: true,
		},

		// gt operations
		{
			op: "gt",
			fieldTypes: []string{
				"<INT>", "<FLOAT>", "*<INT>", "*<FLOAT>",
			},
			valid: true,
		},

		// gte operations
		{
			op: "gte",
			fieldTypes: []string{
				"<INT>", "<FLOAT>", "*<INT>", "*<FLOAT>",
			},
			valid: true,
		},

		// lt operations
		{
			op: "lt",
			fieldTypes: []string{
				"<INT>", "<FLOAT>", "*<INT>", "*<FLOAT>",
			},
			valid: true,
		},

		// lte operations
		{
			op: "lte",
			fieldTypes: []string{
				"<INT>", "<FLOAT>", "*<INT>", "*<FLOAT>",
			},
			valid: true,
		},

		// min operations
		{
			op: "min",
			fieldTypes: []string{
				"<STRING>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"*<STRING>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: true,
		},

		// max operations
		{
			op: "max",
			fieldTypes: []string{
				"<STRING>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"*<STRING>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: true,
		},

		// eq_ignore_case operations
		{
			op: "eq_ignore_case",
			fieldTypes: []string{
				"<STRING>",
				"*<STRING>",
			},
			valid: true,
		},

		// len operations
		{
			op: "len",
			fieldTypes: []string{
				"<STRING>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"*<STRING>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: true,
		},

		// neq operations
		{
			op: "neq",
			fieldTypes: []string{
				"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
				"*<STRING>", "*<INT>", "*<FLOAT>", "*<BOOL>",
			},
			valid: true,
		},

		// neq_ignore_case operations
		{
			op: "neq_ignore_case",
			fieldTypes: []string{
				"<STRING>",
				"*<STRING>",
			},
			valid: true,
		},

		// in operations
		{
			op: "in",
			fieldTypes: []string{
				"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>",
				"*<STRING>", "*<INT>", "*<FLOAT>", "*<BOOL>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*[N]<STRING>", "*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: true,
		},

		// nin operations
		{
			op: "nin",
			fieldTypes: []string{
				"<STRING>", "<INT>", "<FLOAT>", "<BOOL>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"[N]<STRING>", "[N]<INT>", "[N]<FLOAT>", "[N]<BOOL>",
				"*<STRING>", "*<INT>", "*<FLOAT>", "*<BOOL>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*[N]<STRING>", "*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: true,
		},

		// email operations
		{
			op: "email",
			fieldTypes: []string{
				"<STRING>",
				"*<STRING>",
			},
			valid: true,
		},

		// invalid cases
		{
			op: "email",
			fieldTypes: []string{
				"<INT>", "<FLOAT>", "<BOOL>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"*<INT>", "*<FLOAT>", "*<BOOL>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: false,
		},
		{
			op: "gt",
			fieldTypes: []string{
				"<BOOL>",
				"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>",
				"map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]",
				"*<BOOL>",
				"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>",
				"*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]",
			},
			valid: false,
		},
		{
			op:         "neq_ignore_case",
			fieldTypes: []string{"<INT>", "<FLOAT>", "<BOOL>", "[]<STRING>", "map[<STRING>]", "<XPTO>", "*<XPTO>"},
			valid:      false,
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

		for _, fieldType := range tt.fieldTypes {
			testName := fmt.Sprintf("%s %s %s", testType, fieldType, tt.op)

			t.Run(testName, func(t *testing.T) {
				valid := ops.IsValidByType(tt.op, fieldType)
				if valid != tt.valid {
					t.Errorf("IsValidTypeOperation() = %v, want %v", valid, tt.valid)
					return
				}
			})
		}
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
