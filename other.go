package blume

import (
	"bytes"
	"io"
)

func isZero[A comparable](value A) bool { var def A; return value == def }

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

func True[A any](_ A) bool  { return true }
func False[A any](_ A) bool { return false }
