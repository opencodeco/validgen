package main

import (
	"log"
)

type MapString struct {
	FieldRequired map[string]string `valid:"required"`
	FieldMin      map[string]string `valid:"min=2"`
	FieldMax      map[string]string `valid:"max=5"`
	FieldLen      map[string]string `valid:"len=3"`
	FieldIn       map[string]string `valid:"in=a b c"`
	FieldNotIn    map[string]string `valid:"nin=a b c"`
}

func mapStringTests() {
	log.Println("starting map string tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &MapString{
		FieldRequired: map[string]string{},
		FieldMin:      map[string]string{"1": "1"},
		FieldMax:      map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6"},
		FieldLen:      map[string]string{"1": "1", "2": "2"},
		FieldIn:       map[string]string{"d": "d"},
		FieldNotIn:    map[string]string{"a": "a", "b": "b"},
	}
	expectedMsgErrors = []string{
		"FieldRequired must not be empty",
		"FieldMin must have at least 2 elements",
		"FieldMax must have at most 5 elements",
		"FieldLen must have exactly 3 elements",
		"FieldIn elements must be one of 'a' 'b' 'c'",
		"FieldNotIn elements must not be one of 'a' 'b' 'c'",
	}
	errs = MapStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &MapString{
		FieldRequired: map[string]string{"1": "1", "2": "2"},
		FieldMin:      map[string]string{"1": "1", "2": "2"},
		FieldMax:      map[string]string{"1": "1", "2": "2", "3": "3", "4": "4"},
		FieldLen:      map[string]string{"1": "1", "2": "2", "3": "3"},
		FieldIn:       map[string]string{"a": "a", "b": "b", "c": "c"},
		FieldNotIn:    map[string]string{"d": "d", "e": "e", "f": "f"},
	}
	expectedMsgErrors = nil
	errs = MapStringValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("map string tests ok")
}
