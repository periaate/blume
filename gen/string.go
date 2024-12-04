package gen

import "github.com/periaate/blume/gen/T"

var (
	_ T.Str              = String("")
	_ T.Or[string]       = String("")
	_ T.Contains[string] = String("")
)

type String string

func (s String) Contains(args ...string) bool        { return Contains(args...)(string(s)) }
func (s String) HasPrefix(args ...string) bool       { return HasPrefix(args...)(string(s)) }
func (s String) HasSuffix(args ...string) bool       { return HasSuffix(args...)(string(s)) }
func (s String) ReplacePrefix(pats ...string) string { return ReplacePrefix(pats...)(string(s)) }
func (s String) ReplaceSuffix(pats ...string) string { return ReplaceSuffix(pats...)(string(s)) }
func (s String) Replace(pats ...string) string       { return Replace(pats...)(string(s)) }
func (s String) ReplaceRegex(pat string, rep string) string {
	return ReplaceRegex[string](pat, rep)(string(s))
}
func (s String) Shift(count int) string        { return Shift[string](count)(string(s)) }
func (s String) Pop(count int) string          { return Pop[string](count)(string(s)) }
func (s String) Split(pats ...string) []string { return SplitWithAll(string(s), false, pats...) }

func (s String) Or(Default string) string {
	if s == "" {
		return Default
	}
	return string(s)
}

func (s String) Len() int       { return len(string(s)) }
func (s String) String() string { return string(s) }
