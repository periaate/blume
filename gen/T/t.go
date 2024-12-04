package T

import "iter"

type (
	// Niladic is a function that takes no arguments and returns a single value.
	Niladic[A any] func() A
	// Monadic is a function that takes a single argument and returns a single value.
	Monadic[A, B any] func(A) B
	// Dyadic is a function that takes two arguments and returns a single value.
	Dyadic[A, B, C any] func(A, B) C
	// Predicate is a function that takes a single argument and returns a boolean.
	Predicate[A any] Monadic[A, bool]
	// Transformer is a function that takes a single argument and returns a modified value.
	Transformer[A any] Monadic[A, A]
	// Mapper is a function that takes variadic arguments and returns a slice.
	Mapper[A, B any] Monadic[[]A, []B]
	// Reducer is a function that takes variadic arguments and returns a single value.
	Reducer[A, B any] Monadic[[]A, B]
)

type Collect[A any] interface {
	Collect() []A // Collect returns a slice of the values.
}

type Result[A any] struct {
	value A
	err   Error[any]
}

// Unwrap returns the value of the result.
func (r Result[A]) Unwrap() A {
	return r.value
}

// Or returns the value of the result, or a default value if it is an error.
func (r Result[A]) Or(def A) A {
	if r.IsErr() {
		return def
	}
	return r.value
}

// IsOk returns true if the result is a success.
func (r Result[A]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the result is an error.
func (r Result[A]) IsErr() bool {
	return r.err != nil
}

// Match takes two functions and calls the first if the result is a success, and the second if it is an error.
func (r Result[A]) Match(ok func(A), err func(Error[any])) {
	if r.IsOk() {
		ok(r.value)
	} else {
		err(r.err)
	}
}

func Results[A any](val A, err any) Result[A] {
	e := Error[any](nil)
	switch v := err.(type) {
	case Error[any]:
		e = v
	case error:
		e = Err[any]{err: v}
	}
	return Result[A]{value: val, err: e}
}

// Result is a type that represents the result of an operation.
type Resultable[A any] interface {
	// Unwrap returns the value of the result.
	Unwrap() A
	// Or (UnwrapOr) returns the value of the result, or a default value if it is an error.
	Or(A) A

	// IsOk returns true if the result is a success.
	IsOk() bool
	// IsErr returns true if the result is an error.
	IsErr() bool

	// Match takes two functions and calls the first if the result is a success, and the second if it is an error.
	Match(
		ok func(A),
		err func(Error[any]),
	)
}

type (
	Or[A any]       interface{ Or(Default A) A }
	Contains[A any] interface{ Contains(args ...A) bool }
)

type Filters[A any] interface {
	Filter(args ...Predicate[A]) []A
}

type Maps[A, B any] interface {
	Map(args ...Monadic[A, B]) []B
}
type Reduce[A any, B any] interface {
	Reduce(args ...Dyadic[B, A, B]) B
}

type Predicates[A any] interface {
	Any(args ...Predicate[A]) bool
	All(args ...Predicate[A]) bool
	None(args ...Predicate[A]) bool
}

type Getter[A any] interface {
	First(args ...Predicate[A]) A
	Last(args ...Predicate[A]) A
	Nth(n int, args ...Predicate[A]) A
}

type Asserts[A any] interface {
	Assert(args ...Predicate[A]) A
}

// Array is a type constraint that combines the Filter, Predicates, Getter, and Assert interfaces.
// Array does not implement Map or Reduce due to Go's type system limitations.
type Array[A any] interface {
	Filters[A]
	Predicates[A]
	Getter[A]
	Asserts[A]
	Values() iter.Seq[A]
	Iter() iter.Seq2[int, A]
}

type Stringable interface{ String() string }
