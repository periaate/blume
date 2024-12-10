package core

type Err[A any] struct {
	Val A
	Str string
	Err error
}

func (e Err[A]) Reason() string             { return e.Str }
func (e Err[A]) Value() A                   { return e.Val }
func (e Err[A]) Values() (A, string, error) { return e.Val, e.Str, e.Err }
func (e Err[A]) Error() string              {
	if e.Err == nil { return e.Str }
	if e.Str == ""  { return e.Err.Error() }
	return "error: " + e.Err.Error() + "; beacuse: " + e.Str
}

type StrError string

func (e StrError) Error() string { return string(e) }

var _ Result[any, any] = Res[any, any]{}

type Res[V, E any] struct {
	Value V
	Error Error[E]
}

// Unwrap returns the value of the result.
func (r Res[V, E]) Values() (V, Error[E]) { return r.Value, r.Error }

// Unwrap returns the value of the result.
func (r Res[V, E]) Unwrap() V {
	if r.Error != nil { panic(r.Error) }
	return r.Value
}

// Or returns the value of the result, or a default value if it is an error.
func (r Res[V, E]) Or(def V) V {
	if r.IsErr() { return def }
	return r.Value
}

func (r Res[V, E]) IsOk() bool  { return r.Error == nil }
func (r Res[V, E]) IsErr() bool { return r.Error != nil }

// Match takes two functions and calls the first if the result is a success, and the second if it is an error.
func (r Res[V, E]) Match(ok func(V), err func(Error[E])) {
	if r.IsOk() {
		ok(r.Value)
	} else {
		err(r.Error)
	}
}

func (r Res[V, E]) Err(f func(Error[E])) V {
	if f == nil { return r.Value }
	if r.Error != nil { f(r.Error) }
	return r.Value
}

func (r Res[V, E]) Ok(f func(V)) Error[E] {
	if r.Error == nil { f(r.Value) }
	return r.Error
}
