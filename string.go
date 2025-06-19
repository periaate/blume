package blume

import (
	"fmt"
	"io/fs"
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
	return S(strings.Join(Map(StoD[S])(arr), arg))
}

type String string

func (s S) In(a Array[S]) bool { return a.First(Is(s)).IsOk() }

func (s String) Map(args ...func(String) String) String  { return Pipe(args...)(s) }
func (s String) Maps(args ...func(string) string) String { return String(Pipe(args...)(s.String())) }
func (s String) Has(args ...Selector[String]) bool       { return Has(args...)(s) }
func (s String) Del(args ...Selector[String]) String     { return String(Del(args...)(s)) }
func (s String) Rep(args ...any) String                  { return String(Rep[string](args...)(string(s))) }

func (s String) Is(args ...String) bool       { return Is(args...)(String(s)) }
func (s String) Contains(args ...String) bool { return Contains(args...)(String(s)) }

func (s String) EnsurePrefix(fix String) String {
	if HasPrefix(fix)(s) {
		return s
	}
	return fix + s
}

func (s String) EnsureSuffix(fix String) String {
	if HasSuffix(fix)(s) {
		return s
	}
	return s + fix
}

func EnsurePrefix(fix String) func(S) S { return func(s String) S { return s.EnsurePrefix(fix) } }
func EnsureSuffix(fix String) func(S) S { return func(s String) S { return s.EnsureSuffix(fix) } }

func (s String) HasPrefix(args ...S) bool       { return HasPrefix(args...)(s) }
func (s String) HasSuffix(args ...S) bool       { return HasSuffix(args...)(s) }
func (s String) ReplacePrefix(pats ...S) String { return String(ReplacePrefix(pats...)(s)) }
func (s String) ReplaceSuffix(pats ...S) String { return String(ReplaceSuffix(pats...)(s)) }

func (s String) Replace(pats ...S) String { return String(Replace(pats...)((s))) }

func (s String) ReplaceRegex(pat S, rep S) String {
	return ReplaceRegex(pat, rep)(s)
}

func (s String) Shift(count int) String { return Shift(count)(s) }
func (s String) Pop(count int) String   { return Pop(count)(s) }

func (s String) SplitRegex(pat S) Array[String]                { return ToArray(SplitRegex(pat)(s)) }
func (s String) SplitsRegex(pat S) []String                    { return SplitRegex(pat)(s) }
func (s String) Split(keep bool, pats ...String) Array[String] { return ToArray(Split(s, false, pats...)) }
func (s String) Splits(keep bool, pats ...String) []String     { return Split(s, keep, pats...) }

func IsArray[A any](arg any) bool {
	if arg == nil {
		return false
	}
	kind := reflect.TypeOf(arg).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

func IsSymlink(s S) Option[bool] {
	stat := s.Stat()
	if !stat.IsOk() { return None[bool]() }
	return Some(stat.Value.Mode()&fs.ModeSymlink != 0)
}

func (s String)IsSymlink() Option[bool] {
	stat := s.Stat()
	if !stat.IsOk() { return None[bool]() }
	return Some(stat.Value.Mode()&fs.ModeSymlink != 0)
}

type S = String
type E = error

func (f String) Errorf(args ...any) error { return fmt.Errorf("%s", f.S(args...)) }

func IsURL[S ~string](val S) bool {
	return String(val).Contains("://") // giga scuff
}

func GetPath(val S) S { return Del(Rgx(`^([A-z]*://)?[A-z|0-9|\.|-]*`))(val) }
func GetDomain(val S) S {
	return ReplaceRegex(`^([A-z]*://)?([A-z|0-9|\.|-]*).*`, "$2")(val)
}


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
func Base(s String) String { return String(filepath.Base(string(s))) }

func (s String) Dir() String { return String(filepath.Dir(string(s)))+"/" }
func Dir(s S) String { return String(filepath.Dir(string(s)))+"/" }

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

func (s String) Read() Result[String] {
	bar, err := os.ReadFile(string(s.Path()))
	return Auto(S(bar), err)
}

func (s String) Open() Result[*os.File] { return Auto(os.Open(string(s))) }

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

func Whitespaces() []string { return []string{"\r\n", "\n\r", " ", "\t", "\n", "\r"} }

func ToUpper(s String) String { return S(strings.ToUpper(string(s))) }
func ToLower(s String) String { return S(strings.ToLower(string(s))) }

func Trim(delim String) func(S) S {
	return func(s S) S { return S(strings.Trim(s.String(), delim.String())) }
}
func TrimPrefix(prefix string, s string) string { return strings.TrimPrefix(s, prefix) }
func TrimSuffix(suffix string, s string) string { return strings.TrimSuffix(s, suffix) }
func TrimSpace[S ~string](s S) S                { return S(strings.TrimSpace(string(s))) }

func TrimPrefixes(pats ...S) func(S) S {
	return func(inp S) S {
		for _, pat := range pats {
			if HasPrefix(pat)(inp) {
				return S(strings.TrimPrefix(string(inp), string(pat)))
			}
		}
		return inp
	}
}

func TrimSuffixes(pats ...S) func(S) S {
	return func(inp S) S {
		for _, pat := range pats {
			if HasSuffix(pat)(inp) {
				return S(strings.TrimSuffix(string(inp), string(pat)))
			}
		}
		return inp
	}
}
