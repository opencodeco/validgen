package codegenerator

import (
	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/parser"
)

type Pkg struct {
	Name    string
	Path    string
	Imports map[string]parser.Import
	Structs map[string]*Struct
}

type Struct struct {
	*analyzer.Struct
	ValidatorFuncCode string
}
