package core

import "fmt"

func Results[V, E any](value V, err any) (res Res[V, E]) {
	res = Res[V, E]{Value: value}
	if err == nil { return res }

	switch err := err.(type) {
		case Error[E]: res.Error = err
		case E: res.Error = Err[E]{Val: err}
		case string: res.Error = Err[E]{Err: StrError(err)}
		default: panic("invalid error type")
	}

	return res
}

func OK[V, E any](v V) Result[V, E] { return Res[V, E]{Value: v} }
func Errors[V, E any](s string, err error, e E) Result[V, E] {
	return Results[V, E](Zero[V](), Err[E]{Str: s, Err: err, Val: e})
}
func Errorf[E any](format string, args ...any) Error[E] { return Err[E]{Err: StrError(fmt.Sprintf(format, args...))} }

func Some[V any](v V) Option[V] { return Results[V, Nothing](v, nil) }

func None[V any](args ...any) Option[V]    {
	switch len(args) {
	case 0:
	case 1:
		switch err := args[0].(type) {
			case Error[Nothing]: return Results[V, Nothing](Zero[V](), err)
			case Nothing: return Results[V, Nothing](Zero[V](), Err[Nothing]{Val: err})
			case string: return Results[V, Nothing](Zero[V](), Err[Nothing]{Err: StrError(err)})
			case error: return Results[V, Nothing](Zero[V](), Err[Nothing]{Err: err})
		}
	default:
		if f, ok := args[0].(string); ok {
			return Results[V, Nothing](Zero[V](), Errorf[Nothing](f, args[1:]...))
		}
	}
	return Results[V, Nothing](Zero[V](), Err[Nothing]{}) 
}

func Zero[A any]() (a A)        { return }

func Either[V any](v V, err error) Option[V] {
	if err != nil { return Results[V, Nothing](Zero[V](), Err[Nothing]{Err: err}) }
	return Some(v)
}

// Must panics if the error is not nil.
func Must[A any](a A, err error) A {
	if err != nil { panic(err) }
	return a
}

// Ignore returns the first argument and ignores the second.
func Ignore[A, B any](a A, _ B) A { return a }

func Assert(a any, msg string) {
	switch v := a.(type) {
		case bool:        if !v { panic(msg) }
		case error:       if v != nil { panic(msg) }
		case string:      if v == "" { panic(msg) }
		case []any:       if len(v) == 0 { panic(msg) }
		case map[any]any: if len(v) == 0 { panic(msg) }
	}
}

