# gen

Gen defines the generic primitives which Blume is based on.

```go
// Niladic is a function that takes no arguments and returns a single value.
type Niladic[A any] func() A
// Monadic is a function that takes a single argument and returns a single value.
type Monadic[A, B any] func(A) B
// Dyadic is a function that takes two arguments and returns a single value.
type Dyadic[A, B, C any] func(A, B) C

// Predicate is a function that takes a single argument and returns a boolean.
type Predicate[A any] Monadic[A, bool]
// Comparator is a function that compares two arguments.
type Comparator[A any] Dyadic[A, A, bool]

// Mapper is a function that takes variadic arguments and returns a slice.
type Mapper[A, B any] Monadic[[]A, []B]
// Reducer is a function that takes variadic arguments and returns a single value.
type Reducer[A, B any] Monadic[[]A, B]

// Thunk takes a monadic function, its argument, and returns a niladic function.
// When the niladic function is called, it will call the monadic function with the argument.
func Thunk[A, B any](fn Monadic[A, B]) Monadic[A, Niladic[B]]

// All returns true if all arguments pass the predicate.
func All[A any](fn Predicate[A]) Reducer[A, bool]

// Any returns true if any argument passes the predicate.
func Any[A any](fn Predicate[A]) Reducer[A, bool]

// Filter returns a slice of arguments that pass the predicate.
func Filter[A any](fn Predicate[A]) Mapper[A, A]

// Map applies the function to each argument and returns the results.
func Map[A, B any](fn func(A) B) Mapper[A, B]

// Reduce applies the function to each argument and returns the result.
func Reduce[A any, B any](fn Dyadic[B, A, B], init B) Reducer[A, B]

// Not negates a predicate.
func Not[T any](fn Predicate[T]) Predicate[T]

// Negate negates a comparator.
func Negate[T any](fn Comparator[T]) Comparator[T]

// Is returns a predicate that checks if the argument is in the list.
func Is[C comparable](A ...C) Predicate[C]

// Isnt returns a predicate that checks if the argument is not in the list.
func Isnt[C comparable](A ...C) Predicate[C]

// Comp takes a comparator, variadic arguments, and returns a predicate.
// If any of the arguments pass the comparator, the predicate returns true.
func Comp[A any](fn Comparator[A]) func(...A) Predicate[A]

// Or combines variadic predicates with an OR operation.
func Or[A any](fns ...Predicate[A]) Predicate[A]

// And combines variadic predicates with an AND operation.
func And[A any](fns ...Predicate[A]) Predicate[A]
```
