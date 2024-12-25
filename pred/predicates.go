package pred

// Not negates a predicate.
func Not[A any](fn func(A) bool) func(A) bool { return func(arg A) bool { return !fn(arg) } }

// Is returns a predicate that checks if the argument is in the list.
func Is[K comparable](args ...K) func(K) bool {
	in := make(map[K]bool, len(args))
	for _, a := range args {
		in[a] = true
	}
	return func(c K) bool {
		_, ok := in[c]
		return ok
	}
}

// Isnt composes [Not] and [Is].
func Isnt[K comparable](args ...K) func(K) bool { return Not(Is(args...)) }

func And[A any](preds ...func(A) bool) func(A) bool {
	return func(arg A) bool {
		for _, pred := range preds {
			if !pred(arg) {
				return false
			}
		}
		return true
	}
}

func Or[A any](preds ...func(A) bool) func(A) bool {
	return func(arg A) bool {
		for _, pred := range preds {
			if pred(arg) {
				return true
			}
		}
		return false
	}
}

func Every[A any](cond func(A) bool) func([]A) bool {
	return func(args []A) bool {
		for _, arg := range args {
			if !cond(arg) {
				return false
			}
		}
		return true
	}
}

func Any[A any](cond func(A) bool) func([]A) bool {
	return func(args []A) bool {
		for _, arg := range args {
			if cond(arg) {
				return true
			}
		}
		return false
	}
}
