package analyzer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/types"
)

func TestValidParserValidation(t *testing.T) {
	tests := []struct {
		name       string
		validation string
		want       *Validation
	}{
		{
			name:       "tag without value",
			validation: "required",
			want: &Validation{
				Operation:      "required",
				ExpectedValues: common.ZeroValue,
				Values:         []string{},
			},
		},
		{
			name:       "tag with value",
			validation: "eq=abc",
			want: &Validation{
				Operation:      "eq",
				ExpectedValues: common.OneValue,
				Values:         []string{"abc"},
			},
		},
		{
			name:       "tag with multivalue (a b c)",
			validation: "in=a b c",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"a", "b", "c"},
			},
		},
		{
			name:       "tag with multivalue ('abc')",
			validation: "in='abc'",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"abc"},
			},
		},
		{
			name:       "tag with multivalue ('a' 'b' 'c')",
			validation: "in='a' 'b' 'c'",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"a", "b", "c"},
			},
		},
		{
			name:       "tag with multivalue ('a ' 'b ' 'c ')",
			validation: "in='a ' 'b ' 'c '",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"a ", "b ", "c "},
			},
		},
		{
			name:       "tag with multivalue (' a ' ' b ' ' c ')",
			validation: "in=' a ' ' b ' ' c '",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{" a ", " b ", " c "},
			},
		},
		{
			name:       "tag with multivalue ('a b c')",
			validation: "in='a b c'",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"a b c"},
			},
		},
		{
			name:       "tag with multivalue (  a  b  c  )",
			validation: "in=  a  b  c  ",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"a", "b", "c"},
			},
		},
		{
			name:       "tag with multivalue with comma (10,20,30)",
			validation: "in=10,20,30",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: common.ManyValues,
				Values:         []string{"10", "20", "30"},
			},
		},
		{
			name:       "email validation",
			validation: "email",
			want: &Validation{
				Operation:      "email",
				ExpectedValues: common.ZeroValue,
				Values:         []string{},
			},
		},
		{
			name:       "operation between inner fields",
			validation: "eqfield=field123",
			want: &Validation{
				Operation:      "eqfield",
				ExpectedValues: common.OneValue,
				Values:         []string{"field123"},
			},
		},
		{
			name:       "operation between nested fields",
			validation: "eqfield=Nested.field123",
			want: &Validation{
				Operation:      "eqfield",
				ExpectedValues: common.OneValue,
				Values:         []string{"Nested.field123"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := false
			got, err := ParserValidation(tt.validation)
			if (err != nil) != wantErr {
				t.Errorf("ParserValidation() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParserValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParserInvalidValidation(t *testing.T) {
	tests := []struct {
		name        string
		validation  string
		expectedErr error
	}{
		{
			name:        "tag without expected value",
			validation:  "eq",
			expectedErr: types.NewValidationError("expected one target, but has nothing"),
		},
		{
			name:        "tag without expected value",
			validation:  "eq=",
			expectedErr: types.NewValidationError("expected one target, but has nothing"),
		},
		{
			name:        "malformed tag",
			validation:  "eq=aaa=bbb",
			expectedErr: types.NewValidationError("malformed validation eq=aaa=bbb"),
		},
		{
			name:        "undefined validation",
			validation:  "xpto=a",
			expectedErr: types.NewValidationError("unsupported validation xpto"),
		},
		{
			name:        "malformed value",
			validation:  "in='abc",
			expectedErr: types.NewValidationError("invalid quote value in 'abc"),
		},
		{
			name:        "malformed value",
			validation:  "in='a ' b ' 'c '",
			expectedErr: types.NewValidationError("invalid quote value in 'a ' b ' 'c '"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParserValidation(tt.validation)
			var valErr types.ValidationError
			if !errors.As(err, &valErr) {
				t.Errorf("ParserValidation() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}

			if !errors.Is(valErr, tt.expectedErr) {
				t.Errorf("ParserValidation() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
		})
	}
}
