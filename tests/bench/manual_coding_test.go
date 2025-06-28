package benchtests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManualCoding(t *testing.T) {
	data := &StructToValidate{
		FirstName: "First",
		LastName:  "Last",
		Age:       49,
		UserName:  "myusername",
	}

	errors := ManualCodingValidate(data)
	assert.Equal(t, 0, len(errors))
}

func BenchmarkManualCoding(b *testing.B) {
	for b.Loop() {
		data := &StructToValidate{
			FirstName: "First",
			LastName:  "Last",
			Age:       49,
			UserName:  "myusername",
		}

		ManualCodingValidate(data)
	}
}
