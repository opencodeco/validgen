package benchtests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
