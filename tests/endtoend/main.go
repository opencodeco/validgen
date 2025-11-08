package main

import (
	"errors"
	"log"
	"slices"
	"strings"

	"github.com/opencodeco/validgen/tests/endtoend/structsinpkg"
	"github.com/opencodeco/validgen/types"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type AllTypes1 struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"required"`
}

type AllTypes2 struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"gte=18,lte=130"`
	UserName  string `valid:"min=5,max=10"`
	Optional  string
}

type NoValidateInfo struct {
	Name    string
	Address string
}

func main() {
	log.Println("starting tests")

	structInPkgTests()
	nestedStructTests()
	cmpBetweenInnerFieldsTests()
	cmpBetweenNestedFieldsTests()
	boolTests()
	pointerTests()
	noPointerTests()

	log.Println("finishing tests")
}

func structInPkgTests() {
	var expectedMsgErrors []string
	var errs []error

	v1 := &structsinpkg.Type1{}
	expectedMsgErrors = []string{
		"FirstName is required",
		"LastName is required",
		"Age is required",
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

	log.Println("struct_in_pkg tests ok")
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

func assertExpectedErrorMsgs(testName string, errs []error, expectedMsgErrors []string) {
	ok := slices.EqualFunc(errs, expectedMsgErrors, func(e1 error, e2msg string) bool {
		var valErr types.ValidationError
		if errors.As(e1, &valErr) {
			return valErr.Msg == e2msg
		}

		return false
	})

	if ok {
		return
	}

	dmp := diffmatchpatch.New()
	want := strings.Join(expectedMsgErrors, "\n")
	got := ""
	for _, err := range errs {
		got += err.Error() + "\n"
	}

	diffs := dmp.DiffMain(want, got, false)
	if len(diffs) == 0 {
		return
	}

	diff := dmp.DiffPrettyText(diffs)
	log.Fatalf("%s error\nerror = %v\nwantErr = %v\n%s() diff = \n%v", testName, errs, expectedMsgErrors, testName, diff)
}
