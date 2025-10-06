package codegenerator

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithArrayFields(t *testing.T) {
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
			name: "In array string with spaces",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[N]"},
				fieldValidation: "in=a b c",
			},
			want: TestElements{
				conditions:     []string{`types.SliceOnlyContains(obj.myfield[:], []string{"a", "b", "c"})`},
				concatOperator: "",
				errorMessage:   "myfield elements must be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "In array string with '",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[N]"},
				fieldValidation: "in=' a ' ' b ' ' c '",
			},
			want: TestElements{
				conditions:     []string{`types.SliceOnlyContains(obj.myfield[:], []string{" a ", " b ", " c "})`},
				concatOperator: "",
				errorMessage:   "myfield elements must be one of ' a ' ' b ' ' c '",
			},
		},
		{
			name: "Not in array string with spaces",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[N]"},
				fieldValidation: "nin=a b c",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield[:], []string{"a", "b", "c"})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "Not in array string with '",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[N]"},
				fieldValidation: "nin=' a ' ' b ' ' c '",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield[:], []string{" a ", " b ", " c "})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of ' a ' ' b ' ' c '",
			},
		},
		{
			name: "In array int",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "[N]"},
				fieldValidation: "in=1 2 3",
			},
			want: TestElements{
				conditions:     []string{`types.SliceOnlyContains(obj.myfield[:], []uint8{1, 2, 3})`},
				concatOperator: "",
				errorMessage:   "myfield elements must be one of '1' '2' '3'",
			},
		},
		{
			name: "Not in array int",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "[N]"},
				fieldValidation: "nin=1 2 3",
			},
			want: TestElements{
				conditions:     []string{`types.SliceNotContains(obj.myfield[:], []uint8{1, 2, 3})`},
				concatOperator: "",
				errorMessage:   "myfield elements must not be one of '1' '2' '3'",
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
