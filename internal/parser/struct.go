package parser

type Struct struct {
	StructName  string
	Path        string
	PackageName string
	Fields      []Field
}

type Field struct {
	FieldName string
	Type      string
	Tag       string
}
