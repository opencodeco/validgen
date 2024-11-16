package main

import (
	"fmt"
)

var structFileHeaderTpl = `package validators

import (
	"errors"
	"fmt"

	"%[1]s"
)

var ErrValidation = errors.New("validation error")
`

type FileHeader struct {
	ImportPath string
}

func (p *FileHeader) Generate() (string, error) {
	code := fmt.Sprintf(structFileHeaderTpl, p.ImportPath)

	return code, nil
}
