package parser

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestParseCommonValidatorExample(t *testing.T) {
	src := "package main\n" +
		"type User struct {\n" +
		"	FirstName      string     `validate:\"required\"`\n" +
		"	LastName       string     `validate:\"required\"`\n" +
		"	Age            uint8      `validate:\"gte=0,lte=130\"`\n" +
		"	Email          string     `validate:\"required,email\"`\n" +
		"	Gender         string     `validate:\"oneof=male female prefer_not_to\"`\n" +
		"	FavouriteColor string     `validate:\"iscolor\"`\n" +
		"	Addresses      []*Address `validate:\"required,dive,required\"`\n" +
		"}\n" +
		"type Address struct {\n" +
		"	Street string `validate:\"required\"`\n" +
		"	City   string `validate:\"required\"`\n" +
		"	Planet string `validate:\"required\"`\n" +
		"	Phone  string `validate:\"required\"`\n" +
		"}\n"

	got, err := parseStructs("example/main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("struct count = %d, want 2", len(got))
	}

	wantUser := []Field{
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
	}
	if !reflect.DeepEqual(got[0].Fields, wantUser) {
		t.Fatalf("User fields mismatch\ngot:\n%swant:\n%s", formatFields(got[0].Fields), formatFields(wantUser))
	}

	wantAddress := []Field{
		{FieldName: "Street", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
		{FieldName: "City", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
		{FieldName: "Planet", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
		{FieldName: "Phone", Type: common.FieldType{BaseType: "string"}, Tag: `validate:"required"`},
	}
	if !reflect.DeepEqual(got[1].Fields, wantAddress) {
		t.Fatalf("Address fields mismatch\ngot:\n%swant:\n%s", formatFields(got[1].Fields), formatFields(wantAddress))
	}
}
