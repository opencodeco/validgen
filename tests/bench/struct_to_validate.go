package benchtests

type StructToValidate struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=18,lte=130"`
	UserName  string `validate:"min=5,max=10"`
	Optional  string
}
