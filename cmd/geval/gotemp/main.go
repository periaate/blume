package main

import (
	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
)

var _ String

func main() {
	println(len(ToArray(Map(func(s String) int {
		r, err := fsio.ReadDir(s.String())
		if err == nil {
			return len(r)
		}
		return 0
	})(Input[string]("piped").Value)).Filter(func(i int) bool { return i >= 1 }).Value))
}
