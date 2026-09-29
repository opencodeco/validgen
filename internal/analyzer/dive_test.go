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
			name: "missing endkeys",
			field: parser.Field{
				FieldName: "Labels",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue:     &common.FieldType{BaseType: "string"},
				},
				Tag: `valid:"dive,keys,required"`,
			},
			wantErr: types.NewValidationError("operation keys: field Labels: missing endkeys"),
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

func TestAnalyzeMapKeys(t *testing.T) {
	stringMap := common.FieldType{
		BaseType:     "string",
		ComposedType: "map",
		MapValue:     &common.FieldType{BaseType: "string"},
	}
	uintMap := common.FieldType{
		BaseType:     "uint8",
		ComposedType: "map",
		MapValue:     &common.FieldType{BaseType: "string"},
	}
	tests := []struct {
		name    string
		field   parser.Field
		wantErr error
	}{
		{
			name: "string keys and values",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"dive,keys,min=2,endkeys,required"`,
			},
		},
		{
			name: "typed keys and values",
			field: parser.Field{
				FieldName: "Scores",
				Type:      uintMap,
				Tag:       `valid:"dive,keys,gte=1,endkeys,required"`,
			},
		},
		{
			name: "value validation after endkeys",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"required,dive,keys,min=2,endkeys,email"`,
			},
		},
		{
			name: "dive into values after endkeys",
			field: parser.Field{
				FieldName: "Labels",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue:     &common.FieldType{BaseType: "string", ComposedType: "[]"},
				},
				Tag: `valid:"dive,keys,min=2,endkeys,dive,required"`,
			},
		},
		{
			name: "plain dive still validates map values",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"dive,required"`,
			},
		},
		{
			name: "extra endkeys",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"dive,keys,min=2,endkeys,required,endkeys"`,
			},
			wantErr: types.NewValidationError("operation endkeys: field Labels: extra endkeys"),
		},
		{
			name: "endkeys without keys",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"dive,endkeys,required"`,
			},
			wantErr: types.NewValidationError("operation endkeys: field Labels: endkeys without keys"),
		},
		{
			name: "keys does not follow dive",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"dive,required,keys,min=2,endkeys"`,
			},
			wantErr: types.NewValidationError("operation keys: field Labels: keys must immediately follow dive"),
		},
		{
			name: "keys on a slice",
			field: parser.Field{
				FieldName: "Names",
				Type:      common.FieldType{BaseType: "string", ComposedType: "[]"},
				Tag:       `valid:"dive,keys,min=2,endkeys,required"`,
			},
			wantErr: types.NewValidationError("operation keys: field Names: []string is not a map"),
		},
		{
			name: "keys on a string",
			field: parser.Field{
				FieldName: "Name",
				Type:      common.FieldType{BaseType: "string"},
				Tag:       `valid:"dive,keys,min=2,endkeys"`,
			},
			wantErr: types.NewValidationError("operation keys: field Name: string is not a map"),
		},
		{
			name: "dive inside keys",
			field: parser.Field{
				FieldName: "Labels",
				Type:      stringMap,
				Tag:       `valid:"dive,keys,dive,min=2,endkeys,required"`,
			},
			wantErr: types.NewValidationError("operation dive: field Labels: nested dive inside keys is not supported"),
		},
		{
			name: "nested keys",
			field: parser.Field{
				FieldName: "Labels",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue: &common.FieldType{
						BaseType:     "string",
						ComposedType: "map",
						MapValue:     &common.FieldType{BaseType: "string"},
					},
				},
				Tag: `valid:"dive,keys,min=2,endkeys,dive,keys,min=2,endkeys,required"`,
			},
			wantErr: types.NewValidationError("operation keys: field Labels: nested keys are not supported"),
		},
		{
			name: "array key",
			field: parser.Field{
				FieldName: "Labels",
				Type: common.FieldType{
					BaseType:     "string",
					ComposedType: "map[N]",
					Size:         "2",
					MapValue:     &common.FieldType{BaseType: "string"},
				},
				Tag: `valid:"dive,keys,min=2,endkeys,required"`,
			},
			wantErr: types.NewValidationError("operation keys: field Labels: nested map keys are not supported"),
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
