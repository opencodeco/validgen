package parser

import "github.com/opencodeco/validgen/internal/common"

type Struct struct {
	StructName  string
	Path        string
	PackageName string
	Fields      []Field
	Imports     map[string]Import
}

type Field struct {
	FieldName string
	Type      common.FieldType
	Tag       string
}

type Import struct {
	Name string
	Path string
}
