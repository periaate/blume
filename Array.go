package blume

import (
	"math/rand"
	"strings"
)

func Arr[T any](args ...T) A[T] { return args }

// Prepend prepends the arguments before the array.
// [..., arr]
func Prepend[T any](arr []T, args ...T) []T { return append(args, arr...) }

// Append appends the arguments after the array.
// [arr, ...]
// Just use `append` in most cases.
func Append[T any](arr []T, args ...T) []T { return append(arr, args...) }

type Array[T any] []T


func (a Array[T]) Pattern(selector Selector[Array[T]], actor func(Array[T], [][]int) Array[T]) Array[T] {
	return Pattern(selector, actor)(a)
}

func (a Array[T]) Shuffle() Array[T] {
	args := a
	rand.Shuffle(len(args), func(i, j int) {
		temp := args[j]
		args[j] = args[i]
		args[i] = temp
	})
	return Array[T](args)
}

func (arr Array[T]) Len() int { return len(arr) }

func (arr Array[T]) Get(i int) (res Option[T]) {
	if i < 0         { i = len(arr) + i }
	if i < 0         { return res.Fail() }
	if i >= len(arr) { return res.Fail() }
	return res.Pass(arr[i])
}

func (arr Array[T]) Slice(start int, ends ...int) (res Array[T]) {
	l := len(arr)
	if l == 0 { return }
	c := Clamp(0, len(arr))
	if start < 0 { start = l+start }
	if len(ends) == 0 { return Array[T](arr[c(start):]) }
	end := Array[int](ends).Gets(0)
	if end   < 0 { end   = l+end }
	return arr[c(start):c(end)]
}

func (arr Array[T]) Contains(a any) bool { return arr.First(Cat[T](ToString, Is(P.S(a)))).IsOk() }

func (arr Array[T]) Gets(i int) T { return arr.Get(i).Must() }
func (arr Array[T]) Reverse() Array[T] {
	r := make([]T, 0, len(arr))
	for i := len(arr); i > 0; i-- {
		r = append(r, arr[i-1])
	}
	return Array[T](r)
}

func (arr Array[T]) Filter(fn Pred[T]) Array[T] { return Filter(fn)(arr) }
func (arr Array[T]) FilterMap(fn func(T) Option[T]) Array[T] { return FilterMap(fn)(arr) }
func (arr Array[T]) First(fn Pred[T]) (res Option[T]) { return First(fn)(arr) }
func (arr Array[T]) Then(fn func(Array[T]) Array[T]) Array[T] { return fn(arr) }
func (arr Array[T]) Map(fn func(T) T) Array[T] { return Map[T, T](fn)(arr) }
func (arr Array[T]) Each(fn any) Array[T] { OverCoax[T, T](fn)(arr); return arr }

func (arr Array[T]) Join(sep S) S {
	collect := make([]string, 0, len(arr))
	for _, val := range arr {
		collect = append(collect, string(P.S(val)))
	}
	return S(strings.Join(collect, string(sep)))
}

// JoinAfter joins input after this array
// [this, ...]
func (this Array[T]) JoinAfter(input Array[T]) Array[T] { return  append(this, input...) }

// JoinBefore joins input before this array
// [..., this]
func (this Array[T]) JoinBefore(input Array[T]) Array[T] { return  append(input, this...) }

// Append appends the arguments after the array.
// [this, ...] -> Array[T]
func (arr Array[T]) Append(args ...T) Array[T] { return append(arr, args...) }

// Prepend args before Array
// [..., this] -> Array[T]
func (arr Array[T]) Prepend(args ...T) Array[T] { return append(args, arr...) }


func (arr Array[T]) Split(fn Pred[T]) (HasNot Array[T], Has Array[T]) {
	arr_1 := []T{}
	arr_2 := []T{}
	for i, val := range arr {
		if !fn(val) {
			arr_1 = append(arr_1, val)
			continue
		}
		arr_2 = arr[i+1:]
		break
	}

	return Array[T](arr_1), arr_2
}

func (arr Array[T]) From(n int) Array[T] {
	if n <= 0 || len(arr) == 0 { return arr }
	if len(arr) > n            { arr = arr[n:] }
	return arr
}

func (arr Array[T]) Froms(n int) []T { return arr.From(n) }

func Flag(arr Array[String], flags ...String) (Array[String], bool) {
	pred := Is(flags...)
	new_arr := make([]String, 0, len(arr))
	for i, val := range arr {
		if pred(val) {
			return Array[S](append(new_arr, arr[i+1:]...)), true
		}
		new_arr = append(new_arr, val)
	}

	return Array[S](new_arr), false
}

func (arr Array[T]) Flag(fn Pred[T]) (Array[T], bool) {
	new_arr := make([]T, 0, len(arr))
	for i, val := range arr {
		if fn(val) {
			return Array[T](append(new_arr, arr[i+1:]...)), true
		}
		new_arr = append(new_arr, val)
	}

	return Array[T](new_arr), false
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

func SeenBy[T any, K comparable](fn func(T) K) func(T) bool {
	seen := make(map[K]any)
	return func(val T) bool {
		k := fn(val)
		_, ok := seen[k]
		if ok { return true }
		seen[k] = nil
		return false
	}
}

// TODO: add UniqueBy
func Unique[K comparable](slice Array[K]) Array[K] { return Filter(Not(Seen[K]()))(slice) }
func UniqueBy[T any, K comparable](fn func(T) K, slice Array[T]) Array[T] { return Filter(Not(SeenBy(fn)))(slice) }

// TODO: add UniqueBy
func (arr Array[T]) Unique() Array[T] {
	var a T
	switch any(a).(type) {
	case string    : return Cast[Array[T]](Cast[Array[string]]     (arr).Must().Filter(Not(Seen[string]())))    .Must()
	case bool      : return Cast[Array[T]](Cast[Array[bool]]       (arr).Must().Filter(Not(Seen[bool]())))      .Must()
	case int       : return Cast[Array[T]](Cast[Array[int]]        (arr).Must().Filter(Not(Seen[int]())))       .Must()
	case uint      : return Cast[Array[T]](Cast[Array[uint]]       (arr).Must().Filter(Not(Seen[uint]())))      .Must()
	case int8      : return Cast[Array[T]](Cast[Array[int8]]       (arr).Must().Filter(Not(Seen[int8]())))      .Must()
	case uint8     : return Cast[Array[T]](Cast[Array[uint8]]      (arr).Must().Filter(Not(Seen[uint8]())))     .Must()
	case int16     : return Cast[Array[T]](Cast[Array[int16]]      (arr).Must().Filter(Not(Seen[int16]())))     .Must()
	case uint16    : return Cast[Array[T]](Cast[Array[uint16]]     (arr).Must().Filter(Not(Seen[uint16]())))    .Must()
	case int32     : return Cast[Array[T]](Cast[Array[int32]]      (arr).Must().Filter(Not(Seen[int32]())))     .Must()
	case uint32    : return Cast[Array[T]](Cast[Array[uint32]]     (arr).Must().Filter(Not(Seen[uint32]())))    .Must()
	case int64     : return Cast[Array[T]](Cast[Array[int64]]      (arr).Must().Filter(Not(Seen[int64]())))     .Must()
	case uint64    : return Cast[Array[T]](Cast[Array[uint64]]     (arr).Must().Filter(Not(Seen[uint64]())))    .Must()
	case float32   : return Cast[Array[T]](Cast[Array[float32]]    (arr).Must().Filter(Not(Seen[float32]())))   .Must()
	case float64   : return Cast[Array[T]](Cast[Array[float64]]    (arr).Must().Filter(Not(Seen[float64]())))   .Must()
	case complex64 : return Cast[Array[T]](Cast[Array[complex64]]  (arr).Must().Filter(Not(Seen[complex64]()))) .Must()
	case complex128: return Cast[Array[T]](Cast[Array[complex128]] (arr).Must().Filter(Not(Seen[complex128]()))).Must()
	default        : return arr.Filter(Cat[T](ToString, Not(Seen[S]()))) // ¯\_(ツ)_/¯ it works, can't be bothered with reflection
	}
}

func (arr Array[T]) UniqueBy(fnAny any) Array[T] {
	switch fn := any(fnAny).(type) {
	case func(T) string    : return UniqueBy[T, string](fn, arr)
	case func(T) bool      : return UniqueBy[T, bool](fn, arr)
	case func(T) int       : return UniqueBy[T, int](fn, arr)
	case func(T) uint      : return UniqueBy[T, uint](fn, arr)
	case func(T) int8      : return UniqueBy[T, int8](fn, arr)
	case func(T) uint8     : return UniqueBy[T, uint8](fn, arr)
	case func(T) int16     : return UniqueBy[T, int16](fn, arr)
	case func(T) uint16    : return UniqueBy[T, uint16](fn, arr)
	case func(T) int32     : return UniqueBy[T, int32](fn, arr)
	case func(T) uint32    : return UniqueBy[T, uint32](fn, arr)
	case func(T) int64     : return UniqueBy[T, int64](fn, arr)
	case func(T) uint64    : return UniqueBy[T, uint64](fn, arr)
	case func(T) float32   : return UniqueBy[T, float32](fn, arr)
	case func(T) float64   : return UniqueBy[T, float64](fn, arr)
	case func(T) complex64 : return UniqueBy[T, complex64](fn, arr)
	case func(T) complex128: return UniqueBy[T, complex128](fn, arr)
	default                : return arr.Filter(Cat[T](ToString, Not(Seen[S]()))) // ¯\_(ツ)_/¯ it works, can't be bothered with reflection
	}
}

func ToString[T any](a T) S { return P.S(a) }


func Pair[T any](arr Array[T]) (res Result[Array[Array[T]]]) {
	l := len(arr)
	if l%2 != 0 { return res.Fail("pair called with an uneven array") }
	arrs := make([]Array[T], 0, l/2)
	for i := range l/2 {
		n := i*2
		arrs = append(arrs, []T{ arr[n], arr[n+1] })
	}
	return res.Pass(arrs)
}

func Pairs[T any](arr Array[T]) (res Array[Array[T]]) {
	l := len(arr)
	if l%2 != 0 { Exit("pairs called with an uneven array") }
	arrs := make([]Array[T], 0, l/2)
	for i := range l/2 {
		n := i*2
		arrs = append(arrs, []T{ arr[n], arr[n+1] })
	}
	return res
}

