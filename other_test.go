package blume

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Return[A, B any](a A, b B) (A, B) { return a, b }

func TestReturn(t *testing.T) {
	// Test with int and nil error
	a, b := Return[int, error](10, nil)
	assert.Equal(t, 10, a)
	assert.Nil(t, b)

	// Test with string and non-nil error
	str, err := Return[string, error]("test", assert.AnError)
	assert.Equal(t, "test", str)
	assert.Equal(t, assert.AnError, err)

	// Test with custom types
	type Custom struct {
		Value int
	}
	customA, customB := Return[Custom, string](Custom{Value: 42}, "message")
	assert.Equal(t, Custom{Value: 42}, customA)
	assert.Equal(t, "message", customB)
}

func TestOr(t *testing.T) {
	// Basic behavior
	assert.Equal(t, 10, Or(0, 10))
	assert.Equal(t, "default", Or("default", ""))

	// With boolean handle
	assert.Equal(t, 20, Or(0, 20, true))
	assert.Equal(t, 0, Or(0, 20, false))

	// With error handle
	assert.Equal(t, "success", Or("default", "success", nil))
	assert.Equal(t, "default", Or("default", "failure", assert.AnError))

	// Edge cases
	assert.Equal(t, "default", Or("default", "", nil))
	assert.Equal(t, 100, Or(100, 0, false))
}

func TestMust(t *testing.T) {
	// Basic behavior
	assert.Equal(t, 10, Must(10))
	assert.Equal(t, "success", Must("success"))

	// With boolean handle
	assert.Equal(t, 20, Must(20, true))

	assert.Panics(t, func() {
		Must(20, false)
	}, "must called with false bool")

	// With error handle
	assert.Equal(t, "success", Must("success", nil))

	assert.Panics(t, func() {
		Must("failure", assert.AnError)
	}, "must called with non nil error")
}

func TestIntegration(t *testing.T) {
	// Combining Return, Or, and Must
	a, b := Return[int, error](Or(0, 10, true), nil)
	assert.Equal(t, 10, a)
	assert.Nil(t, b)

	// Ensuring panic is handled
	assert.Panics(t, func() {
		Return[int, error](Must(Or(0, 10, false)), assert.AnError)
	})
}
