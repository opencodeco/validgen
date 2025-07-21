package main

import (
	"fmt"
)

type User struct {
	Email1 string `valid:"required,email"`
	Email2 string `valid:"email"`
}

func main() {
	// Test case 1: Empty required email (should fail)
	u1 := &User{
		Email1: "",
		Email2: "",
	}
	if errs := UserValidate(u1); len(errs) > 0 {
		fmt.Printf("User1: %+v Errors: ", u1)
		for _, err := range errs {
			fmt.Printf("%s; ", err)
		}
		fmt.Println()
	} else {
		fmt.Printf("User1: %+v is valid\n", u1)
	}

	// Test case 2: Invalid required email (should fail)
	u2 := &User{
		Email1: "invalid.email",
		Email2: "",
	}
	if errs := UserValidate(u2); len(errs) > 0 {
		fmt.Printf("User2: %+v Errors: ", u2)
		for _, err := range errs {
			fmt.Printf("%s; ", err)
		}
		fmt.Println()
	} else {
		fmt.Printf("User2: %+v is valid\n", u2)
	}

	// Test case 3: Valid required email, empty optional email (should pass)
	u3 := &User{
		Email1: "valid@example.com",
		Email2: "",
	}
	if errs := UserValidate(u3); len(errs) > 0 {
		fmt.Printf("User3: %+v Errors: ", u3)
		for _, err := range errs {
			fmt.Printf("%s; ", err)
		}
		fmt.Println()
	} else {
		fmt.Printf("User3: %+v is valid\n", u3)
	}

	// Test case 4: Valid required email, valid optional email (should pass)
	u4 := &User{
		Email1: "user@domain.com",
		Email2: "optional@test.org",
	}
	if errs := UserValidate(u4); len(errs) > 0 {
		fmt.Printf("User4: %+v Errors: ", u4)
		for _, err := range errs {
			fmt.Printf("%s; ", err)
		}
		fmt.Println()
	} else {
		fmt.Printf("User4: %+v is valid\n", u4)
	}

	// Test case 5: Valid required email, invalid optional email (should fail)
	u5 := &User{
		Email1: "user@domain.com",
		Email2: "invalid.email",
	}
	if errs := UserValidate(u5); len(errs) > 0 {
		fmt.Printf("User5: %+v Errors: ", u5)
		for _, err := range errs {
			fmt.Printf("%s; ", err)
		}
		fmt.Println()
	} else {
		fmt.Printf("User5: %+v is valid\n", u5)
	}
}
