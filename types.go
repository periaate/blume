package blume

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/periaate/blume/is"
)

func Zero[A any]() (a A)              { return }
func Or[C comparable](a, b C) (res C) { if a == res   { return b };   return a }
func Must[A, B any](a A, b B) A    { return Auto(a, b).Must() }

func Auto[A, B any](a A, b B) Either[A, B] { return Either[A, B]{First: a, Second: b} }

func Some[A any](a A) Option[A] { return Option[A]{Value: a, Ok: true} }
func None[A any]() Option[A] { return Option[A]{Value: Zero[A](), Ok: false} }

func Ok[A any](a A) (A, error) { return a, nil }
//fuckyou
func Err[A any](args ...any) (A, error) { return Zero[A](), StrErr(Format(args...)) }

type StrErr string
func (e StrErr) Error() string { return string(e) }

type Option[A any] struct {
	Value A
	Ok    bool
}

func (o Option[A]) Must() A {
	if !o.Ok { panic("Option is empty") }
	return o.Value
}

func (o Option[A]) Or(def A) A {
	if !o.Ok { return def }
	return o.Value
}

type Either[A, B any] struct {
	First A
	Second B
}

func (e Either[A, B]) Ok() bool { return is.Truthy(e.Second) }

func (e Either[A, B]) Or(def A) A {
	if is.Truthy(e.Second) { return e.First }
	return def
}

func (e Either[A, B]) Must(args ...any) A {
	if is.Truthy(e.Second) { return e.First }
	panic(Format(args...))
}


func fixfmt(arg string) string { return ReplaceRegex[string](`\{:\w\}`, "%$1")(arg) }

//fuckyou
func Format(args ...any) string {
	switch len(args) {
		case 0: return ""
		case 1: return fmt.Sprint(args[0])
	default:
		s, ok := args[0].(string)
		if ok {
			if MatchRegex(`\{:\w\}`)(s) { return fmt.Sprintf(fixfmt(s), args[1:]...) }
		}
		return fmt.Sprint(args...)
	}
}

func isFormatString[A any](a A) bool {
	str := fmt.Sprint(a)
	return Contains("%")(str)
}

type (
	Numeric interface{ Unsigned | Signed | Float }
	Integer interface{ Signed | Unsigned }
	Float   interface{ ~float32 | ~float64 }
	Signed  interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 }
	Unsigned interface { ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr }
)

func LT[N Numeric](n N) func(N) bool   { return func(val N) bool { return val < n } }
func GT[N Numeric](n N) func(N) bool   { return func(val N) bool { return val > n } }
func LTE[N Numeric](n N) func(N) bool  { return func(val N) bool { return val <= n } }
func GTE[N Numeric](n N) func(N) bool  { return func(val N) bool { return val >= n } }
func EQ[N Numeric](n N) func(N) bool   { return func(val N) bool { return val == n } }
func NEQ[N Numeric](n N) func(N) bool  { return func(val N) bool { return val != n } }

func Len[A any](input A) int {
	value := reflect.ValueOf(input)
	if Is(reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String)(value.Kind()) {
		return value.Len()
	}
	return 0 // might want to panic :D
}

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
