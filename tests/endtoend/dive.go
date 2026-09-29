package main

import "log"

type DiveUser struct {
	Addresses []Address         `valid:"required,dive"`
	Labels    map[string]string `valid:"dive,required"`
}

type ShallowDiveUser struct {
	Addresses []Address `valid:"required"`
}

type PointerDiveUser struct {
	Addresses []*Address `valid:"required,dive,required"`
}

func diveTests() {
	log.Println("starting dive tests")

	shallowEmpty := &ShallowDiveUser{}
	assertExpectedErrorMsgs("shallow empty", ShallowDiveUserValidate(shallowEmpty), []string{
		"Addresses must not be empty",
	})

	shallowPresent := &ShallowDiveUser{Addresses: []Address{{}}}
	assertExpectedErrorMsgs("shallow does not dive", ShallowDiveUserValidate(shallowPresent), nil)

	missing := &DiveUser{
		Addresses: []Address{{}},
		Labels:    map[string]string{"home": ""},
	}
	assertExpectedErrorMsgs("dive struct and map value", DiveUserValidate(missing), []string{
		"Street is required",
		"City is required",
		"Labels is required",
	})

	ok := &DiveUser{
		Addresses: []Address{{Street: "av 123", City: "city 123"}},
		Labels:    map[string]string{"home": "earth"},
	}
	assertExpectedErrorMsgs("dive valid", DiveUserValidate(ok), nil)

	nilElem := &PointerDiveUser{Addresses: []*Address{nil}}
	assertExpectedErrorMsgs("nil pointer element", PointerDiveUserValidate(nilElem), []string{
		"Addresses is required",
	})

	emptyElem := &PointerDiveUser{Addresses: []*Address{{}}}
	assertExpectedErrorMsgs("pointer element fields", PointerDiveUserValidate(emptyElem), []string{
		"Street is required",
		"City is required",
	})

	emptySlice := &PointerDiveUser{}
	assertExpectedErrorMsgs("pointer slice empty", PointerDiveUserValidate(emptySlice), []string{
		"Addresses must not be empty",
	})

	pointerOK := &PointerDiveUser{Addresses: []*Address{{Street: "av 123", City: "city 123"}}}
	assertExpectedErrorMsgs("pointer dive valid", PointerDiveUserValidate(pointerOK), nil)

	log.Println("dive tests ok")
}
