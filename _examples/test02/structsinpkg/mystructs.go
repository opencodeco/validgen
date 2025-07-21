package structsinpkg

type User struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"required"`
}

type NoValidTag struct {
	Name    string
	Address string
}
