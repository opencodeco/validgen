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

func generateFunctionCodeUnitTests() {
	generateFunctionCodeTestsFile("function_code_test.tpl", "generated_function_code_no_pointer_test.go", false)
	generateFunctionCodeTestsFile("function_code_test.tpl", "generated_function_code_pointer_test.go", true)
}

func generateFunctionCodeTestsFile(tpl, dest string, pointer bool) {
	log.Printf("Generating function code test file: tpl[%s] dest[%s] pointer[%v]\n", tpl, dest, pointer)

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
				fieldName := "Field" + cases.Title(language.Und).String(typeValidation.tag) + fieldType.ToStringName()
				parsedValidation, err := analyzer.ParserValidation(validation)
				if err != nil {
					log.Fatalf("failed to parse validation %q: %v", validation, err)
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
			log.Fatalf("failed to build function validator code for struct %q: %v", newTest.StructName, err)
		}

		newTest.ExpectedCode = expectedCode

		testCases.Tests = append(testCases.Tests, newTest)
	}

	if err := testCases.GenerateFile(tpl, dest); err != nil {
		log.Fatalf("error generating function code tests file %s", err)
	}

	log.Printf("Generating %s done\n", dest)
}

func (tc *FunctionCodeTestCases) GenerateFile(tplFile, output string) error {
	tpl, err := os.ReadFile(tplFile)
	if err != nil {
		return fmt.Errorf("error reading %s: %s", tplFile, err)
	}

	tmpl, err := template.New("ValidationCodeTests").Parse(string(tpl))
	if err != nil {
		return err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, tc); err != nil {
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
