package benchtests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=18,lte=130"`
	UserName  string `validate:"min=5,max=10"`
	Optional  string
}

func TestValidator(t *testing.T) {
	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	data := &StructValidator{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "myusername",
	}

	err := validate.Struct(data)

	assert.NoError(t, err)
}

func BenchmarkValidator(b *testing.B) {
	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	for b.Loop() {

		data := &StructValidator{
			FirstName: "First",
			LastName:  "Last",
			Age:       49,
			UserName:  "myusername",
		}

		validate.Struct(data)
	}
}
