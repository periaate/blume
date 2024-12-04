package traits

import (
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

type Reference string

func (s Reference) Contains(args ...string) bool  { return gen.Contains(args...)(string(s)) }
func (s Reference) HasPrefix(args ...string) bool { return gen.HasPrefix(args...)(string(s)) }
func (s Reference) HasSuffix(args ...string) bool { return gen.HasSuffix(args...)(string(s)) }
func (s Reference) ReplacePrefix(pats ...string) Reference {
	return Reference(gen.ReplacePrefix(pats...)(string(s)))
}

func (s Reference) ReplaceSuffix(pats ...string) Reference {
	return Reference(gen.ReplaceSuffix(pats...)(string(s)))
}

func (s Reference) Replace(pats ...string) Reference {
	return Reference(gen.Replace(pats...)(string(s)))
}

func (s Reference) ReplaceRegex(pat string, rep string) Reference {
	return Reference(gen.ReplaceRegex[string](pat, rep)(string(s)))
}
func (s Reference) Shift(count int) Reference { return Reference(gen.Shift[string](count)(string(s))) }
func (s Reference) Pop(count int) Reference   { return Reference(gen.Pop[string](count)(string(s))) }
func (s Reference) Split(pats ...string) []Reference {
	split := gen.SplitWithAll(string(s), false, pats...)
	res := make([]Reference, len(split))
	for i, v := range split {
		res[i] = Reference(v)
	}
	return res
}

func (s Reference) Or(Default string) Reference {
	if s == "" {
		return Reference(Default)
	}
	return s
}
func (s Reference) Len() int          { return len(string(s)) }
func (s Reference) Reference() string { return string(s) }

func (s Reference) ToInt() T.Result[int]         { return gen.ToInt(string(s)) }
func (s Reference) ToInt8() T.Result[int8]       { return gen.ToInt8(string(s)) }
func (s Reference) ToInt16() T.Result[int16]     { return gen.ToInt16(string(s)) }
func (s Reference) ToInt32() T.Result[int32]     { return gen.ToInt32(string(s)) }
func (s Reference) ToInt64() T.Result[int64]     { return gen.ToInt64(string(s)) }
func (s Reference) ToUint() T.Result[uint]       { return gen.ToUint(string(s)) }
func (s Reference) ToUint8() T.Result[uint8]     { return gen.ToUint8(string(s)) }
func (s Reference) ToUint16() T.Result[uint16]   { return gen.ToUint16(string(s)) }
func (s Reference) ToUint32() T.Result[uint32]   { return gen.ToUint32(string(s)) }
func (s Reference) ToUint64() T.Result[uint64]   { return gen.ToUint64(string(s)) }
func (s Reference) ToFloat32() T.Result[float32] { return gen.ToFloat32(string(s)) }
func (s Reference) ToFloat64() T.Result[float64] { return gen.ToFloat64(string(s)) }

func (s Reference) Colorize(colorCode int) Reference {
	return Reference(Colorize(colorCode, string(s)))
}
func (s Reference) ToUpper() Reference { return Reference(strings.ToUpper(string(s))) }
func (s Reference) ToLower() Reference { return Reference(strings.ToLower(string(s))) }
func (s Reference) Trim() Reference    { return Reference(strings.Trim(string(s), " ")) }
func (s Reference) TrimPrefix(prefix string) Reference {
	return Reference(strings.TrimPrefix(string(s), prefix))
}

func (s Reference) TrimSuffix(suffix string) Reference {
	return Reference(strings.TrimSuffix(string(s), suffix))
}
func (s Reference) TrimSpace() Reference { return Reference(strings.TrimSpace(string(s))) }
