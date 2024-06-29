package gen

type (
	Option[T any]   func(T)
	Pipeable[T any] func(T) T
)

// First returns the first element in the slice that satisfies the given predicate.
// Returns -1 if no element satisfies the predicate.
func First[T any](arr []T, f func(T) bool) (res T, ind int) {
	for i, v := range arr {
		if f(v) {
			return v, i
		}
	}
	return res, -1
}

func ContainsAny[T, K any](comp func(T, ...K) bool, arr []T, tar ...K) bool {
	for _, v := range arr {
		for _, t := range tar {
			if comp(v, t) {
				return true
			}
		}
	}
	return false
}

func RemoveAny[T, K any](comp func(T, ...K) bool, arr []T, tar ...K) (res []T) {
	for _, v := range arr {
		for _, t := range tar {
			if comp(v, t) {
				goto breakl
			}
		}

		res = append(res, v)

	breakl:
	}
	return
}

func Curry[T, K, V any](fn func(T, K) V) func(K) func(T) V {
	return func(k K) func(T) V { return func(t T) V { return fn(t, k) } }
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Vals[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

func All[T any, K comparable](mustBe K, fns ...func(T) K) func(T) K {
	return func(a T) K {
		for _, fn := range fns {
			if r := fn(a); r != mustBe {
				return r
			}
		}
		return mustBe
	}
}

func Any[T any, K comparable](ifAny K, fns ...func(T) K) func(T) bool {
	return func(a T) bool {
		for _, fn := range fns {
			if fn(a) == ifAny {
				return true
			}
		}
		return false
	}
}

func Negate[T any](fn func(T) bool) func(T) bool {
	return func(t T) bool { return !fn(t) }
}

func Pipe[T any](fns ...func(T) T) func(T) T {
	return func(t T) T {
		for _, fn := range fns {
			if fn != nil {
				t = fn(t)
			}
		}
		return t
	}
}

func Filter[T any](arr []T, fn func(T) bool) (res []T) {
	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func Map[T any, K any](arr []T, fn func(T) K) []K {
	res := make([]K, 0, len(arr))
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}

// Combinations returns all possible pairs of elements from two slices.
func Combinations[T any](a []T, B []T) (res [][2]T) {
	res = make([][2]T, 0, len(a)*len(B))
	for _, v := range a {
		for _, w := range B {
			res = append(res, [2]T{v, w})
		}
	}

	return res
}

// Join joins multiple slices into one.
func Join[T any](a ...[]T) (res []T) {
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

func Deduplicate[T any, K comparable](arr []T, fn func(T) K) (res []T) {
	seen := make(map[K]bool)
	for _, v := range arr {
		if _, ok := seen[fn(v)]; !ok {
			res = append(res, v)
			seen[fn(v)] = true
		}
	}
	return res
}

func Reduce[T any](a []T, fn func(T, T) T, def ...T) (r T) {
	if len(a) == 0 {
		return
	}
	if len(a) == 1 {
		return a[0]
	}
	if len(def) > 0 {
		r = def[0]
	}
	for i, v := range a {
		if i == 0 {
			r = v
			continue
		}
		r = fn(r, v)
	}

	return
}
