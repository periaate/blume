package blume

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	/*
		missing
		- structs
		- all primitive types
		- pointer types
		- deep equality
		- etc.
	*/
	assert.Equal(t, Arr("a", "b", "c"), Arr("a", "b", "c").Unique())
	assert.Equal(t, Arr("a", "b", "c"), Arr("a", "b", "c", "c").Unique())
	assert.Equal(t, Arr([2]int{0, 1}, [2]int{1, 2}), Arr([2]int{0, 1}, [2]int{1, 2}).Unique())
	assert.Equal(t, Arr([2]int{0, 1}, [2]int{1, 2}), Arr([2]int{0, 1}, [2]int{1, 2}, [2]int{1, 2}).Unique())
}
