package match

import (
	"fmt"
	"reflect"
	"testing"
)

func stringSize(s1 string) int   { return len(s1) }
func stringEq(s1, s2 string) bool { return s1 == s2 }

func expectSlice[T any](e T) func(i T) (a, b T, ok bool) { return func(i T) (a, b T, ok bool) { return e, i, reflect.DeepEqual(e, i) } }

type comp[T any] func(T) (T, T, bool)

func splitTest(t *testing.T, input string, fn comp[[]string], args ...act[string]) {
	itr, err := ToIter[string, byte](input)
	if err != nil { t.Error(err); return }
	matcher := Is(stringSize, stringEq, args...)
	res := Split(itr, matcher)
	expect, actual, ok := fn(res)
	if !ok { t.Errorf("=== Split Test ===\nInput:\t\"%s\"\nExpect:\t%v\nActual:\t%v\n", input, expect, actual); return }
	fmt.Printf("=== Split Test ===\nInput:\t\"%s\"\nExpect:\t%v\nActual:\t%v\n", input, expect, actual)
}

func TestMatch(t *testing.T) {
	pred := expectSlice([]string{"Hello", "World", "!", "and", "you"})
	splitTest(t, "Hello, World!   and you !", pred, Act(" ", Skip), Act(", ", Skip), Act("   ", Skip), Act("!", Keep))
}
