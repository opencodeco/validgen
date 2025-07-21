package benchtests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructValidGen struct {
	FirstName string `verify:"required"`
	LastName  string `verify:"required"`
	Age       uint8  `verify:"gte=18,lte=130"`
	UserName  string `verify:"min=5,max=10"`
	Optional  string
}

func TestValidGen(t *testing.T) {
	data := &StructValidGen{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "myusername",
	}

	errors := StructValidGenValidate(data)
	assert.Equal(t, 0, len(errors))
}

func BenchmarkValidGen(b *testing.B) {
	for b.Loop() {
		data := &StructValidGen{
			FirstName: "First",
			LastName:  "Last",
			Age:       49,
			UserName:  "myusername",
		}

		StructValidGenValidate(data)
	}
}
