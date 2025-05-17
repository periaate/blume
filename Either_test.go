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
	var v S
	s := Some(v)
	n := None[S]()
	o := Ok(v)
	e := Err[S](v)

	assert.True(t, s.IsOk())
	assert.True(t, o.IsOk())
	assert.False(t, n.IsOk())
	assert.False(t, e.IsOk())

	assert.True(t, Either[S, bool]{}.Auto(s).IsOk())
	assert.True(t, Either[S, bool]{}.Auto(o).IsOk())
	assert.False(t, Either[S, bool]{}.Auto(n).IsOk())
	assert.False(t, Either[S, bool]{}.Auto(e).IsOk())
}

func Zero[A any]() (def A) { return }
