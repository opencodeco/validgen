package codegenerator

import (
	"reflect"
	"testing"
)

func TestDefineTestElementsWithSliceFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       string
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want TestElements
	}{
		{
			name: "Equal string",
			args: args{
				fieldName:       "myfield",
				fieldType:       "[]string",
				fieldValidation: "required",
			},
			want: TestElements{
				leftOperand:   "len(obj.myfield)",
				operator:      ">",
				rightOperands: []string{`0`},
				errorMessage:  "myfield must not be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := false
			got, err := DefineTestElements(tt.args.fieldName, tt.args.fieldType, tt.args.fieldValidation)
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
