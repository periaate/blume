package T

import "iter"

type Arr[A any] interface {
	// Map applies the function to each element in the array.
	Map(fns ...Monadic[A, A]) Arr[A]
	// Filter returns a new array with elements that pass the predicate.
	Filter(fns ...Predicate[A]) Arr[A]
	// First returns the first element that passes the predicate.
	First(fns ...Predicate[A]) Result[A]
	// Reduce reduces the array to a single value.
	Reduce(fn Dyadic[A, A, A], init A) A
	// Values returns a sequence of elements.
	Values() iter.Seq[A]
	// Iter returns a sequence of elements with index.
	Iter() iter.Seq2[int, A]
	// Len returns the length of the array.
	Len() int
}
