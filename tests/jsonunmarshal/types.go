package main

type User struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"gte=18"`
}
