package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
)

func ExtractStructs(path string) ([]*Struct, error) {
	files, err := findFiles(path)
	if err != nil {
		return nil, err
	}

	structs := []*Struct{}

	for _, file := range files {
		parsedStructs, err := parseFile(file)
		if err != nil {
			return nil, err
		}

		structs = append(structs, parsedStructs...)
	}

	return structs, nil
}

func findFiles(path string) ([]string, error) {
	files := []string{}
	if err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		files = append(files, path)

		return nil

	}); err != nil {
		return nil, err
	}

	return files, nil
}

func parseFile(fullpath string) ([]*Struct, error) {
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

func parseStructs(fullpath, src string) ([]*Struct, error) {

	structs := []*Struct{}

	filename := filepath.Base(fullpath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	packageName := ""
	currentStruct := &Struct{}

	ast.Inspect(f, func(n ast.Node) bool {
		if fileInfo, ok := n.(*ast.File); ok {
			packageName = fileInfo.Name.Name
		}

		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			currentStruct = &Struct{
				StructParserInfo: StructParserInfo{
					StructName:  typeSpec.Name.Name,
					Path:        "./" + filepath.Dir(fullpath),
					PackageName: packageName,
				},
			}
		}

		if structType, ok := n.(*ast.StructType); ok {
			for _, field := range structType.Fields.List {
				if ident, ok := field.Type.(*ast.Ident); ok {
					fieldType := ident.Name
					fieldTag := ""
					if field.Tag != nil {
						fieldTag = field.Tag.Value
						fieldTag, _ = strconv.Unquote(fieldTag)
					}

					for _, name := range field.Names {
						currentStruct.Fields = append(currentStruct.Fields, Field{
							FieldParserInfo: FieldParserInfo{
								FieldName: name.Name,
								Type:      fieldType,
								Tag:       fieldTag,
							},
						})
					}
				}
			}

			structs = append(structs, currentStruct)
		}

		return true
	})

	return structs, nil
}
