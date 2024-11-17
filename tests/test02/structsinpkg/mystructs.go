package structsinpkg

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
