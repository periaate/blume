package based

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoding(t *testing.T) {
	cross := func(val int, exp string) {
		assert.Equal(t, Encode(val), exp)
		assert.Equal(t, Decode(exp), val)
		assert.Equal(t, Decode(Encode(val)), val)
		assert.Equal(t, Encode(Decode(exp)), exp)
	}
	cross(0, "0")
	cross(9, "9")
	cross(10, "a")
	cross(60, "10")
	cross(120, "20")
}
