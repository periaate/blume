package core

type (
	// Niladic is a function that takes no arguments and returns a single value.
	Niladic[A any] func() A
	// Monadic is a function that takes a single argument and returns a single value.
	Monadic[A, B any] func(A) B
	// Dyadic is a function that takes two arguments and returns a single value.
	Dyadic[A, B, C any] func(A, B) C
	// Predicate is a function that takes a single argument and returns a boolean.
	Predicate[A any] Monadic[A, bool]
	// Mapper is a function that takes variadic arguments and returns a slice.
	Mapper[A, B any] Monadic[[]A, []B]
	// Reducer is a function that takes variadic arguments and returns a single value.
	Reducer[A, B any] Monadic[[]A, B]
)

type (
	Nothing        any
	Option[V any]  interface{ Result[V, Nothing] }
	Numeric interface       { Unsigned | Signed | Float }
	Integer interface       { Signed | Unsigned }
	Float interface         { ~float32 | ~float64 }
	Signed interface        { ~int | ~int8 | ~int16 | ~int32 | ~int64 }
	Unsigned interface      { ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr }
	Lennable interface      { ~[]any | ~string | ~map[any]any | ~chan any | ~[]string }
)

type Error[A any] interface {
	Value() A
	Reason() string
	Error() string
	Values() (A, string, error)
}

// Result is a type that represents the result of an operation.
type Result[V, E any] interface {
	Unwrap() V // returns the value or panics.
	Or(V) V    // returns the value or the passed value.
	IsOk() bool
	IsErr() bool
	Ok(func(V)) Error[E]  // Ok calls passed function if Ok, returns Error, even if None.
	Err(func(Error[E])) V // Err calls passed function if Err, returns Value, even if None.

	// Match takes two callbacks, calling the first if Ok, second if Err.
	Match(
		ok func(V),
		err func(Error[E]),
	)
	Values() (V, Error[E])
}

type Array[A any] interface {
	Map(fns ...Monadic[A, A]) Array[A]
	Filter(fns ...Predicate[A]) Array[A]
	First(fns ...Predicate[A]) Option[A]
	Reduce(fn Dyadic[A, A, A], init A) A
	Len() int
	Values() []A
	// Slice(from, to int) Option[Array[A]] // Slice returns a new array from the given range.
	// Cut(from, to int) bool // Cut removes the elements from the array in place.
	// Shift() Option[A] // Shift removes the first element from the array.
	// Pop() Option[A] // Pop removes the last element from the array.
}

type Str[S ~string] interface {
	Contains(args ...string) bool
	HasPrefix(args ...string) bool
	HasSuffix(args ...string) bool
	ReplacePrefix(pats ...string) S
	ReplaceSuffix(pats ...string) S
	Replace(pats ...string) S
	ReplaceRegex(pat string, rep string) S
	Shift(count int) S
	Pop(count int) S
	String() string
	ToUpper() S
	ToLower() S
	Trim() S
	TrimPrefix(prefix string) S
	TrimSuffix(suffix string) S
	TrimSpace() S
}

// type Str[S ~string] interface {
// 	Contains(args ...string) bool
// 	HasPrefix(args ...string) bool
// 	HasSuffix(args ...string) bool
// 	ReplacePrefix(pats ...string) S
// 	ReplaceSuffix(pats ...string) S
// 	Replace(pats ...string) S
// 	ReplaceRegex(pat string, rep string) S
// 	Shift(count int) S
// 	Pop(count int) S
// 	String() string
// 	ToUpper() S
// 	ToLower() S
// 	Trim() S
// 	TrimPrefix(prefix string) S
// 	TrimSuffix(suffix string) S
// 	TrimSpace() S
//
// // conv
// ToInt() Result[int]
// ToInt8() Result[int8]
// ToInt16() Result[int16]
// ToInt32() Result[int32]
// ToInt64() Result[int64]
// ToUint() Result[uint]
// ToUint8() Result[uint8]
// ToUint16() Result[uint16]
// ToUint32() Result[uint32]
// ToUint64() Result[uint64]
// ToFloat32() Result[float32]
// ToFloat64() Result[float64]

// // color
// Colorize(colorCode int) S
// Green() S
// Yellow() S
// Red() S
// Blue() S
// LightGreen() S
// LightYellow() S
// LightRed() S
// LightBlue() S
// Cyan() S
// LightCyan() S
// Magenta() S
// LightMagenta() S
// Gray() S
// LightGray() S
// White() S
// Black() S
// Dim() S
// }
