package main

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestFuncValidatorGenerate(t *testing.T) {
	type fields struct {
		Name              string
		FieldsValidations []FieldInfo
		HasValidateTag    bool
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
				Name: "User",
				FieldsValidations: []FieldInfo{
					{
						Name: "FirstName",
						Type: "string",
						Tag:  "validate:\"required\"",
					},
					{
						Name: "MyAge",
						Type: "uint8",
						Tag:  "validate:\"required\"",
					},
				},
				HasValidateTag: true,
			},
			want: `func UserValidate(u *User) []error {
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
			s := &StructInfo{
				Name:           tt.fields.Name,
				FieldsInfo:     tt.fields.FieldsValidations,
				HasValidateTag: tt.fields.HasValidateTag,
			}
			got, err := s.GenerateFuncValidator()
			if (err != nil) != tt.wantErr {
				t.Errorf("FuncValidator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FuncValidator.Generate() = %v, want %v", got, tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.want, got, false)
				if len(diffs) > 1 {
					t.Errorf("FuncValidator.Generate() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}
