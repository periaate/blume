package gen

import (
	"iter"
	"strconv"
	"strings"

	"github.com/periaate/blume/gen/T"
)

type SArray[S ~string] struct {
	value []S
}

func (s SArray[S]) Values() []S {
	return s.value
}

func (s SArray[S]) Iter() iter.Seq[S] {
	return func(yield func(S) bool) {
		for _, v := range s.value {
			if !yield(v) {
				break
			}
		}
	}
}

func (s SArray[S]) From(arr []string) T.SAr[S] {
	res := make([]S, len(arr))
	for i, v := range arr {
		res[i] = S(v)
	}
	return SArray[S]{value: res}
}

func (s SArray[S]) Array() []string {
	res := make([]string, len(s.value))
	for i, v := range s.value {
		res[i] = string(v)
	}
	return res
}

func (s SArray[S]) Join(sep string) S {
	return S(strings.Join(s.Array(), sep))
}

var (
	_ T.SAr[String]      = SArray[String]{}
	_ T.Str[String]      = String("")
	_ T.Contains[string] = String("")
)

type String string

func (s String) Contains(args ...string) bool  { return Contains(args...)(string(s)) }
func (s String) HasPrefix(args ...string) bool { return HasPrefix(args...)(string(s)) }
func (s String) HasSuffix(args ...string) bool { return HasSuffix(args...)(string(s)) }
func (s String) ReplacePrefix(pats ...string) String {
	return String(ReplacePrefix(pats...)(string(s)))
}

func (s String) ReplaceSuffix(pats ...string) String {
	return String(ReplaceSuffix(pats...)(string(s)))
}
func (s String) Replace(pats ...string) String { return String(Replace(pats...)(string(s))) }
func (s String) ReplaceRegex(pat string, rep string) String {
	return String(ReplaceRegex[string](pat, rep)(string(s)))
}
func (s String) Shift(count int) String { return String(Shift[string](count)(string(s))) }
func (s String) Pop(count int) String   { return String(Pop[string](count)(string(s))) }
func (s String) Split(pats ...string) []String {
	split := SplitWithAll(string(s), false, pats...)
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

func (s String) ToInt() T.Result[int]         { return ToInt(string(s)) }
func (s String) ToInt8() T.Result[int8]       { return ToInt8(string(s)) }
func (s String) ToInt16() T.Result[int16]     { return ToInt16(string(s)) }
func (s String) ToInt32() T.Result[int32]     { return ToInt32(string(s)) }
func (s String) ToInt64() T.Result[int64]     { return ToInt64(string(s)) }
func (s String) ToUint() T.Result[uint]       { return ToUint(string(s)) }
func (s String) ToUint8() T.Result[uint8]     { return ToUint8(string(s)) }
func (s String) ToUint16() T.Result[uint16]   { return ToUint16(string(s)) }
func (s String) ToUint32() T.Result[uint32]   { return ToUint32(string(s)) }
func (s String) ToUint64() T.Result[uint64]   { return ToUint64(string(s)) }
func (s String) ToFloat32() T.Result[float32] { return ToFloat32(string(s)) }
func (s String) ToFloat64() T.Result[float64] { return ToFloat64(string(s)) }

func (s String) Colorize(colorCode int) String { return String(Colorize(colorCode, string(s))) }
func (s String) ToUpper() String               { return String(strings.ToUpper(string(s))) }
func (s String) ToLower() String               { return String(strings.ToLower(string(s))) }
func (s String) Trim() String                  { return String(strings.Trim(string(s), " ")) }
func (s String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(string(s), prefix))
}

func (s String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(string(s), suffix))
}
func (s String) TrimSpace() String { return String(strings.TrimSpace(string(s))) }

func ToInt(s string) T.Result[int] { return T.Results(strconv.Atoi(s)) }
func ToInt8(s string) T.Result[int8] {
	i, err := strconv.ParseInt(s, 10, 8)
	return T.Results(int8(i), err)
}

func ToInt16(s string) T.Result[int16] {
	i, err := strconv.ParseInt(s, 10, 16)
	return T.Results(int16(i), err)
}

func ToInt32(s string) T.Result[int32] {
	i, err := strconv.ParseInt(s, 10, 32)
	return T.Results(int32(i), err)
}

func ToInt64(s string) T.Result[int64] {
	i, err := strconv.ParseInt(s, 10, 64)
	return T.Results(int64(i), err)
}

func ToUint(s string) T.Result[uint] {
	i, err := strconv.ParseUint(s, 10, 0)
	return T.Results(uint(i), err)
}

func ToUint8(s string) T.Result[uint8] {
	i, err := strconv.ParseUint(s, 10, 8)
	return T.Results(uint8(i), err)
}

func ToUint16(s string) T.Result[uint16] {
	i, err := strconv.ParseUint(s, 10, 16)
	return T.Results(uint16(i), err)
}

func ToUint32(s string) T.Result[uint32] {
	i, err := strconv.ParseUint(s, 10, 32)
	return T.Results(uint32(i), err)
}

func ToUint64(s string) T.Result[uint64] {
	i, err := strconv.ParseUint(s, 10, 64)
	return T.Results(uint64(i), err)
}

func ToFloat32(s string) T.Result[float32] {
	i, err := strconv.ParseFloat(s, 32)
	return T.Results(float32(i), err)
}

func ToFloat64(s string) T.Result[float64] {
	i, err := strconv.ParseFloat(s, 64)
	return T.Results(float64(i), err)
}

func Colorize(colorCode int, s string) string {
	return "\033[" + strconv.Itoa(colorCode) + "m" + s + "\033[0m"
}

func ToUpper(s string) string { return strings.ToUpper(s) }
func ToLower(s string) string { return strings.ToLower(s) }

func Trim(s string) string { return strings.Trim(s, " ") }
func TrimPrefix(prefix string, s string) string {
	return strings.TrimPrefix(s, prefix)
}

func TrimSuffix(suffix string, s string) string {
	return strings.TrimSuffix(s, suffix)
}
func TrimSpace(s string) string { return strings.TrimSpace(s) }
