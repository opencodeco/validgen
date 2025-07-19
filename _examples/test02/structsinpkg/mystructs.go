package structsinpkg

type User struct {
	FirstName string `verify:"required"`
	LastName  string `verify:"required"`
	Age       uint8  `verify:"required"`
}

type NoVerifyInfo struct {
	Name    string
	Address string
}
