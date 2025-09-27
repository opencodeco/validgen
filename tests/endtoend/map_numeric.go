package main

import (
	"log"
)

type MapUint8 struct {
	FieldRequired map[uint8]string `valid:"required"`
	FieldMin      map[uint8]string `valid:"min=2"`
	FieldMax      map[uint8]string `valid:"max=5"`
	FieldLen      map[uint8]string `valid:"len=3"`
	FieldIn       map[uint8]string `valid:"in=1 2 3"`
	FieldNotIn    map[uint8]string `valid:"nin=1 2 3"`
}

func map_uint8_tests() {
	log.Println("starting map uint8 tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &MapUint8{
		FieldRequired: map[uint8]string{},
		FieldMin:      map[uint8]string{1: "1"},
		FieldMax:      map[uint8]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6"},
		FieldLen:      map[uint8]string{1: "1", 2: "2"},
		FieldIn:       map[uint8]string{4: "d"},
		FieldNotIn:    map[uint8]string{1: "a", 2: "b"},
	}
	expectedMsgErrors = []string{
		"FieldRequired must not be empty",
		"FieldMin must have at least 2 elements",
		"FieldMax must have at most 5 elements",
		"FieldLen must have exactly 3 elements",
		"FieldIn elements must be one of '1' '2' '3'",
		"FieldNotIn elements must not be one of '1' '2' '3'",
	}
	errs = MapUint8Validate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &MapUint8{
		FieldRequired: map[uint8]string{1: "1", 2: "2"},
		FieldMin:      map[uint8]string{1: "1", 2: "2"},
		FieldMax:      map[uint8]string{1: "1", 2: "2", 3: "3", 4: "4"},
		FieldLen:      map[uint8]string{1: "1", 2: "2", 3: "3"},
		FieldIn:       map[uint8]string{1: "1", 2: "2", 3: "3"},
		FieldNotIn:    map[uint8]string{4: "d", 5: "e", 6: "f"},
	}
	expectedMsgErrors = nil
	errs = MapUint8Validate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("map uint8 tests ok")
}
