package blume

import (
	"fmt"
	"math/rand"
)

// func Join(elems []String, sep String) String {
// 	b := Buf()
// 	b.Grow(n)
// 	b.WriteString(string(elems[0]))
// 	for _, s := range elems[1:] {
// 		b.WriteString(string(sep))
// 		b.WriteString(string(s))
// 	}
// 	return String(b.String())
// }

func Prepend[A any](arg A, arr []A) []A { return append([]A{arg}, arr...) }

type Array[A any] struct{ Value []A }

func (a Array[A]) Pattern(selector Selector[Array[A]], actor func(Array[A], [][]int) Array[A]) Array[A] {
	return Pattern(selector, actor)(a)
}

func (a Array[A]) Shuffle() Array[A] {
	args := a.Value
	rand.Shuffle(len(args), func(i, j int) {
		temp := args[j]
		args[j] = args[i]
		args[i] = temp
	})
	return ToArray(args)
}


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

func (arr Array[A]) Slice(start int, ends ...int) (res Array[A]) {
	l := len(arr.Value)
	if l == 0 { return }
	c := Clamp(0, len(arr.Value))
	if start < 0 { start = l+start }
	if len(ends) == 0 { return ToArray(arr.Value[c(start):]) }
	end := ToArray(ends).Gets(0)
	if end   < 0 { end   = l+end }
	return ToArray(arr.Value[c(start):c(end)])
}

func (arr Array[A]) Contains(a any) bool { return arr.First(Cat[A](ToString, Is(P.S(a)))).IsOk() }

func (arr Array[A]) Gets(i int) A { return arr.Get(i).Must() }
func (arr Array[A]) Reverse() Array[A] {
	r := make([]A, 0, len(arr.Value))
	for i := len(arr.Value); i > 0; i-- {
		r = append(r, arr.Value[i-1])
	}
	return ToArray(r)
}

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

func (arr Array[A]) First(fn Pred[A]) (res Option[A]) {
	for _, val := range arr.Value {
		if fn(val) {
			return res.Pass(val)
		}
	}
	return res.Fail()
}

func (arr Array[A]) Then(fn func(Array[A]) Array[A]) Array[A] { return fn(arr) }

func (arr Array[A]) Map(fn func(A) A) Array[A] {
	res := make([]A, len(arr.Value))
	for i, val := range arr.Value {
		res[i] = fn(val)
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) Reduce(fn func(A, A) A, initial A) A {
	for _, val := range arr.Value {
		initial = fn(initial, val)
	}
	return initial
}

func (arr Array[A]) Flat(fn func(A) []A) Array[A] { return ToArray(FlatMap(fn)(arr.Value)) }

func (arr Array[A]) Each(fn func(A)) Array[A] {
	for _, value := range arr.Value {
		fn(value)
	}
	return arr
}

func Each[A any](fn func(A)) func(Array[A]) {
	return func(arr Array[A]) {
		for _, value := range arr.Value {
			fn(value)
		}
	}
}

func ForSafe[A, B any](fn func(A)) Option[func(B) B] {
	var b B
	switch any(b).(type) {
	case Array[A]: return Cast[func(B) B](func(arr Array[A]) Array[A] { for _, value := range arr.Value { fn(value) }; return arr })
	case []A: return Cast[func(B) B](func(arr []A) []A { for _, value := range arr { fn(value) }; return arr })
	}
	return None[func(B) B]()
}

func For[B, A any](fn func(A)) func(B) B {
	var b B
	switch any(b).(type) {
	case Array[A]: return Cast[func(B) B](func(arr Array[A]) Array[A] { for _, value := range arr.Value { fn(value) }; return arr }).Must()
	case []A: return Cast[func(B) B](func(arr []A) []A { for _, value := range arr { fn(value) }; return arr }).Must()
	}
	panic("unsafe call to blume.Each; input type `B` did not match `Array[A]` or `[]A`; input type B must be array-like")
}

func Forn[B, A any](fn func(A)) func(B) { return Ignore(For[B, A](fn)) }

// Join fuck it everything is just strings now
func (arr Array[A]) Join(sep String) String { return Join(sep)(Map[A](Sprint)(arr.Value)) }

func Sprint[A any](a A) String { return S(fmt.Sprint(a)) }

func (arr Array[A]) Append(val A, rest ...A) Array[A] {
	return ToArray(append(arr.Value, Prepend(val, rest)...))
}

// JoinAfter joins input after this array
// [this, ...]
func (this Array[A]) JoinAfter(input Array[A]) Array[A] { return Array[A]{ Value: append(this.Value, input.Value...) }}

// JoinBefore joins input before this array
// [..., this]
func (this Array[A]) JoinBefore(input Array[A]) Array[A] { return Array[A]{ Value: append(input.Value, this.Value...) }}

func (arr Array[A]) Prepend(val A, rest ...A) Array[A] {
	return ToArray(append(Prepend(val, rest), arr.Value...))
}

func (arr Array[A]) Appends(val A, rest ...A) []A {
	return append(arr.Value, Prepend(val, rest)...)
}

func (arr Array[A]) Prepends(val A, rest ...A) []A {
	return append(Prepend(val, rest), arr.Value...)
}

func (arr Array[A]) Split(fn Pred[A]) (HasNot Array[A], Has Array[A]) {
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

func Flag(arr Array[String], flags ...String) (Array[String], bool) {
	pred := Is(flags...)
	new_arr := make([]String, 0, len(arr.Value))
	for i, val := range arr.Value {
		if pred(val) {
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

func Seen[K comparable]() func(K) bool {
	seen := make(map[K]any)
	return func(k K) bool {
		_, ok := seen[k]
		if ok {
			return true
		}
		seen[k] = nil
		return false
	}
}

// TODO: add UniqueBy
func Unique[K comparable](slice []K) []K { return Filter(Not(Seen[K]()))(slice) }

// TODO: add UniqueBy
func (arr Array[A]) Unique() Array[A] {
	var a A
	switch any(a).(type) {
	case string    : return Cast[Array[A]](Cast[Array[string]]     (arr).Must().Filter(Not(Seen[string]())))    .Must()
	case bool      : return Cast[Array[A]](Cast[Array[bool]]       (arr).Must().Filter(Not(Seen[bool]())))      .Must()
	case int       : return Cast[Array[A]](Cast[Array[int]]        (arr).Must().Filter(Not(Seen[int]())))       .Must()
	case uint      : return Cast[Array[A]](Cast[Array[uint]]       (arr).Must().Filter(Not(Seen[uint]())))      .Must()
	case int8      : return Cast[Array[A]](Cast[Array[int8]]       (arr).Must().Filter(Not(Seen[int8]())))      .Must()
	case uint8     : return Cast[Array[A]](Cast[Array[uint8]]      (arr).Must().Filter(Not(Seen[uint8]())))     .Must()
	case int16     : return Cast[Array[A]](Cast[Array[int16]]      (arr).Must().Filter(Not(Seen[int16]())))     .Must()
	case uint16    : return Cast[Array[A]](Cast[Array[uint16]]     (arr).Must().Filter(Not(Seen[uint16]())))    .Must()
	case int32     : return Cast[Array[A]](Cast[Array[int32]]      (arr).Must().Filter(Not(Seen[int32]())))     .Must()
	case uint32    : return Cast[Array[A]](Cast[Array[uint32]]     (arr).Must().Filter(Not(Seen[uint32]())))    .Must()
	case int64     : return Cast[Array[A]](Cast[Array[int64]]      (arr).Must().Filter(Not(Seen[int64]())))     .Must()
	case uint64    : return Cast[Array[A]](Cast[Array[uint64]]     (arr).Must().Filter(Not(Seen[uint64]())))    .Must()
	case float32   : return Cast[Array[A]](Cast[Array[float32]]    (arr).Must().Filter(Not(Seen[float32]())))   .Must()
	case float64   : return Cast[Array[A]](Cast[Array[float64]]    (arr).Must().Filter(Not(Seen[float64]())))   .Must()
	case complex64 : return Cast[Array[A]](Cast[Array[complex64]]  (arr).Must().Filter(Not(Seen[complex64]()))) .Must()
	case complex128: return Cast[Array[A]](Cast[Array[complex128]] (arr).Must().Filter(Not(Seen[complex128]()))).Must()
	default        : return arr.Filter(Cat[A](ToString, Not(Seen[S]()))) // ¯\_(ツ)_/¯ it works, can't be bothered with reflection
	}
}

func ToString[A any](a A) S { return P.S(a) }


func Pair[A any](arr Array[A]) (res Result[Array[Array[A]]]) {
	l := len(arr.Value)
	if l%2 != 0 { return res.Fail("pair called with an uneven array") }
	arrs := make([]Array[A], 0, l/2)
	for i := range l/2 {
		n := i*2
		arrs = append(arrs, Array[A]{ Value: []A{ arr.Value[n], arr.Value[n+1] } })
	}
	return res.Pass(Array[Array[A]]{Value: arrs})
}

func Pairs[A any](arr Array[A]) (res Array[Array[A]]) {
	l := len(arr.Value)
	if l%2 != 0 { Exit("pairs called with an uneven array") }
	arrs := make([]Array[A], 0, l/2)
	for i := range l/2 {
		n := i*2
		arrs = append(arrs, Array[A]{ Value: []A{ arr.Value[n], arr.Value[n+1] } })
	}
	return Array[Array[A]]{Value: arrs}
}

func ArrayFlat[A any](arr Array[Array[A]]) Array[A] {
	res := Arr[A]()
	for _, a := range arr.Value {
		res = res.JoinAfter(a)
	}
	return res
}



