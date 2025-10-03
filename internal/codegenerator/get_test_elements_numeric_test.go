package codegenerator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestDefineTestElementsWithNumericFields(t *testing.T) {
	type args struct {
		fieldValidation string
	}
	tests := []struct {
		name    string
		args    args
		want    TestElements
		wantErr bool
	}{
		{
			name: "required field",
			args: args{
				fieldValidation: "required",
			},
			want: TestElements{
				conditions:   []string{`obj.field != 0`},
				errorMessage: "field is required",
			},
			wantErr: false,
		},
		{
			name: "field = 123",
			args: args{
				fieldValidation: "eq=123",
			},
			want: TestElements{
				conditions:   []string{`obj.field == 123`},
				errorMessage: "field must be equal to 123",
			},
			wantErr: false,
		},
		{
			name: "field != 64",
			args: args{
				fieldValidation: "neq=64",
			},
			want: TestElements{
				conditions:   []string{`obj.field != 64`},
				errorMessage: "field must not be equal to 64",
			},
			wantErr: false,
		},
		{
			name: "field > 10",
			args: args{
				fieldValidation: "gt=10",
			},
			want: TestElements{
				conditions:   []string{`obj.field > 10`},
				errorMessage: "field must be > 10",
			},
			wantErr: false,
		},
		{
			name: "field >= 0",
			args: args{
				fieldValidation: "gte=0",
			},
			want: TestElements{
				conditions:   []string{`obj.field >= 0`},
				errorMessage: "field must be >= 0",
			},
			wantErr: false,
		},
		{
			name: "field < 100",
			args: args{
				fieldValidation: "lt=100",
			},
			want: TestElements{
				conditions:   []string{`obj.field < 100`},
				errorMessage: "field must be < 100",
			},
			wantErr: false,
		},
		{
			name: "field <= 130",
			args: args{
				fieldValidation: "lte=130",
			},
			want: TestElements{
				conditions:   []string{`obj.field <= 130`},
				errorMessage: "field must be <= 130",
			},
			wantErr: false,
		},
		{
			name: "field in 10 20 30",
			args: args{
				fieldValidation: "in=10,20,30",
			},
			want: TestElements{
				conditions:     []string{`obj.field == 10`, `obj.field == 20`, `obj.field == 30`},
				concatOperator: "||",
				errorMessage:   "field must be one of '10' '20' '30'",
			},
			wantErr: false,
		},
		{
			name: "field not in 10 20 30",
			args: args{
				fieldValidation: "nin=10,20,30",
			},
			want: TestElements{
				conditions:     []string{`obj.field != 10`, `obj.field != 20`, `obj.field != 30`},
				concatOperator: "&&",
				errorMessage:   "field must not be one of '10' '20' '30'",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		basicTypes := common.HelperFromNormalizedToBasicTypes("<INT>")
		for _, fieldType := range basicTypes {
			testName := fmt.Sprintf("%s with %s", tt.name, fieldType)
			t.Run(testName, func(t *testing.T) {
				validation := AssertParserValidation(t, tt.args.fieldValidation)
				got, err := DefineTestElements("field", common.FieldType{BaseType: fieldType}, validation)
				if (err != nil) != tt.wantErr {
					t.Errorf("DefineTestElements() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("DefineTestElements() = %+v, want %+v", got, tt.want)
				}
			})
		}
	}
}
