package codegenerator

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithAllTypes(t *testing.T) {
	type check struct {
		types []string
		value string
		want  TestElements
	}
	tests := []struct {
		validation string
		checks     []check
	}{
		// required
		{
			validation: "required",
			checks: []check{
				{
					types: []string{"<STRING>"},
					want: TestElements{
						conditions:     []string{`obj.field != ""`},
						concatOperator: "",
						errorMessage:   "field is required",
					},
				},
				{
					types: []string{"<INT>", "<FLOAT>"},
					want: TestElements{
						conditions:     []string{`obj.field != 0`},
						concatOperator: "",
						errorMessage:   "field is required",
					},
				},
				{
					types: []string{"<BOOL>"},
					want: TestElements{
						conditions:     []string{`obj.field != false`},
						concatOperator: "",
						errorMessage:   "field is required",
					},
				},
				{
					types: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
					want: TestElements{
						conditions:     []string{`len(obj.field) != 0`},
						concatOperator: "",
						errorMessage:   "field must not be empty",
					},
				},
				{
					types: []string{"*<STRING>"},
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != ""`},
						concatOperator: "",
						errorMessage:   "field is required",
					},
				},
				{
					types: []string{"*<INT>", "*<FLOAT>"},
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != 0`},
						concatOperator: "",
						errorMessage:   "field is required",
					},
				},
				{
					types: []string{"*<BOOL>"},
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != false`},
						concatOperator: "",
						errorMessage:   "field is required",
					},
				},
				{
					types: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) != 0`},
						concatOperator: "",
						errorMessage:   "field must not be empty",
					},
				},
				{
					types: []string{"*[N]<STRING>", "*[N]<INT>", "*[N]<FLOAT>", "*[N]<BOOL>"},
					want: TestElements{
						conditions:     []string{`obj.field != nil`},
						concatOperator: "",
						errorMessage:   "field must not be empty",
					},
				},
			},
		},

		// email
		{
			validation: "email",
			checks: []check{
				{
					types: []string{"<STRING>"},
					want: TestElements{
						conditions:     []string{`types.IsValidEmail(obj.field)`},
						concatOperator: "",
						errorMessage:   "field must be a valid email",
					},
				},
			},
		},

		// eq
		{
			validation: "eq={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`obj.field == "abc"`},
						concatOperator: "",
						errorMessage:   "field must be equal to 'abc'",
					},
				},
				{
					types: []string{"<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field == 123`},
						concatOperator: "",
						errorMessage:   "field must be equal to 123",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field == 123.45`},
						concatOperator: "",
						errorMessage:   "field must be equal to 123.45",
					},
				},
				{
					types: []string{"<BOOL>"},
					value: "true",
					want: TestElements{
						conditions:     []string{`obj.field == true`},
						concatOperator: "",
						errorMessage:   "field must be equal to true",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field == "abc"`},
						concatOperator: "",
						errorMessage:   "field must be equal to 'abc'",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field == 123`},
						concatOperator: "",
						errorMessage:   "field must be equal to 123",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field == 123.45`},
						concatOperator: "",
						errorMessage:   "field must be equal to 123.45",
					},
				},
				{
					types: []string{"*<BOOL>"},
					value: "true",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field == true`},
						concatOperator: "",
						errorMessage:   "field must be equal to true",
					},
				},
			},
		},

		// neq
		{
			validation: "neq={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`obj.field != "abc"`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 'abc'",
					},
				},
				{
					types: []string{"<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field != 123`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 123",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != 123.45`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 123.45",
					},
				},
				{
					types: []string{"<BOOL>"},
					value: "true",
					want: TestElements{
						conditions:     []string{`obj.field != true`},
						concatOperator: "",
						errorMessage:   "field must not be equal to true",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != "abc"`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 'abc'",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != 123`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 123",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != 123.45`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 123.45",
					},
				},
				{
					types: []string{"*<BOOL>"},
					value: "true",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field != true`},
						concatOperator: "",
						errorMessage:   "field must not be equal to true",
					},
				},
			},
		},

		// eq_ignore_case
		{
			validation: "eq_ignore_case={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`types.EqualFold(obj.field, "abc")`},
						concatOperator: "",
						errorMessage:   "field must be equal to 'abc'",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.EqualFold(*obj.field, "abc")`},
						concatOperator: "",
						errorMessage:   "field must be equal to 'abc'",
					},
				},
			},
		},

		// neq_ignore_case
		{
			validation: "neq_ignore_case={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`!types.EqualFold(obj.field, "abc")`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 'abc'",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "abc",
					want: TestElements{
						conditions:     []string{`obj.field != nil && !types.EqualFold(*obj.field, "abc")`},
						concatOperator: "",
						errorMessage:   "field must not be equal to 'abc'",
					},
				},
			},
		},

		// gt
		{
			validation: "gt={{.Value}}",
			checks: []check{
				{
					types: []string{"<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field > 123`},
						concatOperator: "",
						errorMessage:   "field must be > 123",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field > 123.45`},
						concatOperator: "",
						errorMessage:   "field must be > 123.45",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field > 123`},
						concatOperator: "",
						errorMessage:   "field must be > 123",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field > 123.45`},
						concatOperator: "",
						errorMessage:   "field must be > 123.45",
					},
				},
			},
		},

		// gte
		{
			validation: "gte={{.Value}}",
			checks: []check{
				{
					types: []string{"<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field >= 123`},
						concatOperator: "",
						errorMessage:   "field must be >= 123",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field >= 123.45`},
						concatOperator: "",
						errorMessage:   "field must be >= 123.45",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field >= 123`},
						concatOperator: "",
						errorMessage:   "field must be >= 123",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field >= 123.45`},
						concatOperator: "",
						errorMessage:   "field must be >= 123.45",
					},
				},
			},
		},

		// lt
		{
			validation: "lt={{.Value}}",
			checks: []check{
				{
					types: []string{"<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field < 123`},
						concatOperator: "",
						errorMessage:   "field must be < 123",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field < 123.45`},
						concatOperator: "",
						errorMessage:   "field must be < 123.45",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field < 123`},
						concatOperator: "",
						errorMessage:   "field must be < 123",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field < 123.45`},
						concatOperator: "",
						errorMessage:   "field must be < 123.45",
					},
				},
			},
		},

		// lte
		{
			validation: "lte={{.Value}}",
			checks: []check{
				{
					types: []string{"<INT>"},
					value: "123",
					want: TestElements{
						conditions:     []string{`obj.field <= 123`},
						concatOperator: "",
						errorMessage:   "field must be <= 123",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field <= 123.45`},
						concatOperator: "",
						errorMessage:   "field must be <= 123.45",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field <= 123.45`},
						concatOperator: "",
						errorMessage:   "field must be <= 123.45",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "123.45",
					want: TestElements{
						conditions:     []string{`obj.field != nil && *obj.field <= 123.45`},
						concatOperator: "",
						errorMessage:   "field must be <= 123.45",
					},
				},
			},
		},

		// len
		{
			validation: "len={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "8",
					want: TestElements{
						conditions:     []string{`len(obj.field) == 8`},
						concatOperator: "",
						errorMessage:   "field length must be 8",
					},
				},
				{
					types: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
					value: "5",
					want: TestElements{
						conditions:     []string{`len(obj.field) == 5`},
						concatOperator: "",
						errorMessage:   "field must have exactly 5 elements",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "8",
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) == 8`},
						concatOperator: "",
						errorMessage:   "field length must be 8",
					},
				},
				{
					types: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
					value: "5",
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) == 5`},
						concatOperator: "",
						errorMessage:   "field must have exactly 5 elements",
					},
				},
			},
		},

		// min
		{
			validation: "min={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "8",
					want: TestElements{
						conditions:     []string{`len(obj.field) >= 8`},
						concatOperator: "",
						errorMessage:   "field length must be >= 8",
					},
				},
				{
					types: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
					value: "5",
					want: TestElements{
						conditions:     []string{`len(obj.field) >= 5`},
						concatOperator: "",
						errorMessage:   "field must have at least 5 elements",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "8",
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) >= 8`},
						concatOperator: "",
						errorMessage:   "field length must be >= 8",
					},
				},
				{
					types: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
					value: "5",
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) >= 5`},
						concatOperator: "",
						errorMessage:   "field must have at least 5 elements",
					},
				},
			},
		},

		// max
		{
			validation: "max={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "20",
					want: TestElements{
						conditions:     []string{`len(obj.field) <= 20`},
						concatOperator: "",
						errorMessage:   "field length must be <= 20",
					},
				},
				{
					types: []string{"[]<STRING>", "[]<INT>", "[]<FLOAT>", "[]<BOOL>", "map[<STRING>]", "map[<INT>]", "map[<FLOAT>]", "map[<BOOL>]"},
					value: "50",
					want: TestElements{
						conditions:     []string{`len(obj.field) <= 50`},
						concatOperator: "",
						errorMessage:   "field must have at most 50 elements",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "20",
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) <= 20`},
						concatOperator: "",
						errorMessage:   "field length must be <= 20",
					},
				},
				{
					types: []string{"*[]<STRING>", "*[]<INT>", "*[]<FLOAT>", "*[]<BOOL>", "*map[<STRING>]", "*map[<INT>]", "*map[<FLOAT>]", "*map[<BOOL>]"},
					value: "50",
					want: TestElements{
						conditions:     []string{`obj.field != nil && len(*obj.field) <= 50`},
						concatOperator: "",
						errorMessage:   "field must have at most 50 elements",
					},
				},
			},
		},

		// in
		{
			validation: "in={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`obj.field == "abc"`, `obj.field == "def"`, `obj.field == "ghi"`},
						concatOperator: "||",
						errorMessage:   "field must be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`obj.field == 123`, `obj.field == 456`, `obj.field == 789`},
						concatOperator: "||",
						errorMessage:   "field must be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`obj.field == 1.23`, `obj.field == 4.56`, `obj.field == 7.89`},
						concatOperator: "||",
						errorMessage:   "field must be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field == true`, `obj.field == false`},
						concatOperator: "||",
						errorMessage:   "field must be one of 'true' 'false'",
					},
				},
				{
					types: []string{"[]<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`types.SliceOnlyContains(obj.field, []string{"abc", "def", "ghi"})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"[]<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`types.SliceOnlyContains(obj.field, []{{.BaseType}}{123, 456, 789})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"[]<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`types.SliceOnlyContains(obj.field, []{{.BaseType}}{1.23, 4.56, 7.89})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"[]<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`types.SliceOnlyContains(obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'true' 'false'",
					},
				},
				{
					types: []string{"map[<STRING>]"},
					value: "key1,key2,key3",
					want: TestElements{
						conditions:     []string{`types.MapOnlyContains(obj.field, []string{"key1", "key2", "key3"})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'key1' 'key2' 'key3'",
					},
				},
				{
					types: []string{"map[<INT>]"},
					value: "1,2,3",
					want: TestElements{
						conditions:     []string{`types.MapOnlyContains(obj.field, []{{.BaseType}}{1, 2, 3})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1' '2' '3'",
					},
				},
				{
					types: []string{"map[<FLOAT>]"},
					value: "1.1,2.2,3.3",
					want: TestElements{
						conditions:     []string{`types.MapOnlyContains(obj.field, []{{.BaseType}}{1.1, 2.2, 3.3})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1.1' '2.2' '3.3'",
					},
				},
				{
					types: []string{"map[<BOOL>]"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`types.MapOnlyContains(obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field == "abc")`, `(obj.field != nil && *obj.field == "def")`, `(obj.field != nil && *obj.field == "ghi")`},
						concatOperator: "||",
						errorMessage:   "field must be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field == 123)`, `(obj.field != nil && *obj.field == 456)`, `(obj.field != nil && *obj.field == 789)`},
						concatOperator: "||",
						errorMessage:   "field must be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field == 1.23)`, `(obj.field != nil && *obj.field == 4.56)`, `(obj.field != nil && *obj.field == 7.89)`},
						concatOperator: "||",
						errorMessage:   "field must be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"*<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field == true)`, `(obj.field != nil && *obj.field == false)`},
						concatOperator: "||",
						errorMessage:   "field must be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*[]<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(*obj.field, []string{"abc", "def", "ghi"})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"*[]<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(*obj.field, []{{.BaseType}}{123, 456, 789})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"*[]<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(*obj.field, []{{.BaseType}}{1.23, 4.56, 7.89})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"*[]<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(*obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*[N]<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(obj.field[:], []string{"abc", "def", "ghi"})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"*[N]<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(obj.field[:], []{{.BaseType}}{123, 456, 789})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"*[N]<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(obj.field[:], []{{.BaseType}}{1.23, 4.56, 7.89})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"*[N]<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceOnlyContains(obj.field[:], []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*map[<STRING>]"},
					value: "key1,key2,key3",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapOnlyContains(*obj.field, []string{"key1", "key2", "key3"})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'key1' 'key2' 'key3'",
					},
				},
				{
					types: []string{"*map[<INT>]"},
					value: "1,2,3",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapOnlyContains(*obj.field, []{{.BaseType}}{1, 2, 3})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1' '2' '3'",
					},
				},
				{
					types: []string{"*map[<FLOAT>]"},
					value: "1.1,2.2,3.3",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapOnlyContains(*obj.field, []{{.BaseType}}{1.1, 2.2, 3.3})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of '1.1' '2.2' '3.3'",
					},
				},
				{
					types: []string{"*map[<BOOL>]"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapOnlyContains(*obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must be one of 'true' 'false'",
					},
				},
			},
		},

		// nin
		{
			validation: "nin={{.Value}}",
			checks: []check{
				{
					types: []string{"<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`obj.field != "abc"`, `obj.field != "def"`, `obj.field != "ghi"`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`obj.field != 123`, `obj.field != 456`, `obj.field != 789`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`obj.field != 1.23`, `obj.field != 4.56`, `obj.field != 7.89`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != true`, `obj.field != false`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of 'true' 'false'",
					},
				},
				{
					types: []string{"[]<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`types.SliceNotContains(obj.field, []string{"abc", "def", "ghi"})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"[]<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`types.SliceNotContains(obj.field, []{{.BaseType}}{123, 456, 789})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"[]<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`types.SliceNotContains(obj.field, []{{.BaseType}}{1.23, 4.56, 7.89})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"[]<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`types.SliceNotContains(obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'true' 'false'",
					},
				},
				{
					types: []string{"map[<STRING>]"},
					value: "key1,key2,key3",
					want: TestElements{
						conditions:     []string{`types.MapNotContains(obj.field, []string{"key1", "key2", "key3"})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'key1' 'key2' 'key3'",
					},
				},
				{
					types: []string{"map[<INT>]"},
					value: "1,2,3",
					want: TestElements{
						conditions:     []string{`types.MapNotContains(obj.field, []{{.BaseType}}{1, 2, 3})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1' '2' '3'",
					},
				},
				{
					types: []string{"map[<FLOAT>]"},
					value: "1.1,2.2,3.3",
					want: TestElements{
						conditions:     []string{`types.MapNotContains(obj.field, []{{.BaseType}}{1.1, 2.2, 3.3})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1.1' '2.2' '3.3'",
					},
				},
				{
					types: []string{"map[<BOOL>]"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`types.MapNotContains(obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field != "abc")`, `(obj.field != nil && *obj.field != "def")`, `(obj.field != nil && *obj.field != "ghi")`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"*<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field != 123)`, `(obj.field != nil && *obj.field != 456)`, `(obj.field != nil && *obj.field != 789)`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"*<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field != 1.23)`, `(obj.field != nil && *obj.field != 4.56)`, `(obj.field != nil && *obj.field != 7.89)`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"*<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`(obj.field != nil && *obj.field != true)`, `(obj.field != nil && *obj.field != false)`},
						concatOperator: "&&",
						errorMessage:   "field must not be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*[]<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(*obj.field, []string{"abc", "def", "ghi"})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"*[]<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(*obj.field, []{{.BaseType}}{123, 456, 789})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"*[]<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(*obj.field, []{{.BaseType}}{1.23, 4.56, 7.89})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"*[]<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(*obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*[N]<STRING>"},
					value: "abc,def,ghi",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(obj.field[:], []string{"abc", "def", "ghi"})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'abc' 'def' 'ghi'",
					},
				},
				{
					types: []string{"*[N]<INT>"},
					value: "123,456,789",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(obj.field[:], []{{.BaseType}}{123, 456, 789})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '123' '456' '789'",
					},
				},
				{
					types: []string{"*[N]<FLOAT>"},
					value: "1.23,4.56,7.89",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(obj.field[:], []{{.BaseType}}{1.23, 4.56, 7.89})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1.23' '4.56' '7.89'",
					},
				},
				{
					types: []string{"*[N]<BOOL>"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.SliceNotContains(obj.field[:], []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'true' 'false'",
					},
				},
				{
					types: []string{"*map[<STRING>]"},
					value: "key1,key2,key3",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapNotContains(*obj.field, []string{"key1", "key2", "key3"})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'key1' 'key2' 'key3'",
					},
				},
				{
					types: []string{"*map[<INT>]"},
					value: "1,2,3",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapNotContains(*obj.field, []{{.BaseType}}{1, 2, 3})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1' '2' '3'",
					},
				},
				{
					types: []string{"*map[<FLOAT>]"},
					value: "1.1,2.2,3.3",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapNotContains(*obj.field, []{{.BaseType}}{1.1, 2.2, 3.3})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of '1.1' '2.2' '3.3'",
					},
				},
				{
					types: []string{"*map[<BOOL>]"},
					value: "true,false",
					want: TestElements{
						conditions:     []string{`obj.field != nil && types.MapNotContains(*obj.field, []bool{true, false})`},
						concatOperator: "",
						errorMessage:   "field elements must not be one of 'true' 'false'",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		for _, check := range tt.checks {
			for _, acceptedType := range check.types {
				fieldTypes, err := common.HelperFromNormalizedToFieldTypes(acceptedType)
				if err != nil {
					t.Errorf("HelperFromNormalizedToFieldTypes() error = %v", err)
					return
				}
				for _, fieldType := range fieldTypes {
					// op, _, _ := strings.Cut(tt.validation, "=")
					validation := replaceValidationValue(tt.validation, check.value)
					want := TestElements{
						concatOperator: check.want.concatOperator,
						errorMessage:   check.want.errorMessage,
					}
					want.conditions = make([]string, len(check.want.conditions))
					copy(want.conditions, check.want.conditions)

					for i := range want.conditions {
						want.conditions[i] = strings.ReplaceAll(want.conditions[i], "{{.BaseType}}", fieldType.BaseType)
					}

					testName := fmt.Sprintf("validation(%s) type(%s)", validation, fieldType)
					t.Run(testName, func(t *testing.T) {
						validation := AssertParserValidation(t, validation)
						got, err := DefineTestElements("field", fieldType, validation)
						if err != nil {
							t.Errorf("DefineTestElements() error = %v", err)
							return
						}
						if !reflect.DeepEqual(got, want) {
							t.Errorf("DefineTestElements() = %+v, want %+v", got, want)
						}
					})
				}
			}
		}
	}
}

func replaceValidationValue(text, value string) string {

	text = strings.ReplaceAll(text, "{{.Value}}", value)

	return text
}
