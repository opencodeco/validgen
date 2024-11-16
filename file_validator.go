package main

import (
	"fmt"
)

var fileValidatorTpl = "%s\n%s"

// FileValidator template has the following structure:
// - File header
// - Function to validate the struct
type FileValidator struct {
	FileHeader FileHeader
	StructInfo StructInfo
}

func (fv *FileValidator) Generate() (string, error) {
	header, err := fv.FileHeader.Generate()
	if err != nil {
		return "", err
	}

	code, err := fv.StructInfo.GenerateFuncValidator()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(fileValidatorTpl, header, code), nil
}
