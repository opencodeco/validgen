package validgen

import (
	"reflect"
	"testing"
)

func TestGetTestElementsWithNumericFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldValidation string
		fieldType       string
	}
	tests := []struct {
		name    string
		args    args
		want    FieldTestElements
		wantErr bool
	}{
		{
			name: "Required uint8",
			args: args{
				fieldName:       "myfield2",
				fieldValidation: "required",
				fieldType:       "uint8",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield2",
				operator:     "!=",
				roperand:     `0`,
				errorMessage: "myfield2 required",
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
			want: FieldTestElements{
				loperand:     "obj.myfield3",
				operator:     ">=",
				roperand:     `0`,
				errorMessage: "myfield3 must be >= 0",
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
			want: FieldTestElements{
				loperand:     "obj.myfield4",
				operator:     "<=",
				roperand:     `130`,
				errorMessage: "myfield4 must be <= 130",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTestElements(tt.args.fieldName, tt.args.fieldValidation, tt.args.fieldType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFieldTestElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFieldTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
