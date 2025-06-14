package blume

import (
	"bytes"
	"io"
	"os"
	"reflect"
)

type Tree[A any] struct {
	Val A
	Arr []Tree[A]
}

type Delimiter struct {
	Start string
	End   string
}

func EmbedDelims(sar []string, delims ...Delimiter) Tree[string] {
	car := make([]Tree[S], len(sar))
	for i, s := range sar { car[i].Val = s }
	res, _ := embeds(car, delims)
	return res
}

func embeds(car []Tree[S], delims []Delimiter) (res Tree[string], v int) {
	for i := 0; len(car) > i; i++ {
		v := car[i]
		matched := false
		for _, delim := range delims {
			switch v.Val {
			case delim.Start:
				r, k := embeds(car[i+1:], delims)
				i += k
				res.Arr = append(res.Arr, r)
				matched = true
			case delim.End:
				return res, i + 1
			}
			if matched {
				break
			}
		}
		if !matched {
			res.Arr = append(res.Arr, v)
		}
	}

	return res, 0
}

func Thunk[A, B any](val A) func(_ B) A { return func(_ B) A { return val } }
func If[A any](ok bool, a A, b A) A { if ok { return a } else { return b } }

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

func Auto[O, I any](value I, handles ...any) O { return From[I, O](value, handles...).Value }

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

type TypeV struct {
	reflect.Type
}

func T(v any) TypeV { return TypeV{reflect.TypeOf(v)} }
func TA[T any]() TypeV { var t T; return TypeV{reflect.TypeOf(t)} }

func (t TypeV) IsArray() bool { return t.Kind() == reflect.Array }
func (t TypeV) IsFunc() bool  { return t.Kind() == reflect.Func }

// func typeOf(v any) reflect.Type { return reflect.TypeOf(v) }
type typeOf func(v any) reflect.Type
var TypeOf = reflect.TypeOf

func (t TypeV) Inputs(arg any) (res Result[A[TypeV]]) {
	if !t.IsFunc() { return res.Fail("Args must be called on a type with kind Func") }
	r := A[TypeV]{}
	for i := range t.NumIn() {
		r = append(r, TypeV{t.In(i)})
	}
	return res.Pass(r)
}

func (t TypeV) Outputs(arg any) (res Result[A[TypeV]]) {
	if !t.IsFunc() { return res.Fail("Args must be called on a type with kind Func") }
	r := A[TypeV]{}
	for i := range t.NumOut() {
		r = append(r, TypeV{t.Out(i)})
	}
	return res.Pass(r)
}

func IsArray(arg any) bool { return T(arg).IsArray() }
func IsFunc(arg any) bool  { return T(arg).IsFunc() }
func TypeOfEl(arg any) (res Option[reflect.Type]) {
	if !IsArray(arg) { return res.Fail() }
	return res.Pass(reflect.TypeOf(arg).Elem())
}

func Into[Target any](arg any) (res Option[Target]) {
	var t Target
	target := reflect.TypeOf(t)

	input := T(arg)
	if input.AssignableTo(target) || input.ConvertibleTo(target) { return Cast[Target](reflect.ValueOf(arg).Convert(target).Interface()) }
	if input.Kind() == target.Kind() && input.Kind() == reflect.Slice {
		te := target.Elem()
		if !(input.Elem().AssignableTo(te) || input.Elem().ConvertibleTo(te)) { return res.Fail() }
		r := reflect.MakeSlice(target, 0, 0)
		for _, value := range reflect.ValueOf(arg).Seq2() {
			r = reflect.Append(r, value.Convert(te))
		}

		return Cast[Target](r.Interface())
	}

	return res.Fail()
}

func From[I, O any](value I, args ...any) (output Option[O]) {
	if TypeIs.SameAs(value, output) { return Cast[O](value) }
	if v := Cast[O](value); v.IsSome() { return v }

	var t O
	target := reflect.TypeOf(t)

	input := T(value)
	if input.AssignableTo(target) || input.ConvertibleTo(target) { return Cast[O](reflect.ValueOf(value).Convert(target).Interface()) }
	if input.Kind() == target.Kind() && input.Kind() == reflect.Slice {
		te := target.Elem()
		if !(input.Elem().AssignableTo(te) || input.Elem().ConvertibleTo(te)) { return output.Fail() }
		r := reflect.MakeSlice(target, 0, 0)
		for _, value := range reflect.ValueOf(value).Seq2() {
			r = reflect.Append(r, value.Convert(te))
		}

		return Cast[O](r.Interface())
	}

	var try any = value

	// switch r := any(output).(type) {
	// case Result[I]     : try = r.Auto(value, args...)
	// case Option[I]     : try = r.Auto(value, args...)
	// case string: try = P.S(value)
	// case int    : switch v := any(value).(type) { case string: try = S(v).ToInt    ().Must(); case String: try = v.ToInt    ().Must() }
	// case int8   : switch v := any(value).(type) { case string: try = S(v).ToInt8   ().Must(); case String: try = v.ToInt8   ().Must() }
	// case int16  : switch v := any(value).(type) { case string: try = S(v).ToInt16  ().Must(); case String: try = v.ToInt16  ().Must() }
	// case int32  : switch v := any(value).(type) { case string: try = S(v).ToInt32  ().Must(); case String: try = v.ToInt32  ().Must() }
	// case int64  : switch v := any(value).(type) { case string: try = S(v).ToInt64  ().Must(); case String: try = v.ToInt64  ().Must() }
	// case uint   : switch v := any(value).(type) { case string: try = S(v).ToUint   ().Must(); case String: try = v.ToUint   ().Must() }
	// case uint8  : switch v := any(value).(type) { case string: try = S(v).ToUint8  ().Must(); case String: try = v.ToUint8  ().Must() }
	// case uint16 : switch v := any(value).(type) { case string: try = S(v).ToUint16 ().Must(); case String: try = v.ToUint16 ().Must() }
	// case uint32 : switch v := any(value).(type) { case string: try = S(v).ToUint32 ().Must(); case String: try = v.ToUint32 ().Must() }
	// case uint64 : switch v := any(value).(type) { case string: try = S(v).ToUint64 ().Must(); case String: try = v.ToUint64 ().Must() }
	// case float32: switch v := any(value).(type) { case string: try = S(v).ToFloat32().Must(); case String: try = v.ToFloat32().Must() }
	// case float64: switch v := any(value).(type) { case string: try = S(v).ToFloat64().Must(); case String: try = v.ToFloat64().Must() }
	// }

	return Cast[O](try)
}

func R[A any](val A, err ...any) (res Result[A]) { return res.Auto(val, err) }
func O[A any](val A, err ...any) (res Result[A]) { return res.Auto(val, err) }

func Cast[T any](a any) (res Option[T]) {
	value, ok := a.(T)
	if !ok { return }
	res.Other = ok
	res.Value = value
	return
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
	case []byte: return bytes.NewBuffer(v)
	case io.Reader:
		buf := bytes.NewBuffer([]byte{})
		io.Copy(buf, v)
		return buf
	default: return bytes.NewBuffer([]byte{})
	}
}

func LookupEnv(arg string) (res Option[string]) { return From[string, S](os.LookupEnv(arg)) }
