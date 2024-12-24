package blume

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func Zero[A any]() (a A) { return }
func Or[C comparable](a, b C) (res C) {
	if a == res {
		return b
	}
	return a
}
func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func Some[A any](a A) Option[A]         { return Option[A]{Value: a, Ok: true} }
func None[A any]() Option[A]            { return Option[A]{Value: Zero[A](), Ok: false} }
func Ok[A any](a A) (A, error)          { return a, nil }
func Err[A any](args ...any) (A, error) { return Zero[A](), StrErr(Format(args...)) }

type StrErr string

func (e StrErr) Error() string { return string(e) }

type Option[A any] struct {
	Value A
	Ok    bool
}

func (o Option[A]) Must(args ...any) A {
	if !o.Ok {
		panic(Or(Format(args...), "Empty Option called with Must"))
	}
	return o.Value
}

func (o Option[A]) Or(def A) A {
	if !o.Ok {
		return def
	}
	return o.Value
}

func Format(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	switch v := args[0].(type) {
	case error:
		return v.Error()
	case string:
		if !MatchRegex(`\{:\w\}`)(v) {
			break
		}
		return fmt.Sprintf(ReplaceRegex[string](`\{:\w\}`, "%$1")(v), args[1:]...)
	}
	return fmt.Sprint(args...)
}

type (
	Numeric interface{ Unsigned | Signed | Float }
	Integer interface{ Signed | Unsigned }
	Float   interface{ ~float32 | ~float64 }
	Signed  interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
)

func LT[N Numeric](n N) func(N) bool  { return func(val N) bool { return val < n } }
func GT[N Numeric](n N) func(N) bool  { return func(val N) bool { return val > n } }
func LTE[N Numeric](n N) func(N) bool { return func(val N) bool { return val <= n } }
func GTE[N Numeric](n N) func(N) bool { return func(val N) bool { return val >= n } }
func EQ[N Numeric](n N) func(N) bool  { return func(val N) bool { return val == n } }
func NEQ[N Numeric](n N) func(N) bool { return func(val N) bool { return val != n } }

func Len[A any](input A) int {
	value := reflect.ValueOf(input)
	if Is(reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String)(value.Kind()) {
		return value.Len()
	}
	return 0 // might want to panic :D
}

func Abs[N Numeric](n N) (zero N) {
	if n < zero {
		return -n
	}
	return n
}

// Clamp returns a function which ensures that the input value is within the specified range.
func Clamp[N Numeric](lower, upper N) func(N) N {
	if lower > upper {
		lower, upper = upper, lower
	}

	return func(val N) N {
		switch {
		case val >= upper:
			return upper
		case val <= lower:
			return lower
		default:
			return val
		}
	}
}

// SameSign returns true if a and b have the same sign.
func SameSign[N Numeric](a, b N) bool { return (a > 0 && b > 0) || (a < 0 && b < 0) }

func StoS[A, B ~string](a A) B { return B(a) } // StoS converts a string type to another string type.
func NtoN[A, B Numeric](a A) B { return B(a) } // NtoN converts a numeric type to another numeric type.

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
