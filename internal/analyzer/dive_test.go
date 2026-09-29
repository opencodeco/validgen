package analyzer

import (
	"testing"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/types"
)

func addressStruct() *parser.Struct {
	return &parser.Struct{
		PackageName: "main",
		StructName:  "Address",
		Fields: []parser.Field{
			{
				FieldName: "Street",
				Type:      common.FieldType{BaseType: "string"},
				Tag:       `valid:"required"`,
			},
		},
	}
}

func TestAnalyzeDiveAccepted(t *testing.T) {
	tests := []struct {
		name  string
		field parser.Field
	}{
		{
			name: "slice of structs",
			field: parser.Field{
				FieldName: "Addresses",
				Type:      common.FieldType{BaseType: "main.Address", ComposedType: "[]"},
				Tag:       `valid:"required,dive"`,
			},
		},
		{
			name: "slice of struct pointers",
			field: parser.Field{
				FieldName: "Addresses",
				Type:      common.FieldType{BaseType: "main.Address", ComposedType: "*[]", ElemPointer: true},
				Tag:       `valid:"required,dive,required"`,
			},
		},
		{
			name: "map value",
			field: parser.Field{
				FieldName: "Labels",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue:     &common.FieldType{BaseType: "string"},
				},
				Tag: `valid:"dive,required"`,
			},
		},
		{
			name: "slice of structs does not dive",
			field: parser.Field{
				FieldName: "Addresses",
				Type:      common.FieldType{BaseType: "main.Address", ComposedType: "[]"},
				Tag:       `valid:"required"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &parser.Struct{
				PackageName: "main",
				StructName:  "User",
				Fields:      []parser.Field{tt.field},
			}
			if _, err := AnalyzeStructs([]*parser.Struct{addressStruct(), user}); err != nil {
				t.Fatalf("AnalyzeStructs() error = %v", err)
			}
		})
	}
}

func TestAnalyzeDiveRejected(t *testing.T) {
	tests := []struct {
		name    string
		field   parser.Field
		wantErr error
	}{
		{
			name: "dive on a string",
			field: parser.Field{
				FieldName: "Name",
				Type:      common.FieldType{BaseType: "string"},
				Tag:       `valid:"dive"`,
			},
			wantErr: types.NewValidationError("operation dive: field Name: string is not a slice, array, or map"),
		},
		{
			name: "email on a map value of int",
			field: parser.Field{
				FieldName: "Counts",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue:     &common.FieldType{BaseType: "int"},
				},
				Tag: `valid:"dive,email"`,
			},
			wantErr: types.NewValidationError("operation email: invalid int(<INT>) type"),
		},
		{
			name: "field comparison after dive",
			field: parser.Field{
				FieldName: "Names",
				Type:      common.FieldType{BaseType: "string", ComposedType: "[]"},
				Tag:       `valid:"dive,eqfield=Name"`,
			},
			wantErr: types.NewValidationError("operation eqfield: field comparisons are not supported after dive"),
		},
		{
			name: "nested arrays",
			field: parser.Field{
				FieldName: "Matrix",
				Type:      common.FieldType{BaseType: "int", ComposedType: "[N][N]", Size: "2"},
				Tag:       `valid:"dive"`,
			},
			wantErr: types.NewValidationError("operation dive: field Matrix: nested arrays are not supported"),
		},
		{
			name: "pointer composition the generator cannot express",
			field: parser.Field{
				FieldName: "Matrix",
				Type:      common.FieldType{BaseType: "string", ComposedType: "*[][]", ElemPointer: true},
				Tag:       `valid:"dive"`,
			},
			wantErr: types.NewValidationError("operation dive: field Matrix: unsupported pointer composition *[]string"),
		},
		{
			name: "keys stays unsupported",
			field: parser.Field{
				FieldName: "Labels",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue:     &common.FieldType{BaseType: "string"},
				},
				Tag: `valid:"dive,keys,required"`,
			},
			wantErr: types.NewValidationError("parser validation keys: unsupported validation keys"),
		},
		{
			name: "in on a slice of pointers",
			field: parser.Field{
				FieldName: "Names",
				Type:      common.FieldType{BaseType: "string", ComposedType: "*[]", ElemPointer: true},
				Tag:       `valid:"in=a b"`,
			},
			wantErr: types.NewValidationError("operation in: cannot apply to a slice or array of pointers"),
		},
		{
			name: "tag other than required on a struct element",
			field: parser.Field{
				FieldName: "Addresses",
				Type:      common.FieldType{BaseType: "main.Address", ComposedType: "[]"},
				Tag:       `valid:"dive,email"`,
			},
			wantErr: types.NewValidationError("operation email: cannot apply to struct main.Address"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &parser.Struct{
				PackageName: "main",
				StructName:  "User",
				Fields:      []parser.Field{tt.field},
			}
			_, err := AnalyzeStructs([]*parser.Struct{addressStruct(), user})
			if err != tt.wantErr {
				t.Fatalf("AnalyzeStructs() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
