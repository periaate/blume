package blume

import (
	"fmt"
	"math/rand"
	"slices"
)

func Slice[T any](input []T, from, to int) (res []T, ok bool) {
	s, l := min(from, to), max(from, to)
	if s < 0 || l < s || len(input) < l { return }
	return input[s:l], true
}

func Get[T any](arr []T, i int) (res T, ok bool) {
	if i < 0         { i = len(arr) + i }
	if i < 0         { return }
	if i >= len(arr) { return }
	return arr[i], true
}

func Logln(args ...any) { fmt.Println(args...) }

// Prepend prepends the arguments before the array; [..., arr]; As opposed to `append`
func Prepend[T any](arr []T, args ...T) []T { return append(args, arr...) }

func Shuffle[T any](args []T) []T {
	rand.Shuffle(len(args), func(i, j int) {
		temp := args[j]
		args[j] = args[i]
		args[i] = temp
	})
	return args
}

func ShuffleBy[T any](seed int64, args []T) []T {
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(args), func(i, j int) {
		temp := args[j]
		args[j] = args[i]
		args[i] = temp
	})
	return args
}

// Reverse reverses a slice in-place and returns itself. Meant for pipelines.
func Reverse[T any](arr []T) []T {
	slices.Reverse(arr)
	return arr
}

// Generalize slice pattern matching into an overloaded pattern matching setup.

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

		return arr_1, arr_2
	}
}

func Flag(arr []string, flags ...string) ([]string, bool) {
	pred := Is(flags...)
	new_arr := make([]string, 0, len(arr))
	for i, val := range arr {
		if pred(val) {
			return append(new_arr, arr[i+1:]...), true
		}
		new_arr = append(new_arr, val)
	}

	return new_arr, false
}

func Seen[K comparable]() func(K) bool {
	seen := make(map[K]any)
	return func(k K) bool {
		if _, ok := seen[k]; ok { return true }
		seen[k] = nil
		return false
	}
}

func SeenBy[T any, K comparable](fn func(T) K) func(T) bool { return Cat(fn, Seen[K]()) }

func Unique[K comparable](slice []K) []K { return Filter(Not(Seen[K]()))(slice) }
func UniqueBy[T any, K comparable](fn func(T) K, slice []T) []T { return Filter(Not(SeenBy(fn)))(slice) }

func Pair[T any](arr []T) (res Result[[][]T]) {
	l := len(arr)
	if l%2 != 0 { return res.Fail("pair called with an uneven array") }
	arrs := make([][]T, 0, l/2)
	for i := range l/2 {
		n := i*2
		arrs = append(arrs, []T{ arr[n], arr[n+1] })
	}
	return res.Pass(arrs)
}

