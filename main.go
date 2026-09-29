package main

import (
	"flag"
	"log"
	"os"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/codegenerator"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/opencodeco/validgen/internal/pkgwriter"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	unmarshalJSON := fs.Bool("unmarshal-json", false, "generate UnmarshalJSON methods that validate after decoding")
	fs.Usage = func() {
		log.Printf("Usage:\n\tvalidgen [-unmarshal-json] <path>\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	if fs.NArg() != 1 {
		fs.Usage()
		os.Exit(1)
	}

	parsedStructs, err := parser.ExtractStructs(fs.Arg(0))
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

	pkgs, err := codegenerator.GenerateCode(analyzedStructs, codegenerator.Options{
		UnmarshalJSON: *unmarshalJSON,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = pkgwriter.Writer(pkgs)
	if err != nil {
		log.Fatal(err)
	}
}
