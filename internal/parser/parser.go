package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/opencodeco/validgen/internal/common"
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

	var err error
	structs := []*Struct{}

	filename := filepath.Base(fullpath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	packageName := ""
	currentStruct := &Struct{}
	imports := map[string]Import{}

	ast.Inspect(f, func(n ast.Node) bool {
		switch v := n.(type) {
		case (*ast.File):
			packageName, imports = extractPkgAndImports(v)
		case (*ast.TypeSpec):
			currentStruct = extractStructDefinition(v.Name.Name, fullpath, packageName, imports)
		case (*ast.StructType):
			err = extractAndAppendStructFields(v, packageName, currentStruct)
			if err != nil {
				return false
			}

			structs = append(structs, currentStruct)
		}

		return true
	})

	return structs, err
}

func extractPkgAndImports(f *ast.File) (string, map[string]Import) {
	packageName := f.Name.Name
	imports := map[string]Import{}

	for _, importLink := range f.Imports {
		path, _ := strconv.Unquote(importLink.Path.Value)
		name := ""
		if importLink.Name == nil {
			idx := strings.LastIndexByte(path, '/')
			name = path[idx+1:]
		} else {
			name = importLink.Name.Name
		}
		imports[name] = Import{
			Name: name,
			Path: path,
		}
	}

	return packageName, imports
}

func extractStructDefinition(name, fullpath, packageName string, imports map[string]Import) *Struct {
	return &Struct{
		StructName:  name,
		Path:        "./" + filepath.Dir(fullpath),
		PackageName: packageName,
		Imports:     imports,
		Fields:      []Field{},
	}
}

func extractAndAppendStructFields(structType *ast.StructType, packageName string, cstruct *Struct) error {
	if structType.Fields == nil {
		return nil
	}

	for _, field := range structType.Fields.List {
		fieldTag := ""
		if field.Tag != nil {
			fieldTag = field.Tag.Value
		}

		fieldTag, err := extractTag(fieldTag)
		if err != nil {
			return err
		}

		fieldType, err := extractCompleteType(common.FieldType{}, field.Type, packageName)
		if err != nil {
			return err
		}

		if fieldType.BaseType != "" {
			for _, name := range field.Names {
				cstruct.Fields = append(cstruct.Fields, Field{
					FieldName: name.Name,
					Type:      fieldType,
					Tag:       fieldTag,
				})
			}
		}
	}

	return nil
}

func extractCompleteType(fType common.FieldType, expr ast.Expr, packageName string) (common.FieldType, error) {
	var err error

	switch v := expr.(type) {
	case *ast.Ident:
		// Single type (string, int, etc.)
		fType.BaseType = v.Name
		if !fType.IsGoType() {
			fType.BaseType = common.KeyPath(packageName, fType.BaseType)
		}
		return fType, nil
	case *ast.ArrayType:
		// Slice or array type
		fType, err = extractCompleteType(fType, v.Elt, packageName)
		if err != nil {
			return common.FieldType{}, err
		}

		fType.Size = ""
		if v.Len != nil {
			// Array with fixed size
			basicLit, ok := v.Len.(*ast.BasicLit)
			if !ok {
				return common.FieldType{}, fmt.Errorf("cannot find the basic literal: %T", v.Len)
			}

			fType.Size = basicLit.Value
			fType.ComposedType += "[N]"
		} else {
			fType.ComposedType += "[]"
		}
		return fType, nil
	case *ast.SelectorExpr:
		// Nested type from another package
		typeID, ok := v.X.(*ast.Ident)
		if !ok {
			return common.FieldType{}, fmt.Errorf("cannot find the type id: %T", v.X)
		}

		nestedPkgName := typeID.Name
		fType, err = extractCompleteType(fType, v.X, nestedPkgName)
		if err != nil {
			return common.FieldType{}, err
		}

		fType = extractNestedFieldType(nestedPkgName, v.Sel.Name)
		return fType, nil

	case *ast.MapType:
		// Map type
		fType.ComposedType += "map"
		fType, err = extractCompleteType(fType, v.Key, packageName)
		if err != nil {
			return common.FieldType{}, err
		}

		return fType, nil
	case *ast.StarExpr:
		// Pointer type
		fType.ComposedType += "*"
		fType, err = extractCompleteType(fType, v.X, packageName)
		if err != nil {
			return common.FieldType{}, err
		}

		return fType, nil
	}

	return common.FieldType{}, nil
}

func extractTag(fieldTag string) (string, error) {

	if fieldTag == "" {
		return "", nil
	}

	rFieldTag, err := strconv.Unquote(fieldTag)
	if err != nil {
		return "", err
	}

	return rFieldTag, nil
}

func extractNestedFieldType(nestedPkgName, baseType string) common.FieldType {

	return common.FieldType{
		BaseType:     common.KeyPath(nestedPkgName, baseType),
		ComposedType: "",
		Size:         "",
	}
}
