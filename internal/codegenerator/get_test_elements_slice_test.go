package codegenerator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithSliceFields(t *testing.T) {
	type args struct {
		fieldName       string
		normalizedType  string
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want TestElements
	}{
		// slice <STRING>
		{
			name: "Required slice string",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) != 0`},
				errorMessage: "myfield must not be empty",
			},
		},
		{
			name: "Min slice string",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "min=2",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) >= 2`},
				errorMessage: "myfield must have at least 2 elements",
			},
		},
		{
			name: "Max slice string",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "max=5",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) <= 5`},
				errorMessage: "myfield must have at most 5 elements",
			},
		},
		{
			name: "Len slice string",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "len=3",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) == 3`},
				errorMessage: "myfield must have exactly 3 elements",
			},
		},
		{
			name: "In slice string with spaces",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "in=a b c",
			},
			want: TestElements{
				conditions:     []string{`types.SliceOnlyContains(obj.myfield, []string{"a", "b", "c"})`},
				concatOperator: "",
				errorMessage:   "myfield elements must be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "In slice string with '",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "in=' a ' ' b ' ' c '",
			},
			want: TestElements{
				conditions:     []string{`types.SliceOnlyContains(obj.myfield, []string{" a ", " b ", " c "})`},
				concatOperator: "",
				errorMessage:   "myfield elements must be one of ' a ' ' b ' ' c '",
			},
		},
		{
			name: "Not in slice string with spaces",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "nin=a b c",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield, []string{"a", "b", "c"})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "Not in slice string with '",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<STRING>",
				fieldValidation: "nin=' a ' ' b ' ' c '",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield, []string{" a ", " b ", " c "})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of ' a ' ' b ' ' c '",
			},
		},

		// slice <INT>
		{
			name: "Required slice int",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<INT>",
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) != 0`},
				errorMessage: "myfield must not be empty",
			},
		},
		{
			name: "Min slice int",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<INT>",
				fieldValidation: "min=2",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) >= 2`},
				errorMessage: "myfield must have at least 2 elements",
			},
		},
		{
			name: "Max slice int",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<INT>",
				fieldValidation: "max=5",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) <= 5`},
				errorMessage: "myfield must have at most 5 elements",
			},
		},
		{
			name: "Len slice int",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<INT>",
				fieldValidation: "len=3",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) == 3`},
				errorMessage: "myfield must have exactly 3 elements",
			},
		},
		{
			name: "In slice int",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<INT>",
				fieldValidation: "in=1 2 3",
			},
			want: TestElements{
				conditions:     []string{`types.SliceOnlyContains(obj.myfield, []<INT>{1, 2, 3})`},
				concatOperator: "",
				errorMessage:   "myfield elements must be one of '1' '2' '3'",
			},
		},
		{
			name: "Not in slice int",
			args: args{
				fieldName:       "myfield",
				normalizedType:  "[]<INT>",
				fieldValidation: "nin=1 2 3",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield, []<INT>{1, 2, 3})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of '1' '2' '3'",
			},
		},
	}

	for _, tt := range tests {
		originalConditions := make([]string, len(tt.want.conditions))
		copy(originalConditions, tt.want.conditions)
		fieldTypes, err := common.HelperFromNormalizedToFieldTypes(tt.args.normalizedType)
		if err != nil {
			t.Errorf("FromNormalizedToFieldTypes() error = %v", err)
			return
		}
		for _, fieldType := range fieldTypes {
			t.Run(tt.name, func(t *testing.T) {
				wantErr := false
				validation := AssertParserValidation(t, tt.args.fieldValidation)
				got, err := DefineTestElements(tt.args.fieldName, fieldType, validation)
				if (err != nil) != wantErr {
					t.Errorf("DefineTestElements() error = %v, wantErr %v", err, wantErr)
					return
				}
				for i := range originalConditions {
					tt.want.conditions[i] = strings.ReplaceAll(originalConditions[i], "<INT>", fieldType.BaseType)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("DefineTestElements() = %+v, want %+v", got, tt.want)
				}
			})
		}
	}
}
