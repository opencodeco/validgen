package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/codegenerator"
	"github.com/opencodeco/validgen/internal/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ValidationCodeTestCases struct {
	FuncName string
	Tests    []ValidationCodeTestCase
}

type ValidationCodeTestCase struct {
	TestName     string
	FieldName    string
	FieldType    common.FieldType
	Validation   string
	ExpectedCode string
}

func generateValidationCodeUnitTests() {
	generateValidationCodeTestsFile("build_validation_code_test.tpl", "generated_validation_code_no_pointer_test.go", false)
	generateValidationCodeTestsFile("build_validation_code_test.tpl", "generated_validation_code_pointer_test.go", true)
}

func generateValidationCodeTestsFile(tpl, dest string, pointer bool) {
	log.Printf("Generating validation code test file: tpl[%s] dest[%s] pointer[%v]\n", tpl, dest, pointer)

	funcName := "TestBuildValidationCode"
	if pointer {
		funcName += "Pointer"
	}

	testCases := ValidationCodeTestCases{
		FuncName: funcName,
	}

	for _, typeValidation := range typesValidation {
		for _, toGenerate := range typeValidation.testCases {
			// Default ("") gen no pointer and pointer test.
			if toGenerate.generateFor != "" {
				if toGenerate.generateFor == "pointer" && !pointer {
					continue
				}
				if toGenerate.generateFor == "nopointer" && pointer {
					continue
				}
			}

			normalizedType := toGenerate.typeClass
			if pointer {
				normalizedType = "*" + normalizedType
			}

			fieldTypes, _ := common.HelperFromNormalizedToFieldTypes(normalizedType)
			for _, fieldType := range fieldTypes {
				validation := typeValidation.tag
				if typeValidation.argsCount != common.ZeroValue {
					validation += "=" + toGenerate.validation
				}
				testName := fmt.Sprintf("%s_%s_%s", typeValidation.tag, cases.Lower(language.Und).String(fieldType.ToStringName()), validation)
				fieldName := "Field" + cases.Title(language.Und).String(typeValidation.tag) + fieldType.ToStringName()
				gv := codegenerator.GenValidations{}
				parsedValidation, err := analyzer.ParserValidation(validation)
				if err != nil {
					log.Fatalf("failed to parse validation %q: %v", validation, err)
				}
				expectedValidationCode, err := gv.BuildValidationCode(fieldName, fieldType, []*analyzer.Validation{parsedValidation})
				if err != nil {
					log.Fatalf("failed to build validation code for %q: %v", fieldName, err)
				}

				testCases.Tests = append(testCases.Tests, ValidationCodeTestCase{
					TestName:     testName,
					FieldName:    fieldName,
					FieldType:    fieldType,
					Validation:   validation,
					ExpectedCode: expectedValidationCode,
				})
			}
		}
	}

	if err := testCases.GenerateFile(tpl, dest); err != nil {
		log.Fatalf("error generation validation code tests file %s", err)
	}

	log.Printf("Generating %s done\n", dest)
}

func (at *ValidationCodeTestCases) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New("ValidationCodeTests").Parse(string(tpl))
	if err != nil {
		return err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, at); err != nil {
		return err
	}

	formattedCode, err := format.Source(code.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(output, formattedCode, 0644); err != nil {
		return err
	}

	return nil
}
