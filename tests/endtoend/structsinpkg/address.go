package structsinpkg

type Address struct {
	Street string `valid:"required"`
	City   string `valid:"required"`
}
