package main

import (
	"fmt"
)

func main() {
	fmt.Println("Generating tests files")

	generateValidationTypesEndToEndTests()
	generateValidationCodeUnitTests()
	generateFunctionCodeUnitTests()

	fmt.Println("Generating done")
}
