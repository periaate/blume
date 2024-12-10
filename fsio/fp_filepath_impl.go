package fsio

import (
	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/gen"
)

func (s FilePath) Is(args ...string) bool  { return Is(args...)(string(s)) }
func (s FilePath) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s FilePath) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s FilePath) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s FilePath) ReplacePrefix(pats ...string) FilePath {
	return FilePath(gen.ReplacePrefix(pats...)(string(s)))
}

func (s FilePath) ReplaceSuffix(pats ...string) FilePath {
	return FilePath(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s FilePath) Replace(pats ...string) FilePath { return FilePath(gen.Replace(pats...)(string(s))) }

func (s FilePath) ReplaceRegex(pat string, rep string) FilePath {
	return FilePath(gen.ReplaceRegex[string](pat, rep)(string(s)))
}

func (s FilePath) Split(pats ...string) Array[FilePath] {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]FilePath, len(split))
	for i, v := range split {
		res[i] = FilePath(v)
	}
	return ToArray(res)
}

func (s FilePath) Or(Default string) FilePath {
	if s == "" { return FilePath(Default) }
	return s
}
func (s FilePath) Len() int       { return len(string(s)) }
func (s FilePath) String() string { return string(s) }

// func (s FilePath) ToInt() Result[int]         { return ToInt(string(s)) }
// func (s FilePath) ToInt8() Result[int8]       { return ToInt8(string(s)) }
// func (s FilePath) ToInt16() Result[int16]     { return ToInt16(string(s)) }
// func (s FilePath) ToInt32() Result[int32]     { return ToInt32(string(s)) }
// func (s FilePath) ToInt64() Result[int64]     { return ToInt64(string(s)) }
// func (s FilePath) ToUint() Result[uint]       { return ToUint(string(s)) }
// func (s FilePath) ToUint8() Result[uint8]     { return ToUint8(string(s)) }
// func (s FilePath) ToUint16() Result[uint16]   { return ToUint16(string(s)) }
// func (s FilePath) ToUint32() Result[uint32]   { return ToUint32(string(s)) }
// func (s FilePath) ToUint64() Result[uint64]   { return ToUint64(string(s)) }
// func (s FilePath) ToFloat32() Result[float32] { return ToFloat32(string(s)) }
// func (s FilePath) ToFloat64() Result[float64] { return ToFloat64(string(s)) }
//
// func (s FilePath) Colorize(colorCode int) FilePath {
// 	return FilePath(Colorize(colorCode, string(s)))
// }
// func (s FilePath) ToUpper() FilePath { return FilePath(strings.ToUpper(string(s))) }
// func (s FilePath) ToLower() FilePath { return FilePath(strings.ToLower(string(s))) }
// func (s FilePath) Trim() FilePath    { return FilePath(strings.Trim(string(s), " ")) }
// func (s FilePath) TrimPrefix(prefix string) FilePath {
// 	return FilePath(strings.TrimPrefix(string(s), prefix))
// }
//
// func (s FilePath) TrimSuffix(suffix string) FilePath {
// 	return FilePath(strings.TrimSuffix(string(s), suffix))
// }
// func (s FilePath) TrimSpace() FilePath { return FilePath(strings.TrimSpace(string(s))) }
//
// func (s FilePath) Green() FilePath { return FilePath(Colorize(Green, string(s))) }
// func (s FilePath) LightGreen() FilePath {
// 	return FilePath(Colorize(LightGreen, string(s)))
// }
// func (s FilePath) Yellow() FilePath { return FilePath(Colorize(Yellow, string(s))) }
// func (s FilePath) LightYellow() FilePath {
// 	return FilePath(Colorize(LightYellow, string(s)))
// }
// func (s FilePath) Red() FilePath       { return Colorize(Red, s) }
// func (s FilePath) LightRed() FilePath  { return Colorize(LightRed, s) }
// func (s FilePath) Blue() FilePath      { return Colorize(Blue, s) }
// func (s FilePath) LightBlue() FilePath { return Colorize(LightBlue, s) }
// func (s FilePath) Cyan() FilePath      { return Colorize(Cyan, s) }
// func (s FilePath) LightCyan() FilePath { return Colorize(LightCyan, s) }
// func (s FilePath) Magenta() FilePath   { return Colorize(Magenta, s) }
// func (s FilePath) LightMagenta() FilePath {
// 	return Colorize(LightMagenta, s)
// }
// func (s FilePath) White() FilePath     { return Colorize(White, s) }
// func (s FilePath) Black() FilePath     { return Colorize(Black, s) }
// func (s FilePath) Gray() FilePath      { return Colorize(DarkGray, s) }
// func (s FilePath) LightGray() FilePath { return Colorize(LightGray, s) }
//
// func (s FilePath) Dim() FilePath  { return Colorize(2, s) }
// func (s FilePath) Bold() FilePath { return Bold(s) }
