package util

import (
	"fmt"

	"github.com/periaate/blume/gen"
)

type Answer[A any] struct {
	Name   string
	Reason string
	Value  A
}

type Condition[A any] func(A) *Answer[A]

func LongerThan[L gen.Lennable](min int) Condition[L] {
	return func(l L) *Answer[L] {
		ans := &Answer[L]{"LongerThan", fmt.Sprintf("length less than %v", min), l}
		if len(l) <= min {
			return ans
		}
		return nil
	}
}

func AtLeast[N gen.Numeric](min N) Condition[N] {
	return func(i N) *Answer[N] {
		ans := &Answer[N]{"AtLeast", "value less than minimum", min}
		if i < min {
			return ans
		}
		return nil
	}
}

func AtMost[N gen.Numeric](max N) Condition[N] {
	return func(i N) *Answer[N] {
		ans := &Answer[N]{"AtMost", "value greater than maximum", max}
		if i > max {
			return ans
		}
		return nil
	}
}

func InRange[N gen.Numeric](min, max N) Condition[N] {
	return func(i N) *Answer[N] {
		ans := &Answer[N]{"InRange", fmt.Sprintf("value not in range [%v:%v]", min, max), 0}
		if i < min || i > max {
			return ans
		}
		return nil
	}
}

func NotZero[K comparable]() Condition[K] {
	var zero K
	return func(i K) *Answer[K] {
		ans := &Answer[K]{"NotZero", "value is zero", zero}
		if i == zero {
			return ans
		}
		return nil
	}
}

func Free[A any](cond gen.Predicate[A], name, reason string, val A) Condition[A] {
	return func(i A) *Answer[A] {
		if !cond(i) {
			return &Answer[A]{"Free", "condition not met", val}
		}
		return nil
	}
}
