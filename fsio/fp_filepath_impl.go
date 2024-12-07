package fsio

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

var (
	_ T.SAr[FilePath]     = gen.SArray[FilePath]{}
	_ T.Str[FilePath]     = FilePath("")
	_ T.Contains[string] = FilePath("")
)

func (s FilePath) Is(args ...string) bool  { return gen.Is(args...)(string(s)) }
func (s FilePath) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s FilePath) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s FilePath) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s FilePath) ReplacePrefix(pats ...string) FilePath {
	return FilePath(gen.ReplacePrefix(pats...)(string(s)))
}

func (s FilePath) ReplaceSuffix(pats ...string) FilePath {
	return FilePath(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s FilePath) Replace(pats ...string) FilePath {
	return FilePath(gen.Replace(pats...)(string(s)))
}

func (s FilePath) ReplaceRegex(pat string, rep string) FilePath {
	return FilePath(gen.ReplaceRegex[string](pat, rep)(string(s)))
}

func (s FilePath) Shift(count int) FilePath {
	return FilePath(gen.Shift[string](count)(string(s)))
}
func (s FilePath) Pop(count int) FilePath { return FilePath(gen.Pop[string](count)(string(s))) }
func (s FilePath) Split(pats ...string) gen.Array[FilePath] {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]FilePath, len(split))
	for i, v := range split {
		res[i] = FilePath(v)
	}
	return gen.ToArray(res)
}

func (s FilePath) Or(Default string) FilePath {
	if s == "" {
		return FilePath(Default)
	}
	return s
}
func (s FilePath) Len() int       { return len(string(s)) }
func (s FilePath) String() string { return string(s) }

func (s FilePath) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s FilePath) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s FilePath) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s FilePath) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s FilePath) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s FilePath) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s FilePath) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s FilePath) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s FilePath) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s FilePath) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s FilePath) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s FilePath) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s FilePath) Colorize(colorCode int) FilePath {
	return FilePath(gen.Colorize(colorCode, string(s)))
}
func (s FilePath) ToUpper() FilePath { return FilePath(strings.ToUpper(string(s))) }
func (s FilePath) ToLower() FilePath { return FilePath(strings.ToLower(string(s))) }
func (s FilePath) Trim() FilePath    { return FilePath(strings.Trim(string(s), " ")) }
func (s FilePath) TrimPrefix(prefix string) FilePath {
	return FilePath(strings.TrimPrefix(string(s), prefix))
}

func (s FilePath) TrimSuffix(suffix string) FilePath {
	return FilePath(strings.TrimSuffix(string(s), suffix))
}
func (s FilePath) TrimSpace() FilePath { return FilePath(strings.TrimSpace(string(s))) }

func (s FilePath) Green() FilePath { return FilePath(gen.Colorize(gen.Green, string(s))) }
func (s FilePath) LightGreen() FilePath {
	return FilePath(gen.Colorize(gen.LightGreen, string(s)))
}
func (s FilePath) Yellow() FilePath { return FilePath(gen.Colorize(gen.Yellow, string(s))) }
func (s FilePath) LightYellow() FilePath {
	return FilePath(gen.Colorize(gen.LightYellow, string(s)))
}
func (s FilePath) Red() FilePath       { return gen.Colorize(gen.Red, s) }
func (s FilePath) LightRed() FilePath  { return gen.Colorize(gen.LightRed, s) }
func (s FilePath) Blue() FilePath      { return gen.Colorize(gen.Blue, s) }
func (s FilePath) LightBlue() FilePath { return gen.Colorize(gen.LightBlue, s) }
func (s FilePath) Cyan() FilePath      { return gen.Colorize(gen.Cyan, s) }
func (s FilePath) LightCyan() FilePath { return gen.Colorize(gen.LightCyan, s) }
func (s FilePath) Magenta() FilePath   { return gen.Colorize(gen.Magenta, s) }
func (s FilePath) LightMagenta() FilePath {
	return gen.Colorize(gen.LightMagenta, s)
}
func (s FilePath) White() FilePath     { return gen.Colorize(gen.White, s) }
func (s FilePath) Black() FilePath     { return gen.Colorize(gen.Black, s) }
func (s FilePath) Gray() FilePath      { return gen.Colorize(gen.DarkGray, s) }
func (s FilePath) LightGray() FilePath { return gen.Colorize(gen.LightGray, s) }

func (s FilePath) Dim() FilePath  { return gen.Colorize(2, s) }
func (s FilePath) Bold() FilePath { return gen.Bold(s) }
