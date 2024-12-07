package T

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

func (r Result[A]) Ignore() A { return r.value }

// Unwrap returns the value of the result.
func (r Result[A]) Values() (A, Error[any]) { return r.value, r.err }

// Unwrap returns the value of the result.
func (r Result[A]) Unwrap() A {
	if r.IsErr() {
		panic(r.err)
	}
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
	case string:
		e = Err[any]{err: errt(v)}
	}
	return Result[A]{value: val, err: e}
}

func (r Result[A]) Err(f func(Error[any])) A {
	if f == nil {
		return r.value
	}
	if r.IsErr() {
		f(r.err)
	}
	return r.value
}

func (r Result[A]) Ok(f func(A)) Error[any] {
	if r.IsOk() {
		f(r.value)
	}
	return r.err
}

var _ Resultable[any] = Result[any]{}

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

	Err(func(Error[any])) A
	Ok(func(A)) Error[any]
}

type (
	Orrable[A any]  interface{ Or(Default A) A }
	Contains[A any] interface{ Contains(args ...A) bool }
)

type Filters[A any] interface {
	Filter(args ...Predicate[A]) []A
}

type Stringable interface{ String() string }

/*

`T`
```go
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
	Err(func(Error[any])) A
	Ok(func(A)) Error[any]
}

type Error[A any] interface {
	Error() string
	Err() error
	Reason() string
	Data() A
}
```

`gen`
```go
// First returns the first value which passes the [T.Predicate].
func First[A any](fns ...T.Predicate[A]) T.Reducer[A, T.Result[A]]

// Filter returns a slice of arguments that pass the [T.Predicate].
func Filter[A any](fns ...T.Predicate[A]) T.Mapper[A, A]

// Map applies the function to each argument and returns the results.
func Map[A, B any](fn T.Monadic[A, B]) T.Mapper[A, B]
```

`fsio`
```go
// ReadDir reads the directory and returns a list of files as typ.String.
// Directories always end with `/`.
func ReadsDir[S ~string](fp S) T.Result[gen.Array[typ.String]]
```


`T`, implemented in `typ.String`
```go
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
	ToInt() Result[int]
	ToInt8() Result[int8]
	ToInt16() Result[int16]
	ToInt32() Result[int32]
	ToInt64() Result[int64]
	ToUint() Result[uint]
	ToUint8() Result[uint8]
	ToUint16() Result[uint16]
	ToUint32() Result[uint32]
	ToUint64() Result[uint64]
	ToFloat32() Result[float32]
	ToFloat64() Result[float64]
	Colorize(colorCode int) S
	ToUpper() S
	ToLower() S
	Trim() S
	TrimPrefix(prefix string) S
	TrimSuffix(suffix string) S
	TrimSpace() S
```

*/
