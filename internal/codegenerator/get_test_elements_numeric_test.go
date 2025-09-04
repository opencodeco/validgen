package codegenerator

import (
	"reflect"
	"testing"
)

func TestDefineTestElementsWithNumericFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       string
		fieldValidation string
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
				fieldType:       "uint8",
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`obj.myfield2 != 0`},
				errorMessage: "myfield2 is required",
			},
			wantErr: false,
		},
		{
			name: "uint8 >= 0",
			args: args{
				fieldName:       "myfield3",
				fieldType:       "uint8",
				fieldValidation: "gte=0",
			},
			want: TestElements{
				conditions:   []string{`obj.myfield3 >= 0`},
				errorMessage: "myfield3 must be >= 0",
			},
			wantErr: false,
		},
		{
			name: "uint8 <= 130",
			args: args{
				fieldName:       "myfield4",
				fieldType:       "uint8",
				fieldValidation: "lte=130",
			},
			want: TestElements{
				conditions:   []string{`obj.myfield4 <= 130`},
				errorMessage: "myfield4 must be <= 130",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := DefineTestElements(tt.args.fieldName, tt.args.fieldType, validation)
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
