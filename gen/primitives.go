package gen

type Lennable interface {
	~[]any | ~string | ~map[any]any | ~chan any
}

// Must panics if the error is not nil.
func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

// Middleware wraps a function with a middleware, calling it before the next function.
func Middleware[A, B any](mw func(A)) Transformer[Monadic[A, B]] {
	return func(next Monadic[A, B]) Monadic[A, B] {
		return func(a A) B {
			mw(a)
			return next(a)
		}
	}
}

// Ignore returns the first argument and ignores the second.
func Ignore[A, B any](a A, _ B) A { return a }

// Ignores transforms a function from F : A -> (B, C) to [Monadic] F : A -> B.
func Ignores[A, B, C any](fn func(A) (B, C)) func(A) B { return func(a A) B { return Ignore(fn(a)) } }

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
