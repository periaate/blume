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

func (s Protocol) Is(args ...string) bool  { return gen.Is(args...)(string(s)) }
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

func (s Protocol) Shift(count int) Protocol {
	return Protocol(gen.Shift[string](count)(string(s)))
}
func (s Protocol) Pop(count int) Protocol { return Protocol(gen.Pop[string](count)(string(s))) }
func (s Protocol) Split(pats ...string) gen.Array[Protocol] {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]Protocol, len(split))
	for i, v := range split {
		res[i] = Protocol(v)
	}
	return gen.ToArray(res)
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

func (s Protocol) Green() Protocol { return Protocol(gen.Colorize(gen.Green, string(s))) }
func (s Protocol) LightGreen() Protocol {
	return Protocol(gen.Colorize(gen.LightGreen, string(s)))
}
func (s Protocol) Yellow() Protocol { return Protocol(gen.Colorize(gen.Yellow, string(s))) }
func (s Protocol) LightYellow() Protocol {
	return Protocol(gen.Colorize(gen.LightYellow, string(s)))
}
func (s Protocol) Red() Protocol       { return gen.Colorize(gen.Red, s) }
func (s Protocol) LightRed() Protocol  { return gen.Colorize(gen.LightRed, s) }
func (s Protocol) Blue() Protocol      { return gen.Colorize(gen.Blue, s) }
func (s Protocol) LightBlue() Protocol { return gen.Colorize(gen.LightBlue, s) }
func (s Protocol) Cyan() Protocol      { return gen.Colorize(gen.Cyan, s) }
func (s Protocol) LightCyan() Protocol { return gen.Colorize(gen.LightCyan, s) }
func (s Protocol) Magenta() Protocol   { return gen.Colorize(gen.Magenta, s) }
func (s Protocol) LightMagenta() Protocol {
	return gen.Colorize(gen.LightMagenta, s)
}
func (s Protocol) White() Protocol     { return gen.Colorize(gen.White, s) }
func (s Protocol) Black() Protocol     { return gen.Colorize(gen.Black, s) }
func (s Protocol) Gray() Protocol      { return gen.Colorize(gen.DarkGray, s) }
func (s Protocol) LightGray() Protocol { return gen.Colorize(gen.LightGray, s) }

func (s Protocol) Dim() Protocol  { return gen.Colorize(2, s) }
func (s Protocol) Bold() Protocol { return gen.Bold(s) }
