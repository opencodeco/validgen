package validgen

import (
	"errors"
	"testing"

	"github.com/opencodeco/validgen/types"
)

func TestGetTestElements(t *testing.T) {
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
			_, err := GetTestElements(tt.args.fieldName, tt.args.fieldValidation, tt.args.fieldType)
			var valErr types.ValidationError
			if !errors.As(err, &valErr) {
				t.Errorf("GetTestElements() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}

			if !errors.Is(valErr, tt.expectedErr) {
				t.Errorf("GetTestElements() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
		})
	}
}
