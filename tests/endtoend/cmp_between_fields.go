package main

import (
	"log"
)

type MyStruct struct {
	FieldStr1 string `valid:"required"`
	FieldStr2 string `valid:"required,eqfield=FieldStr1"`
}

func cmp_between_fields_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &MyStruct{
		FieldStr1: "abc",
		FieldStr2: "def",
	}
	expectedMsgErrors = []string{
		"FieldStr2 must be equal to FieldStr1",
	}
	errs = MyStructValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &MyStruct{
		FieldStr1: "abc",
		FieldStr2: "abc",
	}
	expectedMsgErrors = nil
	errs = MyStructValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between fields tests ok")
}
