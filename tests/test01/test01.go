package main

import (
	"fmt"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"required"`
	// Age       uint8  `validate:"gte=0,lte=130"`
	// Optional  string
}

type NoValidateInfo struct {
	Name    string
	Address string
}

func main() {
	u := &User{}
	if err := UserValidate(u); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u)
	}
}
