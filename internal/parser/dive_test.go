package parser

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
)

func TestParseDiveFieldTypes(t *testing.T) {
	src := "package main\n" +
		"type User struct {\n" +
		"	Addresses []*Address `valid:\"required,dive,required\"`\n" +
		"	Labels map[string]string `valid:\"dive,required\"`\n" +
		"	Homes []Address `valid:\"required\"`\n" +
		"	Maybe *[]Address `valid:\"dive\"`\n" +
		"}\n" +
		"type Address struct {\n" +
		"	Street string `valid:\"required\"`\n" +
		"}\n"

	got, err := parseStructs("example/main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("struct count = %d, want 2", len(got))
	}

	stringValue := &common.FieldType{BaseType: "string"}
	want := []Field{
		{
			FieldName: "Addresses",
			Type:      common.FieldType{BaseType: "main.Address", ComposedType: "*[]", ElemPointer: true},
			Tag:       `valid:"required,dive,required"`,
		},
		{
			FieldName: "Labels",
			Type:      common.FieldType{BaseType: "string", ComposedType: "map", MapValue: stringValue},
			Tag:       `valid:"dive,required"`,
		},
		{
			FieldName: "Homes",
			Type:      common.FieldType{BaseType: "main.Address", ComposedType: "[]"},
			Tag:       `valid:"required"`,
		},
		{
			FieldName: "Maybe",
			Type:      common.FieldType{BaseType: "main.Address", ComposedType: "*[]"},
			Tag:       `valid:"dive"`,
		},
	}
	if !reflect.DeepEqual(got[0].Fields, want) {
		t.Fatalf("fields mismatch\ngot:\n%swant:\n%s", formatStructs(got[:1]), formatFields(want))
	}
}

func formatFields(fields []Field) string {
	s := &Struct{Fields: fields}
	return formatStructs([]*Struct{s})
}
