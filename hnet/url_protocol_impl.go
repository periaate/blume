package hnet

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

var (
	_ T.SAr[Protocol]     = gen.SArray[Protocol]{}
	_ T.Str[Protocol]     = Protocol("")
	_ T.Contains[string] = Protocol("")
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

func (s Protocol) Replace(pats ...string) Protocol {
	return Protocol(gen.Replace(pats...)(string(s)))
}

func (s Protocol) ReplaceRegex(pat string, rep string) Protocol {
	return Protocol(gen.ReplaceRegex[string](pat, rep)(string(s)))
}
func (s Protocol) Shift(count int) Protocol { return Protocol(gen.Shift[string](count)(string(s))) }
func (s Protocol) Pop(count int) Protocol   { return Protocol(gen.Pop[string](count)(string(s))) }
func (s Protocol) Split(pats ...string) []Protocol {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]Protocol, len(split))
	for i, v := range split {
		res[i] = Protocol(v)
	}
	return res
}

func (s Protocol) Or(Default string) Protocol {
	if s == "" {
		return Protocol(Default)
	}
	return s
}
func (s Protocol) Len() int       { return len(string(s)) }
func (s Protocol) String() string { return string(s) }

func (s Protocol) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s Protocol) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s Protocol) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s Protocol) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s Protocol) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s Protocol) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s Protocol) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s Protocol) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s Protocol) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s Protocol) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s Protocol) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s Protocol) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s Protocol) Colorize(colorCode int) Protocol {
	return Protocol(gen.Colorize(colorCode, string(s)))
}
func (s Protocol) ToUpper() Protocol { return Protocol(strings.ToUpper(string(s))) }
func (s Protocol) ToLower() Protocol { return Protocol(strings.ToLower(string(s))) }
func (s Protocol) Trim() Protocol    { return Protocol(strings.Trim(string(s), " ")) }
func (s Protocol) TrimPrefix(prefix string) Protocol {
	return Protocol(strings.TrimPrefix(string(s), prefix))
}

func (s Protocol) TrimSuffix(suffix string) Protocol {
	return Protocol(strings.TrimSuffix(string(s), suffix))
}
func (s Protocol) TrimSpace() Protocol { return Protocol(strings.TrimSpace(string(s))) }

func (s Protocol) Green() Protocol      { return Protocol(gen.Colorize(gen.Green, string(s))) }
func (s Protocol) LightGreen() Protocol { return Protocol(gen.Colorize(gen.LightGreen, string(s))) }
func (s Protocol) Yellow() Protocol     { return Protocol(gen.Colorize(gen.Yellow, string(s))) }
func (s Protocol) LightYellow() Protocol {
	return Protocol(gen.Colorize(gen.LightYellow, string(s)))
}
func (s Protocol) Red() Protocol       { return Protocol(gen.Colorize(gen.Red, string(s))) }
func (s Protocol) LightRed() Protocol  { return Protocol(gen.Colorize(gen.LightRed, string(s))) }
func (s Protocol) Blue() Protocol      { return Protocol(gen.Colorize(gen.Blue, string(s))) }
func (s Protocol) LightBlue() Protocol { return Protocol(gen.Colorize(gen.LightBlue, string(s))) }
func (s Protocol) Cyan() Protocol      { return Protocol(gen.Colorize(gen.Cyan, string(s))) }
func (s Protocol) LightCyan() Protocol { return Protocol(gen.Colorize(gen.LightCyan, string(s))) }
func (s Protocol) Magenta() Protocol   { return Protocol(gen.Colorize(gen.Magenta, string(s))) }
func (s Protocol) LightMagenta() Protocol {
	return Protocol(gen.Colorize(gen.LightMagenta, string(s)))
}
func (s Protocol) White() Protocol     { return Protocol(gen.Colorize(gen.White, string(s))) }
func (s Protocol) Black() Protocol     { return Protocol(gen.Colorize(gen.Black, string(s))) }
func (s Protocol) Gray() Protocol      { return Protocol(gen.Colorize(gen.DarkGray, string(s))) }
func (s Protocol) LightGray() Protocol { return Protocol(gen.Colorize(gen.LightGray, string(s))) }
func (s Protocol) Dim() Protocol       { return Protocol(gen.Colorize(2, string(s))) }
