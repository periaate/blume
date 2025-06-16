package match

import (
	"fmt"
	"reflect"
	"testing"
)

func size[V any, T []V | string](t T) int   { return len(t) }
func stringEq(s1, s2 string) bool { return s1 == s2 }

func slice[T any](s1, s2 []T) bool { return reflect.DeepEqual(s1, s2) }

func expectSlice[T any](e T) func(i T) (a, b T, ok bool) { return func(i T) (a, b T, ok bool) { return e, i, reflect.DeepEqual(e, i) } }

type comp[T any] func(T) (T, T, bool)

func splitTest[Arr string | []Item, Item any](t *testing.T, input Arr, fn comp[[]Arr], matcher Match[Window[Arr], SplitResult[Arr]]) {
	itr, err := ToIter[Arr, Item](input)
	if err != nil { t.Error(err); return }
	expect, actual, ok := fn(Split(itr, matcher))
	val := fmt.Sprintf("=== Split Test ===\nInput:\t\"%v\"\nExpect:\t%v\nActual:\t%v\n", prettyPrintSlice(expect), prettyPrintSlice(expect), prettyPrintSlice(actual))
	if !ok { t.Error(val); return }
	fmt.Print(val)
}

func TestMatch(t *testing.T) {
	splitTest[string, byte](t,
		"aba",
		expectSlice([]string{"a", "a"}),
		Is(size[string], stringEq, Act("b", Skip)),
	)

	splitTest[string, byte](t,
		"abba",
		expectSlice([]string{"a", "a"}),
		Is(size[string], stringEq, Act("b", Skip)),
	)

	splitTest[string, byte](t,
		"abab",
		expectSlice([]string{"a", "a"}),
		Is(size[string], stringEq, Act("b", Skip)),
	)

	splitTest[string, byte](t,
		"Hello, World!   and you !",
		expectSlice([]string{"Hello", "World", "!", "and", "you", "!"}),
		Is(size[string], stringEq,
			Act(" ", Skip),
			Act(", ", Skip),
			Act("   ", Skip),
			Act("!", Keep),
		),
	)

	splitTest[[]int, int](t,
		[]int{1, 3, 5, 8, 10, 12},
		expectSlice([][]int{
			// []int{1, 3, 5},
			{8, 10, 12},
		}),
		IsBy(
			ActFn(1, func(in []int) bool {
				return in[0]%2!=0
			}, Skip),
		),
	)

	splitTest[[]int, int](t,
		[]int{1, 3, 5, 6, 7, 8, 10, 12},
		expectSlice([][]int{
			{1, 3, 5},
			{8, 10, 12},
		}),
		Until(false, func(iar []int) bool {
			for i, v := range iar[1:] {
				if iar[i]+2 != v { return false }
			}
			return true
		}, Keep),
	)
}
