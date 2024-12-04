package gen

import "github.com/periaate/blume/gen/T"

type Lennable interface {
	~[]any | ~string | ~map[any]any | ~chan any | ~[]string
}

// Must panics if the error is not nil.
func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func Assert(a any, msg string) {
	switch v := a.(type) {
	case bool:
		if !v {
			panic(msg)
		}
	case error:
		if v != nil {
			panic(msg)
		}
	case string:
		if v == "" {
			panic(msg)
		}
	case []any:
		if len(v) == 0 {
			panic(msg)
		}
	case map[any]any:
		if len(v) == 0 {
			panic(msg)
		}
	}
}

func ArrayOrDefault[A any](inp []A, def ...A) []A {
	if len(inp) == 0 {
		return def
	}
	return inp
}

// Middleware wraps a function with a middleware, calling it before the next function.
func Middleware[A, B any](mw func(A)) T.Transformer[T.Monadic[A, B]] {
	return func(next T.Monadic[A, B]) T.Monadic[A, B] {
		return func(a A) B {
			mw(a)
			return next(a)
		}
	}
}

// Ignore returns the first argument and ignores the second.
func Ignore[A, B any](a A, _ B) A { return a }

// Ignores transforms a function from F : A -> (B, C) to [T.Monadic] F : A -> B.
func Ignores[A, B, C any](fn func(A) (B, C)) func(A) B { return func(a A) B { return Ignore(fn(a)) } }

// IgnoresNil transforms a function from F : () -> (B, C) to [T.Niladic] F : A -> B.
func IgnoresNil[B, C any](fn func() (B, C)) func() B { return func() B { return Ignore(fn()) } }

// func Repeat(dur time.Duration, fn func()) (stop func()) {
// 	stopch := make(chan struct{})
// 	var stopped bool
// 	go func() {
// 		tck := time.NewTicker(dur)
// 		defer tck.Stop()
// 		defer close(stopch)
// 		defer func() { stopped = true }()
// 		for {
// 			select {
// 			case <-stopch:
// 				return
// 			case <-tck.C:
// 				go fn()
// 			}
// 		}
// 	}()
//
// 	return func() {
// 		if stopped {
// 			return
// 		}
// 		stopch <- struct{}{}
// 	}
// }
