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
			err = appendFields(v, packageName, currentStruct)
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

func appendFields(structType *ast.StructType, packageName string, cstruct *Struct) error {
	if structType.Fields == nil {
		return nil
	}

	for _, field := range structType.Fields.List {
		appendFieldNames := false
		fieldType := common.FieldType{}
		fieldTag := ""
		if field.Tag != nil {
			fieldTag = field.Tag.Value
		}

		switch v := field.Type.(type) {
		case *ast.Ident:
			fieldType.BaseType = v.Name
			fieldType, fieldTag = extractFieldTypeAndTag(packageName, fieldType, fieldTag)
			appendFieldNames = true

		case *ast.ArrayType:
			fieldType.ComposedType = "[]"
			ident, ok := v.Elt.(*ast.Ident)
			if !ok {
				return fmt.Errorf("cannot find the identifier: %T", v.Elt)
			}

			fieldType.BaseType = ident.Name
			fieldType.Size = ""
			if v.Len != nil {
				// Array with fixed size
				basicLit, ok := v.Len.(*ast.BasicLit)
				if !ok {
					return fmt.Errorf("cannot find the basic literal: %T", v.Len)
				}

				fieldType.Size = basicLit.Value
				fieldType.ComposedType = "[N]"
			}
			fieldType, fieldTag = extractFieldTypeAndTag(packageName, fieldType, fieldTag)
			appendFieldNames = true

		case *ast.SelectorExpr:
			ident, ok := v.X.(*ast.Ident)
			if !ok {
				return fmt.Errorf("cannot find the identifier: %T", v.X)
			}

			nestedPkgName := ident.Name
			fieldType, fieldTag = extractNestedFieldTypeAndTag(nestedPkgName, v.Sel.Name, fieldTag)
			appendFieldNames = true

		case *ast.MapType:
			ident, ok := v.Key.(*ast.Ident)
			if !ok {
				return fmt.Errorf("cannot find the identifier: %T", v.Key)
			}

			fieldType.ComposedType = "map"
			fieldType.BaseType = ident.Name
			_, fieldTag = extractFieldTypeAndTag(packageName, fieldType, fieldTag)
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

	return nil
}

func extractFieldTypeAndTag(packageName string, fieldType common.FieldType, fieldTag string) (common.FieldType, string) {
	rFieldType := fieldType

	if !fieldType.IsGoType() {
		rFieldType.BaseType = common.KeyPath(packageName, fieldType.BaseType)
	}

	rFieldTag := ""
	if fieldTag != "" {
		rFieldTag = fieldTag
		rFieldTag, _ = strconv.Unquote(rFieldTag)
	}

	return rFieldType, rFieldTag
}

func extractNestedFieldTypeAndTag(nestedPkgName, baseType, fieldTag string) (common.FieldType, string) {
	rFieldType := common.KeyPath(nestedPkgName, baseType)
	rFieldTag, _ := strconv.Unquote(fieldTag)

	return common.FieldType{
		BaseType:     rFieldType,
		ComposedType: "",
		Size:         "",
	}, rFieldTag
}
