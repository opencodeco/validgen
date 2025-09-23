package main

import "log"

type BoolType struct {
	FieldEqTrue   bool `valid:"eq=true"`
	FieldNeqFalse bool `valid:"neq=false"`
}

func bool_tests() {
	log.Println("starting bool tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &BoolType{
		FieldEqTrue:   false,
		FieldNeqFalse: false,
	}
	expectedMsgErrors = []string{
		"FieldEqTrue must be equal to true",
		"FieldNeqFalse must not be equal to false",
	}
	errs = BoolTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &BoolType{
		FieldEqTrue:   true,
		FieldNeqFalse: true,
	}
	expectedMsgErrors = nil
	errs = BoolTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("bool tests ok")
}
