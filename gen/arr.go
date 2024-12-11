package gen

import (
	. "github.com/periaate/blume/core"
)

// Lim filters the args to be less than or equal to the given Max length.
func Lim[A ~string | ~[]any](Max int) Monadic[[]A, []A] {
	return func(args []A) (res []A) {
		for _, a := range args {
			if len(a) <= Max { res = append(res, a) }
		}
		return
	}
}

func Shifts[A any](a []A) (res []A, popped A, ok bool) {
	if len(a) == 0 { return }
	return a[1:], a[0], true
}

func Pops[A any](a []A) (res []A, popped A, ok bool) {
	if len(a) == 0 { return }
	return a[:len(a)-1], a[len(a)-1], true
}

func Pair[A any](arr []A) Array[[]A] {
	pairs := [][]A{}
	var i int
	for i = 0; i < len(arr); i += 2 {
		cur := []A{}
		if i+1 <= len(arr) { cur = append(cur, arr[i]) }
		if i+2 <= len(arr) { cur = append(cur, arr[i+1]) }
		pairs = append(pairs, cur)
	}
	return ToArray(pairs)
}


