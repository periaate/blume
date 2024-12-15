package main

import (
	"fmt"

	. "github.com/periaate/blume"
	. "github.com/periaate/blume/fsio"
	"github.com/periaate/blume/is"
)

func main() {
	mapper := Patterns[String](
		Callback(handleIs, "is"),
		Callback(Contains, "has", "contains"),
		Callback(HasPrefix, "pre", "hasprefix"),
		Callback(HasSuffix, "suf", "hassuffix"),
		Callback(MatchRegex, "reg", "regex"),
	)
	inputArgs := IArgs[String](is.NotEmpty[[]string]).Must()
	maps := ToArray(Map(mapper)(inputArgs.Val)).Filter(is.NotNil)
	res := PArgs[string](is.NotEmpty[[]string]).Must().Filter(maps.Val...)
	for _, v := range res.Val {
		fmt.Println(v)
	}
}



func Pat(pats ...string) func(string) bool { return Is(append(Map(func(s string) string { return "!" + s })(pats), pats...)...) }
//
// func Callback(pats ...string) Variadic[string, PredConst[string]] {
// 	return func(fn func(args ...string) Pred[string]) PredConst[string] {
// 		return func(args ...string) Pred[string] {
//
//
// 	return func(fn func(args ...string) func(...string)
// 		return func(cmd string) bool {
// 			cmd = String(cmd).ToLower().String()
// 			if !Pat(pats...)(cmd) { return false }
// 			if HasPrefix("!")(cmd) { return Negate(fn) }
// 			return fn
// 		}
// 	}
// }
//
/*
((...S) => pred[S], ...S) => (S) => (S) => Pred[S]
(Var[S, pred[S]], ...S) => S => S => Pred[S]
*/

func Callback(fn func(...string) func(string) bool, pats ...string) func(string) func(string) func(string) bool {
	pred := Pat(pats...)
	return func(cmd string) func(string) func(string) bool {
		if !pred(cmd) { return nil }
		return func(args string)func(string) bool {
			res := fn(args)
			if HasPrefix("!")(cmd) { res = Negate[string](res) }
			return res
		}
	}
}

func Patterns[S ~string](pairs ...func(string) func(string) func(string) bool) func(s S) (res func(string) bool) {
	var match func(string) func(string) bool
	return func(s S) (res func(string) bool) {
		if match != nil {
			res = match(string(s))
			match = nil
			return
		}
		for _, pair := range pairs {
			m := pair(string(s))
			if m != nil {
				match = m
				return nil
			}
		}
		return nil
	}
}

func handleIs(args ...string) func(string) bool {
	if len(args) == 0 { return nil }
	switch {
		case args[0] == "dir" : return IsDir
		case args[0] == "file": return Negate(IsDir[string])
		default: return Is(args...)
	}
}

func split(str string) []string {
	mapper := Map[String, string](StoS)
	return mapper(String(str).ReplaceRegex(`^\[(.*)\]$`, "$1").Split(",").Map(TrimSpace).Val)
}
