package blume

import (
	"strconv"
	"strings"
)

type String string

func (s String) Has(args ...Selector[string]) bool   { return Has(args...)(s.String()) }
func (s String) Del(args ...Selector[string]) String { return String(Del(args...)(s.String())) }
func (s String) Rep(args ...any) String              { return String(Rep[string](args...)(string(s))) }

func (s String) Is(args ...string) bool       { return Is(args...)(string(s)) }
func (s String) Contains(args ...string) bool { return Contains(args...)(string(s)) }

// HasPrefix
// Deprecated: Use [Has] with [Pre] instead.
func (s String) HasPrefix(args ...string) bool { return HasPrefix(args...)(string(s)) }

// HasPrefix
// Deprecated: Use [Has] with [Suf] instead.
func (s String) HasSuffix(args ...string) bool { return HasSuffix(args...)(string(s)) }

// ReplacePrefix
// Deprecated: Use [Rep] with [Pre] instead.
func (s String) ReplacePrefix(pats ...string) String {
	return String(ReplacePrefix(pats...)(string(s)))
}

// ReplaceSuffix
// Deprecated: Use [Has] with [Suf] instead.
func (s String) ReplaceSuffix(pats ...string) String {
	return String(ReplaceSuffix(pats...)(string(s)))
}

func (s String) Replace(pats ...string) String { return String(Replace(pats...)(string(s))) }

func (s String) ReplaceRegex(pat string, rep string) String {
	return String(ReplaceRegex[string](pat, rep)(string(s)))
}

func (s String) Shift(count int) String { return String(Shift[string](count)(string(s))) }
func (s String) Pop(count int) String   { return String(Pop[string](count)(string(s))) }
func (s String) Split(pats ...string) Array[String] {
	split := Split(string(s), false, pats...)
	res := make([]String, len(split))
	for i, v := range split {
		res[i] = String(v)
	}
	return ToArray(res)
}

func (s String) Or(Default string) String {
	if s == "" {
		return String(Default)
	}
	return s
}
func (s String) Len() int       { return len(string(s)) }
func (s String) String() string { return string(s) }

func (s String) Colorize(colorCode int) String { return String(Colorize(colorCode, string(s))) }
func (s String) ToUpper() String               { return String(strings.ToUpper(string(s))) }
func (s String) ToLower() String               { return String(strings.ToLower(string(s))) }
func (s String) Trim(arg string) String        { return String(strings.Trim(string(s), arg)) }
func (s String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(string(s), prefix))
}

func (s String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(string(s), suffix))
}
func (s String) TrimSpace() String { return String(strings.TrimSpace(string(s))) }

func (s String) Green() String        { return String(Colorize(Green, string(s))) }
func (s String) LightGreen() String   { return String(Colorize(LightGreen, string(s))) }
func (s String) Yellow() String       { return String(Colorize(Yellow, string(s))) }
func (s String) LightYellow() String  { return String(Colorize(LightYellow, string(s))) }
func (s String) Red() String          { return Colorize(Red, s) }
func (s String) LightRed() String     { return Colorize(LightRed, s) }
func (s String) Blue() String         { return Colorize(Blue, s) }
func (s String) LightBlue() String    { return Colorize(LightBlue, s) }
func (s String) Cyan() String         { return Colorize(Cyan, s) }
func (s String) LightCyan() String    { return Colorize(LightCyan, s) }
func (s String) Magenta() String      { return Colorize(Magenta, s) }
func (s String) LightMagenta() String { return Colorize(LightMagenta, s) }
func (s String) White() String        { return Colorize(White, s) }
func (s String) Black() String        { return Colorize(Black, s) }
func (s String) Gray() String         { return Colorize(DarkGray, s) }
func (s String) LightGray() String    { return Colorize(LightGray, s) }

func (s String) Dim() String  { return Colorize(2, s) }
func (s String) Bold() String { return Bold(s) }

func Whitespaces() []string { return []string{"\r\n", "\n\r", " ", "\t", "\n", "\r"} }

func (s String) ToInt() Option[int]         { return ToInt(s) }
func (s String) ToInt8() Option[int8]       { return ToInt8(s) }
func (s String) ToInt16() Option[int16]     { return ToInt16(s) }
func (s String) ToInt32() Option[int32]     { return ToInt32(s) }
func (s String) ToInt64() Option[int64]     { return ToInt64(s) }
func (s String) ToUint() Option[uint]       { return ToUint(s) }
func (s String) ToUint8() Option[uint8]     { return ToUint8(s) }
func (s String) ToUint16() Option[uint16]   { return ToUint16(s) }
func (s String) ToUint32() Option[uint32]   { return ToUint32(s) }
func (s String) ToUint64() Option[uint64]   { return ToUint64(s) }
func (s String) ToFloat32() Option[float32] { return ToFloat32(s) }
func (s String) ToFloat64() Option[float64] { return ToFloat64(s) }

func ToInt[S ~string](s S) Option[int] {
	i, err := strconv.Atoi(string(s))
	return Either[int, bool]{Value: int(i), Other: err == nil}
}
func ToInt8[S ~string](s S) Option[int8] {
	i, err := strconv.ParseInt(string(s), 10, 8)
	return Either[int8, bool]{Value: int8(i), Other: err == nil}
}
func ToInt16[S ~string](s S) Option[int16] {
	i, err := strconv.ParseInt(string(s), 10, 16)
	return Either[int16, bool]{Value: int16(i), Other: err == nil}
}
func ToInt32[S ~string](s S) Option[int32] {
	i, err := strconv.ParseInt(string(s), 10, 32)
	return Either[int32, bool]{Value: int32(i), Other: err == nil}
}
func ToInt64[S ~string](s S) Option[int64] {
	i, err := strconv.ParseInt(string(s), 10, 64)
	return Either[int64, bool]{Value: int64(i), Other: err == nil}
}
func ToUint[S ~string](s S) Option[uint] {
	i, err := strconv.ParseUint(string(s), 10, 0)
	return Either[uint, bool]{Value: uint(i), Other: err == nil}
}
func ToUint8[S ~string](s S) Option[uint8] {
	i, err := strconv.ParseUint(string(s), 10, 8)
	return Either[uint8, bool]{Value: uint8(i), Other: err == nil}
}
func ToUint16[S ~string](s S) Option[uint16] {
	i, err := strconv.ParseUint(string(s), 10, 16)
	return Either[uint16, bool]{Value: uint16(i), Other: err == nil}
}
func ToUint32[S ~string](s S) Option[uint32] {
	i, err := strconv.ParseUint(string(s), 10, 32)
	return Either[uint32, bool]{Value: uint32(i), Other: err == nil}
}
func ToUint64[S ~string](s S) Option[uint64] {
	i, err := strconv.ParseUint(string(s), 10, 64)
	return Either[uint64, bool]{Value: uint64(i), Other: err == nil}
}
func ToFloat32[S ~string](s S) Option[float32] {
	i, err := strconv.ParseFloat(string(s), 32)
	return Either[float32, bool]{Value: float32(i), Other: err == nil}
}
func ToFloat64[S ~string](s S) Option[float64] {
	i, err := strconv.ParseFloat(string(s), 64)
	return Either[float64, bool]{Value: float64(i), Other: err == nil}
}

func ToUpper(s string) string { return strings.ToUpper(s) }
func ToLower(s string) string { return strings.ToLower(s) }

func Trim(s string) string                      { return strings.Trim(s, " ") }
func TrimPrefix(prefix string, s string) string { return strings.TrimPrefix(s, prefix) }
func TrimSuffix(suffix string, s string) string { return strings.TrimSuffix(s, suffix) }
func TrimSpace[S ~string](s S) S                { return S(strings.TrimSpace(string(s))) }

func TrimPrefixes[S, A ~string](pats ...A) func(S) S {
	return func(inp S) S {
		for _, pat := range pats {
			if HasPrefix(pat)(A(inp)) {
				return S(strings.TrimPrefix(string(inp), string(pat)))
			}
		}
		return inp
	}
}

func TrimSuffixes[A, S ~string](pats ...A) func(S) S {
	return func(inp S) S {
		for _, pat := range pats {
			if HasSuffix(pat)(A(inp)) {
				return S(strings.TrimSuffix(string(inp), string(pat)))
			}
		}
		return inp
	}
}

// Arguably coloring strings does not belong in blume. A problem for another day.
const (
	reset = "\033[0m"

	Black        = 30
	Red          = 31
	Green        = 32
	Yellow       = 33
	Blue         = 34
	Magenta      = 35
	Cyan         = 36
	LightGray    = 37
	DarkGray     = 90
	LightRed     = 91
	LightGreen   = 92
	LightYellow  = 93
	LightBlue    = 94
	LightMagenta = 95
	LightCyan    = 96
	White        = 97
)

func Colorize[S ~string](colorCode int, s S) S {
	return "\033[" + S(strconv.Itoa(colorCode)) + "m" + s + "\033[0m"
}

func Dim[S ~string](s S) S  { return Colorize(2, s) }
func Bold[S ~string](s S) S { return "\033[1m" + s + "\033[0m" }
