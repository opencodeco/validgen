package main

import (
	"fmt"
)

func main() {
	fmt.Println("Generating tests files")

	generateValidationTypesEndToEndTests()
	generateValidationCodeUnitTests()
	generateFunctionCodeUnitTests()
	generateComparativePerformanceTests()

	fmt.Println("Generating done")
}
