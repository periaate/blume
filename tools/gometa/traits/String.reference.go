package traits

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

type __Ref__ string

var (
	_ T.SAr[__Ref__]     = gen.SArray[__Ref__]{}
	_ T.Str[__Ref__]     = __Ref__("")
	_ T.Contains[string] = __Ref__("")
)

func (s __Ref__) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s __Ref__) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s __Ref__) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s __Ref__) ReplacePrefix(pats ...string) __Ref__ {
	return __Ref__(gen.ReplacePrefix(pats...)(string(s)))
}

func (s __Ref__) ReplaceSuffix(pats ...string) __Ref__ {
	return __Ref__(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s __Ref__) Replace(pats ...string) __Ref__ {
	return __Ref__(gen.Replace(pats...)(string(s)))
}

func (s __Ref__) ReplaceRegex(pat string, rep string) __Ref__ {
	return __Ref__(gen.ReplaceRegex[string](pat, rep)(string(s)))
}
func (s __Ref__) Shift(count int) __Ref__ { return __Ref__(gen.Shift[string](count)(string(s))) }
func (s __Ref__) Pop(count int) __Ref__   { return __Ref__(gen.Pop[string](count)(string(s))) }
func (s __Ref__) Split(pats ...string) []__Ref__ {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]__Ref__, len(split))
	for i, v := range split {
		res[i] = __Ref__(v)
	}
	return res
}

func (s __Ref__) Or(Default string) __Ref__ {
	if s == "" {
		return __Ref__(Default)
	}
	return s
}
func (s __Ref__) Len() int       { return len(string(s)) }
func (s __Ref__) String() string { return string(s) }

func (s __Ref__) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s __Ref__) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s __Ref__) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s __Ref__) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s __Ref__) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s __Ref__) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s __Ref__) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s __Ref__) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s __Ref__) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s __Ref__) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s __Ref__) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s __Ref__) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s __Ref__) Colorize(colorCode int) __Ref__ {
	return __Ref__(gen.Colorize(colorCode, string(s)))
}
func (s __Ref__) ToUpper() __Ref__ { return __Ref__(strings.ToUpper(string(s))) }
func (s __Ref__) ToLower() __Ref__ { return __Ref__(strings.ToLower(string(s))) }
func (s __Ref__) Trim() __Ref__    { return __Ref__(strings.Trim(string(s), " ")) }
func (s __Ref__) TrimPrefix(prefix string) __Ref__ {
	return __Ref__(strings.TrimPrefix(string(s), prefix))
}

func (s __Ref__) TrimSuffix(suffix string) __Ref__ {
	return __Ref__(strings.TrimSuffix(string(s), suffix))
}
func (s __Ref__) TrimSpace() __Ref__ { return __Ref__(strings.TrimSpace(string(s))) }

func (s __Ref__) Green() __Ref__      { return __Ref__(gen.Colorize(gen.Green, string(s))) }
func (s __Ref__) LightGreen() __Ref__ { return __Ref__(gen.Colorize(gen.LightGreen, string(s))) }
func (s __Ref__) Yellow() __Ref__     { return __Ref__(gen.Colorize(gen.Yellow, string(s))) }
func (s __Ref__) LightYellow() __Ref__ {
	return __Ref__(gen.Colorize(gen.LightYellow, string(s)))
}
func (s __Ref__) Red() __Ref__       { return __Ref__(gen.Colorize(gen.Red, string(s))) }
func (s __Ref__) LightRed() __Ref__  { return __Ref__(gen.Colorize(gen.LightRed, string(s))) }
func (s __Ref__) Blue() __Ref__      { return __Ref__(gen.Colorize(gen.Blue, string(s))) }
func (s __Ref__) LightBlue() __Ref__ { return __Ref__(gen.Colorize(gen.LightBlue, string(s))) }
func (s __Ref__) Cyan() __Ref__      { return __Ref__(gen.Colorize(gen.Cyan, string(s))) }
func (s __Ref__) LightCyan() __Ref__ { return __Ref__(gen.Colorize(gen.LightCyan, string(s))) }
func (s __Ref__) Magenta() __Ref__   { return __Ref__(gen.Colorize(gen.Magenta, string(s))) }
func (s __Ref__) LightMagenta() __Ref__ {
	return __Ref__(gen.Colorize(gen.LightMagenta, string(s)))
}
func (s __Ref__) White() __Ref__     { return __Ref__(gen.Colorize(gen.White, string(s))) }
func (s __Ref__) Black() __Ref__     { return __Ref__(gen.Colorize(gen.Black, string(s))) }
func (s __Ref__) Gray() __Ref__      { return __Ref__(gen.Colorize(gen.DarkGray, string(s))) }
func (s __Ref__) LightGray() __Ref__ { return __Ref__(gen.Colorize(gen.LightGray, string(s))) }
func (s __Ref__) Dim() __Ref__       { return __Ref__(gen.Colorize(2, string(s))) }
