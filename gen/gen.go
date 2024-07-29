/*
Package gen implements generic types and functional forms which make use of them.

## TODO

  - [ ] Support for error handling.
  - [ ] Functional forms for nested types.
*/
package gen

type (
	// Niladic is a function that takes no arguments and returns a single value.
	Niladic[A any] func() A
	// Monadic is a function that takes a single argument and returns a single value.
	Monadic[A, B any] func(A) B
	// Dyadic is a function that takes two arguments and returns a single value.
	Dyadic[A, B, C any] func(A, B) C

	// Predicate is a function that takes a single argument and returns a boolean.
	Predicate[A any] Monadic[A, bool]
	// Comparator is a function that compares two arguments.
	Comparator[A any] Dyadic[A, A, bool]
	// Transformer is a function that takes a single argument and returns a modified value.
	Transformer[A any] Monadic[A, A]

	// Mapper is a function that takes variadic arguments and returns a slice.
	Mapper[A, B any] Monadic[[]A, []B]
	// Reducer is a function that takes variadic arguments and returns a single value.
	Reducer[A, B any] Monadic[[]A, B]
)

// Thunk takes a [Monadic] function, its argument, and returns a [Niladic] function.
// When the [Niladic] function is called, it will call the [Monadic] function with the argument.
func Thunk[A, B any](fn Monadic[A, B]) Monadic[A, Niladic[B]] {
	return func(a A) Niladic[B] { return func() B { return fn(a) } }
}

// All returns true if all arguments pass the [Predicate].
func All[A any](fn Predicate[A]) Reducer[A, bool] {
	return func(args []A) bool {
		for _, arg := range args {
			if !fn(arg) {
				return false
			}
		}
		return true
	}
}

// Any returns true if any argument passes the [Predicate].
func Any[A any](fn Predicate[A]) Reducer[A, bool] {
	return func(args []A) bool {
		for _, arg := range args {
			if fn(arg) {
				return true
			}
		}
		return false
	}
}

// Filter returns a slice of arguments that pass the [Predicate].
func Filter[A any](fn Predicate[A]) Mapper[A, A] {
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
func Map[A, B any](fn func(A) B) Mapper[A, B] {
	return func(args []A) (res []B) {
		res = make([]B, 0, len(args))
		for _, arg := range args {
			res = append(res, fn(arg))
		}
		return res
	}
}

// Reduce applies the function to each argument and returns the result.
func Reduce[A any, B any](fn Dyadic[B, A, B], init B) Reducer[A, B] {
	return func(args []A) B {
		res := init
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipe[A any](fns ...Transformer[A]) Transformer[A] {
	return func(a A) A {
		for _, fn := range fns {
			a = fn(a)
		}
		return a
	}
}

// Cat concatenates two [Monadic] functions into a single [Monadic] function.
func Cat[A, B, C any](a Monadic[A, B], b Monadic[B, C]) Monadic[A, C] {
	return func(c A) C { return b(a(c)) }
}

// Not negates a [Predicate].
func Not[T any](fn Predicate[T]) Predicate[T] { return func(t T) bool { return !fn(t) } }

// Negate negates a [Comparator].
func Negate[T any](fn Comparator[T]) Comparator[T] { return func(a, b T) bool { return !fn(a, b) } }

// Is returns a [Predicate] that checks if the argument is in the list.
func Is[C comparable](A ...C) Predicate[C] {
	return func(B C) bool {
		for _, a := range A {
			if a == B {
				return true
			}
		}
		return false
	}
}

// Isnt returns a [Predicate] that checks if the argument is not in the list.
func Isnt[C comparable](A ...C) Predicate[C] {
	return func(B C) bool {
		for _, a := range A {
			if a == B {
				return false
			}
		}
		return true
	}
}

// Comp takes a [Comparator], variadic arguments, and returns a [Predicate].
// If any of the arguments pass the [Comparator], the [Predicate] returns true.
func Comp[A any](fn Comparator[A]) func(...A) Predicate[A] {
	return func(args ...A) Predicate[A] {
		return func(a A) bool {
			for _, arg := range args {
				if fn(arg, a) {
					return true
				}
			}
			return false
		}
	}
}

// Or combines variadic [Predicate]s with an OR operation.
func Or[A any](fns ...Predicate[A]) Predicate[A] {
	return func(a A) bool {
		for _, fn := range fns {
			if fn(a) {
				return true
			}
		}
		return false
	}
}

// And combines variadic [Predicate]s with an AND operation.
func And[A any](fns ...Predicate[A]) Predicate[A] {
	return func(a A) bool {
		for _, fn := range fns {
			if !fn(a) {
				return false
			}
		}
		return true
	}
}
