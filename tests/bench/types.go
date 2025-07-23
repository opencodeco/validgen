package benchtests

type StructValidGen struct {
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Age       uint8  `valid:"gte=18,lte=130"`
	UserName  string `valid:"min=5,max=10"`
	Optional  string
}

type StructManualCoding struct {
	FirstName string
	LastName  string
	Age       uint8
	UserName  string
	Optional  string
}
