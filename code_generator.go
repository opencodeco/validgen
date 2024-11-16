package main

import (
	"fmt"
	"os"
	"strings"
)

func generateCode(structs []StructInfo) error {
	// TODO: validate tags ok?

	for _, structInfo := range structs {
		if !structInfo.HasValidateTag {
			continue
		}

		if err := generatePackageDefinition(structInfo); err != nil {
			return err
		}

		if err := generateFileValidator(structInfo); err != nil {
			return err
		}
	}

	return nil
}

func generatePackageDefinition(structInfo StructInfo) error {
	fmt.Println("Generating package definitions code")

	pd := &PackageDefinitions{
		PackageName: structInfo.PackageName,
	}

	code, err := pd.Generate()
	if err != nil {
		return err
	}

	if err := os.WriteFile(structInfo.Path+"/validators.go", []byte(code), 0644); err != nil {
		return err
	}

	return nil
}

func generateFileValidator(structInfo StructInfo) error {
	fmt.Printf("Generating struct %s validations code\n", structInfo.Name)

	fv := &FileValidator{
		FileHeader: FileHeader{
			PackageName: structInfo.PackageName,
		},
		StructInfo: structInfo,
	}

	code, err := fv.Generate()
	if err != nil {
		return err
	}

	if err := os.WriteFile(structInfo.Path+"/"+strings.ToLower(structInfo.Name)+"_validator.go", []byte(code), 0644); err != nil {
		return err
	}

	return nil
}
