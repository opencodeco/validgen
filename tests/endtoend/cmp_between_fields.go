package main

import (
	"log"
)

func cmp_between_fields_tests() {
	cmp_between_string_fields_tests()
	cmp_between_uint8_fields_tests()

	log.Println("cmp between fields tests ok")
}

type CmpStringFields struct {
	Field1 string `valid:"required"`
	Field2 string `valid:"required,eqfield=Field1"`
	Field3 string `valid:"required,neqfield=Field1"`
}

func cmp_between_string_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpStringFields{
		Field1: "abc",
		Field2: "def",
		Field3: "abc",
	}
	expectedMsgErrors = []string{
		"Field2 must be equal to Field1",
		"Field3 must not be equal to Field1",
	}
	errs = CmpStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpStringFields{
		Field1: "abc",
		Field2: "abc",
		Field3: "def",
	}
	expectedMsgErrors = nil
	errs = CmpStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between string fields tests ok")
}

type CmpUint8Fields struct {
	Field1 uint8 `valid:"required"`
	Field2 uint8 `valid:"required,eqfield=Field1"`
	Field3 uint8 `valid:"required,neqfield=Field1"`
}

func cmp_between_uint8_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpUint8Fields{
		Field1: 1,
		Field2: 2,
		Field3: 1,
	}
	expectedMsgErrors = []string{
		"Field2 must be equal to Field1",
		"Field3 must not be equal to Field1",
	}
	errs = CmpUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpUint8Fields{
		Field1: 1,
		Field2: 1,
		Field3: 2,
	}
	expectedMsgErrors = nil
	errs = CmpUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between uint8 fields tests ok")
}
