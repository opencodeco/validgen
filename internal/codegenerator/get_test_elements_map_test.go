package codegenerator

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithMapFields(t *testing.T) {
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
		// string map operations
		{
			name: "required string map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) != 0`},
				errorMessage: "myfield must not be empty",
			},
		},
		{
			name: "len string map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "len=3",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) == 3`},
				errorMessage: "myfield must have exactly 3 elements",
			},
		},
		{
			name: "min string map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "min=2",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) >= 2`},
				errorMessage: "myfield must have at least 2 elements",
			},
		},
		{
			name: "max string map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "max=5",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) <= 5`},
				errorMessage: "myfield must have at most 5 elements",
			},
		},
		{
			name: "in string map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "in=1 2 3",
			},
			want: TestElements{
				conditions:   []string{`types.MapOnlyContains(obj.myfield, []string{"1", "2", "3"})`},
				errorMessage: "myfield elements must be one of '1' '2' '3'",
			},
		},
		{
			name: "nin string map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "nin=1 2 3",
			},
			want: TestElements{
				conditions:   []string{`types.MapNotContains(obj.myfield, []string{"1", "2", "3"})`},
				errorMessage: "myfield elements must not be one of '1' '2' '3'",
			},
		},

		// uint8 map operations
		{
			name: "required uint8 map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) != 0`},
				errorMessage: "myfield must not be empty",
			},
		},
		{
			name: "len uint8 map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "len=3",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) == 3`},
				errorMessage: "myfield must have exactly 3 elements",
			},
		},
		{
			name: "min uint8 map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "min=2",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) >= 2`},
				errorMessage: "myfield must have at least 2 elements",
			},
		},
		{
			name: "max uint8 map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "max=5",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield) <= 5`},
				errorMessage: "myfield must have at most 5 elements",
			},
		},
		{
			name: "in uint8 map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "in=1 2 3",
			},
			want: TestElements{
				conditions:   []string{`types.MapOnlyContains(obj.myfield, []uint8{1, 2, 3})`},
				errorMessage: "myfield elements must be one of '1' '2' '3'",
			},
		},
		{
			name: "nin uint8 map",
			args: args{
				fieldName:       "myfield",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "nin=1 2 3",
			},
			want: TestElements{
				conditions:   []string{`types.MapNotContains(obj.myfield, []uint8{1, 2, 3})`},
				errorMessage: "myfield elements must not be one of '1' '2' '3'",
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
