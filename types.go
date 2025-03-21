package blume

import (
	"bytes"
	"io"
	"os"
	"strconv"
)

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

type Option[A any] = Either[A, bool]
type Result[A any] = Either[A, error]
type Pred[A any] = func(A) bool
type Selector[A any] = func(A) [][]int

func Opt[A any](a A, other any) Option[A] {
	if IsOk(a, other) {
		return Some(a)
	}
	return None[A]()
}

func Res[A any](a A, other any) Result[A] {
	if IsOk(a, other) {
		return Ok(a)
	}
	return Err[A](other)
}

func Cast[T any](a any) Option[T] {
	value, ok := a.(T)
	return Opt(value, ok)
}

func StoS[A, B ~string](value A) B   { return B(value) }
func ItoI[A, B Numeric](value A) B   { return B(value) }
func StoD[A ~string](value A) string { return string(value) }

func SD(args []String) []string     { return Map(StoD[String])(args) }
func SS[A, B ~string](args []A) []B { return Map(StoS[A, B])(args) }
func DS(args []string) []String     { return Map(StoS[string, String])(args) }

func SAD(args ...String) Array[string]    { return ToArray(Map(StoD[String])(args)) }
func SAS[A, B ~string](args []A) Array[B] { return ToArray(Map(StoS[A, B])(args)) }
func DAS(args ...string) Array[String]    { return ToArray(Map(StoS[string, String])(args)) }

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

func Parse[N Integer | Float](s string, args ...any) Option[N] {
	var a N
	var (
		bitSize int
		base    = Cast[int](ToArray(args).Get(0).Or(10)).Or(10)
	)

	switch any(a).(type) {
	case int8, uint8:
		bitSize = 8
	case int16, uint16:
		bitSize = 16
	case int32, uint32, float32:
		bitSize = 32
	case int64, uint64, float64:
		bitSize = 64
	default:
		return None[N]()
	}
	var value any
	var err error

	switch any(a).(type) {
	case int, int8, int16, int32, int64:
		value, err = strconv.ParseInt(s, base, bitSize)
	case uint, uint8, uint16, uint32, uint64:
		value, err = strconv.ParseUint(s, base, bitSize)
	case float32, float64:
		value, err = strconv.ParseFloat(s, bitSize)
	}
	if err == nil {
		return Cast[N](value)
	}
	return None[N]()
}

func ToInt[S ~string](s S) Option[int] {
	i, err := strconv.Atoi(string(s))
	return Either[int, bool]{Value: int(i), Other: err == nil}
}
func ToInt8[S ~string](s S) Option[int8] {
	i, err := strconv.ParseInt(string(s), 10, 8)
	return Either[int8, bool]{Value: int8(i), Other: err == nil}
}
func ToInt16[S ~string](s S) Option[int16] {
	i, err := strconv.ParseInt(string(s), 10, 16)
	return Either[int16, bool]{Value: int16(i), Other: err == nil}
}
func ToInt32[S ~string](s S) Option[int32] {
	i, err := strconv.ParseInt(string(s), 10, 32)
	return Either[int32, bool]{Value: int32(i), Other: err == nil}
}
func ToInt64[S ~string](s S) Option[int64] {
	i, err := strconv.ParseInt(string(s), 10, 64)
	return Either[int64, bool]{Value: int64(i), Other: err == nil}
}
func ToUint[S ~string](s S) Option[uint] {
	i, err := strconv.ParseUint(string(s), 10, 0)
	return Either[uint, bool]{Value: uint(i), Other: err == nil}
}
func ToUint8[S ~string](s S) Option[uint8] {
	i, err := strconv.ParseUint(string(s), 10, 8)
	return Either[uint8, bool]{Value: uint8(i), Other: err == nil}
}
func ToUint16[S ~string](s S) Option[uint16] {
	i, err := strconv.ParseUint(string(s), 10, 16)
	return Either[uint16, bool]{Value: uint16(i), Other: err == nil}
}
func ToUint32[S ~string](s S) Option[uint32] {
	i, err := strconv.ParseUint(string(s), 10, 32)
	return Either[uint32, bool]{Value: uint32(i), Other: err == nil}
}
func ToUint64[S ~string](s S) Option[uint64] {
	i, err := strconv.ParseUint(string(s), 10, 64)
	return Either[uint64, bool]{Value: uint64(i), Other: err == nil}
}
func ToFloat32[S ~string](s S) Option[float32] {
	i, err := strconv.ParseFloat(string(s), 32)
	return Either[float32, bool]{Value: float32(i), Other: err == nil}
}
func ToFloat64[S ~string](s S) Option[float64] {
	i, err := strconv.ParseFloat(string(s), 64)
	return Either[float64, bool]{Value: float64(i), Other: err == nil}
}

func LookupEnv(arg string) Option[String] {
	r, ok := os.LookupEnv(arg)
	if !ok {
		return None[String]()
	}
	return Some(String(r))
}
