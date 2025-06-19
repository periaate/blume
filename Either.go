package blume

import "fmt"

type Option[A any] = Either[A, bool]
type Result[A any] = Either[A, error]

type Either[A, B any] struct {
	Value A
	Other B
}

func (r Either[A, B]) Unwrap() (A, B) { return r.Value, r.Other }

func IsOk(handle any) (ok bool) {
	switch val := handle.(type) {
	case bool:
		return val
	case error:
		return val == nil
	default: return
	}
}

func Or[A any](def A, in A, handle ...any) (res A) {
	if _, ok := Get(handle, -1); ok { return in }
	return
}

func Must[A any](a A, handle ...any) A {
	last, ok := Get(handle, -1)
	if !ok { last = any(a) }
	switch val := last.(type) {
	case bool:  if !val       { panic("must called with false bool") }
	case error: if val != nil { panic(val) }
	}
	return a
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
	val, ok := Get(args, -1)
	if ok {
		switch v := val.(type) {
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
	case Either[A, bool] : if v.Other { return r.Pass(v.Value)
                         } else        { return r.Fail()        }
	case Either[A, error]: if v.Other == nil { return r.Pass(v.Value)
                         } else        { return r.Fail(v.Other) }
	}
	return Either[A, B]{}
}

func (r Either[A, B]) Must() A              { return Must(r.Value, r.Other) }
func (r Either[A, B]) Or(def A) A           { return Or(def, r.Value, r.Other) }
func (r Either[A, B]) OrDef() (def A)       { return Or(def, r.Value, r.Other) }
func (r Either[A, B]) OrExit(args ...any) A { return OrExit(r, args...) }

func None[A any]() Option[A]          { return Option[A]{Other: false} }
func Some[A any](value A) Option[A]   { return Option[A]{Value: value, Other: true} }
func Err[A any](msg ...any) Result[A] { return Result[A]{Other: fmt.Errorf("%s", msg...)} }
func Ok[A any](value A) Result[A]     { return Result[A]{Value: value} }

func (e Either[A, B]) IsOk() bool { return IsOk(e.Other) }

