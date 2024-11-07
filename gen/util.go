package gen

/*
util.go includes more specialized functional forms.
*/

// First returns the first value which passes the [Predicate].
func First[A any](fn Predicate[A]) Monadic[[]A, A] {
	return func(args []A) (res A) {
		for _, arg := range args {
			if fn(arg) {
				return arg
			}
		}
		return res
	}
}
