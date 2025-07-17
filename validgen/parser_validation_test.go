package validgen

import (
	"errors"
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/types"
)

func Test_ValidParserValidation(t *testing.T) {
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
				ExpectedValues: ZERO_VALUE,
				Values:         []string{},
			},
		},
		{
			name:       "tag with value",
			validation: "eq=abc",
			want: &Validation{
				Operation:      "eq",
				ExpectedValues: ONE_VALUE,
				Values:         []string{"abc"},
			},
		},
		{
			name:       "tag with multivalue (a b c)",
			validation: "in=a b c",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{"a", "b", "c"},
			},
		},
		{
			name:       "tag with multivalue ('abc')",
			validation: "in='abc'",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{"abc"},
			},
		},
		{
			name:       "tag with multivalue ('a' 'b' 'c')",
			validation: "in='a' 'b' 'c'",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{"a", "b", "c"},
			},
		},
		{
			name:       "tag with multivalue ('a ' 'b ' 'c ')",
			validation: "in='a ' 'b ' 'c '",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{"a ", "b ", "c "},
			},
		},
		{
			name:       "tag with multivalue (' a ' ' b ' ' c ')",
			validation: "in=' a ' ' b ' ' c '",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{" a ", " b ", " c "},
			},
		},
		{
			name:       "tag with multivalue ('a b c')",
			validation: "in='a b c'",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{"a b c"},
			},
		},
		{
			name:       "tag with multivalue (  a  b  c  )",
			validation: "in=  a  b  c  ",
			want: &Validation{
				Operation:      "in",
				ExpectedValues: MANY_VALUES,
				Values:         []string{"a", "b", "c"},
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

func Test_ParserInvalidValidation(t *testing.T) {
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
