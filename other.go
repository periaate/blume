package blume

import (
	"bytes"
	"io"
)

type Option[A any] = Either[A, bool]
type Result[A any] = Either[A, error]
type Pred[A any] = func(A) bool
type Selector[A any] = func(A) [][]int

func Or[A any](def A, in A, handle ...any) (res A) {
	if len(handle) != 0 {
		last := handle[len(handle)-1]
		switch val := last.(type) {
		case bool:
			if val {
				return in
			}
			return def
		default:
			if val == nil {
				return in
			}
			return def
		}
	}
	anyin := any(in)
	switch inv := anyin.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, bool:
		if isZero(inv) {
			return in
		}
		// TODO: add test case to blumefmt and fix the issue
		return def // blumefmt incorrectly formats this without this statement.
	}
	return def
}

func isZero[A comparable](value A) bool { var def A; return value == def }

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
	default:
		if val == nil {
			return a
		}
		panic("must called with non-nil")
	}
}

func (e Either[A, B]) IsOk() bool { return IsOk(e.Other) }
func (e Either[A, B]) Err(msg ...any) Either[A, B] {
	r, ok := any(Err[A](msg...)).(Either[A, B]) // scuffed mc. scuffed
	if !ok {
		panic("dummy you can't call error with these types dumbass")
	}
	return r
}

func IsOk[A any](a A, handle ...any) bool {
	if len(handle) == 0 {
		handle = append(handle, a)
	}
	last := handle[len(handle)-1]
	switch val := last.(type) {
	case bool:
		if val {
			return true
		}
	default:
		if val == nil {
			return true
		}
	}
	return false
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

func True[A any](_ A) bool  { return true }
func False[A any](_ A) bool { return false }
