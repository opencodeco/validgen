package main

import (
	"log"
)

func main() {
	fullpath := "./tests/test01/test01.go"

	structs, err := parseFile(fullpath)
	if err != nil {
		log.Fatal(err)
	}

	writeStructsInfo(structs)

	if err := generateCode(structs); err != nil {
		log.Fatal(err)
	}
}

func writeStructsInfo(structs []StructInfo) {
	for _, structInfo := range structs {
		structInfo.PrintInfo()
	}
}
