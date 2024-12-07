package gen

import "github.com/periaate/blume/gen/T"

// Product returns the cartesian product of two slices.
func Product[A any](a []A, B []A) (res [][2]A) {
	res = make([][2]A, 0, len(a)*len(B))
	for _, v := range a {
		for _, w := range B {
			res = append(res, [2]A{v, w})
		}
	}
	return
}

func Partition[A any](fn T.Predicate[A]) T.Mapper[A, []A] {
	return func(args []A) (res [][]A) {
		res = make([][]A, 2)
		for _, arg := range args {
			switch fn(arg) {
			case false:
				res[0] = append(res[0], arg)
			case true:
				res[1] = append(res[1], arg)
			}
		}
		return res
	}
}

// Split returns a predicate that splits the given slice based on the given function.
func Split[A any](fn func(A) bool) T.Mapper[A, []A] {
	return func(args []A) (res [][]A) {
		res = [][]A{}
		temp := []A{}
		for _, arg := range args {
			if fn(arg) {
				res = append(res, temp)
				temp = []A{}
			} else {
				temp = append(temp, arg)
			}
		}
		res = append(res, temp)
		return res
	}
}

// V turns variadic arguments into a slice.
func V[A any](a ...A) []A { return a }

// Join multiple slices into a new slice.
func Join[A any](a ...[]A) (res []A) {
	var l int
	for _, arr := range a {
		l += len(arr)
	}
	res = make([]A, 0, l)
	for _, arr := range a {
		res = append(res, arr...)
	}
	return res
}

// Lim filters the args to be less than or equal to the given Max length.
func Lim[A ~string | ~[]any](Max int) T.Mapper[A, A] {
	return func(args []A) (res []A) {
		for _, a := range args {
			if len(a) <= Max {
				res = append(res, a)
			}
		}
		return
	}
}

// Reverses the given slice in place.
func Reverses[A any](arr []A) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// Reverse copies and reverses the given slice.
func Reverse[A any](arr []A) (res []A) {
	res = make([]A, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return
}

// Deduplicate returns a predicate that filters out duplicates based on the given function.
func Deduplicate[A any, C comparable](fn func(A) C) T.Predicate[A] {
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

func Shifts[A any](a []A) (res []A, popped A, ok bool) {
	if len(a) == 0 {
		return
	}
	return a[1:], a[0], true
}

func Pops[A any](a []A) (res []A, popped A, ok bool) {
	if len(a) == 0 {
		return
	}
	return a[:len(a)-1], a[len(a)-1], true
}
