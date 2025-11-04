package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Generating tests files")

	if err := generateValidationTypesEndToEndTests(); err != nil {
		fmt.Printf("error generating validation types end-to-end tests: %s\n", err)
		os.Exit(1)
	}

	if err := generateValidationCodeUnitTests(); err != nil {
		fmt.Printf("error generating validation code unit tests: %s\n", err)
		os.Exit(1)
	}

	if err := generateFunctionCodeUnitTests(); err != nil {
		fmt.Printf("error generating function code unit tests: %s\n", err)
		os.Exit(1)
	}

	if err := generateComparativePerformanceTests(); err != nil {
		fmt.Printf("error generating comparative performance tests: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Generating done")
}
