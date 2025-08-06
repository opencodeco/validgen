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
			appendFields(v, packageName, currentStruct)
			structs = append(structs, currentStruct)
		}

		return true
	})

	return structs, nil
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

func appendFields(structType *ast.StructType, packageName string, cstruct *Struct) {
	if structType.Fields == nil {
		return
	}

	for _, field := range structType.Fields.List {
		appendFieldNames := false
		fieldType := ""
		fieldTag := ""
		if field.Tag != nil {
			fieldTag = field.Tag.Value
		}

		switch v := field.Type.(type) {
		case *ast.Ident:
			fieldType = v.Name
			fieldType, fieldTag = extractFieldTypeAndTag(packageName, fieldType, fieldTag)
			appendFieldNames = true

		case *ast.ArrayType:
			fieldType = v.Elt.(*ast.Ident).Name
			fieldType, fieldTag = extractSliceFieldTypeAndTag(packageName, fieldType, fieldTag)
			appendFieldNames = true

		case *ast.SelectorExpr:
			nestedPkgName := v.X.(*ast.Ident).Name
			fieldType, fieldTag = extractNestedFieldTypeAndTag(nestedPkgName, v.Sel.Name, fieldTag)
			appendFieldNames = true
		}

		if appendFieldNames {
			for _, name := range field.Names {
				cstruct.Fields = append(cstruct.Fields, Field{
					FieldName: name.Name,
					Type:      fieldType,
					Tag:       fieldTag,
				})
			}
		}
	}
}

func extractFieldTypeAndTag(packageName, fieldType, fieldTag string) (string, string) {
	rFieldType := fieldType

	if !common.IsGoType(fieldType) {
		rFieldType = packageName + "." + fieldType
	}
	rFieldTag := ""
	if fieldTag != "" {
		rFieldTag = fieldTag
		rFieldTag, _ = strconv.Unquote(rFieldTag)
	}

	return rFieldType, rFieldTag
}

func extractSliceFieldTypeAndTag(packageName, fieldType, fieldTag string) (string, string) {
	rFieldType := fieldType

	if !common.IsGoType(fieldType) {
		rFieldType = packageName + "." + fieldType
	}

	rFieldType = "[]" + rFieldType

	rFieldTag := ""
	if fieldTag != "" {
		rFieldTag = fieldTag
		rFieldTag, _ = strconv.Unquote(rFieldTag)
	}

	return rFieldType, rFieldTag
}

func extractNestedFieldTypeAndTag(nestedPkgName, fieldType, fieldTag string) (string, string) {
	rFieldType := nestedPkgName + "." + fieldType
	rFieldTag, _ := strconv.Unquote(fieldTag)

	return rFieldType, rFieldTag
}
