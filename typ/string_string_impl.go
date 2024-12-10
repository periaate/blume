package typ

import (
	"strings"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/gen"
)


func (s String) Is(args ...string) bool  { return Is(args...)(string(s)) }
func (s String) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s String) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s String) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s String) ReplacePrefix(pats ...string) String {
	return String(gen.ReplacePrefix(pats...)(string(s)))
}

func (s String) ReplaceSuffix(pats ...string) String {
	return String(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s String) Replace(pats ...string) String { return String(gen.Replace(pats...)(string(s))) }

func (s String) ReplaceRegex(pat string, rep string) String {
	return String(gen.ReplaceRegex[string](pat, rep)(string(s)))
}

func (s String) Shift(count int) String { return String(gen.Shift[string](count)(string(s))) }
func (s String) Pop(count int) String { return String(gen.Pop[string](count)(string(s))) }
func (s String) Split(pats ...string) Array[String] {
	split := gen.SplitWithAll(string(s), false, pats...)
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

func (s String) Colorize(colorCode int) String { return String(gen.Colorize(colorCode, string(s))) }
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

func (s String) Green() String { return String(gen.Colorize(gen.Green, string(s))) }
func (s String) LightGreen() String { return String(gen.Colorize(gen.LightGreen, string(s))) }
func (s String) Yellow() String { return String(gen.Colorize(gen.Yellow, string(s))) }
func (s String) LightYellow() String { return String(gen.Colorize(gen.LightYellow, string(s))) }
func (s String) Red() String       { return gen.Colorize(gen.Red, s) }
func (s String) LightRed() String  { return gen.Colorize(gen.LightRed, s) }
func (s String) Blue() String      { return gen.Colorize(gen.Blue, s) }
func (s String) LightBlue() String { return gen.Colorize(gen.LightBlue, s) }
func (s String) Cyan() String      { return gen.Colorize(gen.Cyan, s) }
func (s String) LightCyan() String { return gen.Colorize(gen.LightCyan, s) }
func (s String) Magenta() String   { return gen.Colorize(gen.Magenta, s) }
func (s String) LightMagenta() String { return gen.Colorize(gen.LightMagenta, s) }
func (s String) White() String     { return gen.Colorize(gen.White, s) }
func (s String) Black() String     { return gen.Colorize(gen.Black, s) }
func (s String) Gray() String      { return gen.Colorize(gen.DarkGray, s) }
func (s String) LightGray() String { return gen.Colorize(gen.LightGray, s) }

func (s String) Dim() String  { return gen.Colorize(2, s) }
func (s String) Bold() String { return gen.Bold(s) }
