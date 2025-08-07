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
			name: "Required slice string",
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
		{
			name: "Min slice string",
			args: args{
				fieldName:       "myfield",
				fieldType:       "[]string",
				fieldValidation: "min=2",
			},
			want: TestElements{
				leftOperand:   "len(obj.myfield)",
				operator:      ">=",
				rightOperands: []string{`2`},
				errorMessage:  "myfield must have at least 2 elements",
			},
		},
		{
			name: "Max slice string",
			args: args{
				fieldName:       "myfield",
				fieldType:       "[]string",
				fieldValidation: "max=5",
			},
			want: TestElements{
				leftOperand:   "len(obj.myfield)",
				operator:      "<=",
				rightOperands: []string{`5`},
				errorMessage:  "myfield must have at most 5 elements",
			},
		},
		{
			name: "Len slice string",
			args: args{
				fieldName:       "myfield",
				fieldType:       "[]string",
				fieldValidation: "len=3",
			},
			want: TestElements{
				leftOperand:   "len(obj.myfield)",
				operator:      "==",
				rightOperands: []string{`3`},
				errorMessage:  "myfield must have exactly 3 elements",
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
