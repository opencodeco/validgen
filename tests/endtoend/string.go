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
	FieldIn     string `validate:"in=ab bc cd"`
	FieldNotIn  string `validate:"nin=xx yy zz"`
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
		FieldIn:     "abc",
		FieldNotIn:  "zz",
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
		FieldIn:     "bc",
		FieldNotIn:  "xy",
	}
	expectedMsgErrors = nil
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("string tests ok")
}
