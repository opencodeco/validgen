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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := genValidations{}
			got, err := gv.buildValidationCode(tt.args.fieldName, tt.args.fieldType, []string{tt.args.fieldValidation})
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

func TestBuildValidationCodeWithNestedStructs(t *testing.T) {
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
	if !(len(obj.Field) > 0) {
		errs = append(errs, types.NewValidationError("Field must not be empty"))
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
			got, err := gv.buildValidationCode(tt.args.fieldName, tt.args.fieldType, []string{tt.args.fieldValidation})
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
