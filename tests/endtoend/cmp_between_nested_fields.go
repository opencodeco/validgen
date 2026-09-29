package main

import (
	"log"
)

func cmpBetweenNestedFieldsTests() {
	log.Println("starting between nested fields tests")

	cmpBetweenNestedStringFieldsTests()
	cmpBetweenNestedUint8FieldsTests()
	cmpBetweenNestedFloat32FieldsTests()
	cmpBetweenNestedFloat64FieldsTests()
	cmpBetweenNestedComplex64FieldsTests()
	cmpBetweenNestedComplex128FieldsTests()

	log.Println("cmp between nested fields tests ok")
}

type CmpNestedStringFields struct {
	Field1eqNestedField1  string `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 string `valid:"neqfield=Nested.Field1"`
	Nested                NestedStringFields
}

type NestedStringFields struct {
	Field1 string
}

func cmpBetweenNestedStringFieldsTests() {
	log.Println("starting between nested string fields tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpNestedStringFields{
		Field1eqNestedField1:  "def",
		Field2neqNestedField1: "abc",
		Nested: NestedStringFields{
			Field1: "abc",
		},
	}
	expectedMsgErrors = []string{
		"Field1eqNestedField1 must be equal to Nested.Field1",
		"Field2neqNestedField1 must not be equal to Nested.Field1",
	}
	errs = CmpNestedStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpNestedStringFields{
		Field1eqNestedField1:  "abc",
		Field2neqNestedField1: "def",
		Nested: NestedStringFields{
			Field1: "abc",
		},
	}
	expectedMsgErrors = nil
	errs = CmpNestedStringFieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between nested string fields tests ok")
}

type CmpNestedUint8Fields struct {
	Field1eqNestedField1  uint8 `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 uint8 `valid:"neqfield=Nested.Field1"`
	Field3gteNestedField2 uint8 `valid:"gtefield=Nested.Field2"`
	Field4gtNestedField2  uint8 `valid:"gtfield=Nested.Field2"`
	Field5lteNestedField2 uint8 `valid:"ltefield=Nested.Field2"`
	Field6ltNestedField2  uint8 `valid:"ltfield=Nested.Field2"`
	Nested                NestedUint8Fields
}

type NestedUint8Fields struct {
	Field1 uint8
	Field2 uint8
}

func cmpBetweenNestedUint8FieldsTests() {
	log.Println("starting between nested uint8 fields tests")

	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &CmpNestedUint8Fields{
		Field1eqNestedField1:  2,
		Field2neqNestedField1: 1,
		Field3gteNestedField2: 9,
		Field4gtNestedField2:  9,
		Field5lteNestedField2: 11,
		Field6ltNestedField2:  11,
		Nested: NestedUint8Fields{
			Field1: 1,
			Field2: 10,
		},
	}
	expectedMsgErrors = []string{
		"Field1eqNestedField1 must be equal to Nested.Field1",
		"Field2neqNestedField1 must not be equal to Nested.Field1",
		"Field3gteNestedField2 must be >= Nested.Field2",
		"Field4gtNestedField2 must be > Nested.Field2",
		"Field5lteNestedField2 must be <= Nested.Field2",
		"Field6ltNestedField2 must be < Nested.Field2",
	}

	errs = CmpNestedUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &CmpNestedUint8Fields{
		Field1eqNestedField1:  1,
		Field2neqNestedField1: 2,
		Field3gteNestedField2: 10,
		Field4gtNestedField2:  11,
		Field5lteNestedField2: 10,
		Field6ltNestedField2:  9,
		Nested: NestedUint8Fields{
			Field1: 1,
			Field2: 10,
		},
	}
	expectedMsgErrors = nil
	errs = CmpNestedUint8FieldsValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("cmp between nested uint8 fields tests ok")
}

type CmpNestedFloat32Fields struct {
	Field1eqNestedField1  float32 `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 float32 `valid:"neqfield=Nested.Field1"`
	Field3gteNestedField2 float32 `valid:"gtefield=Nested.Field2"`
	Field4gtNestedField2  float32 `valid:"gtfield=Nested.Field2"`
	Field5lteNestedField2 float32 `valid:"ltefield=Nested.Field2"`
	Field6ltNestedField2  float32 `valid:"ltfield=Nested.Field2"`
	Nested                NestedFloat32Fields
}

type NestedFloat32Fields struct {
	Field1 float32
	Field2 float32
}

func cmpBetweenNestedFloat32FieldsTests() {
	log.Println("starting between nested float32 fields tests")
	cmpBetweenNestedFloatFieldsTests(
		func(eq, neq, gte, gt, lte, lt, nested1, nested2 float64) []error {
			v := &CmpNestedFloat32Fields{
				Field1eqNestedField1:  float32(eq),
				Field2neqNestedField1: float32(neq),
				Field3gteNestedField2: float32(gte),
				Field4gtNestedField2:  float32(gt),
				Field5lteNestedField2: float32(lte),
				Field6ltNestedField2:  float32(lt),
				Nested: NestedFloat32Fields{
					Field1: float32(nested1),
					Field2: float32(nested2),
				},
			}
			return CmpNestedFloat32FieldsValidate(v)
		},
	)
	log.Println("cmp between nested float32 fields tests ok")
}

type CmpNestedFloat64Fields struct {
	Field1eqNestedField1  float64 `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 float64 `valid:"neqfield=Nested.Field1"`
	Field3gteNestedField2 float64 `valid:"gtefield=Nested.Field2"`
	Field4gtNestedField2  float64 `valid:"gtfield=Nested.Field2"`
	Field5lteNestedField2 float64 `valid:"ltefield=Nested.Field2"`
	Field6ltNestedField2  float64 `valid:"ltfield=Nested.Field2"`
	Nested                NestedFloat64Fields
}

type NestedFloat64Fields struct {
	Field1 float64
	Field2 float64
}

func cmpBetweenNestedFloat64FieldsTests() {
	log.Println("starting between nested float64 fields tests")
	cmpBetweenNestedFloatFieldsTests(
		func(eq, neq, gte, gt, lte, lt, nested1, nested2 float64) []error {
			v := &CmpNestedFloat64Fields{
				Field1eqNestedField1:  eq,
				Field2neqNestedField1: neq,
				Field3gteNestedField2: gte,
				Field4gtNestedField2:  gt,
				Field5lteNestedField2: lte,
				Field6ltNestedField2:  lt,
				Nested: NestedFloat64Fields{
					Field1: nested1,
					Field2: nested2,
				},
			}
			return CmpNestedFloat64FieldsValidate(v)
		},
	)
	log.Println("cmp between nested float64 fields tests ok")
}

func cmpBetweenNestedFloatFieldsTests(validate func(eq, neq, gte, gt, lte, lt, nested1, nested2 float64) []error) {
	expectedMsgErrors := []string{
		"Field1eqNestedField1 must be equal to Nested.Field1",
		"Field2neqNestedField1 must not be equal to Nested.Field1",
		"Field3gteNestedField2 must be >= Nested.Field2",
		"Field4gtNestedField2 must be > Nested.Field2",
		"Field5lteNestedField2 must be <= Nested.Field2",
		"Field6ltNestedField2 must be < Nested.Field2",
	}
	errs := validate(2.5, 1.5, 9.5, 9.5, 11.5, 11.5, 1.5, 10.5)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	expectedMsgErrors = []string{
		"Field2neqNestedField1 must not be equal to Nested.Field1",
		"Field4gtNestedField2 must be > Nested.Field2",
		"Field6ltNestedField2 must be < Nested.Field2",
	}
	errs = validate(1.5, 1.5, -0.5, -0.5, -0.5, -0.5, 1.5, -0.5)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	errs = validate(1.5, 0, 10.5, 11.5, 10.5, 9.5, 1.5, 10.5)
	if !expectedMsgErrorsOk(errs, nil) {
		log.Fatalf("error = %v, wantErr %v", errs, nil)
	}
}

type CmpNestedComplex64Fields struct {
	Field1eqNestedField1  complex64 `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 complex64 `valid:"neqfield=Nested.Field1"`
	Nested                NestedComplex64Fields
}

type NestedComplex64Fields struct {
	Field1 complex64
}

func cmpBetweenNestedComplex64FieldsTests() {
	log.Println("starting between nested complex64 fields tests")
	cmpBetweenNestedComplexFieldsTests(
		func(eq, neq, nested1 complex128) []error {
			v := &CmpNestedComplex64Fields{
				Field1eqNestedField1:  complex64(eq),
				Field2neqNestedField1: complex64(neq),
				Nested: NestedComplex64Fields{
					Field1: complex64(nested1),
				},
			}
			return CmpNestedComplex64FieldsValidate(v)
		},
	)
	log.Println("cmp between nested complex64 fields tests ok")
}

type CmpNestedComplex128Fields struct {
	Field1eqNestedField1  complex128 `valid:"eqfield=Nested.Field1"`
	Field2neqNestedField1 complex128 `valid:"neqfield=Nested.Field1"`
	Nested                NestedComplex128Fields
}

type NestedComplex128Fields struct {
	Field1 complex128
}

func cmpBetweenNestedComplex128FieldsTests() {
	log.Println("starting between nested complex128 fields tests")
	cmpBetweenNestedComplexFieldsTests(
		func(eq, neq, nested1 complex128) []error {
			v := &CmpNestedComplex128Fields{
				Field1eqNestedField1:  eq,
				Field2neqNestedField1: neq,
				Nested: NestedComplex128Fields{
					Field1: nested1,
				},
			}
			return CmpNestedComplex128FieldsValidate(v)
		},
	)
	log.Println("cmp between nested complex128 fields tests ok")
}

func cmpBetweenNestedComplexFieldsTests(validate func(eq, neq, nested1 complex128) []error) {
	expectedMsgErrors := []string{
		"Field1eqNestedField1 must be equal to Nested.Field1",
		"Field2neqNestedField1 must not be equal to Nested.Field1",
	}
	errs := validate(3+4i, 1+2i, 1+2i)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	expectedMsgErrors = []string{
		"Field2neqNestedField1 must not be equal to Nested.Field1",
	}
	errs = validate(1+2i, 1+2i, 1+2i)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	errs = validate(1+2i, 3+4i, 1+2i)
	if !expectedMsgErrorsOk(errs, nil) {
		log.Fatalf("error = %v, wantErr %v", errs, nil)
	}
}
