package blume

type Option[A any] = Either[A, bool]
type Result[A any] = Either[A, error]

type Either[A, B any] struct {
	Value A
	Other B
}

func (r Either[A, B]) Unwrap() (A, B) { return r.Value, r.Other }

func (r Either[A, B]) Then(fn func(A) A) (res Either[A, B]) {
	if r.IsOk() { return res.Pass(fn(r.Value)) }
	return r
}

func (r Either[A, B]) Else(fn func(B) B) (res Either[A, B]) {
	if !r.IsOk() { return res.Fail(fn(r.Other)) }
	return r
}

func NotNil[A any](inp *A) Option[A] {
	if inp == nil { return None[A]() }
	return Some(*inp)
}

func Or[A any](def A, in A, handle ...any) (res A) {
	if len(handle) != 0 {
		last := handle[len(handle)-1]
		switch val := last.(type) {
		case bool:
			if val {
				return in
			}
			return def
		case error:
			if val == nil {
				return in
			}
			return def
		default:
			return def
		}
	}
	anyin := any(in)
	switch inv := anyin.(type) {
	case String, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, bool:
		if isZero(inv) {
			return in
		}
	}
	return def
}

func Must[A any](a A, handle ...any) A {
	if len(handle) == 0 {
		return a
	}
	last := handle[len(handle)-1]
	switch val := last.(type) {
	case bool:
		if val {
			return a
		}
		panic("must called with false bool")
	default:
		if val == nil {
			return a
		}
		panic(P.S("must called with nil value").S(handle...))
	}
}

func Mustnt[A, B any](a A, handle any) B {
	switch val := handle.(type) {
	case bool:
		if !val {
			return any(val).(B)
		}
		panic("mustnt called with true bool")
	default:
		if val != nil {
			return any(val).(B)
		}
		panic(P.S("mustnt called with non nil value").S(handle))
	}
}

func (r Either[A, B]) Pass(val A) Either[A, B] {
	r.Value = val
	var other B
	switch any(other).(type) {
	case bool:
		r.Other = any(true).(B)
	}
	return r
}

func (r Either[A, B]) Fail(val ...any) Either[A, B] {
	var other B
	switch any(other).(type) {
	case bool:
		return Either[A, B]{}
	case error:
		return Either[A, B]{Other: Cast[B](P.S(val...)).Must()}
	}
	return Either[A, B]{}
}

func Pass[A, B any](val A) (res Either[A, B])      { return res.Pass(val) }
func Fail[A, B any](val ...any) (res Either[A, B]) { return res.Fail(val...) }

func (r Either[A, B]) Must() A              { return Must(r.Value, r.Other) }
func (r Either[A, B]) Mustnt() B            { return Mustnt[A, B](r.Value, r.Other) }
func (r Either[A, B]) Or(def A) A           { return Or(def, r.Value, r.Other) }
func (r Either[A, B]) OrDef() (def A)           { return Or(def, r.Value, r.Other) }
func (r Either[A, B]) OrExit(args ...any) A { return OrExit(r, args...) }

func None[A any]() Option[A]          { return Option[A]{Other: false} }
func Some[A any](value A) Option[A]   { return Option[A]{Value: value, Other: true} }
func Err[A any](msg ...any) Result[A] { return Result[A]{Other: P.Errorf(msg...)} }
func Ok[A any](value A) Result[A]     { return Result[A]{Value: value} }

func Match[A, B, C any](r Either[A, B], value func(A) C, other func(B) C) C {
	switch IsOk(r) {
	case true:
		return value(r.Value)
	default:
		return other(r.Other)
	}
}

func (e Either[A, B]) IsOk() bool { return IsOk(e.Other) }

func IsOk[A any](a A, handle ...any) bool {
	if len(handle) == 0 {
		handle = append(handle, a)
	}
	last := handle[len(handle)-1]
	switch val := last.(type) {
	case bool:
		if val {
			return true
		}
	default:
		if val == nil {
			return true
		}
	}
	return false
}

func AllOk[A, B any](arr []Either[A, B]) bool {
	return Reduce(func(acc bool, cur Either[A, B]) bool { return acc && cur.IsOk() }, true)(arr)
}
