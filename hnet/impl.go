package hnet

import (
	. "github.com/periaate/blume/gen"
)

func (s Header) Contains(args ...string) bool        { return Contains(args...)(string(s)) }
func (s Header) HasPrefix(args ...string) bool       { return HasPrefix(args...)(string(s)) }
func (s Header) HasSuffix(args ...string) bool       { return HasSuffix(args...)(string(s)) }
func (s Header) ReplacePrefix(pats ...string) string { return ReplacePrefix(pats...)(string(s)) }
func (s Header) ReplaceSuffix(pats ...string) string { return ReplaceSuffix(pats...)(string(s)) }
func (s Header) Replace(pats ...string) string       { return Replace(pats...)(string(s)) }
func (s Header) ReplaceRegex(pat string, rep string) string {
	return ReplaceRegex[string](pat, rep)(string(s))
}
func (s Header) Shift(count int) string        { return Shift[string](count)(string(s)) }
func (s Header) Pop(count int) string          { return Pop[string](count)(string(s)) }
func (s Header) Split(pats ...string) []string { return SplitWithAll(string(s), false, pats...) }
func (s Header) Or(Default string) string {
	if s == "" {
		return Default
	}
	return string(s)
}
func (s Header) Len() int       { return len(string(s)) }
func (s Header) String() string { return string(s) }

func (s Protocol) Contains(args ...string) bool        { return Contains(args...)(string(s)) }
func (s Protocol) HasPrefix(args ...string) bool       { return HasPrefix(args...)(string(s)) }
func (s Protocol) HasSuffix(args ...string) bool       { return HasSuffix(args...)(string(s)) }
func (s Protocol) ReplacePrefix(pats ...string) string { return ReplacePrefix(pats...)(string(s)) }
func (s Protocol) ReplaceSuffix(pats ...string) string { return ReplaceSuffix(pats...)(string(s)) }
func (s Protocol) Replace(pats ...string) string       { return Replace(pats...)(string(s)) }
func (s Protocol) ReplaceRegex(pat string, rep string) string {
	return ReplaceRegex[string](pat, rep)(string(s))
}
func (s Protocol) Shift(count int) string        { return Shift[string](count)(string(s)) }
func (s Protocol) Pop(count int) string          { return Pop[string](count)(string(s)) }
func (s Protocol) Split(pats ...string) []string { return SplitWithAll(string(s), false, pats...) }
func (s Protocol) Or(Default string) string {
	if s == "" {
		return Default
	}
	return string(s)
}
func (s Protocol) Len() int       { return len(string(s)) }
func (s Protocol) String() string { return string(s) }

func (s URL) Contains(args ...string) bool        { return Contains(args...)(string(s)) }
func (s URL) HasPrefix(args ...string) bool       { return HasPrefix(args...)(string(s)) }
func (s URL) HasSuffix(args ...string) bool       { return HasSuffix(args...)(string(s)) }
func (s URL) ReplacePrefix(pats ...string) string { return ReplacePrefix(pats...)(string(s)) }
func (s URL) ReplaceSuffix(pats ...string) string { return ReplaceSuffix(pats...)(string(s)) }
func (s URL) Replace(pats ...string) string       { return Replace(pats...)(string(s)) }
func (s URL) ReplaceRegex(pat string, rep string) string {
	return ReplaceRegex[string](pat, rep)(string(s))
}
func (s URL) Shift(count int) string        { return Shift[string](count)(string(s)) }
func (s URL) Pop(count int) string          { return Pop[string](count)(string(s)) }
func (s URL) Split(pats ...string) []string { return SplitWithAll(string(s), false, pats...) }
func (s URL) Or(Default string) string {
	if s == "" {
		return Default
	}
	return string(s)
}
func (s URL) Len() int       { return len(string(s)) }
func (s URL) String() string { return string(s) }
