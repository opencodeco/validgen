package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestBuildFuncValidatorCode(t *testing.T) {
	type fields struct {
		Struct *analyzer.Struct
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Valid struct",
			fields: fields{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
						StructName:  "User",
						Fields: []parser.Field{
							{
								FieldName: "FirstName",
								Type:      "string",
								Tag:       `validate:"required"`,
							},
							{
								FieldName: "MyAge",
								Type:      "uint8",
								Tag:       `validate:"required"`,
							},
						},
					},
					FieldsValidations: []analyzer.FieldValidations{
						{
							Validations: []*analyzer.Validation{AssertParserValidation(t, "required")},
						},
						{
							Validations: []*analyzer.Validation{AssertParserValidation(t, "required")},
						},
					},
				},
			},
			want: `func UserValidate(obj *User) []error {
var errs []error
if !(obj.FirstName != "") {
errs = append(errs, types.NewValidationError("FirstName is required"))
}
if !(obj.MyAge != 0) {
errs = append(errs, types.NewValidationError("MyAge is required"))
}
return errs
}
`,
		},
		{
			name: "FirstName must have 5 characters or more",
			fields: fields{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
						StructName:  "User",
						Fields: []parser.Field{
							{
								FieldName: "FirstName",
								Type:      "string",
								Tag:       `validate:"min=5"`,
							},
						},
					},
					FieldsValidations: []analyzer.FieldValidations{
						{
							Validations: []*analyzer.Validation{AssertParserValidation(t, "min=5")},
						},
					},
				},
			},
			want: `func UserValidate(obj *User) []error {
var errs []error
if !(len(obj.FirstName) >= 5) {
errs = append(errs, types.NewValidationError("FirstName length must be >= 5"))
}
return errs
}
`,
		},
		{
			name: "Field op",
			fields: fields{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
						StructName:  "TestStruct",
						Fields: []parser.Field{
							{
								FieldName: "Field1",
								Type:      "string",
								Tag:       ``,
							},
							{
								FieldName: "Field2",
								Type:      "string",
								Tag:       `validate:"neqfield=Field1"`,
							},
						},
					},
					FieldsValidations: []analyzer.FieldValidations{
						{
							Validations: []*analyzer.Validation{},
						},
						{
							Validations: []*analyzer.Validation{AssertParserValidation(t, "neqfield=Field1")},
						},
					},
				},
			},
			want: `func TestStructValidate(obj *TestStruct) []error {
var errs []error
if !(obj.Field2 != obj.Field1) {
errs = append(errs, types.NewValidationError("Field2 must not be equal to Field1"))
}
return errs
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := genValidations{
				Struct: tt.fields.Struct,
			}
			got, err := gv.BuildFuncValidatorCode()
			if err != nil {
				t.Errorf("FileValidator.GenerateValidator() error = %v, wantErr %v", err, nil)
				return
			}
			if got != tt.want {
				t.Errorf("FileValidator.GenerateValidator() = %v, want %v", got, tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.want, got, false)
				if len(diffs) > 1 {
					t.Errorf("FileValidator.GenerateValidator() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}
