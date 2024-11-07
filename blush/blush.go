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
	"github.com/periaate/blume/clog"
	. "github.com/periaate/blume/gen"
	"github.com/periaate/blume/str"
)

func Eval(inp string) (val string, err error) {
	// splits := str.SplitWithAll(val, false, "|")
	splits := []string{inp}
	clog.Debug("evaluating blush code", "splits", splits)

	for _, split := range splits {
		delims := []string{"(", ")", " "}

		res := str.SplitWithAll(split, true, delims...)

		res = Filter(Isnt(" "))(res)

		ebd, _ := str.EmbedDelims(res, [2]string{"(", ")"})

		clog.Debug("split parsed", "IDENT", ebd.Arr[0].Arr[0].Str, "ARGS", ebd.Arr[0].Arr[1:])
	}

	return
}
