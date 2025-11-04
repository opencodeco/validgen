package main

import (
	"fmt"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/codegenerator"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FunctionCodeTestCases struct {
	FuncName string
	Tests    []FunctionCodeTestCase
}

type FunctionCodeTestCase struct {
	TestName     string
	StructName   string
	Fields       []FunctionCodeTestField
	ExpectedCode string
}

type FunctionCodeTestField struct {
	Name string
	Type common.FieldType
	Tag  string
}

func generateFunctionCodeUnitTests() error {
	if err := generateFunctionCodeUnitTest("function_code_test.tpl", "generated_function_code_no_pointer_test.go", false); err != nil {
		return err
	}

	if err := generateFunctionCodeUnitTest("function_code_test.tpl", "generated_function_code_pointer_test.go", true); err != nil {
		return err
	}

	return nil
}

func generateFunctionCodeUnitTest(tplFile, outputFile string, pointer bool) error {
	fmt.Printf("Generating function code test file: tplFile[%s] outputFile[%s] pointer[%v]\n", tplFile, outputFile, pointer)

	funcName := "TestBuildFunctionCode"
	if pointer {
		funcName += "Pointer"
	}

	testCases := FunctionCodeTestCases{
		FuncName: funcName,
	}

	for _, typeValidation := range typesValidation {
		newTest := FunctionCodeTestCase{
			TestName:   typeValidation.tag + "Struct",
			StructName: typeValidation.tag + "Struct",
		}

		structInfo := &analyzer.Struct{
			Struct: parser.Struct{
				PackageName: "main",
				StructName:  newTest.StructName,
			},
		}

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
				fieldName := "Field" + cases.Title(language.Und).String(typeValidation.tag) + fieldType.ToStringName()
				parsedValidation, err := analyzer.ParserValidation(validation)
				if err != nil {
					return fmt.Errorf("failed to parse validation %q: %v", validation, err)
				}

				newTest.Fields = append(newTest.Fields, FunctionCodeTestField{
					Name: fieldName,
					Type: fieldType,
					Tag:  validation,
				})

				structInfo.Fields = append(structInfo.Fields, parser.Field{
					FieldName: fieldName,
					Type:      fieldType,
					Tag:       `validate:"` + validation + `"`,
				})
				structInfo.FieldsValidations = append(structInfo.FieldsValidations, analyzer.FieldValidations{
					Validations: []*analyzer.Validation{parsedValidation},
				})
			}
		}

		gv := codegenerator.GenValidations{
			Struct: structInfo,
		}

		expectedCode, err := gv.BuildFuncValidatorCode()
		if err != nil {
			return fmt.Errorf("failed to build function validator code for struct %q: %v", newTest.StructName, err)
		}

		newTest.ExpectedCode = expectedCode

		testCases.Tests = append(testCases.Tests, newTest)
	}

	if err := ExecTemplate("FunctionCodeTests", tplFile, outputFile, testCases); err != nil {
		return fmt.Errorf("generating function code tests file %s", err)
	}

	fmt.Printf("Generating %s done\n", outputFile)

	return nil
}
