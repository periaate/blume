package main

import (
	"fmt"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/is"
	"github.com/periaate/blume/yap"
)

// `is` keyword is heavily overloaded; consider separating it into separate keywords.

func main() {
	inputArgs := fsio.IArgs[string](is.NotEmpty[[]string]).Must()
	if Any(Is("-d", "--debug"))(inputArgs.Val) {
		yap.SetLevel(yap.L_Debug)
		inputArgs.Val = Filter(Not(Is("-d", "--debug")))(inputArgs.Val)
	}
	res := fsio.PArgs[string](is.NotEmpty[[]string]).Must()
	for _, v := range res.Val {
		fmt.Println(v)
	}
}

func Parse(args []string) func(input []string) []string {
	mapper := Patterns(
		Callback(Is, "is"),
		Callback(Contains, "has", "contains"),
		Callback(HasPrefix, "pre", "hasprefix"),
		Callback(HasSuffix, "suf", "hassuffix"),
		Callback(MatchRegex, "reg", "regex"),
	)

	filters := Filter(func(f func(string) bool) bool { return f != nil })(Map(mapper)(args))

	return func(input []string) []string {
		return Filter(filters...)(input)
	}
}

// func Pat(pats ...string) func(string) bool {}

func Callback(fn func(...string) func(string) bool, pats ...string) func(string) func(string) func(string) bool {
	pred := Is(append(Map(func(s string) string { return "!" + s })(pats), pats...)...)
	return func(cmd string) func(string) func(string) bool {
		if !pred(cmd) {
			return nil
		}
		return func(args string) func(string) bool {
			res := fn(split(args)...)
			if HasPrefix("!")(cmd) {
				yap.Debug("negating", cmd)
				res = Negate(res)
			}
			return res
		}
	}
}

func Patterns(pairs ...func(string) func(string) func(string) bool) func(s string) (res func(string) bool) {
	var match func(string) func(string) bool
	return func(s string) (res func(string) bool) {
		if match != nil {
			res = match(s)
			match = nil
			return
		}
		for _, pair := range pairs {
			m := pair(string(s))
			if m != nil {
				yap.Debug("matched", s)
				match = m
				return nil
			}
		}
		return nil
	}
}

func split(str string) []string {
	mapper := Map[String, string](StoS)
	sar := String(str).ReplaceRegex(`\[(.*)\]`, "$1").Split(",").Map(TrimSpace).Val
	for _, v := range sar {
		yap.Debug("argument", v)
	}
	return mapper(sar)
}
