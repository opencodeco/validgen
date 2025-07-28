package main

import "log"

type User struct {
	FirstName string  `valid:"required"`
	Age       uint8   `valid:"gte=18,lte=130"`
	Address   Address `valid:"required"`
}

type Address struct {
	Street string `valid:"required"`
	City   string `valid:"required"`
}

func nested_struct_tests() {
	var expectedMsgErrors []string
	var errs []error

	// Test case 1: All failure scenarios
	v := &User{
		FirstName: "",
		Age:       0,
		Address:   Address{},
	}
	expectedMsgErrors = []string{
		"FirstName is required",
		"Age must be >= 18",
		"Street is required",
		"City is required",
	}
	errs = UserValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	// Test case 2: All valid input
	v = &User{
		FirstName: "Myname",
		Age:       22,
		Address: Address{
			Street: "av 123",
			City:   "city 123",
		},
	}
	expectedMsgErrors = nil
	errs = UserValidate(v)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("nested struct tests ok")
}
