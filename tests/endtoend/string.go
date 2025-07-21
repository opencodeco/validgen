package main

import "log"

type StringType struct {
	FieldReq    string `valid:"required"`
	FieldEq     string `valid:"eq=aabbcc"`
	FieldEqIC   string `valid:"eq_ignore_case=YeS"`
	FieldMinMax string `valid:"min=5,max=10"`
	FieldLen    string `valid:"len=8"`
	FieldNeq    string `valid:"neq=cba"`
	FieldNeqIC  string `valid:"neq_ignore_case=YeS"`
	FieldIn     string `valid:"in=ab bc cd"`
	FieldNotIn  string `valid:"nin=xx yy zz"`
	EmailReq    string `valid:"required,email"`
	EmailOpt    string `valid:"email"`
}

func string_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &StringType{
		FieldEq:     "123",
		FieldEqIC:   "abc",
		FieldMinMax: "1",
		FieldLen:    "abcde",
		FieldNeq:    "cba",
		FieldNeqIC:  "yeS",
		FieldIn:     "abc",
		FieldNotIn:  "zz",
		EmailReq:    "invalid.email.format", // Invalid required email
		EmailOpt:    "invalid",              // Invalid optional email
	}
	expectedMsgErrors = []string{
		"FieldReq is required",
		"FieldEq must be equal to 'aabbcc'",
		"FieldEqIC must be equal to 'yes'",
		"FieldMinMax length must be >= 5",
		"FieldLen length must be 8",
		"FieldNeq must not be equal to 'cba'",
		"FieldNeqIC must not be equal to 'yes'",
		"FieldIn must be one of 'ab' 'bc' 'cd'",
		"FieldNotIn must not be one of 'xx' 'yy' 'zz'",
		"EmailReq must be a valid email",
		"EmailOpt must be a valid email",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &StringType{
		FieldReq:    "123",
		FieldEq:     "aabbcc",
		FieldEqIC:   "yEs",
		FieldMinMax: "12345678",
		FieldLen:    "abcdefgh",
		FieldNeq:    "ops",
		FieldNeqIC:  "No",
		FieldIn:     "bc",
		FieldNotIn:  "xy",
		EmailReq:    "user@example.com", // Valid required email
		EmailOpt:    "",                 // Empty optional email (valid)
	}
	expectedMsgErrors = nil
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("string tests ok")
}
