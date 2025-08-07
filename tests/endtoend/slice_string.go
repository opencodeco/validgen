package main

import (
	"log"
)

type SliceString struct {
	TypesRequired []string `valid:"required"`
	TypesMin      []string `valid:"min=2"`
	TypesMax      []string `valid:"max=5"`
	TypesLen      []string `valid:"len=3"`
	TypesIn       []string `valid:"in=a b c"`
}

func slice_string_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &SliceString{
		TypesRequired: []string{},
		TypesMin:      []string{"1"},
		TypesMax:      []string{"1", "2", "3", "4", "5", "6"},
		TypesLen:      []string{"1", "2"},
		TypesIn:       []string{"d"},
	}
	expectedMsgErrors = []string{
		"TypesRequired must not be empty",
		"TypesMin must have at least 2 elements",
		"TypesMax must have at most 5 elements",
		"TypesLen must have exactly 3 elements",
		"TypesIn elements must be one of 'a' 'b' 'c'",
	}
	errs = SliceStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &SliceString{
		TypesRequired: []string{"Type1", "Type2"},
		TypesMin:      []string{"1", "2"},
		TypesMax:      []string{"1", "2", "3", "4"},
		TypesLen:      []string{"1", "2", "3"},
		TypesIn:       []string{"a", "b", "c", "b", "a"},
	}
	expectedMsgErrors = nil
	errs = SliceStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("slice string tests ok")
}
