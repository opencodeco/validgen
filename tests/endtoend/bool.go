package main

import "log"

type BoolType struct {
	FieldEqTrue         bool `valid:"eq=true"`
	FieldNeqFalse       bool `valid:"neq=false"`
	FieldEqFieldEqTrue  bool `valid:"eqfield=FieldEqTrue"`
	FieldNeqFieldEqTrue bool `valid:"neqfield=FieldEqTrue"`
}

func boolTests() {
	log.Println("starting bool tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &BoolType{
		FieldEqTrue:         false,
		FieldNeqFalse:       false,
		FieldEqFieldEqTrue:  true,
		FieldNeqFieldEqTrue: false,
	}
	expectedMsgErrors = []string{
		"FieldEqTrue must be equal to true",
		"FieldNeqFalse must not be equal to false",
		"FieldEqFieldEqTrue must be equal to FieldEqTrue",
		"FieldNeqFieldEqTrue must not be equal to FieldEqTrue",
	}
	errs = BoolTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &BoolType{
		FieldEqTrue:         true,
		FieldNeqFalse:       true,
		FieldEqFieldEqTrue:  true,
		FieldNeqFieldEqTrue: false,
	}
	expectedMsgErrors = nil
	errs = BoolTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("bool tests ok")
}
