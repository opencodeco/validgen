package benchtests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

func TestValidator(t *testing.T) {
	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	data := &StructToValidate{
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

		data := &StructToValidate{
			FirstName: "First",
			LastName:  "Last",
			Age:       49,
			UserName:  "myusername",
		}

		validate.Struct(data)
	}
}
