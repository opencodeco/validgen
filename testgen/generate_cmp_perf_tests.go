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

type CmpBenchTests struct {
	Tests []CmpBenchTest
}

type CmpBenchTest struct {
	TestName     string
	FieldType    string
	BasicType    string
	ValidGenTag  string
	ValidatorTag string
	ValidInput   string
}

func generateComparativePerformanceTests() {
	generateComparativePerformanceTest("cmp_perf_no_pointer_tests.tpl", "generated_cmp_perf_no_pointer_test.go", false)
	generateComparativePerformanceTest("cmp_perf_pointer_tests.tpl", "generated_cmp_perf_pointer_test.go", true)
}

func generateComparativePerformanceTest(tpl, dest string, pointer bool) {
	log.Printf("Generating comparative performance tests file: tpl[%s] dest[%s] pointer[%v]\n", tpl, dest, pointer)

	benchTests := CmpBenchTests{}

	for _, typeVal := range typesValidation {
		if typeVal.validatorTag == "" {
			log.Printf("Skipping tag %s: go-validator tag not defined\n", typeVal.tag)
			continue
		}

		for _, testCase := range typeVal.testCases {
			if testCase.excludeIf&cmpBenchTests != 0 {
				log.Printf("Skipping test: tag %s type %s\n", typeVal.tag, testCase.typeClass)
				continue
			}
			if testCase.excludeIf&noPointer != 0 && !pointer {
				log.Printf("Skipping no pointer: tag %s type %s\n", typeVal.tag, testCase.typeClass)
				continue
			}

			normalizedType := testCase.typeClass
			if pointer {
				normalizedType = "*" + normalizedType
			}

			fTypes := common.HelperFromNormalizedToBasicTypes(normalizedType)
			sNames := common.HelperFromNormalizedToStringNames(normalizedType)

			for i := range fTypes {
				validGenTag := typeVal.tag
				if typeVal.argsCount != common.ZeroValue {
					validGenTag += "=" + testCase.validation
				}
				goValidatorTag := typeVal.validatorTag
				if typeVal.argsCount != common.ZeroValue {
					goValidatorTag += "=" + testCase.validation
				}
				testName := cases.Title(language.Und).String(typeVal.tag) + sNames[i]

				basicType, _ := strings.CutPrefix(fTypes[i], "*")

				benchTests.Tests = append(benchTests.Tests, CmpBenchTest{
					TestName:     testName,
					FieldType:    fTypes[i],
					BasicType:    basicType,
					ValidGenTag:  validGenTag,
					ValidatorTag: goValidatorTag,
					ValidInput:   strings.ReplaceAll(testCase.validCase, "{{.BasicType}}", basicType),
				})
			}
		}
	}

	log.Printf("%d test cases were generated\n", len(benchTests.Tests))

	if err := benchTests.GenerateFile(tpl, dest); err != nil {
		log.Fatalf("error generating comparative performance tests file %s", err)
	}

	log.Println("Generating done")
}

func (cbt *CmpBenchTests) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New("BenchTest").Parse(string(tpl))
	if err != nil {
		return fmt.Errorf("error parsing template %s: %s", tplFile, err)
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, cbt); err != nil {
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
