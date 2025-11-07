package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
)

func TestBuildValidationCodeFieldOperations(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := GenValidations{}
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := gv.BuildValidationCode(tt.args.fieldName, tt.args.fieldType, []*analyzer.Validation{validation})
			if err != nil {
				t.Errorf("BuildValidationCode() error = %v, wantErr %v", err, nil)
				return
			}
			if got != tt.want {
				t.Errorf("BuildValidationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildValidationCodeWithNestedStructs(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := GenValidations{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{
						PackageName: "main",
					},
				},
				StructsWithValidation: map[string]struct{}{},
			}
			gv.StructsWithValidation[tt.args.fieldType.BaseType] = struct{}{}
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			got, err := gv.BuildValidationCode(tt.args.fieldName, tt.args.fieldType, []*analyzer.Validation{validation})
			if err != nil {
				t.Errorf("BuildValidationCode() error = %v, wantErr %v", err, nil)
				return
			}
			if got != tt.want {
				t.Errorf("BuildValidationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
