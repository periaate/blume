package blume

import (
	"bytes"
	"io"
	"os"
)

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

func ItoI[A, B Numeric](value A) B   { return B(value) }
func StoS[A, B ~string](value A) B   { return B(value) }
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

func LookupEnv(arg string) Option[String] {
	r, ok := os.LookupEnv(arg)
	if !ok {
		return None[String]()
	}
	return Some(String(r))
}
