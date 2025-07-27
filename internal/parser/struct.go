package parser

import "fmt"

type Struct struct {
	StructParserInfo
	StructAnalyzerInfo
}

type StructParserInfo struct {
	StructName  string
	Path        string
	PackageName string
	Fields      []Field
}

type StructAnalyzerInfo struct {
	HasValidTag bool
}

type Field struct {
	FieldParserInfo
	AnalyzerInfo
}

type FieldParserInfo struct {
	FieldName string
	Type      string
	Tag       string
}

type AnalyzerInfo struct {
	Validations []string
}

func (s *Struct) PrintInfo() {
	fmt.Println("Struct:", s.StructName)
	fmt.Println("\tHasValidTag:", s.HasValidTag)

	for _, f := range s.Fields {
		fmt.Println("\tField:", f.FieldName, f.Type, f.Tag)
	}
}
