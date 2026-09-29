package main

import (
	"log"

	"github.com/opencodeco/validgen/tests/endtoend/commonexample"
)

func commonValidatorExampleTests() {
	log.Println("starting common validator example tests")

	invalid := &commonexample.User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Gender:         "male",
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses: []*commonexample.Address{{
			Street: "Eavesdown Docks",
			Planet: "Persphone",
			Phone:  "none",
		}},
	}
	assertExpectedErrorMsgs("playground example", commonexample.UserValidate(invalid), []string{
		"Age must be <= 130",
		"FavouriteColor must be a valid color",
		"City is required",
	})

	ok := validCommonUser()
	assertExpectedErrorMsgs("playground example valid", commonexample.UserValidate(ok), nil)

	for _, color := range []string{
		"#000",
		"#aabbcc",
		"#aabbccdd",
		"rgb(255,255,255)",
		"rgba(0,0,0,1)",
		"hsl(360,100%,50%)",
		"hsla(120,100%,50%,0.5)",
	} {
		user := validCommonUser()
		user.FavouriteColor = color
		assertExpectedErrorMsgs("color "+color, commonexample.UserValidate(user), nil)
	}

	for _, gender := range []string{"male", "female", "prefer_not_to"} {
		user := validCommonUser()
		user.Gender = gender
		assertExpectedErrorMsgs("gender "+gender, commonexample.UserValidate(user), nil)
	}

	badGender := validCommonUser()
	badGender.Gender = "other"
	assertExpectedErrorMsgs("gender other", commonexample.UserValidate(badGender), []string{
		"Gender must be one of 'male' 'female' 'prefer_not_to'",
	})

	badEmail := validCommonUser()
	badEmail.Email = "joeybloggs.gmail.com"
	assertExpectedErrorMsgs("email without at", commonexample.UserValidate(badEmail), []string{
		"Email must be a valid email",
	})

	empty := &commonexample.User{}
	assertExpectedErrorMsgs("empty user", commonexample.UserValidate(empty), []string{
		"FirstName is required",
		"LastName is required",
		"Email is required",
		"Email must be a valid email",
		"Gender must be one of 'male' 'female' 'prefer_not_to'",
		"FavouriteColor must be a valid color",
		"Addresses must not be empty",
	})

	nilElem := validCommonUser()
	nilElem.Addresses = []*commonexample.Address{nil}
	assertExpectedErrorMsgs("nil address element", commonexample.UserValidate(nilElem), []string{
		"Addresses is required",
	})

	missingPhone := validCommonUser()
	missingPhone.Addresses = append(missingPhone.Addresses, &commonexample.Address{
		Street: "second",
		City:   "city",
		Planet: "Persphone",
	})
	assertExpectedErrorMsgs("second address phone", commonexample.UserValidate(missingPhone), []string{
		"Phone is required",
	})

	colors := &commonexample.Palette{
		Colors: []string{"#fff", "nope"},
		Labels: map[string]string{"home": "red"},
	}
	assertExpectedErrorMsgs("dive iscolor", commonexample.PaletteValidate(colors), []string{
		"Colors must be a valid color",
	})

	paletteOK := &commonexample.Palette{
		Colors: []string{"#fff", "rgb(0,0,0)"},
		Labels: map[string]string{"not-a-color": "green"},
	}
	assertExpectedErrorMsgs("dive color and map value", commonexample.PaletteValidate(paletteOK), nil)

	badLabel := &commonexample.Palette{
		Colors: []string{"#000"},
		Labels: map[string]string{"home": "purple"},
	}
	assertExpectedErrorMsgs("dive oneof map value", commonexample.PaletteValidate(badLabel), []string{
		"Labels must be one of 'red' 'green' 'blue'",
	})

	log.Println("common validator example tests ok")
}

func validCommonUser() *commonexample.User {
	return &commonexample.User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            30,
		Email:          "Badger.Smith@gmail.com",
		Gender:         "male",
		FavouriteColor: "#000",
		Addresses: []*commonexample.Address{{
			Street: "Eavesdown Docks",
			City:   "London",
			Planet: "Persphone",
			Phone:  "none",
		}},
	}
}
