package main

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestFileValidatorGenerate(t *testing.T) {
	type fields struct {
		FileHeader FileHeader
		StructInfo FuncValidator
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Valid struct",
			fields: fields{
				FileHeader: FileHeader{
					ImportPath: "github.com/alexgarzao/myvalidator_samples/ex1/structs",
				},
				StructInfo: FuncValidator{
					Name: "User",
					FieldsValidations: []FieldValidation{
						{
							Name: "FirstName",
							Type: "string",
							Tag:  `validate:"required"`,
						},
						{
							Name: "MyAge",
							Type: "uint8",
							Tag:  `validate:"required"`,
						},
					},
					HasValidateTag: true,
				},
			},
			want: `package validators

import (
	"errors"
	"fmt"

	"github.com/alexgarzao/myvalidator_samples/ex1/structs"
)

var ErrValidation = errors.New("validation error")

func UserValidate(u *structs.User) []error {
	var errs []error

	if u.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}

	if u.MyAge == 0 {
		errs = append(errs, fmt.Errorf("%w: MyAge required", ErrValidation))
	}

	return errs
}
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := &FileValidator{
				FileHeader: tt.fields.FileHeader,
				StructInfo: tt.fields.StructInfo,
			}
			got, err := fv.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileValidator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileValidator.Generate() = %v, want %v", got, tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.want, got, false)
				if len(diffs) > 1 {
					t.Errorf("FileValidator.Generate() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}
