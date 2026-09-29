package codegenerator

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsBetweenInnerFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want TestElements
	}{
		// <STRING>
		{
			name: "inner string fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "eqfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 == obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to myfield2",
			},
		},
		{
			name: "inner string fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "neqfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 != obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must not be equal to myfield2",
			},
		},

		// <INT>
		{
			name: "inner uint8 fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "eqfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 == obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to myfield2",
			},
		},
		{
			name: "inner uint8 fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "neqfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 != obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must not be equal to myfield2",
			},
		},
		{
			name: "inner uint8 field must be greater than or equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "gtefield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 >= obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be >= myfield2",
			},
		},
		{
			name: "inner uint8 field must be greater than",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "gtfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 > obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be > myfield2",
			},
		},
		{
			name: "inner uint8 fields must less than or equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "ltefield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 <= obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be <= myfield2",
			},
		},
		{
			name: "inner uint8 fields must less than",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "ltfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 < obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be < myfield2",
			},
		},

		// <BOOL>
		{
			name: "inner bool fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "bool"},
				fieldValidation: "eqfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 == obj.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to myfield2",
			},
		},
		{
			name: "inner bool fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "bool"},
				fieldValidation: "neqfield=myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 != obj.myfield2`},
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

	testFloatFieldComparisons(t, "myfield2", "")
	testComplexFieldComparisons(t, "myfield2", "")
}

func testFloatFieldComparisons(t *testing.T, target, namePrefix string) {
	t.Helper()

	ops := []struct {
		op       string
		operator string
		message  string
	}{
		{"eqfield", "==", "must be equal to"},
		{"neqfield", "!=", "must not be equal to"},
		{"gtefield", ">=", "must be >="},
		{"gtfield", ">", "must be >"},
		{"ltefield", "<=", "must be <="},
		{"ltfield", "<", "must be <"},
	}

	for _, baseType := range []string{"float32", "float64"} {
		for _, op := range ops {
			t.Run(namePrefix+baseType+" fields "+op.op, func(t *testing.T) {
				validation := AssertParserValidation(t, op.op+"="+target)
				got, err := DefineTestElements("myfield1", common.FieldType{BaseType: baseType}, validation)
				if err != nil {
					t.Errorf("DefineTestElements() error = %v", err)
					return
				}

				want := TestElements{
					conditions:     []string{"obj.myfield1 " + op.operator + " obj." + target},
					concatOperator: "",
					errorMessage:   "myfield1 " + op.message + " " + target,
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("DefineTestElements() = %+v, want %+v", got, want)
				}
			})
		}
	}
}

func testComplexFieldComparisons(t *testing.T, target, namePrefix string) {
	t.Helper()

	ops := []struct {
		op       string
		operator string
		message  string
	}{
		{"eqfield", "==", "must be equal to"},
		{"neqfield", "!=", "must not be equal to"},
	}

	for _, baseType := range []string{"complex64", "complex128"} {
		for _, op := range ops {
			t.Run(namePrefix+baseType+" fields "+op.op, func(t *testing.T) {
				validation := AssertParserValidation(t, op.op+"="+target)
				got, err := DefineTestElements("myfield1", common.FieldType{BaseType: baseType}, validation)
				if err != nil {
					t.Errorf("DefineTestElements() error = %v", err)
					return
				}

				want := TestElements{
					conditions:     []string{"obj.myfield1 " + op.operator + " obj." + target},
					concatOperator: "",
					errorMessage:   "myfield1 " + op.message + " " + target,
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("DefineTestElements() = %+v, want %+v", got, want)
				}
			})
		}
	}
}

func TestDefineTestElementsBetweenNestedFields(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want TestElements
	}{
		{
			name: "nested string fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "eqfield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 == obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to nested.myfield2",
			},
		},
		{
			name: "nested string fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "neqfield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 != obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must not be equal to nested.myfield2",
			},
		},
		{
			name: "nested uint8 fields must be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "eqfield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 == obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be equal to nested.myfield2",
			},
		},
		{
			name: "nested uint8 fields must not be equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "neqfield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 != obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must not be equal to nested.myfield2",
			},
		},
		{
			name: "nested uint8 field must be greater than or equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "gtefield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 >= obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be >= nested.myfield2",
			},
		},
		{
			name: "nested uint8 field must be greater than",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "gtfield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 > obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be > nested.myfield2",
			},
		},
		{
			name: "nested uint8 fields must less than or equal",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "ltefield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 <= obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be <= nested.myfield2",
			},
		},
		{
			name: "nested uint8 fields must less than",
			args: args{
				fieldName:       "myfield1",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "ltfield=nested.myfield2",
			},
			want: TestElements{
				conditions:     []string{`obj.myfield1 < obj.nested.myfield2`},
				concatOperator: "",
				errorMessage:   "myfield1 must be < nested.myfield2",
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

	testFloatFieldComparisons(t, "nested.myfield2", "nested ")
	testComplexFieldComparisons(t, "nested.myfield2", "nested ")
}
