package codegenerator

import (
	"reflect"
	"testing"
)

func TestDefineTestElementsBetweenFields(t *testing.T) {
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
			name: "string fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       "string",
				fieldValidation: "eqfield=myfield2",
			},
			want: TestElements{
				leftOperand:    "obj.myfield1",
				operator:       "==",
				rightOperands:  []string{"obj.myfield2"},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to myfield2",
			},
		},
		{
			name: "string fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       "string",
				fieldValidation: "neqfield=myfield2",
			},
			want: TestElements{
				leftOperand:    "obj.myfield1",
				operator:       "!=",
				rightOperands:  []string{"obj.myfield2"},
				concatOperator: "",
				errorMessage:   "myfield1 must not be equal to myfield2",
			},
		},
		{
			name: "uint8 fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       "uint8",
				fieldValidation: "eqfield=myfield2",
			},
			want: TestElements{
				leftOperand:    "obj.myfield1",
				operator:       "==",
				rightOperands:  []string{"obj.myfield2"},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to myfield2",
			},
		},
		{
			name: "uint8 fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       "uint8",
				fieldValidation: "neqfield=myfield2",
			},
			want: TestElements{
				leftOperand:    "obj.myfield1",
				operator:       "!=",
				rightOperands:  []string{"obj.myfield2"},
				concatOperator: "",
				errorMessage:   "myfield1 must not be equal to myfield2",
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
