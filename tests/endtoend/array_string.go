package main

import (
	"log"
)

type ArrayString struct {
	TypesIn    [8]string `valid:"in=a b c"`
	TypesNotIn [8]string `valid:"nin=a b c"`
}

func arrayStringTests() {
	log.Println("starting array string tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &ArrayString{
		TypesIn:    [8]string{"d"},
		TypesNotIn: [8]string{"a", "b", "c", "d", "e", "f", "g", "h"},
	}
	expectedMsgErrors = []string{
		"TypesIn elements must be one of 'a' 'b' 'c'",
		"TypesNotIn elements must not be one of 'a' 'b' 'c'",
	}
	errs = ArrayStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &ArrayString{
		TypesIn:    [8]string{"a", "b", "c", "a", "b", "c", "a", "b"},
		TypesNotIn: [8]string{"d", "e", "f", "d", "e", "f", "d", "e"},
	}
	expectedMsgErrors = nil
	errs = ArrayStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("array string tests ok")
}
