package blume

func Opt[A any](a A, other any) Option[A] {
	if IsOk(a, other) {
		return Some(a)
	}
	return None[A]()
}

func Res[A any](a A, other any) Result[A] {
	if IsOk(a, other) {
		return Ok(a)
	}
	return Err[A](other)
}

func Cast[T any](a any) Option[T] {
	value, ok := a.(T)
	return Opt(value, ok)
}
