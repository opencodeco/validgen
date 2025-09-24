package codegenerator

import (
	"reflect"
	"testing"
)

func TestDefineTestElementsWithBoolFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want TestElements
	}{
		{
			name: "equal bool",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "eq=true",
			},
			want: TestElements{
				conditions:   []string{`obj.myfield1 == true`},
				errorMessage: "myfield1 must be equal to true",
			},
		},
		{
			name: "not equal bool",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq=false",
			},
			want: TestElements{
				conditions:   []string{`obj.MyFieldNotEqual != false`},
				errorMessage: "MyFieldNotEqual must not be equal to false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldType := "bool"
			wantErr := false
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := DefineTestElements(tt.args.fieldName, fieldType, validation)
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
