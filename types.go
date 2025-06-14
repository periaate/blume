package blume

import (
	"bytes"
	"io"
	"os"
)

type Selector[A any] func(A) [][]int

func (s Selector[A]) Pred(input A) bool {
	return len(s(input)) > 0 
}

func IsNillable[A any](val A) bool {
	switch any(val).(type) {
	case error, uintptr, map[any]any, []any, chan any: return true
	default: return false
	}
}

func IsNil(val any) bool { return val == nil }

func IsOk[A any](a A, handle ...any) bool {
	if len(handle) == 0 { handle = append(handle, a) }
	switch val := handle[len(handle)-1].(type) {
	case bool: return val
	default  : return val == nil
	}
}

func Match[A, B, C any](r Either[A, B], value func(A) C, other func(B) C) C {
	switch IsOk(r) {
	case true: return value(r.Value)
	default:   return other(r.Other)
	}
}

func Auto[V any](value V, handles ...any) Result[V] {
	if IsOk(value, handles...) { return Ok(value) }
	if IsNil(value) { return Ok(value) }
	return Err[V](handles...)
}

func AutoRes[V any](value V, handles ...any) Result[V] {
	if IsOk(value, handles...) { return Ok(value) }
	if IsNil(value) { return Ok(value) }
	return Err[V](handles...)
}

func AutoOpt[V any](value V, handles ...any) Option[V] {
	if IsOk(value, handles...) { return Some(value) }
	if IsNil(value) { return Some(value) }
	return None[V]()
}

func From[V, E any](value V, handles ...any) (res E) {
	switch any(res).(type) {
	case Result[V]: return Cast[E](AutoRes(value, handles...)).OrDef()
	case Option[V]: return Cast[E](AutoOpt(value, handles...)).OrDef()
	default:        return
	}
}

func Opt[A any](a A, other any) Option[A] {
	if IsOk(a, other) { return Some(a) }
	return None[A]()
}

func Res[A any](a A, other any) Result[A] {
	if IsOk(a, other) { return Ok(a) }
	return Err[A](other)
}

func Cast[T any](a any) Option[T] {
	value, ok := a.(T)
	return Opt(value, ok)
}

func CastR[T any](a any) Result[T] {
	value, ok := a.(T)
	return Res(value, ok)
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
	if len(args) == 0 { return bytes.NewBuffer([]byte{}) }
	arg := args[0]
	switch v := arg.(type) {
	case string: return bytes.NewBufferString(v)
	case String: return bytes.NewBufferString(string(v))
	case []byte: return bytes.NewBuffer(v)
	case io.Reader:
		buf := bytes.NewBuffer([]byte{})
		io.Copy(buf, v)
		return buf
	default: return bytes.NewBuffer([]byte{})
	}
}

func LookupEnv(arg String) (res Option[S]) { return From[string, Option[S]](os.LookupEnv(arg.String())) }
