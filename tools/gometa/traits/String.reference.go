package traits

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

type __String__ string

//blume:trait
var (
	_ T.SAr[__String__]  = gen.SArray[__String__]{}
	_ T.Str[__String__]  = __String__("")
	_ T.Contains[string] = __String__("")
)

func (s __String__) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s __String__) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s __String__) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s __String__) ReplacePrefix(pats ...string) __String__ {
	return __String__(gen.ReplacePrefix(pats...)(string(s)))
}

func (s __String__) ReplaceSuffix(pats ...string) __String__ {
	return __String__(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s __String__) Replace(pats ...string) __String__ {
	return __String__(gen.Replace(pats...)(string(s)))
}

func (s __String__) ReplaceRegex(pat string, rep string) __String__ {
	return __String__(gen.ReplaceRegex[string](pat, rep)(string(s)))
}

func (s __String__) Shift(count int) __String__ {
	return __String__(gen.Shift[string](count)(string(s)))
}
func (s __String__) Pop(count int) __String__ { return __String__(gen.Pop[string](count)(string(s))) }
func (s __String__) Split(pats ...string) []__String__ {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]__String__, len(split))
	for i, v := range split {
		res[i] = __String__(v)
	}
	return res
}

func (s __String__) Or(Default string) __String__ {
	if s == "" {
		return __String__(Default)
	}
	return s
}
func (s __String__) Len() int       { return len(string(s)) }
func (s __String__) String() string { return string(s) }

func (s __String__) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s __String__) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s __String__) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s __String__) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s __String__) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s __String__) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s __String__) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s __String__) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s __String__) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s __String__) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s __String__) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s __String__) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s __String__) Colorize(colorCode int) __String__ {
	return __String__(gen.Colorize(colorCode, string(s)))
}
func (s __String__) ToUpper() __String__ { return __String__(strings.ToUpper(string(s))) }
func (s __String__) ToLower() __String__ { return __String__(strings.ToLower(string(s))) }
func (s __String__) Trim() __String__    { return __String__(strings.Trim(string(s), " ")) }
func (s __String__) TrimPrefix(prefix string) __String__ {
	return __String__(strings.TrimPrefix(string(s), prefix))
}

func (s __String__) TrimSuffix(suffix string) __String__ {
	return __String__(strings.TrimSuffix(string(s), suffix))
}
func (s __String__) TrimSpace() __String__ { return __String__(strings.TrimSpace(string(s))) }

func (s __String__) Green() __String__ { return __String__(gen.Colorize(gen.Green, string(s))) }
func (s __String__) LightGreen() __String__ {
	return __String__(gen.Colorize(gen.LightGreen, string(s)))
}
func (s __String__) Yellow() __String__ { return __String__(gen.Colorize(gen.Yellow, string(s))) }
func (s __String__) LightYellow() __String__ {
	return __String__(gen.Colorize(gen.LightYellow, string(s)))
}
func (s __String__) Red() __String__       { return gen.Colorize(gen.Red, s) }
func (s __String__) LightRed() __String__  { return gen.Colorize(gen.LightRed, s) }
func (s __String__) Blue() __String__      { return gen.Colorize(gen.Blue, s) }
func (s __String__) LightBlue() __String__ { return gen.Colorize(gen.LightBlue, s) }
func (s __String__) Cyan() __String__      { return gen.Colorize(gen.Cyan, s) }
func (s __String__) LightCyan() __String__ { return gen.Colorize(gen.LightCyan, s) }
func (s __String__) Magenta() __String__   { return gen.Colorize(gen.Magenta, s) }
func (s __String__) LightMagenta() __String__ {
	return gen.Colorize(gen.LightMagenta, s)
}
func (s __String__) White() __String__     { return gen.Colorize(gen.White, s) }
func (s __String__) Black() __String__     { return gen.Colorize(gen.Black, s) }
func (s __String__) Gray() __String__      { return gen.Colorize(gen.DarkGray, s) }
func (s __String__) LightGray() __String__ { return gen.Colorize(gen.LightGray, s) }

func (s __String__) Dim() __String__  { return gen.Colorize(2, s) }
func (s __String__) Bold() __String__ { return gen.Bold(s) }
