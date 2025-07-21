package structsinpkg

type Type1 struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"required"`
}

type NoValidateInfo struct {
	Name    string
	Address string
}
