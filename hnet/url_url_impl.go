package hnet

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

var (
	_ T.SAr[URL]     = gen.SArray[URL]{}
	_ T.Str[URL]     = URL("")
	_ T.Contains[string] = URL("")
)

func (s URL) Is(args ...string) bool  { return gen.Is(args...)(string(s)) }
func (s URL) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s URL) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s URL) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s URL) ReplacePrefix(pats ...string) URL {
	return URL(gen.ReplacePrefix(pats...)(string(s)))
}

func (s URL) ReplaceSuffix(pats ...string) URL {
	return URL(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s URL) Replace(pats ...string) URL {
	return URL(gen.Replace(pats...)(string(s)))
}

func (s URL) ReplaceRegex(pat string, rep string) URL {
	return URL(gen.ReplaceRegex[string](pat, rep)(string(s)))
}

func (s URL) Shift(count int) URL {
	return URL(gen.Shift[string](count)(string(s)))
}
func (s URL) Pop(count int) URL { return URL(gen.Pop[string](count)(string(s))) }
func (s URL) Split(pats ...string) gen.Array[URL] {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]URL, len(split))
	for i, v := range split {
		res[i] = URL(v)
	}
	return gen.ToArray(res)
}

func (s URL) Or(Default string) URL {
	if s == "" {
		return URL(Default)
	}
	return s
}
func (s URL) Len() int       { return len(string(s)) }
func (s URL) String() string { return string(s) }

func (s URL) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s URL) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s URL) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s URL) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s URL) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s URL) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s URL) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s URL) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s URL) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s URL) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s URL) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s URL) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s URL) Colorize(colorCode int) URL {
	return URL(gen.Colorize(colorCode, string(s)))
}
func (s URL) ToUpper() URL { return URL(strings.ToUpper(string(s))) }
func (s URL) ToLower() URL { return URL(strings.ToLower(string(s))) }
func (s URL) Trim() URL    { return URL(strings.Trim(string(s), " ")) }
func (s URL) TrimPrefix(prefix string) URL {
	return URL(strings.TrimPrefix(string(s), prefix))
}

func (s URL) TrimSuffix(suffix string) URL {
	return URL(strings.TrimSuffix(string(s), suffix))
}
func (s URL) TrimSpace() URL { return URL(strings.TrimSpace(string(s))) }

func (s URL) Green() URL { return URL(gen.Colorize(gen.Green, string(s))) }
func (s URL) LightGreen() URL {
	return URL(gen.Colorize(gen.LightGreen, string(s)))
}
func (s URL) Yellow() URL { return URL(gen.Colorize(gen.Yellow, string(s))) }
func (s URL) LightYellow() URL {
	return URL(gen.Colorize(gen.LightYellow, string(s)))
}
func (s URL) Red() URL       { return gen.Colorize(gen.Red, s) }
func (s URL) LightRed() URL  { return gen.Colorize(gen.LightRed, s) }
func (s URL) Blue() URL      { return gen.Colorize(gen.Blue, s) }
func (s URL) LightBlue() URL { return gen.Colorize(gen.LightBlue, s) }
func (s URL) Cyan() URL      { return gen.Colorize(gen.Cyan, s) }
func (s URL) LightCyan() URL { return gen.Colorize(gen.LightCyan, s) }
func (s URL) Magenta() URL   { return gen.Colorize(gen.Magenta, s) }
func (s URL) LightMagenta() URL {
	return gen.Colorize(gen.LightMagenta, s)
}
func (s URL) White() URL     { return gen.Colorize(gen.White, s) }
func (s URL) Black() URL     { return gen.Colorize(gen.Black, s) }
func (s URL) Gray() URL      { return gen.Colorize(gen.DarkGray, s) }
func (s URL) LightGray() URL { return gen.Colorize(gen.LightGray, s) }

func (s URL) Dim() URL  { return gen.Colorize(2, s) }
func (s URL) Bold() URL { return gen.Bold(s) }
