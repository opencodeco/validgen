package main

import (
	"fmt"
)

type User struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"required"`
}

type NoValidTag struct {
	Name    string
	Address string
}

func main() {
	u1 := &User{}
	if err := UserValidate(u1); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u1, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u1)
	}

	u2 := &User{
		FirstName: "First",
		LastName:  "Last",
		Age:       18,
	}

	if err := UserValidate(u2); err != nil {
		fmt.Printf("User: %+v Error: %s\n", u2, err)
	} else {
		fmt.Printf("User: %+v is valid\n", u2)
	}
}
