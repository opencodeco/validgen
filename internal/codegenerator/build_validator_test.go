package codegenerator

import (
	"fmt"
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
)

func TestBuildValidationCode(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "if code with string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "eq=abc",
			},
			want: `if !(obj.strField == "abc") {
errs = append(errs, types.NewValidationError("strField must be equal to 'abc'"))
}
`,
		},
		{
			name: "if code with uint8",
			args: args{
				fieldName:       "intField",
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "gte=123",
			},
			want: `if !(obj.intField >= 123) {
errs = append(errs, types.NewValidationError("intField must be >= 123"))
}
`,
		},
		{
			name: "if code with string and in",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "in=a b c",
			},
			want: `if !(obj.strField == "a" || obj.strField == "b" || obj.strField == "c") {
errs = append(errs, types.NewValidationError("strField must be one of 'a' 'b' 'c'"))
}
`,
		},
		{
			name: "if code with string and not in",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "nin=a b c",
			},
			want: `if !(obj.strField != "a" && obj.strField != "b" && obj.strField != "c") {
errs = append(errs, types.NewValidationError("strField must not be one of 'a' 'b' 'c'"))
}
`,
		},
		{
			name: "Required slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "required",
			},
			want: `if !(len(obj.strField) != 0) {
errs = append(errs, types.NewValidationError("strField must not be empty"))
}
`,
		},
		{
			name: "Min slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "min=2",
			},
			want: `if !(len(obj.strField) >= 2) {
errs = append(errs, types.NewValidationError("strField must have at least 2 elements"))
}
`,
		},
		{
			name: "Max slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "max=5",
			},
			want: `if !(len(obj.strField) <= 5) {
errs = append(errs, types.NewValidationError("strField must have at most 5 elements"))
}
`,
		},
		{
			name: "Len slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "len=3",
			},
			want: `if !(len(obj.strField) == 3) {
errs = append(errs, types.NewValidationError("strField must have exactly 3 elements"))
}
`,
		},
		{
			name: "In slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "in=a b c",
			},
			want: `if !(types.SliceOnlyContains(obj.strField, []string{"a", "b", "c"})) {
errs = append(errs, types.NewValidationError("strField elements must be one of 'a' 'b' 'c'"))
}
`,
		},
		{
			name: "Not in slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "nin=a b c",
			},
			want: `if !(types.SliceNotContains(obj.strField, []string{"a", "b", "c"})) {
errs = append(errs, types.NewValidationError("strField elements must not be one of 'a' 'b' 'c'"))
}
`,
		},

		{
			name: "In array string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[N]"},
				fieldValidation: "in=a b c",
			},
			want: `if !(types.SliceOnlyContains(obj.strField[:], []string{"a", "b", "c"})) {
errs = append(errs, types.NewValidationError("strField elements must be one of 'a' 'b' 'c'"))
}
`,
		},
		{
			name: "Not in array string",
			args: args{
				fieldName:       "strField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[N]"},
				fieldValidation: "nin=a b c",
			},
			want: `if !(types.SliceNotContains(obj.strField[:], []string{"a", "b", "c"})) {
errs = append(errs, types.NewValidationError("strField elements must not be one of 'a' 'b' 'c'"))
}
`,
		},

		{
			name: "inner field operation",
			args: args{
				fieldName:       "field1",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "eqfield=field2",
			},
			want: `if !(obj.field1 == obj.field2) {
errs = append(errs, types.NewValidationError("field1 must be equal to field2"))
}
`,
		},
		{
			name: "nested field operation",
			args: args{
				fieldName:       "field1",
				fieldType:       common.FieldType{BaseType: "string"},
				fieldValidation: "eqfield=Nested.field2",
			},
			want: `if !(obj.field1 == obj.Nested.field2) {
errs = append(errs, types.NewValidationError("field1 must be equal to Nested.field2"))
}
`,
		},

		{
			name: "if code with bool",
			args: args{
				fieldName:       "boolField",
				fieldType:       common.FieldType{BaseType: "bool"},
				fieldValidation: "eq=true",
			},
			want: `if !(obj.boolField == true) {
errs = append(errs, types.NewValidationError("boolField must be equal to true"))
}
`,
		},

		{
			name: "if code with int",
			args: args{
				fieldName:       "intField",
				fieldType:       common.FieldType{BaseType: "int"},
				fieldValidation: "eq=5",
			},
			want: `if !(obj.intField == 5) {
errs = append(errs, types.NewValidationError("intField must be equal to 5"))
}
`,
		},

		{
			name: "if code with float32",
			args: args{
				fieldName:       "floatField",
				fieldType:       common.FieldType{BaseType: "float32"},
				fieldValidation: "eq=5.0",
			},
			want: `if !(obj.floatField == 5.0) {
errs = append(errs, types.NewValidationError("floatField must be equal to 5.0"))
}
`,
		},

		// Map type
		{
			name: "if code with string map",
			args: args{
				fieldName:       "mapField",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "map"},
				fieldValidation: "len=3",
			},
			want: `if !(len(obj.mapField) == 3) {
errs = append(errs, types.NewValidationError("mapField must have exactly 3 elements"))
}
`,
		},
		{
			name: "if code with uint8 map",
			args: args{
				fieldName:       "mapField",
				fieldType:       common.FieldType{BaseType: "uint8", ComposedType: "map"},
				fieldValidation: "len=3",
			},
			want: `if !(len(obj.mapField) == 3) {
errs = append(errs, types.NewValidationError("mapField must have exactly 3 elements"))
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := genValidations{}
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := gv.buildValidationCode(tt.args.fieldName, tt.args.fieldType, []*analyzer.Validation{validation})
			if err != nil {
				t.Errorf("buildValidationCode() error = %v, wantErr %v", err, nil)
				return
			}
			if got != tt.want {
				t.Errorf("buildValidationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildValidationCodeWithNestedStructsAndSlices(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test code with inner struct",
			args: args{
				fieldName:       "Field",
				fieldType:       common.FieldType{BaseType: "main.InnerStructType"},
				fieldValidation: "required",
			},
			want: "errs = append(errs, InnerStructTypeValidate(&obj.Field)...)\n",
		},
		{
			name: "test code with inner struct in another package",
			args: args{
				fieldName:       "Field",
				fieldType:       common.FieldType{BaseType: "mypkg.InnerStructType"},
				fieldValidation: "required",
			},
			want: "errs = append(errs, mypkg.InnerStructTypeValidate(&obj.Field)...)\n",
		},
		{
			name: "test code with required slice of string",
			args: args{
				fieldName:       "Field",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "required",
			},
			want: `if !(len(obj.Field) != 0) {
errs = append(errs, types.NewValidationError("Field must not be empty"))
}
`,
		},
		{
			name: "test code with min slice of string",
			args: args{
				fieldName:       "Field",
				fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
				fieldValidation: "min=2",
			},
			want: `if !(len(obj.Field) >= 2) {
errs = append(errs, types.NewValidationError("Field must have at least 2 elements"))
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := genValidations{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
					},
				},
				StructsWithValidation: map[string]struct{}{},
			}
			gv.StructsWithValidation[tt.args.fieldType.BaseType] = struct{}{}
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := gv.buildValidationCode(tt.args.fieldName, tt.args.fieldType, []*analyzer.Validation{validation})
			if err != nil {
				t.Errorf("buildValidationCode() error = %v, wantErr %v", err, nil)
				return
			}
			if got != tt.want {
				t.Errorf("buildValidationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildValidationCodeWithPointerTypes(t *testing.T) {
	tests := []struct {
		fieldType  common.FieldType
		validation string
		want       string
	}{
		// Pointer basic types.
		{
			fieldType:  common.FieldType{BaseType: "string", ComposedType: "*"},
			validation: "eq=abc",
			want: `if !(obj.field != nil && *obj.field == "abc") {
errs = append(errs, types.NewValidationError("field must be equal to 'abc'"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "bool", ComposedType: "*"},
			validation: "eq=true",
			want: `if !(obj.field != nil && *obj.field == true) {
errs = append(errs, types.NewValidationError("field must be equal to true"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "int", ComposedType: "*"},
			validation: "eq=5",
			want: `if !(obj.field != nil && *obj.field == 5) {
errs = append(errs, types.NewValidationError("field must be equal to 5"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "float32", ComposedType: "*"},
			validation: "eq=5.0",
			want: `if !(obj.field != nil && *obj.field == 5.0) {
errs = append(errs, types.NewValidationError("field must be equal to 5.0"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "bool", ComposedType: "*"},
			validation: "eq=true",
			want: `if !(obj.field != nil && *obj.field == true) {
errs = append(errs, types.NewValidationError("field must be equal to true"))
}
`,
		},

		// Slice pointer types.
		{
			fieldType:  common.FieldType{BaseType: "string", ComposedType: "*[]"},
			validation: "required",
			want: `if !(obj.field != nil && len(*obj.field) != 0) {
errs = append(errs, types.NewValidationError("field must not be empty"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "int", ComposedType: "*[]"},
			validation: "min=10",
			want: `if !(obj.field != nil && len(*obj.field) >= 10) {
errs = append(errs, types.NewValidationError("field must have at least 10 elements"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "float32", ComposedType: "*[]"},
			validation: "in=10.0 12.0 14.0",
			want: `if !(obj.field != nil && types.SliceOnlyContains(*obj.field, []float32{10.0, 12.0, 14.0})) {
errs = append(errs, types.NewValidationError("field elements must be one of '10.0' '12.0' '14.0'"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "bool", ComposedType: "*[]"},
			validation: "min=3",
			want: `if !(obj.field != nil && len(*obj.field) >= 3) {
errs = append(errs, types.NewValidationError("field must have at least 3 elements"))
}
`,
		},

		// Array pointer types.
		{
			fieldType:  common.FieldType{BaseType: "string", ComposedType: "*[N]"},
			validation: "in=a b c",
			want: `if !(obj.field != nil && types.SliceOnlyContains(obj.field[:], []string{"a", "b", "c"})) {
errs = append(errs, types.NewValidationError("field elements must be one of 'a' 'b' 'c'"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "int", ComposedType: "*[N]"},
			validation: "in=1 2 3",
			want: `if !(obj.field != nil && types.SliceOnlyContains(obj.field[:], []int{1, 2, 3})) {
errs = append(errs, types.NewValidationError("field elements must be one of '1' '2' '3'"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "float32", ComposedType: "*[N]"},
			validation: "in=1.1 2.2 3.3",
			want: `if !(obj.field != nil && types.SliceOnlyContains(obj.field[:], []float32{1.1, 2.2, 3.3})) {
errs = append(errs, types.NewValidationError("field elements must be one of '1.1' '2.2' '3.3'"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "bool", ComposedType: "*[N]"},
			validation: "in=true",
			want: `if !(obj.field != nil && types.SliceOnlyContains(obj.field[:], []bool{true})) {
errs = append(errs, types.NewValidationError("field elements must be one of 'true'"))
}
`,
		},

		// Map pointer types.
		{
			fieldType:  common.FieldType{BaseType: "string", ComposedType: "*map"},
			validation: "len=3",
			want: `if !(obj.field != nil && len(*obj.field) == 3) {
errs = append(errs, types.NewValidationError("field must have exactly 3 elements"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "int", ComposedType: "*map"},
			validation: "len=3",
			want: `if !(obj.field != nil && len(*obj.field) == 3) {
errs = append(errs, types.NewValidationError("field must have exactly 3 elements"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "float32", ComposedType: "*map"},
			validation: "len=3",
			want: `if !(obj.field != nil && len(*obj.field) == 3) {
errs = append(errs, types.NewValidationError("field must have exactly 3 elements"))
}
`,
		},
		{
			fieldType:  common.FieldType{BaseType: "bool", ComposedType: "*map"},
			validation: "len=3",
			want: `if !(obj.field != nil && len(*obj.field) == 3) {
errs = append(errs, types.NewValidationError("field must have exactly 3 elements"))
}
`,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("validation: %s with %s (%s)", tt.validation, tt.fieldType.ToGenericType(), tt.fieldType.ToNormalizedString())
		t.Run(testName, func(t *testing.T) {
			gv := genValidations{}
			validation := AssertParserValidation(t, tt.validation)
			got, err := gv.buildValidationCode("field", tt.fieldType, []*analyzer.Validation{validation})
			if err != nil {
				t.Errorf("buildValidationCode() error = %v, wantErr %v", err, nil)
				return
			}
			if got != tt.want {
				t.Errorf("buildValidationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
