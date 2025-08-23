package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/parser"
)

func TestBuildValidationCode(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       string
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
				fieldType:       "string",
				fieldValidation: "eq=abc",
			},
			want: `
	if !(obj.strField == "abc") {
		errs = append(errs, types.NewValidationError("strField must be equal to 'abc'"))
	}
`,
		},
		{
			name: "if code with uint8",
			args: args{
				fieldName:       "intField",
				fieldType:       "uint8",
				fieldValidation: "gte=123",
			},
			want: `
	if !(obj.intField >= 123) {
		errs = append(errs, types.NewValidationError("intField must be >= 123"))
	}
`,
		},
		{
			name: "if code with string and in",
			args: args{
				fieldName:       "strField",
				fieldType:       "string",
				fieldValidation: "in=a b c",
			},
			want: `
	if !(obj.strField == "a" || obj.strField == "b" || obj.strField == "c") {
		errs = append(errs, types.NewValidationError("strField must be one of 'a' 'b' 'c'"))
	}
`,
		},
		{
			name: "if code with string and not in",
			args: args{
				fieldName:       "strField",
				fieldType:       "string",
				fieldValidation: "nin=a b c",
			},
			want: `
	if !(obj.strField != "a" && obj.strField != "b" && obj.strField != "c") {
		errs = append(errs, types.NewValidationError("strField must not be one of 'a' 'b' 'c'"))
	}
`,
		},
		{
			name: "Required slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       "[]string",
				fieldValidation: "required",
			},
			want: `
	if !(len(obj.strField) != 0) {
		errs = append(errs, types.NewValidationError("strField must not be empty"))
	}
`,
		},
		{
			name: "Min slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       "[]string",
				fieldValidation: "min=2",
			},
			want: `
	if !(len(obj.strField) >= 2) {
		errs = append(errs, types.NewValidationError("strField must have at least 2 elements"))
	}
`,
		},
		{
			name: "Max slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       "[]string",
				fieldValidation: "max=5",
			},
			want: `
	if !(len(obj.strField) <= 5) {
		errs = append(errs, types.NewValidationError("strField must have at most 5 elements"))
	}
`,
		},
		{
			name: "Len slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       "[]string",
				fieldValidation: "len=3",
			},
			want: `
	if !(len(obj.strField) == 3) {
		errs = append(errs, types.NewValidationError("strField must have exactly 3 elements"))
	}
`,
		},
		{
			name: "In slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       "[]string",
				fieldValidation: "in=a b c",
			},
			want: `
	if !(types.SlicesContains(obj.strField, "a") || types.SlicesContains(obj.strField, "b") || types.SlicesContains(obj.strField, "c")) {
		errs = append(errs, types.NewValidationError("strField elements must be one of 'a' 'b' 'c'"))
	}
`,
		},
		{
			name: "Not in slice string",
			args: args{
				fieldName:       "strField",
				fieldType:       "[]string",
				fieldValidation: "nin=a b c",
			},
			want: `
	if !(!types.SlicesContains(obj.strField, "a") && !types.SlicesContains(obj.strField, "b") && !types.SlicesContains(obj.strField, "c")) {
		errs = append(errs, types.NewValidationError("strField elements must not be one of 'a' 'b' 'c'"))
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
		fieldType       string
		fieldValidation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "if code with nested struct",
			args: args{
				fieldName:       "Field",
				fieldType:       "main.NestedStructType",
				fieldValidation: "required",
			},
			want: `
	errs = append(errs, NestedStructTypeValidate(&obj.Field)...)
`,
		},
		{
			name: "if code with nested struct in another package",
			args: args{
				fieldName:       "Field",
				fieldType:       "mypkg.NestedStructType",
				fieldValidation: "required",
			},
			want: `
	errs = append(errs, mypkg.NestedStructTypeValidate(&obj.Field)...)
`,
		},
		{
			name: "if code with required slice of string",
			args: args{
				fieldName:       "Field",
				fieldType:       "[]string",
				fieldValidation: "required",
			},
			want: `
	if !(len(obj.Field) != 0) {
		errs = append(errs, types.NewValidationError("Field must not be empty"))
	}
`,
		},
		{
			name: "if code with min slice of string",
			args: args{
				fieldName:       "Field",
				fieldType:       "[]string",
				fieldValidation: "min=2",
			},
			want: `
	if !(len(obj.Field) >= 2) {
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
			gv.StructsWithValidation[tt.args.fieldType] = struct{}{}
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
