/*
This shit fucking sucks

# Better abstractions for data and functions

- Reflection is necessary for functions
- Data should be typed, and type safety should be ensured during parsing.
- "Process" based model likely better than "routine" based model
  - communication based off of unix pipes, which are between procecces
*/
package blush

import (
	"fmt"
	"strings"

	"github.com/periaate/blume/clog"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/str"
)

type Node struct {
	gen.Tree[string]
	Value
}

type AST gen.Tree[Node]

func Eval(inp string) (val string, err error) {
	splits := []string{inp}
	clog.Debug("evaluating blush code", "splits", splits)

	for _, split := range splits {
		delims := []string{"(", ")", " "}
		res := str.SplitWithAll(split, true, delims...)
		res = gen.Filter(gen.Isnt(" "))(res)
		ebd := str.EmbedDelims(res, [2]string{"(", ")"})

		ebd = unwrap(ebd, ebd)
		clog.Debug("split parsed", "IDENT", ebd.Value, "ARGS", ebd.Nodes[1:])
		traverse(ebd, 0)
	}

	return
}

func unwrap(sh, up gen.Tree[string]) (res gen.Tree[string]) {
	if len(sh.Nodes) == 0 {
		return up
	}

	return unwrap(sh.Nodes[0], sh)
}

func traverse(sh gen.Tree[string], depth int) {
	r := strings.Repeat(" ", depth)

	if len(sh.Nodes) == 0 {
		fmt.Printf("%s ", sh.Value)
		return
	}

	fmt.Printf("%s\n%s( ", r, r)
	for _, v := range sh.Nodes {
		traverse(v, depth+4)
	}
	fmt.Printf(")\n")

	return
}

type Value struct {
	V    any
	Type string
}

func (v *Value) ToInt() int {
	i, ok := v.V.(int)
	if ok {
		return i
	}

	return 0
}
