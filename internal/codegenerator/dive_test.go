package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
)

func TestBuildDiveValidationCode(t *testing.T) {
	address := common.FieldType{BaseType: "main.Address"}
	tests := []struct {
		name            string
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
		want            string
	}{
		{
			name:            "slice of structs",
			fieldName:       "Addresses",
			fieldType:       common.FieldType{BaseType: "main.Address", ComposedType: "[]"},
			fieldValidation: "dive",
			want: `for _, elem1 := range obj.Addresses {
errs = append(errs, AddressValidate(&elem1)...)
}
`,
		},
		{
			name:            "slice of struct pointers",
			fieldName:       "Addresses",
			fieldType:       common.FieldType{BaseType: "main.Address", ComposedType: "*[]", ElemPointer: true},
			fieldValidation: "required,dive,required",
			want: `if !(len(obj.Addresses) != 0) {
errs = append(errs, types.NewValidationError("Addresses must not be empty"))
}
for _, elem1 := range obj.Addresses {
if !(elem1 != nil) {
errs = append(errs, types.NewValidationError("Addresses is required"))
}
if elem1 != nil {
errs = append(errs, AddressValidate(elem1)...)
}
}
`,
		},
		{
			name:      "map value",
			fieldName: "Labels",
			fieldType: common.FieldType{
				BaseType:     "string",
				ComposedType: "map",
				MapValue:     &common.FieldType{BaseType: "string"},
			},
			fieldValidation: "dive,required",
			want: `for _, elem1 := range obj.Labels {
if !(elem1 != "") {
errs = append(errs, types.NewValidationError("Labels is required"))
}
}
`,
		},
		{
			name:            "slice of structs does not dive",
			fieldName:       "Addresses",
			fieldType:       common.FieldType{BaseType: "main.Address", ComposedType: "[]"},
			fieldValidation: "required",
			want: `if !(len(obj.Addresses) != 0) {
errs = append(errs, types.NewValidationError("Addresses must not be empty"))
}
`,
		},
		{
			name:            "pointer to a slice of structs",
			fieldName:       "Addresses",
			fieldType:       common.FieldType{BaseType: "main.Address", ComposedType: "*[]"},
			fieldValidation: "dive",
			want: `if obj.Addresses != nil {
for _, elem1 := range *obj.Addresses {
errs = append(errs, AddressValidate(&elem1)...)
}
}
`,
		},
		{
			name:            "nested string slice",
			fieldName:       "Matrix",
			fieldType:       common.FieldType{BaseType: "string", ComposedType: "[][]"},
			fieldValidation: "dive,dive,required",
			want: `for _, elem1 := range obj.Matrix {
for _, elem2 := range elem1 {
if !(elem2 != "") {
errs = append(errs, types.NewValidationError("Matrix is required"))
}
}
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := GenValidations{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{PackageName: "main"},
				},
				StructsWithValidation: map[string]struct{}{
					address.BaseType: {},
				},
			}
			var validations []*analyzer.Validation
			for _, part := range splitValidations(tt.fieldValidation) {
				validations = append(validations, AssertParserValidation(t, part))
			}
			got, err := gv.BuildValidationCode(tt.fieldName, tt.fieldType, validations)
			if err != nil {
				t.Fatalf("BuildValidationCode() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("BuildValidationCode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildMapKeyValidationCode(t *testing.T) {
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
		name            string
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
		want            string
		wantErr         string
	}{
		{
			name:            "string keys and values",
			fieldName:       "Labels",
			fieldType:       stringMap,
			fieldValidation: "dive,keys,min=2,endkeys,required",
			want: `for key1, elem1 := range obj.Labels {
if !(len(key1) >= 2) {
errs = append(errs, types.NewValidationError("Labels length must be >= 2"))
}
if !(elem1 != "") {
errs = append(errs, types.NewValidationError("Labels is required"))
}
}
`,
		},
		{
			name:            "typed keys and values",
			fieldName:       "Scores",
			fieldType:       uintMap,
			fieldValidation: "dive,keys,gte=1,endkeys,required",
			want: `for key1, elem1 := range obj.Scores {
if !(key1 >= 1) {
errs = append(errs, types.NewValidationError("Scores must be >= 1"))
}
if !(elem1 != "") {
errs = append(errs, types.NewValidationError("Scores is required"))
}
}
`,
		},
		{
			name:            "value validation after endkeys",
			fieldName:       "Labels",
			fieldType:       stringMap,
			fieldValidation: "required,dive,keys,min=2,endkeys,email",
			want: `if !(len(obj.Labels) != 0) {
errs = append(errs, types.NewValidationError("Labels must not be empty"))
}
for key1, elem1 := range obj.Labels {
if !(len(key1) >= 2) {
errs = append(errs, types.NewValidationError("Labels length must be >= 2"))
}
if !(types.IsValidEmail(elem1)) {
errs = append(errs, types.NewValidationError("Labels must be a valid email"))
}
}
`,
		},
		{
			name:      "dive into values after endkeys",
			fieldName: "Labels",
			fieldType: common.FieldType{
				BaseType:     "string",
				ComposedType: "map",
				MapValue:     &common.FieldType{BaseType: "string", ComposedType: "[]"},
			},
			fieldValidation: "dive,keys,min=2,endkeys,dive,required",
			want: `for key1, elem1 := range obj.Labels {
if !(len(key1) >= 2) {
errs = append(errs, types.NewValidationError("Labels length must be >= 2"))
}
for _, elem2 := range elem1 {
if !(elem2 != "") {
errs = append(errs, types.NewValidationError("Labels is required"))
}
}
}
`,
		},
		{
			name:            "keys only",
			fieldName:       "Labels",
			fieldType:       stringMap,
			fieldValidation: "dive,keys,min=2,endkeys",
			want: `for key1 := range obj.Labels {
if !(len(key1) >= 2) {
errs = append(errs, types.NewValidationError("Labels length must be >= 2"))
}
}
`,
		},
		{
			name:            "pointer to a map",
			fieldName:       "Labels",
			fieldType:       common.FieldType{BaseType: "string", ComposedType: "*map", MapValue: &common.FieldType{BaseType: "string"}},
			fieldValidation: "dive,keys,min=2,endkeys,required",
			want: `if obj.Labels != nil {
for key1, elem1 := range *obj.Labels {
if !(len(key1) >= 2) {
errs = append(errs, types.NewValidationError("Labels length must be >= 2"))
}
if !(elem1 != "") {
errs = append(errs, types.NewValidationError("Labels is required"))
}
}
}
`,
		},
		{
			name:            "missing endkeys",
			fieldName:       "Labels",
			fieldType:       stringMap,
			fieldValidation: "dive,keys,required",
			wantErr:         "operation keys: field Labels: missing endkeys",
		},
		{
			name:            "extra endkeys",
			fieldName:       "Labels",
			fieldType:       stringMap,
			fieldValidation: "dive,keys,min=2,endkeys,endkeys",
			wantErr:         "operation endkeys: field Labels: extra endkeys",
		},
		{
			name:            "keys on a slice",
			fieldName:       "Names",
			fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
			fieldValidation: "dive,keys,min=2,endkeys,required",
			wantErr:         "operation keys: field Names: []string is not a map",
		},
		{
			name:      "nested keys",
			fieldName: "Labels",
			fieldType: common.FieldType{
				BaseType:     "string",
				ComposedType: "map",
				MapValue: &common.FieldType{
					BaseType:     "string",
					ComposedType: "map",
					MapValue:     &common.FieldType{BaseType: "string"},
				},
			},
			fieldValidation: "dive,keys,min=2,endkeys,dive,keys,min=2,endkeys,required",
			wantErr:         "operation keys: field Labels: nested keys are not supported",
		},
		{
			name:            "dive inside keys",
			fieldName:       "Labels",
			fieldType:       stringMap,
			fieldValidation: "dive,keys,dive,min=2,endkeys,required",
			wantErr:         "operation dive: field Labels: nested dive inside keys is not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := GenValidations{
				Struct: &analyzer.Struct{
					Struct: parser.Struct{PackageName: "main"},
				},
			}
			var validations []*analyzer.Validation
			for _, part := range splitValidations(tt.fieldValidation) {
				validations = append(validations, AssertParserValidation(t, part))
			}
			got, err := gv.BuildValidationCode(tt.fieldName, tt.fieldType, validations)
			if tt.wantErr != "" {
				if err == nil || err.Error() != tt.wantErr {
					t.Fatalf("BuildValidationCode() error = %v, want %s", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("BuildValidationCode() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("BuildValidationCode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func splitValidations(tag string) []string {
	if tag == "" {
		return nil
	}
	parts := []string{}
	start := 0
	for i := 0; i <= len(tag); i++ {
		if i == len(tag) || tag[i] == ',' {
			parts = append(parts, tag[start:i])
			start = i + 1
		}
	}
	return parts
}
