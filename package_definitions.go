package main

import "fmt"

var packageDefinitionTpl = `package %[1]s

import (
	"errors"
)

var ErrValidation = errors.New("validation error")
`

type PackageDefinitions struct {
	PackageName string
}

func (mp *PackageDefinitions) Generate() (string, error) {
	code := fmt.Sprintf(packageDefinitionTpl, mp.PackageName)

	return code, nil
}
