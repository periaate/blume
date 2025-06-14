package blume

type Option[A any] = Either[A, bool]
type Result[A any] = Either[A, error]

type Either[A, B any] struct {
	Value A
	Other B
}

func (r Either[A, B]) Unwrap() (A, B) { return r.Value, r.Other }

func (r Either[A, B]) Map(fn any, args ...any) (res Either[A, B]) {
	if !r.IsOk() { return r }
	if f, ok := fn.(func(A) A); ok { return res.Pass(f(r.Value)) }
	Function(fn).Call(args...)
	return r
}

func (r Either[A, B]) Then(fn any, args ...any) (res Either[A, B]) {
	if !r.IsOk() { return r }
	switch {
	case len(args) > 0: Function(fn).Call(args...)
	default: switch f := fn.(type) {
	case func(A) A: return res.Pass(f(r.Value))
	case func(Either[A, B]) Either[A, B]: return f(r)
	case func(A): f(r.Value)
	}}
	return r
}

func (r Either[A, B]) Else(fn any, args ...any) (res Either[A, B]) {
	if !r.IsOk() { return r }
	switch {
	case len(args) > 0: Function(fn).Call(args...)
	default: switch f := fn.(type) {
	case func(A) B: return res.Fail(f(r.Value))
	case func(Either[A, B]) Either[A, B]: return f(r)
	case func(B): f(r.Other)
	}}
	return r
}

// func (r Either[A, B]) Then(fn func(A) A) (res Either[A, B]) {
// 	if r.IsOk() { return res.Pass(fn(r.Value)) }
// 	return r
// }
//
// func (r Either[A, B]) Else(fn func(B) B) (res Either[A, B]) {
// 	if !r.IsOk() { return res.Fail(fn(r.Other)) }
// 	return r
// }

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
			if val != nil {
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

func HasTo[A, B, C any](fn func(A) Either[B, C]) func(A) B { return func(a A) B { return fn(a).Must() } }

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
	case error:
		if val == nil { return a }
		panic(val)
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
	switch any(r).(type) {
	case Either[A, bool]: r.Other = any(true).(B)
	case Either[A, error]: return Either[A, B]{Value: val}
	}
	return r
}

func (r Either[A, B]) Fail(val ...any) (res Either[A, B]) {
	switch any(r).(type) {
	case Either[A, bool]: return
	case Either[A, error]:
		res, ok := any(Err[A](val...)).(Either[A, B])
		if !ok { panic("Either[A, error] could not be cast to Either[A, error]; impossible invariant") }
		return res
	}
	return
}


func (r Either[A, B]) Auto(arg any, args ...any) Either[A, B] {
	val := Arr(args...).Get(-1)
	if val.IsOk() {
		switch v := val.Value.(type) {
		case bool:  if v        {
			value, ok := arg.(A)
			if !ok { return r.Fail() }
			return r.Pass(value)
		} else { return r.Fail()  }
		case error: if v == nil {
			value, ok := arg.(A)
			if !ok { return r.Fail() }
			return r.Pass(value)
		} else { return r.Fail(v) } }
	}

	switch v := any(arg).(type) {
	case Either[A, bool] : if v.IsOk() { return r.Pass(v.Value)
                         } else        { return r.Fail()        }
	case Either[A, error]: if v.IsOk() { return r.Pass(v.Value)
                         } else        { return r.Fail(v.Other) }
	}
	return Either[A, B]{} // alternatively panic; illegal invariant
}

func Pass[A, B any](val A) (res Either[A, B])      { return res.Pass(val) }
func Fail[A, B any](val ...any) (res Either[A, B]) { return res.Fail(val...) }

func (r Either[A, B]) Must() A              { return Must(r.Value, r.Other) }
func (r Either[A, B]) Mustnt() B            { return Mustnt[A, B](r.Value, r.Other) }
func (r Either[A, B]) Or(def A) A           { return Or(def, r.Value, r.Other) }
func (r Either[A, B]) OrDef() A             { return r.Value }
func (r Either[A, B]) OrExit(args ...any) A { return OrExit(r, args...) }
func (r Either[A, B]) OrExits() A { return OrExits(r) }

func None[A any]() Option[A]          { return Option[A]{Other: false} }
func Some[A any](value A) Option[A]   { return Option[A]{Value: value, Other: true} }
func Err[A any](msg ...any) Result[A] { return Result[A]{Other: P.Errorf(msg...)} }
func Ok[A any](value A) Result[A]     { return Result[A]{Value: value} }

func (e Either[A, B]) IsOk() bool { return IsOk(e.Other) }
func (e Either[A, B]) IsSome() bool { return IsOk(e.Other) }
func (e Either[A, B]) IsNone() bool { return !IsOk(e.Other) }
func (e Either[A, B]) IsErr() bool { return !IsOk(e.Other) }

func AllOk[A, B any](arr []Either[A, B]) bool {
	return Reduce(func(acc bool, cur Either[A, B]) bool { return acc && cur.IsOk() }, true)(arr)
}

func (r Either[A, B]) Expect(pred Pred[A], args ...any) Either[A, B] {
	if r.IsOk() && !pred(r.Value) { return r.Fail(args) }
	return r
}

