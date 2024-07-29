package typ

import "github.com/periaate/blume/core/gen"

// Product returns the cartesian product of two slices.
func Product[T any](a []T, B []T) (res [][2]T) {
	res = make([][2]T, 0, len(a)*len(B))
	for _, v := range a {
		for _, w := range B {
			res = append(res, [2]T{v, w})
		}
	}

	return
}

// Join multiple slices into a new slice.
func Join[T any](a ...[]T) (res []T) {
	var l int
	for _, arr := range a {
		l += len(arr)
	}
	res = make([]T, 0, l)
	for _, arr := range a {
		res = append(res, arr...)
	}
	return res
}

// Reverse reverses the given slice in place.
func Reverse[T any](arr []T) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// ReverseC copies and reverses the given slice.
func ReverseC[T any](arr []T) (res []T) {
	res = make([]T, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return
}

// Deduplicate returns a predicate that filters out duplicates based on the given function.
func Deduplicate[A any, C comparable](fn func(A) C) gen.Predicate[A] {
	seen := map[C]struct{}{}
	return func(a A) bool {
		c := fn(a)
		if _, ok := seen[c]; ok {
			return false
		}
		seen[c] = struct{}{}
		return true
	}
}
