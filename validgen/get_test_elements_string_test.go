package validgen

import (
	"reflect"
	"testing"
)

func TestGetTestElementsWithStringFields(t *testing.T) {
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
				leftOperand:   "obj.myfield1",
				operator:      "==",
				rightOperands: []string{`"abc"`},
				errorMessage:  "myfield1 must be equal to 'abc'",
			},
		},
		{
			name: "Required string",
			args: args{
				fieldName:       "myfield1",
				fieldValidation: "required",
			},
			want: TestElements{
				leftOperand:   "obj.myfield1",
				operator:      "!=",
				rightOperands: []string{`""`},
				errorMessage:  "myfield1 is required",
			},
		},
		{
			name: "String size >= 5",
			args: args{
				fieldName:       "myfield5",
				fieldValidation: "min=5",
			},
			want: TestElements{
				leftOperand:   "len(obj.myfield5)",
				operator:      ">=",
				rightOperands: []string{`5`},
				errorMessage:  "myfield5 length must be >= 5",
			},
		},
		{
			name: "String size <= 10",
			args: args{
				fieldName:       "myfield6",
				fieldValidation: "max=10",
			},
			want: TestElements{
				leftOperand:   "len(obj.myfield6)",
				operator:      "<=",
				rightOperands: []string{`10`},
				errorMessage:  "myfield6 length must be <= 10",
			},
		},
		{
			name: "Equal string ignore case",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "eq_ignore_case=AbC",
			},
			want: TestElements{
				leftOperand:   "types.ToLower(obj.myStrField)",
				operator:      "==",
				rightOperands: []string{`"abc"`},
				errorMessage:  "myStrField must be equal to 'abc'",
			},
		},
		{
			name: "Len string",
			args: args{
				fieldName:       "myStrField",
				fieldValidation: "len=8",
			},
			want: TestElements{
				leftOperand:   "len(obj.myStrField)",
				operator:      "==",
				rightOperands: []string{`8`},
				errorMessage:  "myStrField length must be 8",
			},
		},
		{
			name: "Not equal string",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq=abc",
			},
			want: TestElements{
				leftOperand:   "obj.MyFieldNotEqual",
				operator:      "!=",
				rightOperands: []string{`"abc"`},
				errorMessage:  "MyFieldNotEqual must not be equal to 'abc'",
			},
		},
		{
			name: "Not equal string ignore case",
			args: args{
				fieldName:       "MyFieldNotEqual",
				fieldValidation: "neq_ignore_case=AbC",
			},
			want: TestElements{
				leftOperand:   "types.ToLower(obj.MyFieldNotEqual)",
				operator:      "!=",
				rightOperands: []string{`"abc"`},
				errorMessage:  "MyFieldNotEqual must not be equal to 'abc'",
			},
		},
		{
			name: "In string with spaces",
			args: args{
				fieldName:       "InField",
				fieldValidation: "in=a b c",
			},
			want: TestElements{
				leftOperand:   "obj.InField",
				operator:      "==",
				rightOperands: []string{`"a"`, `"b"`, `"c"`},
				errorMessage:  "InField must be one of 'a' 'b' 'c'",
			},
		},
		{
			name: "In string with '",
			args: args{
				fieldName:       "InField",
				fieldValidation: "in=' a ' ' b ' ' c '",
			},
			want: TestElements{
				leftOperand:   "obj.InField",
				operator:      "==",
				rightOperands: []string{`" a "`, `" b "`, `" c "`},
				errorMessage:  "InField must be one of ' a ' ' b ' ' c '",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fieldType := "string"
			wantErr := false
			got, err := GetTestElements(tt.args.fieldName, tt.args.fieldValidation, fieldType)
			if (err != nil) != wantErr {
				t.Errorf("GetTestElements() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestElements() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
