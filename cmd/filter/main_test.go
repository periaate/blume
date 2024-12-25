package main

import (
	"strings"
	"testing"

	. "github.com/periaate/blume"
)

func TestFilter(t *testing.T) {
	args := []string{
		"a", "b",
		"A", "B",
		"abc", "bca", "cab",
		"ABC", "BCA", "CAB",
		"ababc", "baabc", "abbac", "cabba", "cbaab",
		"ABABC", "BAABC", "ABBAC", "CABBA", "CBAAB",
	}

	exps := []struct {
		pred  func(string) bool
		input []string
	}{
		{Is("a"), []string{"is", "a"}},
		{Contains("ba"), []string{"has", "ll"}},
		{HasPrefix("a"), []string{"pre", "a"}},
		{HasPrefix("ab"), []string{"pre", "ab"}},
	}

	for _, run := range exps {
		t.Run(strings.Join(run.input, " "), func(t *testing.T) {
			res := Parse(run.input)(args)
			if len(res) == 0 {
				t.Fail()
			}
			if !All(run.pred)(res) {
				t.Fail()
			}
		})
	}
}
