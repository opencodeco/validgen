package benchtests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidGen(t *testing.T) {
	data := &StructToValidate{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "myusername",
	}

	errors := ValidGenValidate(data)
	assert.Equal(t, 0, len(errors))
}

func BenchmarkValidGen(b *testing.B) {
	for b.Loop() {
		data := &StructToValidate{
			FirstName: "First",
			LastName:  "Last",
			Age:       49,
			UserName:  "myusername",
		}

		ValidGenValidate(data)
	}
}
