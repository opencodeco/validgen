package main

import (
	"log"
	"os"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/codegenerator"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/internal/pkgwriter"
)

func main() {
	argsWithoutCmd := os.Args[1:]
	if len(argsWithoutCmd) != 1 {
		log.Fatal("Invalid parameters:\n\tvalidgen <path>\n")
	}

	parsedStructs, err := parser.ExtractStructs(argsWithoutCmd[0])
	if err != nil {
		log.Fatal(err)
	}

	analyzedStructs, err := analyzer.AnalyzeStructs(parsedStructs)
	if err != nil {
		log.Fatal(err)
	}

	for _, st := range analyzedStructs {
		st.PrintInfo()
	}

	pkgs, err := codegenerator.GenerateCode(analyzedStructs)
	if err != nil {
		log.Fatal(err)
	}

	err = pkgwriter.Writer(pkgs)
	if err != nil {
		log.Fatal(err)
	}
}
