package benchtests

type StructToValidate struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
	UserName  string `validate:"gte=5,lte=10"`
	Optional  string
}
