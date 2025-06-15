package blume

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
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

func Gt[N Numeric](arg N) Pred[N] { return func(n N) bool { return n > arg } }
func Ge[N Numeric](arg N) Pred[N] { return func(n N) bool { return n >= arg } }
func Lt[N Numeric](arg N) Pred[N] { return func(n N) bool { return n < arg } }
func Le[N Numeric](arg N) Pred[N] { return func(n N) bool { return n <= arg } }
func Eq[K comparable](arg K) Pred[K] { return func(n K) bool { return n == arg } }
func Ne[K comparable](arg K) Pred[K] { return func(n K) bool { return n != arg } }

type Tree[A any] struct {
	Val A
	Arr []Tree[A]
}

type Delimiter struct {
	Start string
	End   string
}

func EmbedDelims(sar []string, delims ...Delimiter) Tree[string] {
	car := make([]Tree[string], len(sar))
	for i, s := range sar { car[i].Val = s }
	res, _ := embeds(car, delims)
	return res
}

func embeds(car []Tree[string], delims []Delimiter) (res Tree[string], v int) {
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

func Into[Target any](arg any) (res Option[Target]) {
	if v, ok := any(arg).(Target); ok { return res.Pass(v) }

	var try any = arg
	var ok bool = true
	var output Target

	switch any(output).(type) {
	case string: try = fmt.Sprint(arg)
	case int    : switch v := any(arg).(type) { case string: try, ok = ToInt    (v).Unwrap() }
	case int8   : switch v := any(arg).(type) { case string: try, ok = ToInt8   (v).Unwrap() }
	case int16  : switch v := any(arg).(type) { case string: try, ok = ToInt16  (v).Unwrap() }
	case int32  : switch v := any(arg).(type) { case string: try, ok = ToInt32  (v).Unwrap() }
	case int64  : switch v := any(arg).(type) { case string: try, ok = ToInt64  (v).Unwrap() }
	case uint   : switch v := any(arg).(type) { case string: try, ok = ToUint   (v).Unwrap() }
	case uint8  : switch v := any(arg).(type) { case string: try, ok = ToUint8  (v).Unwrap() }
	case uint16 : switch v := any(arg).(type) { case string: try, ok = ToUint16 (v).Unwrap() }
	case uint32 : switch v := any(arg).(type) { case string: try, ok = ToUint32 (v).Unwrap() }
	case uint64 : switch v := any(arg).(type) { case string: try, ok = ToUint64 (v).Unwrap() }
	case float32: switch v := any(arg).(type) { case string: try, ok = ToFloat32(v).Unwrap() }
	case float64: switch v := any(arg).(type) { case string: try, ok = ToFloat64(v).Unwrap() } }
	if ok { return Cast[Target](try) }

	target := reflect.TypeOf(output)
	input := reflect.TypeOf(arg)
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

	return Cast[Target](try)
}

func Cast[T any](a any) (res Option[T]) {
	value, ok := a.(T)
	if !ok { return }
	res.Other = ok
	res.Value = value
	return
}

// NewCast will work the same way as an option would do to us using reflection based function construction
func NewCast[T any](a any) (res T, ok bool) {
	res, ok = a.(T)
	return
}

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
