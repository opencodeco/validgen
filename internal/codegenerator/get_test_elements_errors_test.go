package codegenerator

import (
	"errors"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/types"
)

func TestDefineTestElementsWithInvalidOperations(t *testing.T) {
	type args struct {
		fieldName       string
		fieldType       common.FieldType
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
				fieldType:       common.FieldType{BaseType: "uint8"},
				fieldValidation: "min=1 2 3",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation min type <INT> (uint8)"),
		},
		{
			name: "unsupported eqfield on float slice",
			args: args{
				fieldName:       "values",
				fieldType:       common.FieldType{ComposedType: "[]", BaseType: "float32"},
				fieldValidation: "eqfield=other",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation eqfield type []<FLOAT> (float32)"),
		},
		{
			name: "unsupported gtfield on float map",
			args: args{
				fieldName:       "values",
				fieldType:       common.FieldType{ComposedType: "map", BaseType: "float64"},
				fieldValidation: "gtfield=other",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation gtfield type map[<FLOAT>] (float64)"),
		},
		{
			name: "unsupported gt on complex64",
			args: args{
				fieldName:       "value",
				fieldType:       common.FieldType{BaseType: "complex64"},
				fieldValidation: "gt=1+2i",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation gt type <COMPLEX> (complex64)"),
		},
		{
			name: "unsupported gte on complex128",
			args: args{
				fieldName:       "value",
				fieldType:       common.FieldType{BaseType: "complex128"},
				fieldValidation: "gte=1+2i",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation gte type <COMPLEX> (complex128)"),
		},
		{
			name: "unsupported lt on complex64",
			args: args{
				fieldName:       "value",
				fieldType:       common.FieldType{BaseType: "complex64"},
				fieldValidation: "lt=1+2i",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation lt type <COMPLEX> (complex64)"),
		},
		{
			name: "unsupported lte on complex128",
			args: args{
				fieldName:       "value",
				fieldType:       common.FieldType{BaseType: "complex128"},
				fieldValidation: "lte=1+2i",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation lte type <COMPLEX> (complex128)"),
		},
		{
			name: "unsupported gtfield on complex128",
			args: args{
				fieldName:       "values",
				fieldType:       common.FieldType{BaseType: "complex128"},
				fieldValidation: "gtfield=other",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation gtfield type <COMPLEX> (complex128)"),
		},
		{
			name: "unsupported eq on complex pointer",
			args: args{
				fieldName:       "value",
				fieldType:       common.FieldType{ComposedType: "*", BaseType: "complex64"},
				fieldValidation: "eq=1+2i",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation eq type *<COMPLEX> (complex64)"),
		},
		{
			name: "unsupported eq on complex slice",
			args: args{
				fieldName:       "value",
				fieldType:       common.FieldType{ComposedType: "[]", BaseType: "complex128"},
				fieldValidation: "eq=1+2i",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation eq type []<COMPLEX> (complex128)"),
		},
		{
			name: "unsupported eqfield on complex slice",
			args: args{
				fieldName:       "values",
				fieldType:       common.FieldType{ComposedType: "[]", BaseType: "complex64"},
				fieldValidation: "eqfield=other",
			},
			expectedErr: types.NewValidationError("INTERNAL ERROR: unsupported operation eqfield type []<COMPLEX> (complex64)"),
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
