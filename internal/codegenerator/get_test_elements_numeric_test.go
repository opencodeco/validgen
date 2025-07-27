package codegenerator

import (
	"reflect"
	"testing"
)

func TestDefineTestElementsWithNumericFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldValidation string
		fieldType       string
	}
	tests := []struct {
		name    string
		args    args
		want    TestElements
		wantErr bool
	}{
		{
			name: "Required uint8",
			args: args{
				fieldName:       "myfield2",
				fieldValidation: "required",
				fieldType:       "uint8",
			},
			want: TestElements{
				leftOperand:   "obj.myfield2",
				operator:      "!=",
				rightOperands: []string{`0`},
				errorMessage:  "myfield2 is required",
			},
			wantErr: false,
		},
		{
			name: "uint8 >= 0",
			args: args{
				fieldName:       "myfield3",
				fieldValidation: "gte=0",
				fieldType:       "uint8",
			},
			want: TestElements{
				leftOperand:   "obj.myfield3",
				operator:      ">=",
				rightOperands: []string{`0`},
				errorMessage:  "myfield3 must be >= 0",
			},
			wantErr: false,
		},
		{
			name: "uint8 <= 130",
			args: args{
				fieldName:       "myfield4",
				fieldValidation: "lte=130",
				fieldType:       "uint8",
			},
			want: TestElements{
				leftOperand:   "obj.myfield4",
				operator:      "<=",
				rightOperands: []string{`130`},
				errorMessage:  "myfield4 must be <= 130",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefineTestElements(tt.args.fieldName, tt.args.fieldType, tt.args.fieldValidation)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefineTestElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefineTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
