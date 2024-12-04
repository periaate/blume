package hnet

import (
	"github.com/periaate/blume/gen/T"
	"github.com/periaate/blume/gen"
)


func (s URL) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s URL) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s URL) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s URL) ReplacePrefix(pats ...string) URL {
	return URL(gen.ReplacePrefix(pats...)(string(s)))
}

func (s URL) ReplaceSuffix(pats ...string) URL {
	return URL(gen.ReplaceSuffix(pats...)(string(s)))
}
func (s URL) Replace(pats ...string) URL { return URL(gen.Replace(pats...)(string(s))) }
func (s URL) ReplaceRegex(pat string, rep string) URL {
	return URL(gen.ReplaceRegex[string](pat, rep)(string(s)))
}
func (s URL) Shift(count int) URL { return URL(gen.Shift[string](count)(string(s))) }
func (s URL) Pop(count int) URL   { return URL(gen.Pop[string](count)(string(s))) }
func (s URL) Split(pats ...string) T.SAr[URL] {
	return gen.SArray[URL]{}.From(gen.SplitWithAll(string(s), false, pats...))
}

func (s URL) Or(Default string) URL {
	if s == "" {
		return URL(Default)
	}
	return s
}
func (s URL) Len() int       { return len(string(s)) }
func (s URL) string() string { return string(s) }
