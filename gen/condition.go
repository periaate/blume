package gen

import (
	. "github.com/periaate/blume/core"
)

type Condition[A any] func(A) Error[string]

func AtLeast[N Numeric](n N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i < n {
			er = Errorf[string]("condition [AtLeast] failed because input [%v] is less than [%v]", i, n)
		}
		return nil
	}
}

func AtMost[N Numeric](n N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i > n {
			er = Errorf[string]("condition [AtMost] failed because input [%v] is greater than [%v]", i, n)
		}
		return nil
	}
}

func Between[N Numeric](min, max N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i < min {
			er = Errorf[string]("condition [Between] failed because input [%v] is less than [%v]", i, min)
		} else if i > max {
			er = Errorf[string]("condition [Between] failed because input [%v] is greater than [%v]", i, max)
		}
		return nil
	}
}

func Exactly[N Numeric](n N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i != n {
			er = Errorf[string]("condition [Exactly] failed because input [%v] is not equal to [%v]", i, n)
		}
		return nil
	}
}

func NotZero[K comparable](k K) Error[string] {
	var zero K
	if k == zero { return Errorf[string]("condition [NotZero] failed because input [%v] is zero", k) }
	return nil
}

func Len[A any](cond Condition[int]) Condition[[]A] {
	return func(ar []A) Error[string] { return cond(len(ar)) }
}
