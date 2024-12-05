package typ

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

var (
	_ T.SAr[String]     = gen.SArray[String]{}
	_ T.Str[String]     = String("")
	_ T.Contains[string] = String("")
)

func (s String) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s String) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s String) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s String) ReplacePrefix(pats ...string) String {
	return String(gen.ReplacePrefix(pats...)(string(s)))
}

func (s String) ReplaceSuffix(pats ...string) String {
	return String(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s String) Replace(pats ...string) String {
	return String(gen.Replace(pats...)(string(s)))
}

func (s String) ReplaceRegex(pat string, rep string) String {
	return String(gen.ReplaceRegex[string](pat, rep)(string(s)))
}
func (s String) Shift(count int) String { return String(gen.Shift[string](count)(string(s))) }
func (s String) Pop(count int) String   { return String(gen.Pop[string](count)(string(s))) }
func (s String) Split(pats ...string) []String {
	split := gen.SplitWithAll(string(s), false, pats...)
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

func (s String) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s String) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s String) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s String) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s String) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s String) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s String) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s String) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s String) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s String) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s String) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s String) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s String) Colorize(colorCode int) String {
	return String(gen.Colorize(colorCode, string(s)))
}
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

func (s String) Green() String      { return String(gen.Colorize(gen.Green, string(s))) }
func (s String) LightGreen() String { return String(gen.Colorize(gen.LightGreen, string(s))) }
func (s String) Yellow() String     { return String(gen.Colorize(gen.Yellow, string(s))) }
func (s String) LightYellow() String {
	return String(gen.Colorize(gen.LightYellow, string(s)))
}
func (s String) Red() String       { return String(gen.Colorize(gen.Red, string(s))) }
func (s String) LightRed() String  { return String(gen.Colorize(gen.LightRed, string(s))) }
func (s String) Blue() String      { return String(gen.Colorize(gen.Blue, string(s))) }
func (s String) LightBlue() String { return String(gen.Colorize(gen.LightBlue, string(s))) }
func (s String) Cyan() String      { return String(gen.Colorize(gen.Cyan, string(s))) }
func (s String) LightCyan() String { return String(gen.Colorize(gen.LightCyan, string(s))) }
func (s String) Magenta() String   { return String(gen.Colorize(gen.Magenta, string(s))) }
func (s String) LightMagenta() String {
	return String(gen.Colorize(gen.LightMagenta, string(s)))
}
func (s String) White() String     { return String(gen.Colorize(gen.White, string(s))) }
func (s String) Black() String     { return String(gen.Colorize(gen.Black, string(s))) }
func (s String) Gray() String      { return String(gen.Colorize(gen.DarkGray, string(s))) }
func (s String) LightGray() String { return String(gen.Colorize(gen.LightGray, string(s))) }
func (s String) Dim() String       { return String(gen.Colorize(2, string(s))) }
