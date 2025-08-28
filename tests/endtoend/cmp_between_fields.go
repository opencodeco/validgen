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
	Field1     string
	Field2eq1  string `valid:"eqfield=Field1"`
	Field3neq1 string `valid:"neqfield=Field1"`
}

func cmp_between_string_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpStringFields{
		Field1:     "abc",
		Field2eq1:  "def",
		Field3neq1: "abc",
	}
	expectedMsgErrors = []string{
		"Field2eq1 must be equal to Field1",
		"Field3neq1 must not be equal to Field1",
	}
	errs = CmpStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpStringFields{
		Field1:     "abc",
		Field2eq1:  "abc",
		Field3neq1: "def",
	}
	expectedMsgErrors = nil
	errs = CmpStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between string fields tests ok")
}

type CmpUint8Fields struct {
	Field1     uint8
	Field2eq1  uint8 `valid:"eqfield=Field1"`
	Field3neq1 uint8 `valid:"neqfield=Field1"`
	Field4     uint8
	Field5gte4 uint8 `valid:"gtefield=Field4"`
	Field6gt4  uint8 `valid:"gtfield=Field4"`
	Field7lte4 uint8 `valid:"ltefield=Field4"`
	Field8lt4  uint8 `valid:"ltfield=Field4"`
}

func cmp_between_uint8_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpUint8Fields{
		Field1:     1,
		Field2eq1:  2,
		Field3neq1: 1,
		Field4:     10,
		Field5gte4: 9,
		Field6gt4:  9,
		Field7lte4: 11,
		Field8lt4:  11,
	}
	expectedMsgErrors = []string{
		"Field2eq1 must be equal to Field1",
		"Field3neq1 must not be equal to Field1",
		"Field5gte4 must be >= Field4",
		"Field6gt4 must be > Field4",
		"Field7lte4 must be <= Field4",
		"Field8lt4 must be < Field4",
	}
	errs = CmpUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpUint8Fields{
		Field1:     1,
		Field2eq1:  1,
		Field3neq1: 2,
		Field4:     10,
		Field5gte4: 11,
		Field6gt4:  11,
		Field7lte4: 9,
		Field8lt4:  9,
	}
	expectedMsgErrors = nil
	errs = CmpUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between uint8 fields tests ok")
}
