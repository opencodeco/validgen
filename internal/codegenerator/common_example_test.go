package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestCommonValidatorExampleCode(t *testing.T) {
	analyzed, err := analyzer.AnalyzeStructs(commonExampleStructs())
	if err != nil {
		t.Fatal(err)
	}

	gv := GenValidations{
		StructsWithValidation: map[string]struct{}{
			"main.User":    {},
			"main.Address": {},
		},
		Struct: analyzed[0],
	}
	got, err := gv.BuildFuncValidatorCode()
	if err != nil {
		t.Fatal(err)
	}

	want := `func UserValidate(obj *User) []error {
var errs []error
if !(obj.FirstName != "") {
errs = append(errs, types.NewValidationError("FirstName is required"))
}
if !(obj.LastName != "") {
errs = append(errs, types.NewValidationError("LastName is required"))
}
if !(obj.Age >= 0) {
errs = append(errs, types.NewValidationError("Age must be >= 0"))
}
if !(obj.Age <= 130) {
errs = append(errs, types.NewValidationError("Age must be <= 130"))
}
if !(obj.Email != "") {
errs = append(errs, types.NewValidationError("Email is required"))
}
if !(types.IsValidEmail(obj.Email)) {
errs = append(errs, types.NewValidationError("Email must be a valid email"))
}
if !(obj.Gender == "male" || obj.Gender == "female" || obj.Gender == "prefer_not_to") {
errs = append(errs, types.NewValidationError("Gender must be one of 'male' 'female' 'prefer_not_to'"))
}
if !(types.IsColor(obj.FavouriteColor)) {
errs = append(errs, types.NewValidationError("FavouriteColor must be a valid color"))
}
if !(len(obj.Addresses) != 0) {
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
return errs
}
`
	if got != want {
		dmp := diffmatchpatch.New()
		t.Fatalf("UserValidate mismatch\n%s", dmp.DiffPrettyText(dmp.DiffMain(want, got, false)))
	}

	gv.Struct = analyzed[1]
	got, err = gv.BuildFuncValidatorCode()
	if err != nil {
		t.Fatal(err)
	}
	wantAddress := `func AddressValidate(obj *Address) []error {
var errs []error
if !(obj.Street != "") {
errs = append(errs, types.NewValidationError("Street is required"))
}
if !(obj.City != "") {
errs = append(errs, types.NewValidationError("City is required"))
}
if !(obj.Planet != "") {
errs = append(errs, types.NewValidationError("Planet is required"))
}
if !(obj.Phone != "") {
errs = append(errs, types.NewValidationError("Phone is required"))
}
return errs
}
`
	if got != wantAddress {
		dmp := diffmatchpatch.New()
		t.Fatalf("AddressValidate mismatch\n%s", dmp.DiffPrettyText(dmp.DiffMain(wantAddress, got, false)))
	}
}

func TestDiveOneofAndColor(t *testing.T) {
	tests := []struct {
		name            string
		fieldName       string
		fieldType       common.FieldType
		fieldValidation string
		want            string
	}{
		{
			name:            "dive iscolor",
			fieldName:       "Colors",
			fieldType:       common.FieldType{BaseType: "string", ComposedType: "[]"},
			fieldValidation: "dive,iscolor",
			want: `for _, elem1 := range obj.Colors {
if !(types.IsColor(elem1)) {
errs = append(errs, types.NewValidationError("Colors must be a valid color"))
}
}
`,
		},
		{
			name:      "dive oneof on map values",
			fieldName: "Labels",
			fieldType: common.FieldType{
				BaseType:     "string",
				ComposedType: "map",
				MapValue:     &common.FieldType{BaseType: "string"},
			},
			fieldValidation: "dive,oneof=red green blue",
			want: `for _, elem1 := range obj.Labels {
if !(elem1 == "red" || elem1 == "green" || elem1 == "blue") {
errs = append(errs, types.NewValidationError("Labels must be one of 'red' 'green' 'blue'"))
}
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gv := GenValidations{}
			var validations []*analyzer.Validation
			for _, part := range splitValidations(tt.fieldValidation) {
				validations = append(validations, AssertParserValidation(t, part))
			}
			got, err := gv.BuildValidationCode(tt.fieldName, tt.fieldType, validations)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("BuildValidationCode() = %q, want %q", got, tt.want)
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
