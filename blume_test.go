package blume

import (
	"fmt"
	"testing"
)

func Test_blume(t *testing.T) {
	val := Arr("hello", "world", "aab", "aba", "baa", "bab").
		Each(Logs).
		Map(ReplacePrefix("a", "ba"))
	fmt.Println()
	val = val.Each(Logs).
		Filter(HasSuffix("b"))
	fmt.Println()
	val.Each(Logs)
}
