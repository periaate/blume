package gen

import (
	"iter"
	"strings"

	"github.com/periaate/blume/gen/T"
)

type SArray[S ~string] struct {
	value []S
}

func (s SArray[S]) Values() []S {
	return s.value
}

func (s SArray[S]) Iter() iter.Seq[S] {
	return func(yield func(S) bool) {
		for _, v := range s.value {
			if !yield(v) {
				break
			}
		}
	}
}

func (s SArray[S]) From(arr []string) T.SAr[S] {
	res := make([]S, len(arr))
	for i, v := range arr {
		res[i] = S(v)
	}
	return SArray[S]{value: res}
}

func (s SArray[S]) Array() []string {
	res := make([]string, len(s.value))
	for i, v := range s.value {
		res[i] = string(v)
	}
	return res
}

func (s SArray[S]) Join(sep string) S {
	return S(strings.Join(s.Array(), sep))
}

var (
	_ T.SAr[String]      = SArray[String]{}
	_ T.Str[String]      = String("")
	_ T.Contains[string] = String("")
)

type String string

func (s String) Contains(args ...string) bool  { return Contains(args...)(string(s)) }
func (s String) HasPrefix(args ...string) bool { return HasPrefix(args...)(string(s)) }
func (s String) HasSuffix(args ...string) bool { return HasSuffix(args...)(string(s)) }
func (s String) ReplacePrefix(pats ...string) String {
	return String(ReplacePrefix(pats...)(string(s)))
}

func (s String) ReplaceSuffix(pats ...string) String {
	return String(ReplaceSuffix(pats...)(string(s)))
}
func (s String) Replace(pats ...string) String { return String(Replace(pats...)(string(s))) }
func (s String) ReplaceRegex(pat string, rep string) String {
	return String(ReplaceRegex[string](pat, rep)(string(s)))
}
func (s String) Shift(count int) String { return String(Shift[string](count)(string(s))) }
func (s String) Pop(count int) String   { return String(Pop[string](count)(string(s))) }
func (s String) Split(pats ...string) []String {
	split := SplitWithAll(string(s), false, pats...)
	res := make([]String, len(split))
	for i, v := range split {
		res[i] = String(v)
	}
	return res
}

func (s String) Or(Default string) String {
	if s == "" {
		return String(Default)
	}
	return s
}
func (s String) Len() int       { return len(string(s)) }
func (s String) String() string { return string(s) }
