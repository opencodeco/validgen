package main

import (
	"fmt"

	"github.com/alexgarzao/myvalidator/tests/test02/structsinpkg"
)

func main() {

	u := &structsinpkg.User{}
	if err := structsinpkg.UserValidate(u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User is valid")
}
