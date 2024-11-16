package main

import (
	"fmt"
)

var structFileHeaderTpl = `package %[1]s

import (
	"fmt"
)
`

type FileHeader struct {
	PackageName string
}

func (p *FileHeader) Generate() (string, error) {
	code := fmt.Sprintf(structFileHeaderTpl, p.PackageName)

	return code, nil
}
