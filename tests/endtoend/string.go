package main

import "log"

type StringType struct {
	FieldReq    string `validate:"required"`
	FieldEq     string `validate:"eq=aabbcc"`
	FieldEqIC   string `validate:"eq_ignore_case=YeS"`
	FieldMinMax string `validate:"min=5,max=10"`
	FieldLen    string `validate:"len=8"`
	FieldNeq    string `validate:"neq=cba"`
	FieldNeqIC  string `validate:"neq_ignore_case=YeS"`
}

func string_tests() {
	var expectedMsgErrors []string
	var errs []error

	v := &StringType{
		FieldEq:     "123",
		FieldEqIC:   "abc",
		FieldMinMax: "1",
		FieldLen:    "abcde",
		FieldNeq:    "cba",
		FieldNeqIC:  "yeS",
	}
	expectedMsgErrors = []string{
		"FieldReq required",
		"FieldEq must be equal to 'aabbcc'",
		"FieldEqIC must be equal to 'yes'",
		"FieldMinMax length must be >= 5",
		"FieldLen length must be 8",
		"FieldNeq must be not equal to 'cba'",
		"FieldNeqIC must be not equal to 'yes'",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v = &StringType{
		FieldReq:    "123",
		FieldEq:     "aabbcc",
		FieldEqIC:   "yEs",
		FieldMinMax: "12345678901",
	}
	expectedMsgErrors = []string{
		"FieldMinMax length must be <= 10",
		"FieldLen length must be 8",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v = &StringType{
		FieldReq:    "123",
		FieldEq:     "aabbcc",
		FieldEqIC:   "yEs",
		FieldMinMax: "12345678",
		FieldLen:    "abcdefgh",
		FieldNeq:    "ops",
		FieldNeqIC:  "No",
	}
	expectedMsgErrors = nil
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("string tests ok")
}
