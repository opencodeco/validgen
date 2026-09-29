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

type KeyedDiveUser struct {
	Labels map[string]string `valid:"dive,keys,min=2,endkeys,required"`
	Scores map[uint8]string  `valid:"dive,keys,gte=1,endkeys,required"`
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

	keyTests()

	log.Println("dive tests ok")
}

func keyTests() {
	log.Println("starting keys tests")

	shortKey := &KeyedDiveUser{
		Labels: map[string]string{"a": "earth"},
		Scores: map[uint8]string{1: "ok"},
	}
	assertExpectedErrorMsgs("string key failure", KeyedDiveUserValidate(shortKey), []string{
		"Labels length must be >= 2",
	})

	lowScore := &KeyedDiveUser{
		Labels: map[string]string{"home": "earth"},
		Scores: map[uint8]string{0: "ok"},
	}
	assertExpectedErrorMsgs("typed key failure", KeyedDiveUserValidate(lowScore), []string{
		"Scores must be >= 1",
	})

	emptyValue := &KeyedDiveUser{
		Labels: map[string]string{"home": ""},
		Scores: map[uint8]string{1: ""},
	}
	assertExpectedErrorMsgs("value validation after endkeys", KeyedDiveUserValidate(emptyValue), []string{
		"Labels is required",
		"Scores is required",
	})

	both := &KeyedDiveUser{
		Labels: map[string]string{"a": ""},
		Scores: map[uint8]string{0: ""},
	}
	assertExpectedErrorMsgs("key and value failure", KeyedDiveUserValidate(both), []string{
		"Labels length must be >= 2",
		"Labels is required",
		"Scores must be >= 1",
		"Scores is required",
	})

	ok := &KeyedDiveUser{
		Labels: map[string]string{"home": "earth"},
		Scores: map[uint8]string{1: "ok"},
	}
	assertExpectedErrorMsgs("keys valid", KeyedDiveUserValidate(ok), nil)

	log.Println("keys tests ok")
}
