package gen

import (
	"fmt"

	. "github.com/periaate/blume/core"
)

type Condition[A any] func(A) Error[string]

func AtLeast[N Numeric](n N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i < n {
			er = Errorf[string]("at least "+fmt.Sprint(n), "input "+fmt.Sprint(i)+" is less than "+fmt.Sprint(n), "")
		}
		return nil
	}
}

func AtMost[N Numeric](n N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i > n {
			er = Errorf[string]("at most "+fmt.Sprint(n), "input "+fmt.Sprint(i)+" is greater than "+fmt.Sprint(n), "")
		}
		return nil
	}
}

func Between[N Numeric](min, max N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i < min {
			er = Errorf[string]("between "+fmt.Sprint(min)+" and "+fmt.Sprint(max), "input "+fmt.Sprint(i)+" is less than "+fmt.Sprint(min), "")
		} else if i > max {
			er = Errorf[string]("between "+fmt.Sprint(min)+" and "+fmt.Sprint(max), "input "+fmt.Sprint(i)+" is greater than "+fmt.Sprint(max), "")
		}
		return nil
	}
}

func Exactly[N Numeric](n N) Condition[N] {
	return func(i N) (er Error[string]) {
		if i != n {
			er = Errorf[string]("exactly "+fmt.Sprint(n), "input "+fmt.Sprint(i)+" is not equal to "+fmt.Sprint(n), "")
		}
		return nil
	}
}

func NotZero[K comparable](k K) Error[string] {
	var zero K
	if k == zero { return Errorf[string]("not zero", "input is zero", "") }
	return nil
}

func Len[A any](cond Condition[int]) Condition[[]A] {
	return func(ar []A) Error[string] { return cond(len(ar)) }
}