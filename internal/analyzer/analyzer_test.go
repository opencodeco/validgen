package analyzer

import (
	"fmt"
	"testing"

	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/types"
)

func TestAnalyzeStructsWithValidFieldOperations(t *testing.T) {
	tests := []struct {
		name  string
		fType string
		op    string
	}{
		{
			name:  "valid eqfield between strings",
			fType: "string",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between strings",
			fType: "string",
			op:    "neqfield",
		},
		{
			name:  "valid eqfield between uint8",
			fType: "uint8",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between uint8",
			fType: "uint8",
			op:    "neqfield",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := []*parser.Struct{
				{
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      tt.fType,
							Tag:       fmt.Sprintf(`valid:"%s=Field2"`, tt.op),
						},
						{
							FieldName: "Field2",
							Type:      tt.fType,
							Tag:       ``,
						},
					},
				},
			}

			_, err := AnalyzeStructs(arg)
			if err != nil {
				t.Errorf("AnalyzeStructs() error = %v, wantErr %v", err, nil)
				return
			}
		})
	}
}

func TestAnalyzeStructsWithInvalidFieldOperations(t *testing.T) {
	tests := []struct {
		name    string
		arg     []*parser.Struct
		wantErr error
	}{
		{
			name: "mismatched types between fields",
			arg: []*parser.Struct{
				{
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       `valid:"eqfield=Field2"`,
						},
						{
							FieldName: "Field2",
							Type:      "uint8",
							Tag:       ``,
						},
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: mismatched types between Field1 and Field2"),
		},
		{
			name: "undefined field",
			arg: []*parser.Struct{
				{
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       `valid:"eqfield=Field2"`,
						},
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: undefined field Field2"),
		},
		{
			name: "invalid operation to string type",
			arg: []*parser.Struct{
				{
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       `valid:"ltfield=Field2"`,
						},
						{
							FieldName: "Field2",
							Type:      "uint8",
							Tag:       ``,
						},
					},
				},
			},
			wantErr: types.NewValidationError("invalid operation ltfield to string type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AnalyzeStructs(tt.arg)
			if err != tt.wantErr {
				t.Errorf("AnalyzeStructs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
