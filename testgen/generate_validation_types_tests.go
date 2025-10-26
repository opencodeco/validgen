package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/opencodeco/validgen/internal/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type AllTestCasesToGenerate struct {
	TestCases []TestCaseToGenerate
}

type TestCaseToGenerate struct {
	StructName string
	Tests      []TestCase
}

type TestCase struct {
	FieldName    string
	Validation   string
	FieldType    string
	BasicType    string
	ValidCase    string
	InvalidCase  string
	ErrorMessage string
}

func generateValidationTypesTests() {
	generateValidationTypesTestsFile("no_pointer_tests.tpl", "generated_endtoend_no_pointer_tests.go", false)
	generateValidationTypesTestsFile("pointer_tests.tpl", "generated_endtoend_pointer_tests.go", true)
}

func generateValidationTypesTestsFile(tpl, dest string, pointer bool) {
	log.Printf("Generating validation types test file: tpl[%s] dest[%s] pointer[%v]\n", tpl, dest, pointer)

	allTestsToGenerate := AllTestCasesToGenerate{}

	for _, testCase := range typesValidation {
		structName := testCase.tag + "StructFields"
		if pointer {
			structName += "Pointer"
		}
		allTestsToGenerate.TestCases = append(allTestsToGenerate.TestCases, TestCaseToGenerate{
			StructName: structName,
		})
		for _, toGenerate := range testCase.testCases {
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
			fTypes := common.HelperFromNormalizedToBasicTypes(normalizedType)
			sNames := common.HelperFromNormalizedToStringNames(normalizedType)
			for i := range fTypes {
				validation := testCase.tag
				if testCase.argsCount != common.ZeroValue {
					validation += "=" + toGenerate.validation
				}
				fieldName := "Field" + cases.Title(language.Und).String(testCase.tag) + sNames[i]
				basicType, _ := strings.CutPrefix(fTypes[i], "*")

				errorMessage := toGenerate.errorMessage
				errorMessage = strings.ReplaceAll(errorMessage, "{{.FieldName}}", fieldName)
				errorMessage = strings.ReplaceAll(errorMessage, "{{.Target}}", toGenerate.validation)
				errorMessage = strings.ReplaceAll(errorMessage, "{{.Targets}}", targetsInMessage(toGenerate.validation))

				allTestsToGenerate.TestCases[len(allTestsToGenerate.TestCases)-1].Tests = append(allTestsToGenerate.TestCases[len(allTestsToGenerate.TestCases)-1].Tests, TestCase{
					FieldName:    fieldName,
					Validation:   validation,
					FieldType:    fTypes[i],
					BasicType:    basicType,
					ValidCase:    strings.ReplaceAll(toGenerate.validCase, "{{.BasicType}}", basicType),
					InvalidCase:  strings.ReplaceAll(toGenerate.invalidCase, "{{.BasicType}}", basicType),
					ErrorMessage: errorMessage,
				})
			}
		}
	}

	if err := allTestsToGenerate.GenerateFile(tpl, dest); err != nil {
		log.Fatalf("error generation validation types file %s", err)
	}

	log.Printf("Generating %s done\n", dest)
}

func (at *AllTestCasesToGenerate) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New("ValidationTypesTests").Parse(string(tpl))
	if err != nil {
		return err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, at); err != nil {
		return err
	}

	// if err := os.WriteFile(output, code.Bytes(), 0644); err != nil {
	// 	return err
	// }

	formattedCode, err := format.Source(code.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(output, formattedCode, 0644); err != nil {
		return err
	}

	return nil
}

func targetsInMessage(validation string) string {
	tokens := strings.Split(validation, " ")

	msg := ""
	for _, token := range tokens {
		msg += "'" + token + "' "
	}

	return strings.TrimSpace(msg)
}
