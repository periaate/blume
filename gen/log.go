package gen

import (
	"fmt"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/types"
)

var _ = Zero[any]


func Logf[A any](f string) func([]A) []A {
	sar := SplitWithAll(f, true, "{", "}", ":")

	ebd := EmbedDelims(sar, [2]string{"{", "}"})
	res := ebd.Array.Filter(func(t types.Tree[string]) bool {
		return t.Array.Len() != 0
	})

	f = ReplaceRegex[string](`{.*:(.*)}`,`$1`)(f)

	getters := []Monadic[[]A, A]{}
	for _, t := range res.Values() {
		i := ToInt(t.Array.Values()[0].Value).Unwrap()
		getters = append(getters, get[A](i))
	}

	return func(args []A) []A {
		values := []any{}
		for _, getter := range getters {
			values = append(values, getter(args))
		}
		fmt.Printf(f, values...)
		return args
	}
}


func Logs[A, B any](f string, args ...Monadic[[]A, B]) Monadic[[]A, []A] {
	return func(a []A) []A {
		toLog := []any{}
		for _, arg := range args {
			toLog = append(toLog, arg(a))
		}

		fmt.Printf(f, toLog...)
		return a
	}
}

func get[A any](i int) Monadic[[]A, A] { return func(a []A) A { return a[i] } }
