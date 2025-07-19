package validgen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const verifyTag = "verify"

func parseFile(fullpath string) ([]Struct, error) {
	fmt.Printf("Parsing %s\n", fullpath)

	src, err := os.ReadFile(fullpath)
	if err != nil {
		return nil, err
	}

	structs, err := parseStructs(fullpath, string(src))
	if err != nil {
		return nil, err
	}

	return structs, nil
}

func parseStructs(fullpath, src string) ([]Struct, error) {

	filename := filepath.Base(fullpath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	var structs []Struct
	packageName := ""

	ast.Inspect(f, func(n ast.Node) bool {
		if fileInfo, ok := n.(*ast.File); ok {
			packageName = fileInfo.Name.Name
		}

		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			structs = append(structs, Struct{
				Name:        typeSpec.Name.Name,
				Path:        "./" + filepath.Dir(fullpath),
				PackageName: packageName,
			})
		}

		if structType, ok := n.(*ast.StructType); ok {
			currentStruct := &structs[len(structs)-1]

			for _, field := range structType.Fields.List {
				if ident, ok := field.Type.(*ast.Ident); ok {
					fieldType := ident.Name
					fieldTag := ""
					if field.Tag != nil {
						fieldTag = field.Tag.Value
						fieldTag, _ = strconv.Unquote(fieldTag)
					}

					fieldValidations, hasVerifyTag := parseFieldValidations(fieldTag)
					if hasVerifyTag {
						currentStruct.HasVerifyTag = true
					}

					for _, name := range field.Names {
						currentStruct.Fields = append(currentStruct.Fields, Field{
							Name:        name.Name,
							Type:        fieldType,
							Tag:         fieldTag,
							Validations: fieldValidations,
						})
					}
				}
			}
		}

		return true
	})

	return structs, nil
}

func parseFieldValidations(fieldTag string) ([]string, bool) {
	fieldValidations := []string{}
	hasVerifyTag := false
	prefixToSearch := verifyTag + ":"

	if strings.HasPrefix(fieldTag, prefixToSearch) {
		hasVerifyTag = true
		tagWithoutPrefix, _ := strings.CutPrefix(fieldTag, prefixToSearch)
		tagWithoutQuotes, _ := strconv.Unquote(tagWithoutPrefix)
		fieldValidations = strings.Split(tagWithoutQuotes, ",")
	}

	return fieldValidations, hasVerifyTag
}
