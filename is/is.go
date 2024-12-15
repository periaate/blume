package is

import "reflect"

func Empty[A any](input A) bool {
	a := interface{}(input)
	if a == nil { return true }
	val := reflect.ValueOf(a)
	kind := val.Kind()
	cond := kind == reflect.Array ||
	kind == reflect.Slice ||
	kind == reflect.Map ||
	kind == reflect.Chan ||
	kind == reflect.String
	
	if cond { return reflect.ValueOf(a).Len() == 0 }
	switch v := a.(type) {
		case interface{ Len() int }: return v.Len() == 0
		case interface{ String() string }: return len(v.String()) == 0
		default: return false
	}
}

func Nil[A any](input A) bool {
	value := reflect.ValueOf(input)
	ref := reflect.ValueOf(value)
	switch ref.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String: return ref.Len() == 0
		case reflect.Ptr: return ref.IsNil()
		case reflect.Interface: return ref.IsNil()
		case reflect.Func: return ref.IsNil()
	}

	return true
}


func Zero[A any](value A) bool {
	a := any(value)
	switch a.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr: return a == 0
		case float32, float64: return a == 0.0
		case string: return a == ""
		case bool: return a == false
		case complex64, complex128: return a == 0+0i
	}

	val := reflect.ValueOf(a)
	if val.Kind() == reflect.Ptr { val = val.Elem() }
	if val.Kind() != reflect.Struct { return false }
	for i := 0; i < val.NumField(); i++ {
		if !val.Field(i).IsZero() { return false }
	}
	return true
}

func Truthy[A any](input A) bool { return Empty(input) || Nil(input) || Empty(input) }

func NotEmpty[A any](input A) bool { return !Empty(input) }
func NotNil[A any](input A) bool { return !Nil(input) }
func NotZero[A any](input A) bool { return !Zero(input) }
func Falsy[A any](input A) bool { return !Truthy(input) }
