package main

import (
	"fmt"
)

var structFileHeaderTpl = `package validators

import (
	"fmt"

	"%[1]s"
)
`

type FileHeader struct {
	ImportPath string
}

func (p *FileHeader) Generate() (string, error) {
	code := fmt.Sprintf(structFileHeaderTpl, p.ImportPath)

	return code, nil
}
