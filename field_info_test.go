package main

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestFieldValidationGenerate(t *testing.T) {
	type fields struct {
		Name string
		Type string
		Tag  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "String field required",
			fields: fields{
				Name: "FirstName",
				Type: "string",
				Tag:  "validate:\"required\"",
			},
			want: `	if u.FirstName == "" {
		errs = append(errs, fmt.Errorf("%w: FirstName required", ErrValidation))
	}
`,
			wantErr: false,
		},
		{
			name: "Integer field required",
			fields: fields{
				Name: "MyNumber",
				Type: "uint8",
				Tag:  "validate:\"required\"",
			},
			want: `	if u.MyNumber == 0 {
		errs = append(errs, fmt.Errorf("%w: MyNumber required", ErrValidation))
	}
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FieldInfo{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
				Tag:  tt.fields.Tag,
			}
			got, err := f.GenerateTestField()
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldValidation.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("FieldValidation.Generate() = %v, want %v", got, tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.want, got, false)
				if len(diffs) > 1 {
					t.Errorf("FieldValidation.Generate() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}
