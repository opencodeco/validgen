package validgen

import (
	"reflect"
	"testing"
)

func TestGetFieldTestElementsWithStringFields(t *testing.T) {
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
			name: "Equal string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "eq=abc",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield1",
				operator:     "==",
				roperand:     `"abc"`,
				errorMessage: "myfield1 must be equal to 'abc'",
			},
			wantErr: false,
		},
		{
			name: "Required string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "required",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield1",
				operator:     "!=",
				roperand:     `""`,
				errorMessage: "myfield1 required",
			},
			wantErr: false,
		},
		{
			name: "String size >= 5",
			args: args{
				fieldName:       "myfield5",
				fieldValidation: "min=5",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myfield5)",
				operator:     ">=",
				roperand:     `5`,
				errorMessage: "myfield5 length must be >= 5",
			},
			wantErr: false,
		},
		{
			name: "String size <= 10",
			args: args{
				fieldName:       "myfield6",
				fieldValidation: "max=10",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myfield6)",
				operator:     "<=",
				roperand:     `10`,
				errorMessage: "myfield6 length must be <= 10",
			},
			wantErr: false,
		},
		{
			name: "Equal string ignore case",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "eq_ignore_case=AbC",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "types.ToLower(obj.myStrField)",
				operator:     "==",
				roperand:     `"abc"`,
				errorMessage: "myStrField must be equal to 'abc'",
			},
			wantErr: false,
		},
		{
			name: "Len string",
			args: args{
				fieldName:       "myStrField",
				fieldType:       "string",
				fieldValidation: "len=8",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myStrField)",
				operator:     "==",
				roperand:     `8`,
				errorMessage: "myStrField length must be 8",
			},
			wantErr: false,
		},
		{
			name: "Not equal string",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldType:       "string",
				fieldValidation: "neq=abc",
			},
			want: FieldTestElements{
				loperand:     "obj.MyFieldNotEqual",
				operator:     "!=",
				roperand:     `"abc"`,
				errorMessage: "MyFieldNotEqual must be not equal to 'abc'",
			},
			wantErr: false,
		},
		{
			name: "Not equal string ignore case",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq_ignore_case=AbC",
				fieldType:       "string",
			},
			want: FieldTestElements{
				loperand:     "types.ToLower(obj.MyFieldNotEqual)",
				operator:     "!=",
				roperand:     `"abc"`,
				errorMessage: "MyFieldNotEqual must be not equal to 'abc'",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFieldTestElements(tt.args.fieldName, tt.args.fieldValidation, tt.args.fieldType)
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
