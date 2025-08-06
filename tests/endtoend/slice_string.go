package main

import (
	"log"
)

type SliceString struct {
	FirstName     string   `valid:"required"`
	TypesRequired []string `valid:"required"`
}

func slice_string_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &SliceString{
		FirstName:     "",
		TypesRequired: []string{},
	}
	expectedMsgErrors = []string{
		"FirstName is required",
		"TypesRequired must not be empty",
	}
	errs = SliceStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &SliceString{
		FirstName:     "Myname",
		TypesRequired: []string{"Type1", "Type2"},
	}
	expectedMsgErrors = nil
	errs = SliceStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("slice string tests ok")
}
