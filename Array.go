package blume

import "fmt"

func Prepend[A any](arg A, arr []A) []A { return append([]A{arg}, arr...) }

type Array[A any] struct{ Value []A }

type Length int

func (l Length) Is(i int) bool { return int(l) == i }
func (l Length) Gt(i int) bool { return int(l) > i }
func (l Length) Lt(i int) bool { return int(l) < i }
func (l Length) Ge(i int) bool { return int(l) >= i }
func (l Length) Le(i int) bool { return int(l) <= i }
func (l Length) Eq(i int) bool { return int(l) == i }
func (l Length) Ne(i int) bool { return int(l) != i }

func (arr Array[A]) Len() Length { return Length(len(arr.Value)) }

func (arr Array[A]) Get(i int) (res Option[A]) {
	if i < 0 {
		i = len(arr.Value) + i
	}
	if i < 0 {
		return res.Fail()
	}
	if i >= len(arr.Value) {
		return res.Fail()
	}
	return res.Pass(arr.Value[i])
}

func (arr Array[A]) Gets(i int) A { return arr.Get(i).Must() }

func Arr[A any](args ...A) Array[A] { return Array[A]{Value: args} }
func ToArray[A any](a []A) Array[A] { return Array[A]{a} }

func (arr Array[A]) Filter(fn Pred[A]) Array[A] {
	res := []A{}
	for _, val := range arr.Value {
		if fn(val) {
			res = append(res, val)
		}
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) Filter_map(fn func(A) Option[A]) Array[A] {
	res := []A{}
	for _, val := range arr.Value {
		if ret := fn(val); ret.Other {
			res = append(res, ret.Value)
		}
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) First(fn Pred[A]) (res Option[A]) {
	for _, val := range arr.Value {
		if fn(val) {
			return res.Pass(val)
		}
	}
	return res.Fail()
}

func (arr Array[A]) Map(fn func(A) A) Array[A] {
	res := make([]A, len(arr.Value))
	for i, val := range arr.Value {
		res[i] = fn(val)
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) Each(fn func(A)) Array[A] {
	for _, value := range arr.Value {
		fn(value)
	}
	return arr
}

// Join fuck it everything is just strings now
func (arr Array[A]) Join(sep String) String { return Join(sep)(Map[A](Sprint)(arr.Value)) }

func Sprint[A any](a A) String { return S(fmt.Sprint(a)) }

func (arr Array[A]) Append(val A, rest ...A) Array[A] {
	return ToArray(append(arr.Value, Prepend(val, rest)...))
}

func (arr Array[A]) Prepend(val A, rest ...A) Array[A] {
	return ToArray(append(Prepend(val, rest), arr.Value...))
}

func (arr Array[A]) Appends(val A, rest ...A) []A {
	return append(arr.Value, Prepend(val, rest)...)
}

func (arr Array[A]) Prepends(val A, rest ...A) []A {
	return append(Prepend(val, rest), arr.Value...)
}

func (arr Array[A]) Split(fn Pred[A]) (Array[A], Array[A]) {
	arr_1 := []A{}
	arr_2 := []A{}
	for i, val := range arr.Value {
		if !fn(val) {
			arr_1 = append(arr_1, val)
			continue
		}
		arr_2 = arr.Value[i+1:]
		break
	}

	return ToArray(arr_1), ToArray(arr_2)
}

func (arr Array[A]) From(n int) Array[A] {
	if n <= 0 || len(arr.Value) == 0 {
		return arr
	}
	if len(arr.Value) > n {
		arr.Value = arr.Value[n:]
	}
	return arr
}

func (arr Array[A]) Froms(n int) []A { return arr.From(n).Value }

func Flag(arr Array[String], flag String, alt ...String) (Array[String], bool) {
	new_arr := make([]String, 0, len(arr.Value))
	for i, val := range arr.Value {
		if val == flag {
			return ToArray(append(new_arr, arr.Value[i+1:]...)), true
		}
		new_arr = append(new_arr, val)
	}

	return ToArray(new_arr), false
}

func (arr Array[A]) Flag(fn Pred[A]) (Array[A], bool) {
	new_arr := make([]A, 0, len(arr.Value))
	for i, val := range arr.Value {
		if fn(val) {
			return ToArray(append(new_arr, arr.Value[i+1:]...)), true
		}
		new_arr = append(new_arr, val)
	}

	return ToArray(new_arr), false
}
