package blume

func All[A any](fn func(A) bool) func([]A) bool {
	return func(args []A) bool {
		for _, arg := range args {
			if !fn(arg) {
				return false
			}
		}
		return true
	}
}

func Any[A any](fn func(A) bool) func([]A) bool {
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
func Filter[A any](fns ...func(A) bool) func([]A) []A {
	fn := PredAnd(fns...)
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
func Map[A, B any](fn func(A) B) func([]A) []B {
	return func(args []A) (res []B) {
		res = make([]B, 0, len(args))
		for _, arg := range args {
			res = append(res, fn(arg))
		}
		return res
	}
}

// Reduce applies the function to each argument and returns the result.
func Reduce[A any, B any](fn func(B, A) B, init B) func([]A) B {
	return func(args []A) B {
		res := init
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [Predicate].
func Not[A any](fn func(A) bool) func(A) bool { return func(t A) bool { return !fn(t) } }

// Is returns a [Predicate] that checks if the argument is in the list.
func Is[C comparable](A ...C) func(C) bool {
	in := make(map[C]bool, len(A))
	for _, a := range A {
		in[a] = true
	}
	return func(c C) bool {
		_, ok := in[c]
		return ok
	}
}

// First returns the first value which passes the [Predicate].
func First[A any](fns ...func(A) bool) func([]A) Option[A] {
	fn := PredOr(fns...)
	return func(args []A) Option[A] {
		for _, arg := range args {
			if fn(arg) {
				return Some(arg)
			}
		}
		return None[A]()
	}
}

func StrIs[A, B ~string](vals ...A) func(B) bool {
	is := Is(vals...)
	return func(b B) bool { return is(A(b)) }
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipe[A any](fns ...func(A) A) func(A) A {
	return func(a A) A {
		for _, fn := range fns {
			a = fn(a)
		}
		return a
	}
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipes[A any](fns ...func(A) A) func([]A) []A { return Map(Pipe(fns...)) }

// Cat concatenates two [FnA] functions into a single [FnA] function.
func Cat[A, B, C any](a func(A) B, b func(B) C) func(A) C { return func(c A) C { return b(a(c)) } }

func PredAnd[A any](preds ...func(A) bool) func(A) bool {
	return func(a A) bool {
		for _, pred := range preds {
			if !pred(a) {
				return false
			}
		}
		return true
	}
}

func PredOr[A any](preds ...func(A) bool) func(A) bool {
	return func(a A) bool {
		for _, pred := range preds {
			if pred(a) {
				return true
			}
		}
		return false
	}
}

func Lazy[A any](fn func() A) func() A {
	var loaded bool
	var value A
	return func() A {
		if !loaded {
			value, loaded = fn(), true
		}
		return value
	}
}

func Memo[K comparable, V any](fn func(K) V) func(K) V {
	values := map[K]V{}
	return func(input K) (res V) {
		res, ok := values[input]
		if !ok {
			res = fn(input)
			values[input] = res
		}
		return res
	}
}

func Negate[A any](fn func(A) bool) func(A) bool { return func(a A) bool { return !fn(a) } }
