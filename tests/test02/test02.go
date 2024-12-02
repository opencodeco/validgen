package main

import (
	"fmt"

	"github.com/alexgarzao/myvalidator/tests/test02/structsinpkg"
)

func main() {
	u := &structsinpkg.User{}
	if err := structsinpkg.UserValidate(u); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u)
	}
}
