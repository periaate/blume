package blume

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEitherFail(t *testing.T) {
	var base S
	b := Ok(base)
	assert.True(t, b.IsOk(), "Ok(\"\") returned false")
	assert.False(t, b.Fail().IsOk(), "result.Fail returned true")
}

func TestEitherPass(t *testing.T) {
	var base S
	b := Err[S](base)
	assert.False(t, b.IsOk(), "Err(\"\") returned true")
	assert.True(t, b.Pass("").IsOk(), "result.Pass returned false")
}

func TestEitherAuto(t *testing.T) {
	var base S
	b := Err[S](base)
	assert.False(t, b.IsOk(), "Err(\"\") returned true")
	assert.True(t, b.Pass("").IsOk(), "result.Pass returned false")
}
