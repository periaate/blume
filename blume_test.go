package blume

import (
	"testing"

	"github.com/periaate/blume/test"
)

func TestMap(tt *testing.T) {
	v := []int{1, 2, 3}
	t := test.Expect(tt, v)
	t.Is(Map[string, int](func(s string) int { return len(s) })([]string{"a", "ab", "abc"}))
	t.Is(Map[string, int](func(s ...any) int { return len(s[0].(string)) })([]string{"a", "ab", "abc"}))
}
