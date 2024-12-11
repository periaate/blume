package core

// All returns true if all arguments pass the [Predicate].
func All[A any](fns ...Monadic[A, bool]) Monadic[[]A, bool] {
	fn := PredAnd(fns...)
	return func(args []A) bool {
		for _, arg := range args {
			if !fn(arg) { return false }
		}
		return true
	}
}

// Any returns true if any argument passes the [Predicate].
func Any[A any](fns ...Monadic[A, bool]) Monadic[[]A, bool] {
	fn := PredOr(fns...)
	return func(args []A) bool {
		for _, arg := range args {
			if fn(arg) { return true }
		}
		return false
	}
}

// Filter returns a slice of arguments that pass the [Predicate].
func Filter[A any](fns ...Monadic[A, bool]) Monadic[[]A, []A] {
	fn := PredAnd(fns...)
	return func(args []A) (res []A) {
		res = make([]A, 0, len(args))
		for _, arg := range args {
			if fn(arg) { res = append(res, arg) }
		}
		return res
	}
}

// Map applies the function to each argument and returns the results.
func Map[A, B any](fn Monadic[A, B]) Monadic[[]A, []B] {
	return func(args []A) (res []B) {
		res = make([]B, 0, len(args))
		for _, arg := range args {
			res = append(res, fn(arg))
		}
		return res
	}
}

// Reduce applies the function to each argument and returns the result.
func Reduce[A any, B any](fn Dyadic[B, A, B], init B) Monadic[[]A, B] {
	return func(args []A) B {
		res := init
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [Predicate].
func Not[A any](fn Monadic[A, bool]) Monadic[A, bool] { return func(t A) bool { return !fn(t) } }

// Is returns a [Predicate] that checks if the argument is in the list.
func Is[C comparable](A ...C) Monadic[C, bool] {
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
func First[A any](fns ...Monadic[A, bool]) Monadic[[]A, Option[A]] {
	fn := PredOr[A](fns...)
	return func(args []A) Option[A] {
		for _, arg := range args {
			if fn(arg) { return Some(arg) }
		}
		return None[A]()
	}
}

func StrIs[A, B ~string](vals ...A) Monadic[B, bool] {
	is := Is(vals...)
	return func(b B) bool { return is(A(b)) }
}


// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipe[A any](fns ...Monadic[A, A]) Monadic[A, A] {
	return func(a A) A {
		for _, fn := range fns {
			a = fn(a)
		}
		return a
	}
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipes[A any](fns ...Monadic[A, A]) Monadic[[]A, []A] { return Map(Pipe(fns...)) }

// Cat concatenates two [Monadic] functions into a single [Monadic] function.
func Cat[A, B, C any](a Monadic[A, B], b Monadic[B, C]) Monadic[A, C] {
	return func(c A) C { return b(a(c)) }
}

func Con[A, B any](a Monadic[A, B], b Monadic[B, bool]) Monadic[A, bool] {
	return func(c A) bool { return b(a(c)) }
}
