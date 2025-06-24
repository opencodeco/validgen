package main

import (
	"log"
	"os"

	"github.com/opencodeco/validgen/validgen"
)

func main() {
	argsWithoutCmd := os.Args[1:]
	if len(argsWithoutCmd) == 0 {
		log.Fatal("Invalid parameters:\n\tvalidatorgen <path>\n")
	}

	if err := validgen.FindFiles(argsWithoutCmd[0]); err != nil {
		log.Fatal(err)
	}
}
