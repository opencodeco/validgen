package validgen

import (
	"reflect"
	"testing"
)

func TestGetFieldTestElementsWithStringFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want FieldTestElements
	}{
		{
			name: "Equal string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "eq=abc",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield1",
				operator:     "==",
				roperand:     `"abc"`,
				errorMessage: "myfield1 must be equal to 'abc'",
			},
		},
		{
			name: "Required string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "required",
			},
			want: FieldTestElements{
				loperand:     "obj.myfield1",
				operator:     "!=",
				roperand:     `""`,
				errorMessage: "myfield1 required",
			},
		},
		{
			name: "String size >= 5",
			args: args{
				fieldName:       "myfield5",
				fieldValidation: "min=5",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myfield5)",
				operator:     ">=",
				roperand:     `5`,
				errorMessage: "myfield5 length must be >= 5",
			},
		},
		{
			name: "String size <= 10",
			args: args{
				fieldName:       "myfield6",
				fieldValidation: "max=10",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myfield6)",
				operator:     "<=",
				roperand:     `10`,
				errorMessage: "myfield6 length must be <= 10",
			},
		},
		{
			name: "Equal string ignore case",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "eq_ignore_case=AbC",
			},
			want: FieldTestElements{
				loperand:     "types.ToLower(obj.myStrField)",
				operator:     "==",
				roperand:     `"abc"`,
				errorMessage: "myStrField must be equal to 'abc'",
			},
		},
		{
			name: "Len string",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "len=8",
			},
			want: FieldTestElements{
				loperand:     "len(obj.myStrField)",
				operator:     "==",
				roperand:     `8`,
				errorMessage: "myStrField length must be 8",
			},
		},
		{
			name: "Not equal string",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq=abc",
			},
			want: FieldTestElements{
				loperand:     "obj.MyFieldNotEqual",
				operator:     "!=",
				roperand:     `"abc"`,
				errorMessage: "MyFieldNotEqual must be not equal to 'abc'",
			},
		},
		{
			name: "Not equal string ignore case",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq_ignore_case=AbC",
			},
			want: FieldTestElements{
				loperand:     "types.ToLower(obj.MyFieldNotEqual)",
				operator:     "!=",
				roperand:     `"abc"`,
				errorMessage: "MyFieldNotEqual must be not equal to 'abc'",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldType := "string"
			wantErr := false
			got, err := GetFieldTestElements(tt.args.fieldName, tt.args.fieldValidation, fieldType)
			if (err != nil) != false {
				t.Errorf("GetFieldTestElements() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFieldTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
