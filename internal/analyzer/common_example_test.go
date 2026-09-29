package analyzer

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
)

func TestAnalyzeCommonValidatorExample(t *testing.T) {
	structs := commonExampleStructs()
	got, err := AnalyzeStructs(structs)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("struct count = %d, want 2", len(got))
	}
	if !got[0].HasValidTag || !got[1].HasValidTag {
		t.Fatal("example structs should be validated")
	}

	wantUser := [][]string{
		{"required"},
		{"required"},
		{"gte", "lte"},
		{"required", "email"},
		{"oneof"},
		{"iscolor"},
		{"required", "dive", "required"},
	}
	if ops := operationNames(got[0]); !reflect.DeepEqual(ops, wantUser) {
		t.Fatalf("User operations = %#v, want %#v", ops, wantUser)
	}

	gender := got[0].FieldsValidations[4].Validations[0]
	wantValues := []string{"male", "female", "prefer_not_to"}
	if !reflect.DeepEqual(gender.Values, wantValues) {
		t.Fatalf("oneof values = %#v, want %#v", gender.Values, wantValues)
	}

	wantAddress := [][]string{
		{"required"},
		{"required"},
		{"required"},
		{"required"},
	}
	if ops := operationNames(got[1]); !reflect.DeepEqual(ops, wantAddress) {
		t.Fatalf("Address operations = %#v, want %#v", ops, wantAddress)
	}
}

func TestValidationTagKey(t *testing.T) {
	stringField := common.FieldType{BaseType: "string"}
	tests := []struct {
		name string
		tag  string
		want []string
	}{
		{
			name: "valid tag",
			tag:  `valid:"required,email"`,
			want: []string{"required", "email"},
		},
		{
			name: "validate tag after json",
			tag:  `json:"email" validate:"required,email"`,
			want: []string{"required", "email"},
		},
		{
			name: "valid wins when both keys are present",
			tag:  `valid:"required" validate:"email"`,
			want: []string{"required"},
		},
		{
			name: "unrelated tag",
			tag:  `json:"email"`,
			want: nil,
		},
		{
			name: "empty validate value",
			tag:  `validate:""`,
			want: nil,
		},
		{
			name: "blank piece between commas",
			tag:  `valid:"required, ,email"`,
			want: []string{"required", "email"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnalyzeStructs([]*parser.Struct{{
				PackageName: "main",
				StructName:  "User",
				Fields: []parser.Field{{
					FieldName: "Email",
					Type:      stringField,
					Tag:       tt.tag,
				}},
			}})
			if err != nil {
				t.Fatal(err)
			}
			ops := operationNames(got[0])
			var flat []string
			if len(ops) == 1 {
				flat = ops[0]
			}
			if !reflect.DeepEqual(flat, tt.want) {
				t.Fatalf("operations = %#v, want %#v", flat, tt.want)
			}
		})
	}
}

func commonExampleStructs() []*parser.Struct {
	return []*parser.Struct{
		{
			PackageName: "main",
			StructName:  "User",
			Fields: []parser.Field{
				{FieldName: "FirstName", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
				{FieldName: "LastName", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
				{FieldName: "Age", Type: common.FieldType{BaseType: "uint8"}, Tag: `validate:"gte=0,lte=130"`},
				{FieldName: "Email", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required,email"`},
				{FieldName: "Gender", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"oneof=male female prefer_not_to"`},
				{FieldName: "FavouriteColor", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"iscolor"`},
				{
					FieldName: "Addresses",
					Type:      common.FieldType{BaseType: "main.Address", ComposedType: "*[]", ElemPointer: true},
					Tag:       `validate:"required,dive,required"`,
				},
			},
		},
		{
			PackageName: "main",
			StructName:  "Address",
			Fields: []parser.Field{
				{FieldName: "Street", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
				{FieldName: "City", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
				{FieldName: "Planet", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
				{FieldName: "Phone", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
			},
		},
	}
}

func operationNames(st *Struct) [][]string {
	names := make([][]string, len(st.FieldsValidations))
	for i, field := range st.FieldsValidations {
		for _, validation := range field.Validations {
			names[i] = append(names[i], validation.Operation)
		}
	}
	return names
}
