package codegenerator

import (
	"errors"
	"testing"

	"github.com/opencodeco/validgen/types"
)

func TestDefineTestElements(t *testing.T) {
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
			name: "invalid operation",
			args: args{
				fieldName:       "xpto",
				fieldType:       "string",
				fieldValidation: "xy=123",
			},
			expectedErr: types.NewValidationError("parser validation xy=123 type string unsupported validation xy"),
		},
		{
			name: "invalid string operation",
			args: args{
				fieldName:       "xpto",
				fieldType:       "string",
				fieldValidation: "lt=123",
			},
			expectedErr: types.NewValidationError("parser validation lt=123 type string unsupported validation lt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DefineTestElements(tt.args.fieldName, tt.args.fieldType, tt.args.fieldValidation)
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
