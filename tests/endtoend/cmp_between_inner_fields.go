package main

import (
	"log"
)

func cmpBetweenInnerFieldsTests() {
	log.Println("starting between inner fields tests")

	cmpBetweenInnerStringFieldsTests()
	cmpBetweenInnerUint8FieldsTests()
	cmpBetweenInnerBoolFieldsTests()

	log.Println("cmp between inner fields tests ok")
}

type CmpInnerStringFields struct {
	Field1     string
	Field2eq1  string `valid:"eqfield=Field1"`
	Field3neq1 string `valid:"neqfield=Field1"`
}

func cmpBetweenInnerStringFieldsTests() {
	log.Println("starting between inner string fields tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpInnerStringFields{
		Field1:     "abc",
		Field2eq1:  "def",
		Field3neq1: "abc",
	}
	expectedMsgErrors = []string{
		"Field2eq1 must be equal to Field1",
		"Field3neq1 must not be equal to Field1",
	}
	errs = CmpInnerStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpInnerStringFields{
		Field1:     "abc",
		Field2eq1:  "abc",
		Field3neq1: "def",
	}
	expectedMsgErrors = nil
	errs = CmpInnerStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between inner string fields tests ok")
}

type CmpInnerUint8Fields struct {
	Field1     uint8
	Field2eq1  uint8 `valid:"eqfield=Field1"`
	Field3neq1 uint8 `valid:"neqfield=Field1"`
	Field4     uint8
	Field5gte4 uint8 `valid:"gtefield=Field4"`
	Field6gt4  uint8 `valid:"gtfield=Field4"`
	Field7lte4 uint8 `valid:"ltefield=Field4"`
	Field8lt4  uint8 `valid:"ltfield=Field4"`
}

func cmpBetweenInnerUint8FieldsTests() {
	log.Println("starting between inner uint8 fields tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpInnerUint8Fields{
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
	errs = CmpInnerUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpInnerUint8Fields{
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
	errs = CmpInnerUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between inner uint8 fields tests ok")
}

type CmpInnerBoolFields struct {
	Field1     bool
	Field2eq1  bool `valid:"eqfield=Field1"`
	Field3neq1 bool `valid:"neqfield=Field1"`
}

func cmpBetweenInnerBoolFieldsTests() {
	log.Println("starting between inner bool fields tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpInnerBoolFields{
		Field1:     true,
		Field2eq1:  false,
		Field3neq1: true,
	}
	expectedMsgErrors = []string{
		"Field2eq1 must be equal to Field1",
		"Field3neq1 must not be equal to Field1",
	}
	errs = CmpInnerBoolFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpInnerBoolFields{
		Field1:     true,
		Field2eq1:  true,
		Field3neq1: false,
	}
	expectedMsgErrors = nil
	errs = CmpInnerBoolFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between inner bool fields tests ok")
}
