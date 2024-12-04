// Package gen implements generic types and functional forms which make use of them.
package gen

import "github.com/periaate/blume/gen/T"

func Or[C comparable](a, b C) (res C) {
	if a == res {
		return b
	}
	return a
}

func Tern[C comparable, A any](c C, a, b A) A {
	var zero C
	if c == zero {
		return b
	}
	return a
}

func PredAnd[A any](preds ...T.Predicate[A]) T.Predicate[A] {
	return func(a A) bool {
		for _, pred := range preds {
			if !pred(a) {
				return false
			}
		}
		return true
	}
}

func PredOr[A any](preds ...T.Predicate[A]) T.Predicate[A] {
	return func(a A) bool {
		for _, pred := range preds {
			if pred(a) {
				return true
			}
		}
		return false
	}
}

// All returns true if all arguments pass the [T.Predicate].
func All[A any](fns ...T.Predicate[A]) T.Reducer[A, bool] {
	fn := PredAnd[A](fns...)
	return func(args []A) bool {
		for _, arg := range args {
			if !fn(arg) {
				return false
			}
		}
		return true
	}
}

// Any returns true if any argument passes the [T.Predicate].
func Any[A any](fns ...T.Predicate[A]) T.Reducer[A, bool] {
	fn := PredOr[A](fns...)
	return func(args []A) bool {
		for _, arg := range args {
			if fn(arg) {
				return true
			}
		}
		return false
	}
}

// First returns the first value which passes the [T.Predicate].
func First[A any](fn T.Predicate[A]) T.Reducer[A, A] {
	return func(args []A) (res A) {
		for _, arg := range args {
			if fn(arg) {
				return arg
			}
		}
		return res
	}
}

// Filter returns a slice of arguments that pass the [T.Predicate].
func Filter[A any](fns ...T.Predicate[A]) T.Mapper[A, A] {
	fn := PredAnd[A](fns...)
	return func(args []A) (res []A) {
		res = make([]A, 0, len(args))
		for _, arg := range args {
			if fn(arg) {
				res = append(res, arg)
			}
		}
		return res
	}
}

// Map applies the function to each argument and returns the results.
func Map[A, B any](fn T.Monadic[A, B]) T.Mapper[A, B] {
	return func(args []A) (res []B) {
		res = make([]B, 0, len(args))
		for _, arg := range args {
			res = append(res, fn(arg))
		}
		return res
	}
}

// Reduce applies the function to each argument and returns the result.
func Reduce[A any, B any](fn T.Dyadic[B, A, B], init B) T.Reducer[A, B] {
	return func(args []A) B {
		res := init
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [T.Predicate].
func Not[A any](fn T.Predicate[A]) T.Predicate[A] { return func(t A) bool { return !fn(t) } }

// Is returns a [T.Predicate] that checks if the argument is in the list.
func Is[C comparable](A ...C) T.Predicate[C] {
	in := make(map[C]bool, len(A))
	for _, a := range A {
		in[a] = true
	}

	return T.Map[C, bool]{M: in}.Contains
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipe[A any](fns ...T.Transformer[A]) T.Transformer[A] {
	return func(a A) A {
		for _, fn := range fns {
			a = fn(a)
		}
		return a
	}
}

// Cat concatenates two [T.Monadic] functions into a single [T.Monadic] function.
func Cat[A, B, C any](a T.Monadic[A, B], b T.Monadic[B, C]) T.Monadic[A, C] {
	return func(c A) C { return b(a(c)) }
}
