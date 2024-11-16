package main

var packageDefinitionTpl = `package validators

import (
	"errors"
)

var ErrValidation = errors.New("validation error")
`

type PackageDefinitions struct {
}

func (mp *PackageDefinitions) Generate() (string, error) {
	code := packageDefinitionTpl

	return code, nil
}
