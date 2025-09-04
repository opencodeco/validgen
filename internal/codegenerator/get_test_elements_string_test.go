package codegenerator

import (
	"reflect"
	"testing"
)

func TestDefineTestElementsWithStringFields(t *testing.T) {
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
			name: "Equal string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "eq=abc",
			},
			want: TestElements{
				conditions:   []string{`obj.myfield1 == "abc"`},
				errorMessage: "myfield1 must be equal to 'abc'",
			},
		},
		{
			name: "Required string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`obj.myfield1 != ""`},
				errorMessage: "myfield1 is required",
			},
		},
		{
			name: "String size >= 5",
			args: args{
				fieldName:       "myfield5",
				fieldValidation: "min=5",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield5) >= 5`},
				errorMessage: "myfield5 length must be >= 5",
			},
		},
		{
			name: "String size <= 10",
			args: args{
				fieldName:       "myfield6",
				fieldValidation: "max=10",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myfield6) <= 10`},
				errorMessage: "myfield6 length must be <= 10",
			},
		},
		{
			name: "Equal string ignore case",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "eq_ignore_case=AbC",
			},
			want: TestElements{
				conditions:   []string{`types.EqualFold(obj.myStrField, "AbC")`},
				errorMessage: "myStrField must be equal to 'AbC'",
			},
		},
		{
			name: "Len string",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "len=8",
			},
			want: TestElements{
				conditions:   []string{`len(obj.myStrField) == 8`},
				errorMessage: "myStrField length must be 8",
			},
		},
		{
			name: "Not equal string",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq=abc",
			},
			want: TestElements{
				conditions:   []string{`obj.MyFieldNotEqual != "abc"`},
				errorMessage: "MyFieldNotEqual must not be equal to 'abc'",
			},
		},
		{
			name: "Not equal string ignore case",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq_ignore_case=AbC",
			},
			want: TestElements{
				conditions:   []string{`!types.EqualFold(obj.MyFieldNotEqual, "AbC")`},
				errorMessage: "MyFieldNotEqual must not be equal to 'AbC'",
			},
		},
		{
			name: "In string with spaces",
			args: args{
				fieldName:       "InField",
				fieldValidation: "in=a b c",
			},
			want: TestElements{
				conditions:     []string{`obj.InField == "a"`, `obj.InField == "b"`, `obj.InField == "c"`},
				concatOperator: "||",
				errorMessage:   "InField must be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "In string with '",
			args: args{
				fieldName:       "InField",
				fieldValidation: "in=' a ' ' b ' ' c '",
			},
			want: TestElements{
				conditions:     []string{`obj.InField == " a "`, `obj.InField == " b "`, `obj.InField == " c "`},
				concatOperator: "||",
				errorMessage:   "InField must be one of ' a ' ' b ' ' c '",
			},
		},
		{
			name: "NotIn string with spaces",
			args: args{
				fieldName:       "NotInField",
				fieldValidation: "nin=a b c",
			},
			want: TestElements{
				conditions:     []string{`obj.NotInField != "a"`, `obj.NotInField != "b"`, `obj.NotInField != "c"`},
				concatOperator: "&&",
				errorMessage:   "NotInField must not be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "Email validation",
			args: args{
				fieldName:       "EmailField",
				fieldValidation: "email",
			},
			want: TestElements{
				conditions:   []string{`types.IsValidEmail(obj.EmailField)`},
				errorMessage: "EmailField must be a valid email",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldType := "string"
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
