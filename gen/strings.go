package gen

import (
	"strings"

	. "github.com/periaate/blume/core"
)

var _ = Zero[any]

type String string


func (s String) Is(args ...string) bool  { return Is(args...)(string(s)) }
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
func (s String) Pop(count int) String { return String(Pop[string](count)(string(s))) }
func (s String) Split(pats ...string) Array[String] {
	split := SplitWithAll(string(s), false, pats...)
	res := make([]String, len(split))
	for i, v := range split {
		res[i] = String(v)
	}
	return ToArray(res)
}

func (s String) Or(Default string) String {
	if s == "" { return String(Default) }
	return s
}
func (s String) Len() int       { return len(string(s)) }
func (s String) String() string { return string(s) }

func (s String) Colorize(colorCode int) String { return String(Colorize(colorCode, string(s))) }
func (s String) ToUpper() String { return String(strings.ToUpper(string(s))) }
func (s String) ToLower() String { return String(strings.ToLower(string(s))) }
func (s String) Trim() String    { return String(strings.Trim(string(s), " ")) }
func (s String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(string(s), prefix))
}

func (s String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(string(s), suffix))
}
func (s String) TrimSpace() String { return String(strings.TrimSpace(string(s))) }

func (s String) Green() String { return String(Colorize(Green, string(s))) }
func (s String) LightGreen() String { return String(Colorize(LightGreen, string(s))) }
func (s String) Yellow() String { return String(Colorize(Yellow, string(s))) }
func (s String) LightYellow() String { return String(Colorize(LightYellow, string(s))) }
func (s String) Red() String       { return Colorize(Red, s) }
func (s String) LightRed() String  { return Colorize(LightRed, s) }
func (s String) Blue() String      { return Colorize(Blue, s) }
func (s String) LightBlue() String { return Colorize(LightBlue, s) }
func (s String) Cyan() String      { return Colorize(Cyan, s) }
func (s String) LightCyan() String { return Colorize(LightCyan, s) }
func (s String) Magenta() String   { return Colorize(Magenta, s) }
func (s String) LightMagenta() String { return Colorize(LightMagenta, s) }
func (s String) White() String     { return Colorize(White, s) }
func (s String) Black() String     { return Colorize(Black, s) }
func (s String) Gray() String      { return Colorize(DarkGray, s) }
func (s String) LightGray() String { return Colorize(LightGray, s) }

func (s String) Dim() String  { return Colorize(2, s) }
func (s String) Bold() String { return Bold(s) }
