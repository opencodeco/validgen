package structsinpkg

type Type1 struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"required"`
}

type NoValidateInfo struct {
	Name    string
	Address string
}
