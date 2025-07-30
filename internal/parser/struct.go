package parser

type Struct struct {
	StructName  string
	Path        string
	PackageName string
	Fields      []Field
	Imports     map[string]Import
}

type Field struct {
	FieldName string
	Type      string
	Tag       string
}

type Import struct {
	Name string
	Path string
}
