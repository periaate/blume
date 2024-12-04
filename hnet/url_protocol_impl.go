package hnet

import (
	"github.com/periaate/blume/gen/T"
	"github.com/periaate/blume/gen"
)


func (s Protocol) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s Protocol) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s Protocol) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s Protocol) ReplacePrefix(pats ...string) Protocol {
	return Protocol(gen.ReplacePrefix(pats...)(string(s)))
}

func (s Protocol) ReplaceSuffix(pats ...string) Protocol {
	return Protocol(gen.ReplaceSuffix(pats...)(string(s)))
}
func (s Protocol) Replace(pats ...string) Protocol { return Protocol(gen.Replace(pats...)(string(s))) }
func (s Protocol) ReplaceRegex(pat string, rep string) Protocol {
	return Protocol(gen.ReplaceRegex[string](pat, rep)(string(s)))
}
func (s Protocol) Shift(count int) Protocol { return Protocol(gen.Shift[string](count)(string(s))) }
func (s Protocol) Pop(count int) Protocol   { return Protocol(gen.Pop[string](count)(string(s))) }
func (s Protocol) Split(pats ...string) T.SAr[Protocol] {
	return gen.SArray[Protocol]{}.From(gen.SplitWithAll(string(s), false, pats...))
}

func (s Protocol) Or(Default string) Protocol {
	if s == "" {
		return Protocol(Default)
	}
	return s
}
func (s Protocol) Len() int       { return len(string(s)) }
func (s Protocol) string() string { return string(s) }
