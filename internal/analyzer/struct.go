package analyzer

import (
	"fmt"

	"github.com/opencodeco/validgen/internal/parser"
)

type Struct struct {
	parser.Struct
	HasValidTag       bool
	FieldsValidations []Validations
}

type Validations struct {
	Validations []string
}

func (s *Struct) PrintInfo() {
	fmt.Println("Struct:", s.StructName)
	fmt.Println("\tHasValidTag:", s.HasValidTag)

	for _, f := range s.Fields {
		fmt.Println("\tField:", f.FieldName, f.Type, f.Tag)
	}
}
