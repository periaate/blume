package fsio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	r := args[string]([]string{})
	assert.True(t, r.Ok)
}
