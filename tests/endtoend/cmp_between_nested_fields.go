package main

import (
	"log"
)

func cmp_between_nested_fields_tests() {
	cmp_between_nested_string_fields_tests()
	cmp_between_nested_uint8_fields_tests()

	log.Println("cmp between nested fields tests ok")
}

type CmpNestedStringFields struct {
	Field1eqNestedField1  string `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 string `valid:"neqfield=Nested.Field1"`
	Nested                NestedStringFields
}

type NestedStringFields struct {
	Field1 string
}

func cmp_between_nested_string_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpNestedStringFields{
		Field1eqNestedField1:  "def",
		Field2neqNestedField1: "abc",
		Nested: NestedStringFields{
			Field1: "abc",
		},
	}
	expectedMsgErrors = []string{
		"Field1eqNestedField1 must be equal to Nested.Field1",
		"Field2neqNestedField1 must not be equal to Nested.Field1",
	}
	errs = CmpNestedStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpNestedStringFields{
		Field1eqNestedField1:  "abc",
		Field2neqNestedField1: "def",
		Nested: NestedStringFields{
			Field1: "abc",
		},
	}
	expectedMsgErrors = nil
	errs = CmpNestedStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between nested string fields tests ok")
}

type CmpNestedUint8Fields struct {
	Field1eqNestedField1  uint8 `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 uint8 `valid:"neqfield=Nested.Field1"`
	Field3gteNestedField2 uint8 `valid:"gtefield=Nested.Field2"`
	Field4gtNestedField2  uint8 `valid:"gtfield=Nested.Field2"`
	Field5lteNestedField2 uint8 `valid:"ltefield=Nested.Field2"`
	Field6ltNestedField2  uint8 `valid:"ltfield=Nested.Field2"`
	Nested                NestedUint8Fields
}

type NestedUint8Fields struct {
	Field1 uint8
	Field2 uint8
}

func cmp_between_nested_uint8_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpNestedUint8Fields{
		Field1eqNestedField1:  2,
		Field2neqNestedField1: 1,
		Field3gteNestedField2: 9,
		Field4gtNestedField2:  9,
		Field5lteNestedField2: 11,
		Field6ltNestedField2:  11,
		Nested: NestedUint8Fields{
			Field1: 1,
			Field2: 10,
		},
	}
	expectedMsgErrors = []string{
		"Field1eqNestedField1 must be equal to Nested.Field1",
		"Field2neqNestedField1 must not be equal to Nested.Field1",
		"Field3gteNestedField2 must be >= Nested.Field2",
		"Field4gtNestedField2 must be > Nested.Field2",
		"Field5lteNestedField2 must be <= Nested.Field2",
		"Field6ltNestedField2 must be < Nested.Field2",
	}

	errs = CmpNestedUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpNestedUint8Fields{
		Field1eqNestedField1:  1,
		Field2neqNestedField1: 2,
		Field3gteNestedField2: 10,
		Field4gtNestedField2:  11,
		Field5lteNestedField2: 10,
		Field6ltNestedField2:  9,
		Nested: NestedUint8Fields{
			Field1: 1,
			Field2: 10,
		},
	}
	expectedMsgErrors = nil
	errs = CmpNestedUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between nested uint8 fields tests ok")
}
