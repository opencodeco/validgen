package main

import "log"

type StringType struct {
	FieldReq    string `validate:"required"`
	FieldEq     string `validate:"eq=aabbcc"`
	FieldEqIC   string `validate:"eq_ignore_case=YeS"`
	FieldGteLte string `validate:"gte=5,lte=10"`
}

func string_test() {
	var expectedMsgErrors []string
	var errs []error

	v := &StringType{
		FieldEq:     "123",
		FieldEqIC:   "abc",
		FieldGteLte: "1",
	}
	expectedMsgErrors = []string{
		"FieldReq required",
		"FieldEq must be equal to 'aabbcc'",
		"FieldEqIC must be equal to 'yes'",
		"FieldGteLte length must be >= 5",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v = &StringType{
		FieldReq:    "123",
		FieldEq:     "aabbcc",
		FieldEqIC:   "yEs",
		FieldGteLte: "12345678901",
	}
	expectedMsgErrors = []string{
		"FieldGteLte length must be <= 10",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v = &StringType{
		FieldReq:    "123",
		FieldEq:     "aabbcc",
		FieldEqIC:   "yEs",
		FieldGteLte: "12345678",
	}
	expectedMsgErrors = nil
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("string tests ok")
}
