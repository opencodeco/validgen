package codegenerator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithStringPointerFields(t *testing.T) {
	tests := []struct {
		validation string
		want       TestElements
	}{
		{
			validation: "eq=abc",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && *obj.Field == "abc"`},
				errorMessage: "Field must be equal to 'abc'",
			},
		},
		{
			validation: "required",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && *obj.Field != ""`},
				errorMessage: "Field is required",
			},
		},
		{
			validation: "min=5",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && len(*obj.Field) >= 5`},
				errorMessage: "Field length must be >= 5",
			},
		},
		{
			validation: "max=10",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && len(*obj.Field) <= 10`},
				errorMessage: "Field length must be <= 10",
			},
		},
		{
			validation: "eq_ignore_case=AbC",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && types.EqualFold(*obj.Field, "AbC")`},
				errorMessage: "Field must be equal to 'AbC'",
			},
		},
		{
			validation: "len=8",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && len(*obj.Field) == 8`},
				errorMessage: "Field length must be 8",
			},
		},
		{
			validation: "neq=abc",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && *obj.Field != "abc"`},
				errorMessage: "Field must not be equal to 'abc'",
			},
		},
		{
			validation: "neq_ignore_case=AbC",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && !types.EqualFold(*obj.Field, "AbC")`},
				errorMessage: "Field must not be equal to 'AbC'",
			},
		},
		{
			validation: "in=a b c",
			want: TestElements{
				conditions:     []string{`(obj.Field != nil && *obj.Field == "a")`, `(obj.Field != nil && *obj.Field == "b")`, `(obj.Field != nil && *obj.Field == "c")`},
				concatOperator: "||",
				errorMessage:   "Field must be one of 'a' 'b' 'c'",
			},
		},
		{
			validation: "nin=a b c",
			want: TestElements{
				conditions:     []string{`(obj.Field != nil && *obj.Field != "a")`, `(obj.Field != nil && *obj.Field != "b")`, `(obj.Field != nil && *obj.Field != "c")`},
				concatOperator: "&&",
				errorMessage:   "Field must not be one of 'a' 'b' 'c'",
			},
		},
		{
			validation: "email",
			want: TestElements{
				conditions:   []string{`obj.Field != nil && types.IsValidEmail(*obj.Field)`},
				errorMessage: "Field must be a valid email",
			},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("validation: %s with *string", tt.validation)
		t.Run(testName, func(t *testing.T) {
			fieldType := common.FieldType{BaseType: "string", ComposedType: "*"}
			validation := AssertParserValidation(t, tt.validation)
			got, err := DefineTestElements("Field", fieldType, validation)
			if err != nil {
				t.Errorf("DefineTestElements() error = %v, wantErr %v", err, nil)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefineTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
