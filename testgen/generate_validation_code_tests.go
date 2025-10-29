package main

import (
	"fmt"

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

func generateValidationCodeUnitTests() error {
	if err := generateValidationCodeUnitTest("build_validation_code_test.tpl", "generated_validation_code_no_pointer_test.go", false); err != nil {
		return err
	}

	if err := generateValidationCodeUnitTest("build_validation_code_test.tpl", "generated_validation_code_pointer_test.go", true); err != nil {
		return err
	}

	return nil
}

func generateValidationCodeUnitTest(tplFile, outputFile string, pointer bool) error {
	fmt.Printf("Generating validation code test file: tplFile[%s] outputFile[%s] pointer[%v]\n", tplFile, outputFile, pointer)

	funcName := "TestBuildValidationCode"
	if pointer {
		funcName += "Pointer"
	}

	testCases := ValidationCodeTestCases{
		FuncName: funcName,
	}

	for _, typeValidation := range typesValidation {
		for _, toGenerate := range typeValidation.testCases {
			if toGenerate.excludeIf&noPointer != 0 && !pointer {
				fmt.Printf("Skipping no pointer: tag %s type %s\n", typeValidation.tag, toGenerate.typeClass)
				continue
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
					return fmt.Errorf("failed to parse validation %q: %v", validation, err)
				}
				expectedValidationCode, err := gv.BuildValidationCode(fieldName, fieldType, []*analyzer.Validation{parsedValidation})
				if err != nil {
					return fmt.Errorf("failed to build validation code for %q: %v", fieldName, err)
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

	if err := ExecTemplate("ValidationCodeTests", tplFile, outputFile, testCases); err != nil {
		return fmt.Errorf("error generating validation code tests file %s", err)
	}

	fmt.Printf("Generating %s done\n", outputFile)

	return nil
}
