package parser

import "fmt"

type Struct struct {
	StructName  string
	Path        string
	PackageName string
	Fields      []Field
	StructAnalyzerInfo
}

type StructAnalyzerInfo struct {
	HasValidTag bool
}

type Field struct {
	FieldName string
	Type      string
	Tag       string
	FieldAnalyzerInfo
}

type FieldAnalyzerInfo struct {
	Validations []string
}

func (s *Struct) PrintInfo() {
	fmt.Println("Struct:", s.StructName)
	fmt.Println("\tHasValidTag:", s.HasValidTag)

	for _, f := range s.Fields {
		fmt.Println("\tField:", f.FieldName, f.Type, f.Tag)
	}
}
