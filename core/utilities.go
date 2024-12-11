package core

func PredAnd[A any](preds ...Monadic[A, bool]) Monadic[A, bool] {
	return func(a A) bool {
		for _, pred := range preds {
			if !pred(a) { return false }
		}
		return true
	}
}

func PredOr[A any](preds ...Monadic[A, bool]) Monadic[A, bool] {
	return func(a A) bool {
		for _, pred := range preds {
			if pred(a) { return true }
		}
		return false
	}
}

func Or[C comparable](a, b C) (res C) {
	if a == res { return b }
	return a
}
