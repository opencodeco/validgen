package main

import "log"

type StringType struct {
	FieldReq    string `validate:"required"`
	FieldEqIC   string `validate:"eq_ignore_case=YeS"`
	FieldGteLte string `validate:"gte=5,lte=10"`
}

func string_test() {
	var expectedMsgErrors []string
	var errs []error

	v := &StringType{
		FieldEqIC:   "abc",
		FieldGteLte: "1",
	}
	expectedMsgErrors = []string{
		"FieldReq required",
		"FieldEqIC must be equal to 'yes'",
		"length FieldGteLte must be >= 5",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v = &StringType{
		FieldReq:    "123",
		FieldEqIC:   "yEs",
		FieldGteLte: "12345678901",
	}
	expectedMsgErrors = []string{
		"length FieldGteLte must be <= 10",
	}
	errs = StringTypeValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v = &StringType{
		FieldReq:    "123",
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
