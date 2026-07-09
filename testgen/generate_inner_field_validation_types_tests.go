package main

import (
	"fmt"
	"strings"

	"github.com/opencodeco/validgen/internal/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type AllInnerFieldTestCasesToGenerate struct {
	TestCases []InnerFieldTestCaseToGenerate
}

type InnerFieldTestCaseToGenerate struct {
	StructName string
	Tests      []InnerFieldTestCase
}

type InnerFieldTestCase struct {
	FieldName1   string
	FieldName2   string
	Validation   string
	FieldType    string
	BasicType    string
	ValidCase    string
	InvalidCase  string
	ErrorMessage string
}

func generateInnerFieldValidationTypesEndToEndTests() error {
	if err := generateInnerFieldValidationTypesEndToEndTest("inner_field_no_pointer_tests.tpl", "generated_endtoend_inner_field_no_pointer_tests.go", false); err != nil {
		return err
	}

	// if err := generateInnerFieldValidationTypesEndToEndTest("inner_field_pointer_tests.tpl", "generated_endtoend_inner_field_pointer_tests.go", true); err != nil {
	// 	return err
	// }

	return nil
}

func generateInnerFieldValidationTypesEndToEndTest(tplFile, outputFile string, pointer bool) error {
	fmt.Printf("Generating inner field validation types test file: tplFile[%s] outputFile[%s] pointer[%v]\n", tplFile, outputFile, pointer)

	allTestsToGenerate := AllInnerFieldTestCasesToGenerate{}

	for _, testCase := range typesValidation {
		if !testCase.isFieldValidation {
			fmt.Printf("Skipping non-field validation: tag %s\n", testCase.tag)
			continue
		}

		structName := testCase.tag + "StructFields"
		if pointer {
			structName += "Pointer"
		}
		allTestsToGenerate.TestCases = append(allTestsToGenerate.TestCases, InnerFieldTestCaseToGenerate{
			StructName: structName,
		})
		for _, toGenerate := range testCase.testCases {
			if toGenerate.excludeIf&noPointer != 0 && !pointer {
				fmt.Printf("Skipping no pointer: tag %s type %s\n", testCase.tag, toGenerate.typeClass)
				continue
			}

			normalizedType := toGenerate.typeClass
			if pointer {
				normalizedType = "*" + normalizedType
			}
			fTypes := common.HelperFromNormalizedToBasicTypes(normalizedType)
			sNames := common.HelperFromNormalizedToStringNames(normalizedType)
			for i := range fTypes {
				fieldName1 := "Field1" + cases.Title(language.Und).String(testCase.tag) + sNames[i]
				fieldName2 := "Field2" + cases.Title(language.Und).String(testCase.tag) + sNames[i]
				validation := testCase.tag
				if testCase.argsCount != common.ZeroValue {
					validation += "=" + toGenerate.validation
				}
				validation = strings.ReplaceAll(validation, "{{.FieldName1}}", fieldName1)
				validation = strings.ReplaceAll(validation, "{{.FieldName2}}", fieldName2)
				basicType, _ := strings.CutPrefix(fTypes[i], "*")
				errorMessage := toGenerate.errorMessage
				errorMessage = strings.ReplaceAll(errorMessage, "{{.FieldName1}}", fieldName1)
				errorMessage = strings.ReplaceAll(errorMessage, "{{.FieldName2}}", fieldName2)
				errorMessage = strings.ReplaceAll(errorMessage, "{{.Target}}", toGenerate.validation)
				errorMessage = strings.ReplaceAll(errorMessage, "{{.Targets}}", targetsInMessage(toGenerate.validation))

				allTestsToGenerate.TestCases[len(allTestsToGenerate.TestCases)-1].Tests = append(allTestsToGenerate.TestCases[len(allTestsToGenerate.TestCases)-1].Tests, InnerFieldTestCase{
					FieldName1:   fieldName1,
					FieldName2:   fieldName2,
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

	if err := ExecTemplate("InnerFieldValidationTypesTests", tplFile, outputFile, allTestsToGenerate); err != nil {
		return fmt.Errorf("generating inner field validation types tests file %s", err)
	}

	fmt.Printf("Generating %s done\n", outputFile)

	return nil
}
