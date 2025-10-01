package codegenerator

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithSliceFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want TestElements
	}{
		{
			name: "Required slice string",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
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
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "nin=' a ' ' b ' ' c '",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield, []string{" a ", " b ", " c "})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of ' a ' ' b ' ' c '",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := false
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := DefineTestElements(tt.args.fieldName, tt.args.fieldType, validation)
			if (err != nil) != wantErr {
				t.Errorf("DefineTestElements() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefineTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
