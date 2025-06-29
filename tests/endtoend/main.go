package main

import (
	"errors"
	"log"
	"slices"

	"github.com/opencodeco/validgen/tests/endtoend/structsinpkg"
	"github.com/opencodeco/validgen/types"
)

type AllTypes1 struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"required"`
}

type AllTypes2 struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=18,lte=130"`
	UserName  string `validate:"gte=5,lte=10"`
	Optional  string
}

type NoValidateInfo struct {
	Name    string
	Address string
}

func main() {
	log.Println("starting all_types tests")

	all_types1_test()
	all_types2_test()
	struct_in_pkg_test()

	log.Println("all_types tests ok")
}

func all_types1_test() {
	var expectedMsgErrors []string
	var errs []error

	v1 := &AllTypes1{}
	expectedMsgErrors = []string{
		"FirstName required",
		"LastName required",
		"Age required",
	}
	errs = AllTypes1Validate(v1)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v2 := &AllTypes1{
		FirstName: "First",
		LastName:  "Last",
		Age:       18,
	}
	expectedMsgErrors = nil
	errs = AllTypes1Validate(v2)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("all_types1 ok")
}

func all_types2_test() {
	var expectedMsgErrors []string
	var errs []error

	v1 := &AllTypes2{
		FirstName: "",
		LastName:  "",
		Age:       135,
		UserName:  "abc",
	}
	expectedMsgErrors = []string{
		"FirstName required",
		"LastName required",
		"Age must be <= 130",
		"length UserName must be >= 5",
	}
	errs = AllTypes2Validate(v1)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v2 := &AllTypes2{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "mylongusername",
	}
	expectedMsgErrors = []string{
		"length UserName must be <= 10",
	}
	errs = AllTypes2Validate(v2)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v3 := &AllTypes2{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "myusername",
	}
	expectedMsgErrors = nil
	errs = AllTypes2Validate(v3)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("all_types2 ok")
}

func struct_in_pkg_test() {
	var expectedMsgErrors []string
	var errs []error

	v1 := &structsinpkg.Type1{}
	expectedMsgErrors = []string{
		"FirstName required",
		"LastName required",
		"Age required",
	}
	errs = structsinpkg.Type1Validate(v1)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	v2 := &structsinpkg.Type1{
		FirstName: "First",
		LastName:  "Last",
		Age:       18,
	}
	expectedMsgErrors = nil
	errs = structsinpkg.Type1Validate(v2)
	if !expectedMsgErrorsOk(errs, expectedMsgErrors) {
		log.Fatalf("error = %v, wantErr %v", errs, expectedMsgErrors)
	}

	log.Println("struct_in_pkg ok")
}

func expectedMsgErrorsOk(errs []error, expectedMsgErrors []string) bool {
	return slices.EqualFunc(errs, expectedMsgErrors, func(e1 error, e2msg string) bool {
		var valErr types.ValidationError
		if errors.As(e1, &valErr) {
			return valErr.Msg == e2msg
		}

		return false
	})
}
