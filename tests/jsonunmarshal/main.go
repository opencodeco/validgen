package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/opencodeco/validgen/types"
)

func main() {
	log.Println("starting json unmarshal tests")

	validJSON()
	invalidValues()
	malformedJSON()
	recursionAvoidance()

	log.Println("finishing json unmarshal tests")
}

func validJSON() {
	var u User
	err := json.Unmarshal([]byte(`{"FirstName":"Ada","LastName":"Lovelace","Age":36}`), &u)
	if err != nil {
		log.Fatalf("validJSON: unexpected error: %v", err)
	}
	if u.FirstName != "Ada" || u.LastName != "Lovelace" || u.Age != 36 {
		log.Fatalf("validJSON: unexpected value: %+v", u)
	}
	log.Println("validJSON ok")
}

func invalidValues() {
	var u User
	err := json.Unmarshal([]byte(`{"FirstName":"","LastName":"Lovelace","Age":10}`), &u)
	if err == nil {
		log.Fatal("invalidValues: expected validation error")
	}

	msgs := collectValidationMsgs(err)
	if !containsMsg(msgs, "FirstName is required") {
		log.Fatalf("invalidValues: missing FirstName error in %v", msgs)
	}
	if !containsMsg(msgs, "Age must be >= 18") {
		log.Fatalf("invalidValues: missing Age error in %v", msgs)
	}
	log.Println("invalidValues ok")
}

func malformedJSON() {
	var u User
	err := json.Unmarshal([]byte(`{"FirstName":`), &u)
	if err == nil {
		log.Fatal("malformedJSON: expected syntax error")
	}
	var valErr types.ValidationError
	if errors.As(err, &valErr) {
		log.Fatalf("malformedJSON: got validation error, want JSON decode error: %v", err)
	}
	if !strings.Contains(err.Error(), "JSON") && !strings.Contains(err.Error(), "json") {
		// Still accept any non-validation decode failure from encoding/json.
		var syntax *json.SyntaxError
		if !errors.As(err, &syntax) {
			log.Fatalf("malformedJSON: unexpected error type: %v", err)
		}
	}
	log.Println("malformedJSON ok")
}

func recursionAvoidance() {
	// If the alias cast were missing, UnmarshalJSON would recurse until stack overflow.
	var u User
	err := json.Unmarshal([]byte(`{"FirstName":"Grace","LastName":"Hopper","Age":40}`), &u)
	if err != nil {
		log.Fatalf("recursionAvoidance: unexpected error: %v", err)
	}
	if u.FirstName != "Grace" {
		log.Fatalf("recursionAvoidance: unexpected value: %+v", u)
	}
	log.Println("recursionAvoidance ok")
}

func collectValidationMsgs(err error) []string {
	var msgs []string
	for _, e := range flatten(err) {
		var valErr types.ValidationError
		if errors.As(e, &valErr) {
			msgs = append(msgs, valErr.Msg)
		}
	}
	return msgs
}

func flatten(err error) []error {
	if err == nil {
		return nil
	}
	type unwrapper interface {
		Unwrap() []error
	}
	if u, ok := err.(unwrapper); ok {
		var out []error
		for _, e := range u.Unwrap() {
			out = append(out, flatten(e)...)
		}
		return out
	}
	return []error{err}
}

func containsMsg(msgs []string, want string) bool {
	for _, m := range msgs {
		if m == want {
			return true
		}
	}
	return false
}
