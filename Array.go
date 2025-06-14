package blume

import (
	"math/rand"
)

// JoinAfter joins input after this array
// [this, ...]
func JoinAfter[T any](this []T) func(input []T) []T { return func(input []T) []T { return append(this, input...) } }

// JoinBefore joins input before this array
// [..., this]
func JoinBefore[T any](this []T) func(input []T) []T { return func(input []T) []T { return append(input, this...) } }

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

func Get[T any](i int) func(arr []T) (res Option[T]) {
	return func(arr []T) (res Option[T]) {
		if i < 0         { i = len(arr) + i }
		if i < 0         { return res.Fail() }
		if i >= len(arr) { return res.Fail() }
		return res.Pass(arr[i])
	}
}

func  Slice[T any](start int, ends ...int) func(arr []T)(res []T) {
	return func(arr []T) (res []T) {
		l := len(arr)
		if l == 0 { return }
		c := Clamp(0, len(arr))
		if start < 0 { start = l+start }
		if len(ends) == 0 { return Array[T](arr[c(start):]) }
		end := ends[0]
		if end < 0 { end = l+end }
		return arr[c(start):c(end)]
	}
}

func Reverse[T any](arr []T) Array[T] {
	r := make([]T, 0, len(arr))
	for i := len(arr); i > 0; i-- {
		r = append(r, arr[i-1])
	}
	return Array[T](r)
}

func ArrSplit[T any](fn Pred[T]) func(arr []T) (HasNot []T, Has []T) {
	return func(arr []T) (HasNot []T, Has []T) {
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
}

func ArrFrom[T any](n int) func(arr []T) []T {
	return func(arr []T) []T {
		if n <= 0 || len(arr) == 0 { return arr }
		if len(arr) > n            { arr = arr[n:] }
		return arr
	}
}

func Flag(arr []string, flags ...string) ([]string, bool) {
	pred := Is(flags...)
	new_arr := make([]string, 0, len(arr))
	for i, val := range arr {
		if pred(val) {
			return []S(append(new_arr, arr[i+1:]...)), true
		}
		new_arr = append(new_arr, val)
	}

	return []S(new_arr), false
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

func Unique[K comparable](slice []K) []K { return Filter(Not(Seen[K]()))(slice) }
func UniqueBy[T any, K comparable](fn func(T) K, slice []T) []T { return Filter(Not(SeenBy(fn)))(slice) }

func Pair[T any](arr []T) (res Result[Array[[]T]]) {
	l := len(arr)
	if l%2 != 0 { return res.Fail("pair called with an uneven array") }
	arrs := make([][]T, 0, l/2)
	for i := range l/2 {
		n := i*2
		arrs = append(arrs, []T{ arr[n], arr[n+1] })
	}
	return res.Pass(arrs)
}

