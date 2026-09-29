package analyzer

import (
	"fmt"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
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
			name:  "valid eqfield between int",
			fType: "int",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between uint8",
			fType: "uint8",
			op:    "neqfield",
		},
		{
			name:  "valid neqfield between int",
			fType: "int",
			op:    "neqfield",
		},
		{
			name:  "valid gtefield between uint8",
			fType: "uint8",
			op:    "gtefield",
		},
		{
			name:  "valid gtefield between int",
			fType: "int",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between uint8",
			fType: "uint8",
			op:    "gtfield",
		},
		{
			name:  "valid gtfield between int",
			fType: "int",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between uint8",
			fType: "uint8",
			op:    "ltefield",
		},
		{
			name:  "valid ltefield between int",
			fType: "int",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between uint8",
			fType: "uint8",
			op:    "ltfield",
		},
		{
			name:  "valid ltfield between int",
			fType: "int",
			op:    "ltfield",
		},
		{
			name:  "valid eqfield between float32",
			fType: "float32",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between float32",
			fType: "float32",
			op:    "neqfield",
		},
		{
			name:  "valid gtefield between float32",
			fType: "float32",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between float32",
			fType: "float32",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between float32",
			fType: "float32",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between float32",
			fType: "float32",
			op:    "ltfield",
		},
		{
			name:  "valid eqfield between float64",
			fType: "float64",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between float64",
			fType: "float64",
			op:    "neqfield",
		},
		{
			name:  "valid gtefield between float64",
			fType: "float64",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between float64",
			fType: "float64",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between float64",
			fType: "float64",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between float64",
			fType: "float64",
			op:    "ltfield",
		},
		{
			name:  "valid eqfield between complex64",
			fType: "complex64",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between complex64",
			fType: "complex64",
			op:    "neqfield",
		},
		{
			name:  "valid eqfield between complex128",
			fType: "complex128",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between complex128",
			fType: "complex128",
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
							Type:      common.FieldType{BaseType: tt.fType},
							Tag:       fmt.Sprintf(`valid:"%s=Field2"`, tt.op),
						},
						{
							FieldName: "Field2",
							Type:      common.FieldType{BaseType: tt.fType},
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
		{
			name:  "valid eqfield between nested float32",
			fType: "float32",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between nested float32",
			fType: "float32",
			op:    "neqfield",
		},
		{
			name:  "valid gtefield between nested float32",
			fType: "float32",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between nested float32",
			fType: "float32",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between nested float32",
			fType: "float32",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between nested float32",
			fType: "float32",
			op:    "ltfield",
		},
		{
			name:  "valid eqfield between nested float64",
			fType: "float64",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between nested float64",
			fType: "float64",
			op:    "neqfield",
		},
		{
			name:  "valid gtefield between nested float64",
			fType: "float64",
			op:    "gtefield",
		},
		{
			name:  "valid gtfield between nested float64",
			fType: "float64",
			op:    "gtfield",
		},
		{
			name:  "valid ltefield between nested float64",
			fType: "float64",
			op:    "ltefield",
		},
		{
			name:  "valid ltfield between nested float64",
			fType: "float64",
			op:    "ltfield",
		},
		{
			name:  "valid eqfield between nested complex64",
			fType: "complex64",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between nested complex64",
			fType: "complex64",
			op:    "neqfield",
		},
		{
			name:  "valid eqfield between nested complex128",
			fType: "complex128",
			op:    "eqfield",
		},
		{
			name:  "valid neqfield between nested complex128",
			fType: "complex128",
			op:    "neqfield",
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
							Type:      common.FieldType{BaseType: tt.fType},
							Tag:       fmt.Sprintf(`valid:"%s=Nested.Field2"`, tt.op),
						},
						{
							FieldName: "Nested",
							Type:      common.FieldType{BaseType: "main.NestedStruct"},
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
							Type:      common.FieldType{BaseType: tt.fType},
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
						Type:      common.FieldType{BaseType: "string"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "uint8"},
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
						Type:      common.FieldType{BaseType: "string"},
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
						Type:      common.FieldType{BaseType: "string"},
						Tag:       `valid:"ltfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "string"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation ltfield: invalid string(<STRING>) type"),
		},
		{
			name: "mismatched float32 and float64",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "float32"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "float64"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: mismatched types between Field1 and Field2"),
		},
		{
			name: "mismatched float64 and int",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "float64"},
						Tag:       `valid:"gtfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "int"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation gtfield: mismatched types between Field1 and Field2"),
		},
		{
			name: "unsupported eqfield on float slice",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{ComposedType: "[]", BaseType: "float64"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{ComposedType: "[]", BaseType: "float64"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: invalid float64([]<FLOAT>) type"),
		},
		{
			name: "mismatched complex64 and complex128",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "complex64"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "complex128"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: mismatched types between Field1 and Field2"),
		},
		{
			name: "mismatched complex128 and float64",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "complex128"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "float64"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: mismatched types between Field1 and Field2"),
		},
		{
			name: "unsupported gtfield on complex128",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "complex128"},
						Tag:       `valid:"gtfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "complex128"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation gtfield: invalid complex128(<COMPLEX>) type"),
		},
		{
			name: "unsupported gtefield on complex64",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "complex64"},
						Tag:       `valid:"gtefield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "complex64"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation gtefield: invalid complex64(<COMPLEX>) type"),
		},
		{
			name: "unsupported ltefield on complex128",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "complex128"},
						Tag:       `valid:"ltefield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "complex128"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation ltefield: invalid complex128(<COMPLEX>) type"),
		},
		{
			name: "unsupported ltfield on complex64",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{BaseType: "complex64"},
						Tag:       `valid:"ltfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{BaseType: "complex64"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation ltfield: invalid complex64(<COMPLEX>) type"),
		},
		{
			name: "unsupported eqfield on complex slice",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{ComposedType: "[]", BaseType: "complex128"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{ComposedType: "[]", BaseType: "complex128"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: invalid complex128([]<COMPLEX>) type"),
		},
		{
			name: "unsupported eqfield on complex map",
			arg: &parser.Struct{
				Fields: []parser.Field{
					{
						FieldName: "Field1",
						Type:      common.FieldType{ComposedType: "map", BaseType: "complex64"},
						Tag:       `valid:"eqfield=Field2"`,
					},
					{
						FieldName: "Field2",
						Type:      common.FieldType{ComposedType: "map", BaseType: "complex64"},
						Tag:       ``,
					},
				},
			},
			wantErr: types.NewValidationError("operation eqfield: invalid complex64(map[<COMPLEX>]) type"),
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
							Type:      common.FieldType{BaseType: "string"},
							Tag:       `valid:"eqfield=Nested.Field2"`,
						},
						{
							FieldName: "Nested",
							Type:      common.FieldType{BaseType: "main.NestedStruct"},
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
							Type:      common.FieldType{BaseType: "uint8"},
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
							Type:      common.FieldType{BaseType: "string"},
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
							Type:      common.FieldType{BaseType: "string"},
							Tag:       `valid:"ltfield=Nested.Field2"`,
						},
						{
							FieldName: "Nested",
							Type:      common.FieldType{BaseType: "main.NestedStruct"},
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
							Type:      common.FieldType{BaseType: "string"},
							Tag:       ``,
						},
					},
				},
			},
			wantErr: types.NewValidationError("operation ltfield: invalid string(<STRING>) type"),
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

func TestAnalyzeComplexScalarOperations(t *testing.T) {
	validTags := []string{
		`valid:"required"`,
		`valid:"eq=1+2i"`,
		`valid:"neq=3-4i"`,
		`valid:"in=1+2i 5+6i"`,
		`valid:"nin=7+8i 9+0i"`,
	}
	invalidTags := []struct {
		tag string
		op  string
	}{
		{tag: `valid:"gt=1+2i"`, op: "gt"},
		{tag: `valid:"gte=1+2i"`, op: "gte"},
		{tag: `valid:"lt=1+2i"`, op: "lt"},
		{tag: `valid:"lte=1+2i"`, op: "lte"},
	}

	for _, baseType := range []string{"complex64", "complex128"} {
		for _, tag := range validTags {
			t.Run(baseType+" "+tag, func(t *testing.T) {
				_, err := AnalyzeStructs([]*parser.Struct{{
					Fields: []parser.Field{{
						FieldName: "Value",
						Type:      common.FieldType{BaseType: baseType},
						Tag:       tag,
					}},
				}})
				if err != nil {
					t.Errorf("AnalyzeStructs() error = %v", err)
				}
			})
		}

		for _, tt := range invalidTags {
			t.Run(baseType+" "+tt.tag, func(t *testing.T) {
				_, err := AnalyzeStructs([]*parser.Struct{{
					Fields: []parser.Field{{
						FieldName: "Value",
						Type:      common.FieldType{BaseType: baseType},
						Tag:       tt.tag,
					}},
				}})
				wantErr := types.NewValidationError("operation %s: invalid %s(<COMPLEX>) type", tt.op, baseType)
				if err != wantErr {
					t.Errorf("AnalyzeStructs() error = %v, wantErr %v", err, wantErr)
				}
			})
		}
	}

	composed := []struct {
		name     string
		composed string
		norm     string
		tag      string
		op       string
	}{
		{name: "slice", composed: "[]", norm: "[]<COMPLEX>", tag: `valid:"eq=1+2i"`, op: "eq"},
		{name: "map", composed: "map", norm: "map[<COMPLEX>]", tag: `valid:"eq=1+2i"`, op: "eq"},
		{name: "pointer", composed: "*", norm: "*<COMPLEX>", tag: `valid:"eq=1+2i"`, op: "eq"},
		{name: "pointer required", composed: "*", norm: "*<COMPLEX>", tag: `valid:"required"`, op: "required"},
	}
	for _, shape := range composed {
		t.Run(shape.name, func(t *testing.T) {
			_, err := AnalyzeStructs([]*parser.Struct{{
				Fields: []parser.Field{{
					FieldName: "Value",
					Type:      common.FieldType{BaseType: "complex128", ComposedType: shape.composed},
					Tag:       shape.tag,
				}},
			}})
			wantErr := types.NewValidationError("operation %s: invalid complex128(%s) type", shape.op, shape.norm)
			if err != wantErr {
				t.Errorf("AnalyzeStructs() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}
