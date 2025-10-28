package main

import (
	"fmt"
)

func main() {
	fmt.Println("Generating tests files")

	generateValidationTypesTests()
	generateValidationCodeTests()

	fmt.Println("Generating done")
}
