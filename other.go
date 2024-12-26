package blume

import (
	"bytes"
	"io"

	"github.com/periaate/blume/pred/is"
)

func Or[A any](def A, in A, handle ...any) (res A) {
	if len(handle) != 0 {
		last := handle[len(handle)-1]
		switch val := last.(type) {
		case bool:
			if val {
				return in
			}
			return def
		case error:
			if val == nil {
				return in
			}
			return def
		}
	}
	anyin := any(in)
	switch inv := anyin.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, bool:
		if !is.Zero(inv) {
			return in
		}
		// TODO: add test case to blumefmt and fix the issue
		return def // blumefmt incorrectly formats this without this statement.
	}
	return def
}

func Must[A any](a A, handle ...any) A {
	if len(handle) == 0 {
		return a
	}
	last := handle[len(handle)-1]
	switch val := last.(type) {
	case bool:
		if val {
			return a
		}
		panic("must called with false bool")
	case error:
		if val == nil {
			return a
		}
		panic("must called with non nil error")
	default:
		panic("must called with unsupported handle")
	}
}

func Buf(args ...any) *bytes.Buffer {
	if len(args) == 0 {
		return bytes.NewBuffer([]byte{})
	}
	arg := args[0]
	switch v := arg.(type) {
	case string:
		return bytes.NewBufferString(v)
	case []byte:
		return bytes.NewBuffer(v)
	case io.Reader:
		buf := bytes.NewBuffer([]byte{})
		io.Copy(buf, v)
		return buf
	default:
		return bytes.NewBuffer([]byte{})
	}
}
