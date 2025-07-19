package structsinpkg

type Type1 struct {
	FirstName string `verify:"required"`
	LastName  string `verify:"required"`
	Age       uint8  `verify:"required"`
}

type NoValidateInfo struct {
	Name    string
	Address string
}
