package util

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/periaate/blume/gen"
)

func IsLen[L gen.Lennable](n int) gen.Predicate[L] { return func(l L) bool { return len(l) == n } }
func IsT[T any]() gen.Predicate[any] {
	return func(a any) bool {
		_, ok := a.(T)
		return ok
	}
}

// To encodes a [Coder] to JSON.
func To(c Coder) (rw *bytes.Buffer, err error) {
	rw = bytes.NewBuffer([]byte{})
	err = c.Encode(rw)
	return
}

// From attempts to decode any type from a JSON [io.Reader].
func From[A any](r io.Reader) (a A, err error) {
	err = json.NewDecoder(r).Decode(&a)
	return
}

type Coder interface {
	Encode(io.Writer) error
	Decode(io.Reader) error
}
