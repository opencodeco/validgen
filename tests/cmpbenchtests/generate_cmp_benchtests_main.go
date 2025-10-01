package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

type BenchTests struct {
	Tests []Test
}

type Test struct {
	TestType          string
	FieldType         string
	ValidGenValidate  string
	ValidatorValidate string
	ValidInput        string
}

func main() {
	log.Println("Generating benchtest files")

	benchTests := BenchTests{}
	benchTests.Tests = []Test{
		{
			TestType:          "StringRequired",
			FieldType:         "string",
			ValidGenValidate:  "required",
			ValidatorValidate: "required",
			ValidInput:        "xpto",
		},
		{
			TestType:          "StringEq",
			FieldType:         "string",
			ValidGenValidate:  "eq=abc",
			ValidatorValidate: "eq=abc",
			ValidInput:        "abc",
		},
		{
			TestType:          "StringEqIC",
			FieldType:         "string",
			ValidGenValidate:  "eq_ignore_case=abc",
			ValidatorValidate: "eq_ignore_case=abc",
			ValidInput:        "AbC",
		},
		{
			TestType:          "StringNeq",
			FieldType:         "string",
			ValidGenValidate:  "neq=abc",
			ValidatorValidate: "ne=abc",
			ValidInput:        "123",
		},
		{
			TestType:          "StringNeqIC",
			FieldType:         "string",
			ValidGenValidate:  "neq_ignore_case=abc",
			ValidatorValidate: "ne_ignore_case=abc",
			ValidInput:        "123",
		},
		{
			TestType:          "StringLen",
			FieldType:         "string",
			ValidGenValidate:  "len=5",
			ValidatorValidate: "len=5",
			ValidInput:        "abcde",
		},
		{
			TestType:          "StringMax",
			FieldType:         "string",
			ValidGenValidate:  "max=5",
			ValidatorValidate: "max=5",
			ValidInput:        "abcde",
		},
		{
			TestType:          "StringMin",
			FieldType:         "string",
			ValidGenValidate:  "min=3",
			ValidatorValidate: "min=3",
			ValidInput:        "abcd",
		},
		{
			TestType:          "StringIn",
			FieldType:         "string",
			ValidGenValidate:  "in=ab cd ef",
			ValidatorValidate: "oneof=ab cd ef",
			ValidInput:        "ef",
		},
		{
			TestType:          "StringEmail",
			FieldType:         "string",
			ValidGenValidate:  "email",
			ValidatorValidate: "email",
			ValidInput:        "aaa@example.com",
		},

		// nin
	}

	if err := benchTests.GenerateFile("types.tpl", "./generated_tests/types.go"); err != nil {
		log.Fatalf("error generation types file %s", err)
	}

	if err := benchTests.GenerateFile("validgen_vs_validator_test.tpl", "./generated_tests/validgen_vs_validator_test.go"); err != nil {
		log.Fatalf("error generation test file %s", err)
	}

	log.Println("Generating done")
}

func (bt *BenchTests) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		log.Fatalf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New("BenchTest").Parse(string(tpl))
	if err != nil {
		return err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, bt); err != nil {
		return err
	}

	if err := os.WriteFile(output, code.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
