package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var structValidatorTpl = `package {{.PackageName}}

import (
	"fmt"
)

func {{.Name}}Validate(obj *{{.Name}}) []error {
	var errs []error
{{range .FieldsInfo}}
{{condition .Name .Type .Tag}}
{{end}}
	return errs
}
`

var packageDefinitionTpl = `package {{.PackageName}}

import (
	"errors"
)

var ErrValidation = errors.New("validation error")
`

type StructInfo struct {
	Name           string
	Path           string
	PackageName    string
	FieldsInfo     []FieldInfo
	HasValidateTag bool
}

// TODO: NewFieldValidation to validate params and build the object.
type FieldInfo struct {
	Name string
	Type string
	Tag  string
}

func (fv *StructInfo) GenerateValidator() (string, error) {
	funcMap := template.FuncMap{
		"condition": condition,
	}

	tmpl, err := template.New("FileValidator").Funcs(funcMap).Parse(structValidatorTpl)
	if err != nil {
		return "", err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, fv); err != nil {
		return "", err
	}

	return code.String(), nil
}

func condition(fieldName, fieldType, fieldTag string) (string, error) {
	tag := fieldTag
	tag, _ = strings.CutPrefix(tag, "validate:")

	if tag != `"required"` {
		return "", fmt.Errorf("unsupported tag %s", fieldTag)
	}

	switch fieldType {
	case "string":
		return fmt.Sprintf(
			`	if obj.%s == "" {
		errs = append(errs, fmt.Errorf("%%w: %s required", ErrValidation))
	}`, fieldName, fieldName), nil
	case "uint8":
		return fmt.Sprintf(
			`	if obj.%s == 0 {
		errs = append(errs, fmt.Errorf("%%w: %s required", ErrValidation))
	}`, fieldName, fieldName), nil
	default:
		return "", fmt.Errorf("unsupported type %s", fieldType)
	}
}

func (s *StructInfo) GenerateFileValidator() error {
	fmt.Printf("Generating struct %s validations code\n", s.Name)

	code, err := s.GenerateValidator()
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.Path+"/"+strings.ToLower(s.Name)+"_validator.go", []byte(code), 0644); err != nil {
		return err
	}

	return nil
}

func (s *StructInfo) GenerateFilePackageDefinition() error {
	fmt.Println("Generating package definitions code")

	code, err := s.Generate()
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.Path+"/validators.go", []byte(code), 0644); err != nil {
		return err
	}

	return nil
}

func (s *StructInfo) Generate() (string, error) {
	tmpl, err := template.New("PkgDef").Parse(packageDefinitionTpl)
	if err != nil {
		return "", err
	}

	code := new(bytes.Buffer)
	if err := tmpl.Execute(code, s); err != nil {
		return "", err
	}

	return code.String(), nil
}

func (s *StructInfo) PrintInfo() {
	fmt.Println("Struct:", s.Name)
	fmt.Println("\tHasValidateTag:", s.HasValidateTag)

	for _, f := range s.FieldsInfo {
		fmt.Println("\tField:", f.Name, f.Type, f.Tag)
	}
}
