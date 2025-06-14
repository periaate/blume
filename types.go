package blume

import (
	"bytes"
	"io"
	"os"
	"reflect"
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

type EitherT[T any] interface { Result[T] | Option[T] }

func Auto[I, O any](value I, handles ...any) O { return From[I, O](value, handles...).Value }

type typeIs int
const TypeIs typeIs = 10

func (typeIs) Numeric (v any) (res bool) { switch v.(type) { case float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr: res = true }; return }
func (typeIs) Integer (v any) (res bool) { switch v.(type) { case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr: res = true }; return }
func (typeIs) Unsigned(v any) (res bool) { switch v.(type) { case uint, uint8, uint16, uint32, uint64, uintptr: res = true }; return }
func (typeIs) Signed  (v any) (res bool) { switch v.(type) { case int, int8, int16, int32, int64: res = true }; return }
func (typeIs) Float   (v any) (res bool) { switch v.(type) { case float32, float64: res = true }; return }
func (typeIs) Int8    (v any) (res bool) { switch v.(type) { case int8   : res = true }; return }
func (typeIs) Int16   (v any) (res bool) { switch v.(type) { case int16  : res = true }; return }
func (typeIs) Int32   (v any) (res bool) { switch v.(type) { case int32  : res = true }; return }
func (typeIs) Int64   (v any) (res bool) { switch v.(type) { case int64  : res = true }; return }
func (typeIs) Uint8   (v any) (res bool) { switch v.(type) { case uint8  : res = true }; return }
func (typeIs) Uint16  (v any) (res bool) { switch v.(type) { case uint16 : res = true }; return }
func (typeIs) Uint32  (v any) (res bool) { switch v.(type) { case uint32 : res = true }; return }
func (typeIs) Uint64  (v any) (res bool) { switch v.(type) { case uint64 : res = true }; return }
func (typeIs) Float32 (v any) (res bool) { switch v.(type) { case float32: res = true }; return }
func (typeIs) Float64 (v any) (res bool) { switch v.(type) { case float64: res = true }; return }
func (typeIs) Uintptr (v any) (res bool) { switch v.(type) { case uintptr: res = true }; return }

func (typeIs) Bool  (v any) (ok bool) { _, ok = v.(bool)  ; return }
func (typeIs) String(v any) (ok bool) { _, ok = v.(string); return }

func (typeIs) SameAs(a, b any) (ok bool) { return reflect.TypeOf(a).String() == reflect.TypeOf(a).String() }

func From[I, O any](value I, args ...any) (output Option[O]) {
	if TypeIs.SameAs(value, output) { return Cast[O](value) }
	if v := Cast[O](value); v.IsSome() { return v }

	var try any = value

	switch r := any(output).(type) {
	case Result[I]     : try = r.Auto(value, args...)
	case Option[I]     : try = r.Auto(value, args...)
	case string, String: try = P.S(value)
	case int    : switch v := any(value).(type) { case string: try = S(v).ToInt    ().Must(); case String: try = v.ToInt    ().Must() }
	case int8   : switch v := any(value).(type) { case string: try = S(v).ToInt8   ().Must(); case String: try = v.ToInt8   ().Must() }
	case int16  : switch v := any(value).(type) { case string: try = S(v).ToInt16  ().Must(); case String: try = v.ToInt16  ().Must() }
	case int32  : switch v := any(value).(type) { case string: try = S(v).ToInt32  ().Must(); case String: try = v.ToInt32  ().Must() }
	case int64  : switch v := any(value).(type) { case string: try = S(v).ToInt64  ().Must(); case String: try = v.ToInt64  ().Must() }
	case uint   : switch v := any(value).(type) { case string: try = S(v).ToUint   ().Must(); case String: try = v.ToUint   ().Must() }
	case uint8  : switch v := any(value).(type) { case string: try = S(v).ToUint8  ().Must(); case String: try = v.ToUint8  ().Must() }
	case uint16 : switch v := any(value).(type) { case string: try = S(v).ToUint16 ().Must(); case String: try = v.ToUint16 ().Must() }
	case uint32 : switch v := any(value).(type) { case string: try = S(v).ToUint32 ().Must(); case String: try = v.ToUint32 ().Must() }
	case uint64 : switch v := any(value).(type) { case string: try = S(v).ToUint64 ().Must(); case String: try = v.ToUint64 ().Must() }
	case float32: switch v := any(value).(type) { case string: try = S(v).ToFloat32().Must(); case String: try = v.ToFloat32().Must() }
	case float64: switch v := any(value).(type) { case string: try = S(v).ToFloat64().Must(); case String: try = v.ToFloat64().Must() }
	}

	return Cast[O](try)
}

func Cast[T any](a any) Option[T] {
	value, ok := a.(T)
	return Auto[T, Option[T]](value, ok)
}


func ItoI[A, B Numeric](value A) B   { return B(value) }
func StoS[A, B ~string](value A) B   { return B(value) }
func StoD[A ~string](value A) string { return string(value) }

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

func LookupEnv(arg String) (res Option[S]) { return From[string, S](os.LookupEnv(arg.String())) }
