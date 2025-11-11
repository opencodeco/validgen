package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestBuildFuncValidatorCodeFieldOperations(t *testing.T) {
	type fields struct {
		Struct *analyzer.Struct
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Field inner op",
			fields: fields{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
						StructName:  "TestStruct",
						Fields: []parser.Field{
							{
								FieldName: "Field1",
								Type:      common.FieldType{BaseType: "string"},
								Tag:       ``,
							},
							{
								FieldName: "Field2",
								Type:      common.FieldType{BaseType: "string"},
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
		{
			name: "Field nested op",
			fields: fields{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
						StructName:  "TestStruct",
						Fields: []parser.Field{
							{
								FieldName: "Field1",
								Type:      common.FieldType{BaseType: "string"},
								Tag:       `validate:"neqfield=Nested.Field2"`,
							},
							{
								FieldName: "Nested",
								Type:      common.FieldType{BaseType: "NestedStruct"},
								Tag:       ``,
							},
						},
					},
					FieldsValidations: []analyzer.FieldValidations{
						{
							Validations: []*analyzer.Validation{AssertParserValidation(t, "neqfield=Nested.Field2")},
						},
						{
							Validations: []*analyzer.Validation{},
						},
					},
				},
			},
			want: `func TestStructValidate(obj *TestStruct) []error {
var errs []error
if !(obj.Field1 != obj.Nested.Field2) {
errs = append(errs, types.NewValidationError("Field1 must not be equal to Nested.Field2"))
}
return errs
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := GenValidations{
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
