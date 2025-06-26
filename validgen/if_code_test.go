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
		errs = append(errs, fmt.Errorf("%w: strField must be equal to 'abc'", types.ErrValidation))
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
