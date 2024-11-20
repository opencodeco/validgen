package main

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestStructInfoGenerateValidator(t *testing.T) {
	type fields struct {
		StructInfo StructInfo
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
				StructInfo: StructInfo{
					Name: "User",
					FieldsInfo: []FieldInfo{
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
					PackageName:    "main",
				},
			},
			want: `package main

import (
	"fmt"
)

func UserValidate(obj *User) []error {
	var errs []error

	if obj.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}

	if obj.MyAge == 0 {
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
			fv := tt.fields.StructInfo
			got, err := fv.GenerateValidator()
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
