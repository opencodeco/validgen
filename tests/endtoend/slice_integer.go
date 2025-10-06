package main

import (
	"log"
)

type SliceInteger struct {
	TypesRequired []int `valid:"required"`
	TypesMin      []int `valid:"min=2"`
	TypesMax      []int `valid:"max=5"`
	TypesLen      []int `valid:"len=3"`
	TypesIn       []int `valid:"in=1 2 3"`
	TypesNotIn    []int `valid:"nin=1 2 3"`
}

func slice_integer_tests() {
	log.Println("starting slice integer tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &SliceInteger{
		TypesRequired: []int{},
		TypesMin:      []int{1},
		TypesMax:      []int{1, 2, 3, 4, 5, 6},
		TypesLen:      []int{1, 2},
		TypesIn:       []int{4},
		TypesNotIn:    []int{2, 3},
	}
	expectedMsgErrors = []string{
		"TypesRequired must not be empty",
		"TypesMin must have at least 2 elements",
		"TypesMax must have at most 5 elements",
		"TypesLen must have exactly 3 elements",
		"TypesIn elements must be one of '1' '2' '3'",
		"TypesNotIn elements must not be one of '1' '2' '3'",
	}
	errs = SliceIntegerValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &SliceInteger{
		TypesRequired: []int{1, 2},
		TypesMin:      []int{1, 2},
		TypesMax:      []int{1, 2, 3, 4},
		TypesLen:      []int{1, 2, 3},
		TypesIn:       []int{1, 2, 3, 1, 2, 3},
		TypesNotIn:    []int{4, 5, 6},
	}
	expectedMsgErrors = nil
	errs = SliceIntegerValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("slice integer tests ok")
}
