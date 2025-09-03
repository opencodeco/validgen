package analyzer

import (
	"fmt"
	"testing"

	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/types"
)

func TestAnalyzeStructsWithValidInnerFieldOperations(t *testing.T) {
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
		{
			name:  "valid gtefield between uint8",
			fType: "uint8",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between uint8",
			fType: "uint8",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between uint8",
			fType: "uint8",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between uint8",
			fType: "uint8",
			op:    "ltfield",
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

func TestAnalyzeStructsWithValidNestedFieldOperations(t *testing.T) {
	tests := []struct {
		name  string
		fType string
		op    string
	}{
		{
			name:  "valid eqfield between nested strings",
			fType: "string",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between nested strings",
			fType: "string",
			op:    "neqfield",
		},
		{
			name:  "valid eqfield between nested uint8",
			fType: "uint8",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between nested uint8",
			fType: "uint8",
			op:    "neqfield",
		},
		{
			name:  "valid gtefield between nested uint8",
			fType: "uint8",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between nested uint8",
			fType: "uint8",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between nested uint8",
			fType: "uint8",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between nested uint8",
			fType: "uint8",
			op:    "ltfield",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := []*parser.Struct{
				{
					PackageName: "main",
					StructName:  "MyStruct",
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      tt.fType,
							Tag:       fmt.Sprintf(`valid:"%s=Nested.Field2"`, tt.op),
						},
						{
							FieldName: "Nested",
							Type:      "main.NestedStruct",
							Tag:       ``,
						},
					},
				},
				{
					PackageName: "main",
					StructName:  "NestedStruct",
					Fields: []parser.Field{
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

func TestAnalyzeStructsWithInvalidInnerFieldOperations(t *testing.T) {
	tests := []struct {
		name    string
		arg     *parser.Struct
		wantErr error
	}{
		{
			name: "mismatched types between inner fields",
			arg: &parser.Struct{
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
			wantErr: types.NewValidationError("operation eqfield: mismatched types between Field1 and Field2"),
		},
		{
			name: "undefined inner field",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      "string",
						Tag:       `valid:"eqfield=Field2"`,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: undefined field Field2"),
		},
		{
			name: "invalid inner operation for string type",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      "string",
						Tag:       `valid:"ltfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      "string",
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation ltfield: invalid string type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AnalyzeStructs([]*parser.Struct{tt.arg})
			if err != tt.wantErr {
				t.Errorf("AnalyzeStructs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAnalyzeStructsWithInvalidNestedFieldOperations(t *testing.T) {
	tests := []struct {
		name    string
		arg     []*parser.Struct
		wantErr error
	}{
		{
			name: "mismatched types between nested fields",
			arg: []*parser.Struct{
				{
					PackageName: "main",
					StructName:  "Struct",
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       `valid:"eqfield=Nested.Field2"`,
						},
						{
							FieldName: "Nested",
							Type:      "main.NestedStruct",
							Tag:       ``,
						},
					},
				},
				{
					PackageName: "main",
					StructName:  "NestedStruct",
					Fields: []parser.Field{
						{
							FieldName: "Field2",
							Type:      "uint8",
							Tag:       ``,
						},
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: mismatched types between Field1 and Nested.Field2"),
		},
		{
			name: "undefined nested field",
			arg: []*parser.Struct{
				{
					PackageName: "main",
					StructName:  "Struct",
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       `valid:"eqfield=Nested.Field2"`,
						},
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: undefined nested field Nested"),
		},
		{
			name: "invalid operation to nested string type",
			arg: []*parser.Struct{
				{
					PackageName: "main",
					StructName:  "Struct",
					Fields: []parser.Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       `valid:"ltfield=Nested.Field2"`,
						},
						{
							FieldName: "Nested",
							Type:      "main.NestedStruct",
							Tag:       ``,
						},
					},
				},
				{
					PackageName: "main",
					StructName:  "NestedStruct",
					Fields: []parser.Field{
						{
							FieldName: "Field2",
							Type:      "string",
							Tag:       ``,
						},
					},
				},
			},
			wantErr: types.NewValidationError("operation ltfield: invalid string type"),
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
