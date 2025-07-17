package validgen

import "testing"

func TestIfCode(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       string
		fieldValidation string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IfCode(tt.args.fieldName, tt.args.fieldValidation, tt.args.fieldType)
			if (err != nil) != tt.wantErr {
				t.Errorf("IfCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IfCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
