package codegenerator

import (
	"errors"
	"testing"

	"github.com/opencodeco/validgen/types"
)

func TestDefineTestElementsWithInvalidOperations(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       string
		fieldValidation string
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "invalid uint8 operation",
			args: args{
				fieldName:       "xpto",
				fieldType:       "uint8",
				fieldValidation: "in=1 2 3",
			},
			expectedErr: types.NewValidationError("unsupported operation in type uint8"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validation := AssertParserValidation(t, tt.args.fieldValidation)
			_, err := DefineTestElements(tt.args.fieldName, tt.args.fieldType, validation)
			var valErr types.ValidationError
			if !errors.As(err, &valErr) {
				t.Errorf("DefineTestElements() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}

			if !errors.Is(valErr, tt.expectedErr) {
				t.Errorf("DefineTestElements() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
		})
	}
}
