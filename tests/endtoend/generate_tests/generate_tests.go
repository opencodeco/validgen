package main

import (
	"log"
)

func main() {
	log.Println("Generating tests files")

	generateNumericTests()
	generateTestCases()

	log.Println("Generating done")
}
