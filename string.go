package blume

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/periaate/blume/color"
	"github.com/periaate/blume/fsio"
)

func Join(arg String) func(args []String) String {
	return func(args []String) String { return S(strings.Join(SD(args), arg.String())) }
}

func Joins[S ~string](arr Array[S], arg string) S {
	return S(strings.Join(Map(StoD[S])(arr.Value), arg))
}

type String string

func (s String) Map(args ...func(String) String) String  { return Pipe(args...)(s) }
func (s String) Maps(args ...func(string) string) String { return String(Pipe(args...)(s.String())) }
func (s String) Has(args ...Selector[string]) bool       { return Has(args...)(s.String()) }
func (s String) Del(args ...Selector[string]) String     { return String(Del(args...)(s.String())) }
func (s String) Rep(args ...any) String                  { return String(Rep[string](args...)(string(s))) }

func (s String) Is(args ...String) bool       { return Is(args...)(String(s)) }
func (s String) Contains(args ...String) bool { return Contains(args...)(String(s)) }

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

func IsArray[A any](arg any) bool {
	if arg == nil {
		return false
	}
	kind := reflect.TypeOf(arg).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

type S = String

func (f String) Errorf(args ...any) error { return fmt.Errorf("%s", f.S(args...)) }

func IsURL[S ~string](val S) bool {
	return String(val).Contains("://") // giga scuff
}

func GetPath[S ~string](val S) S { return Del(Rgx[S](`^([A-z]*://)?[A-z|0-9|\.|-]*`))(val) }
func GetDomain[S ~string](val S) S {
	return ReplaceRegex[S](`^([A-z]*://)?([A-z|0-9|\.|-]*).*`, "$2")(val)
}

func (s String) Entries() Result[Array[String]]    { return Entries(s) }
func Entries[S ~string](s S) Result[Array[String]] { return Dir(s).Read() }

func (s String) Path(args ...String) String { return Path(Prepend(s, args)...) }

func Exists(s String) bool { return fsio.Exists(string(s)) }
func Chdir(s String) Result[String] {
	switch err := os.Chdir(string(s)); err {
	case nil:
		return Ok(s)
	default:
		return Err[String](s.W().Errorf("couldn't chdir %w", err))
	}
}

func (s String) Exists() bool { return fsio.Exists(string(s)) }
func (s String) Chdir() Result[String] {
	switch err := os.Chdir(string(s)); err {
	case nil:
		return Ok(s)
	default:
		return Err[String](s.W().Errorf("couldn't chdir %w", err))
	}
}

func (s String) Base() String { return String(filepath.Base(string(s))) }
func Base(s String) String    { return String(filepath.Base(string(s))) }

func (s String) IsDir() bool { return fsio.IsDir(string(s)) }
func IsDir(s String) bool    { return fsio.IsDir(string(s)) }

func (s String) Or(Default string) String {
	if s == "" {
		return String(Default)
	}
	return s
}
func (s String) Len() int       { return len(string(s)) }
func (s String) String() string { return string(s) }

func (s String) Colorize(colorCode int) String { return String(color.Colorize(colorCode, string(s))) }
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

func (s String) Green() String        { return String(color.Colorize(color.Green, string(s))) }
func (s String) LightGreen() String   { return String(color.Colorize(color.LightGreen, string(s))) }
func (s String) Yellow() String       { return String(color.Colorize(color.Yellow, string(s))) }
func (s String) LightYellow() String  { return String(color.Colorize(color.LightYellow, string(s))) }
func (s String) Red() String          { return color.Colorize(color.Red, s) }
func (s String) LightRed() String     { return color.Colorize(color.LightRed, s) }
func (s String) Blue() String         { return color.Colorize(color.Blue, s) }
func (s String) LightBlue() String    { return color.Colorize(color.LightBlue, s) }
func (s String) Cyan() String         { return color.Colorize(color.Cyan, s) }
func (s String) LightCyan() String    { return color.Colorize(color.LightCyan, s) }
func (s String) Magenta() String      { return color.Colorize(color.Magenta, s) }
func (s String) LightMagenta() String { return color.Colorize(color.LightMagenta, s) }
func (s String) White() String        { return color.Colorize(color.White, s) }
func (s String) Black() String        { return color.Colorize(color.Black, s) }
func (s String) Gray() String         { return color.Colorize(color.DarkGray, s) }
func (s String) LightGray() String    { return color.Colorize(color.LightGray, s) }

func (s String) Dim() String  { return color.Colorize(2, s) }
func (s String) Bold() String { return color.Bold(s) }

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

func ToUpper(s String) String { return S(strings.ToUpper(string(s))) }
func ToLower(s String) String { return S(strings.ToLower(string(s))) }

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
