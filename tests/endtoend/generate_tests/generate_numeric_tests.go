package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/opencodeco/validgen/internal/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type NumericTests struct {
	FieldTypes []string
}

func generateNumericTests() {
	log.Println("Generating numeric test files")

	numericTests := NumericTests{}
	numericTests.FieldTypes = common.HelperFromNormalizedToBasicTypes("<INT>")

	if err := numericTests.GenerateFile("numeric_int.tpl", "./generated_numeric_int_tests.go"); err != nil {
		log.Fatalf("error generating numeric int file %s", err)
	}

	numericTests = NumericTests{}
	numericTests.FieldTypes = common.HelperFromNormalizedToBasicTypes("<FLOAT>")

	if err := numericTests.GenerateFile("numeric_float.tpl", "./generated_numeric_float_tests.go"); err != nil {
		log.Fatalf("error generating numeric float file %s", err)
	}

	log.Println("Generating done")
}

func (bt *NumericTests) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	funcMap := template.FuncMap{
		"title": cases.Title(language.Und).String,
	}

	tmpl, err := template.New("NumericTest").Funcs(funcMap).Parse(string(tpl))
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
