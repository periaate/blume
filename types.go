package blume

import (
	"bytes"
	"fmt"
)

type (
	Fn[R any]           func() R
	FnA[A, R any]       func(A) R
	FnB[A, B, R any]    func(A, B) R
	FnC[A, B, C, R any] func(A, B, C) R
)

func Zero[A any]() (a A)              { return }
func Or[C comparable](a, b C) (res C) { if a == res   { return b };   return a }
func Must[A any](a A, err error) A    { if err != nil { panic(err) }; return a }

type Option[A any] interface {
	Unwrap() A
	Or(A) A
	Ok() bool
}

type Result[A any] interface {
	Option[A]
	Err() error
}

func Ok[A any](value A)    Result[A] { return Res[A]{Opt: Opt[A]{value: value, isOk: true}} }
func Err[A any](err error) Result[A] { return Res[A]{Opt: Opt[A]{isOk: false}, err: err} }
func Some[A any](value A)  Option[A] { return Opt[A]{value: value, isOk: true} }
func None[A any]()         Option[A] { return Opt[A]{isOk: false} }
func AsOpt[A any](value A, other any) Option[A] {
	switch v := other.(type) {
	case error: if v != nil  { return None[A]() }
	case bool:  if v != true { return None[A]() }
	}
	return Some(value)
}

func Errf[A any](format string, args ...any) Result[A] {
	return Res[A]{Opt: Opt[A]{isOk: false}, err: fmt.Errorf(format, args...)}
}

func AsRes[A any](value A, err error) Result[A] {
	if err != nil { return Err[A](err) }
	return Ok(value)
}

type Opt[A any] struct {
	value A
	isOk  bool
	meta  string
}

func (o Opt[A]) Or(value A) A {
	if o.isOk { return o.value }
	return value
}

func (o Opt[A]) Ok() bool { return o.isOk }
func (o Opt[A]) Unwrap() A {
	if o.isOk { return o.value }
	panic("method Unwrap called on an empty Option")
}

type Res[A any] struct {
	Opt[A]
	err error
}

func (r Res[A]) Err() error { return r.err }
func (r Res[A]) Unwrap() A {
	if r.isOk { return r.value }
	panic("method Unwrap called on an empty Result")
}

type (
	Numeric interface{ Unsigned | Signed | Float }
	Integer interface{ Signed | Unsigned }
	Float   interface{ ~float32 | ~float64 }
	Signed  interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 }
	Unsigned interface { ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr }
)

func LT[N Numeric](n N) FnA[N, bool]   { return func(val N) bool { return val < n } }
func GT[N Numeric](n N) FnA[N, bool]   { return func(val N) bool { return val > n } }
func LTE[N Numeric](n N) FnA[N, bool]  { return func(val N) bool { return val <= n } }
func GTE[N Numeric](n N) FnA[N, bool]  { return func(val N) bool { return val >= n } }
func EQ[N Numeric](n N) FnA[N, bool]   { return func(val N) bool { return val == n } }
func NEQ[N Numeric](n N) FnA[N, bool]  { return func(val N) bool { return val != n } }
func Len(l interface{ Len() int }) int { return l.Len() }

func Abs[N Numeric](n N) (zero N) {
	if n < zero { return -n }
	return n
}

// Clamp returns a function which ensures that the input value is within the specified range.
func Clamp[N Numeric](lower, upper N) func(N) N {
	if lower > upper { lower, upper = upper, lower }

	return func(val N) N {
		switch {
			case val >= upper: return upper
			case val <= lower: return lower
			default: return val
		}
	}
}

// SameSign returns true if a and b have the same sign.
func SameSign[N Numeric](a, b N) bool { return (a > 0 && b > 0) || (a < 0 && b < 0) }

func StoS[A, B ~string](a A) B { return B(a) } // StoS converts a string type to another string type.
func NtoN[A, B Numeric](a A) B { return B(a) } // NtoN converts a numeric type to another numeric type.

func Buf(args ...any) *bytes.Buffer {
	if len(args) == 0 { return bytes.NewBuffer([]byte{}) }
	arg := args[0]
	switch v := arg.(type) {
	case string: return bytes.NewBufferString(v)
	case []byte: return bytes.NewBuffer(v)
	default: return bytes.NewBuffer([]byte{})
	}
}
