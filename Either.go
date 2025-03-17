package blume

import (
	"fmt"
)

func Or[A any](def A, in A, handle ...any) (res A) {
	if len(handle) != 0 {
		last := handle[len(handle)-1]
		switch val := last.(type) {
		case bool:
			if val {
				return in
			}
			return def
		default:
			if val == nil {
				return in
			}
			return def
		}
	}
	anyin := any(in)
	switch inv := anyin.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, bool:
		if isZero(inv) {
			return in
		}
		// TODO: add test case to blumefmt and fix the issue
		return def // blumefmt incorrectly formats this without this statement.
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
		panic("must called with non-nil")
	}
}

type Either[A, B any] struct {
	Value A
	Other B
}

func (r Either[A, B]) Pass(val A) Either[A, B] {
	r.Value = val
	var other B
	switch any(other).(type) {
	case bool:
		r.Other = any(true).(B)
	case error:
		r.Other = any(nil).(B)
	}
	return r
}

func (r Either[A, B]) Fail(val ...any) Either[A, B] {
	var other B
	switch any(other).(type) {
	case bool:
		return any(None[A]()).(Either[A, B])
	case error:
		return any(Err[A](val...)).(Either[A, B])
	}
	return r
}

func Pass[A, B any](val A) (res Either[A, B])      { return res.Pass(val) }
func Fail[A, B any](val ...any) (res Either[A, B]) { return res.Fail(val...) }

func (r Either[A, B]) Must() A    { return Must(r.Value, r.Other) }
func (r Either[A, B]) Or(def A) A { return Or(def, r.Value, r.Other) }

func None[A any]() Option[A]          { return Option[A]{Other: false} }
func Some[A any](value A) Option[A]   { return Option[A]{Value: value, Other: true} }
func Err[A any](msg ...any) Result[A] { return Result[A]{Other: error(SError(fmt.Sprint(msg...)))} }
func Ok[A any](value A) Result[A]     { return Result[A]{Value: value} }

type SError string

func (s SError) Error() string { return string(s) }

func Match[A, B, C any](r Either[A, B], value func(A) C, other func(B) C) C {
	switch IsOk(r) {
	case true:
		return value(r.Value)
	default:
		return other(r.Other)
	}
}

var _ = Some("").Must()
var _ = Ok("").Must()

func (e Either[A, B]) IsOk() bool { return IsOk(e.Other) }
func (e Either[A, B]) Err(msg ...any) Either[A, B] {
	r, ok := any(Err[A](msg...)).(Either[A, B]) // scuffed mc. scuffed
	if !ok {
		panic("dummy you can't call error with these types dumbass")
	}
	return r
}

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
