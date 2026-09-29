package commonexample

// User is the common go-playground/validator example.
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"` // alias for hexcolor|rgb|rgba|hsl|hsla
	Addresses      []*Address `validate:"required,dive,required"`
}

// Address houses a user's address information.
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// Palette checks oneof and iscolor on elements after dive.
// Map keys are not validated.
type Palette struct {
	Colors []string          `validate:"dive,iscolor"`
	Labels map[string]string `validate:"dive,oneof=red green blue"`
}
