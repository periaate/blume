package blume

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	// "github.com/periaate/blume/color"
	"github.com/periaate/blume/fsio"
)

func Join(sep string) func(args []any) string {
	return func(args []any) string {
		res := []string{}
		for _, arg := range args {
			res = append(res, fmt.Sprint(arg))
		}
		return strings.Join(res, sep)
	}
}

func EnsurePrefix(fix string) func(string) string {
	return func(s string) string {
		if HasPrefix(fix)(s) { return s }
		return fix + s
	}
}

func EnsureSuffix(fix string) func(string) string {
	return func(s string) string {
		if HasSuffix(fix)(s) { return s }
		return s + fix
	}
}

func IsSymlink(s S) (res Result[bool]) {
	stat, err := os.Stat(s)
	if err != nil { return res.Fail(err) }
	return res.Pass(stat.Mode()&fs.ModeSymlink != 0)
}


type S = string
type E = error

func GetPath(val string) string { return Del(Rgx(`^([A-z]*://)?[A-z|0-9|\.|-]*`))(val) }
func GetDomain(val string) string { return ReplaceRegex(`^([A-z]*://)?([A-z|0-9|\.|-]*).*`, "$2")(val) }


func Exists(s string) bool { return fsio.Exists(s) }
func Chdir(s string) Result[string] {
	switch err := os.Chdir(s); err {
	case nil: return Ok(s)
	default:  return Err[string](err)
	}
}

func Base(s string) string { return filepath.Base(string(s)) }
func Dir(s string) string  { return filepath.Dir(string(s))+"/" }
func IsDir(s string) bool  { return fsio.IsDir(s) }

func ReadFile(s string) (res Result[string]) {
	bar, err := os.ReadFile(Path(s))
	return res.Auto(S(bar), err)
}

func OpenFile(s string) (res Result[*os.File]) { return res.Auto(os.Open(Path(s))) }

// func  Colorize(colorCode int) string { return color.Colorize(colorCode, s) }

// func (s String) ToUpper() String               { return String(strings.ToUpper(string(s))) }
// func (s String) ToLower() String               { return String(strings.ToLower(string(s))) }
// func (s String) Trim(arg string) String        { return String(strings.Trim(string(s), arg)) }
// func (s String) TrimPrefix(prefix string) String {
// 	return String(strings.TrimPrefix(string(s), prefix))
// }

// func (s String) TrimSuffix(suffix string) String {
// 	return String(strings.TrimSuffix(string(s), suffix))
// }
// func (s String) TrimSpace() String { return String(strings.TrimSpace(string(s))) }

// func Whitespaces() []string { return []string{"\r\n", "\n\r", " ", "\t", "\n", "\r"} }

// func ToUpper(s String) String { return S(strings.ToUpper(string(s))) }
// func ToLower(s String) String { return S(strings.ToLower(string(s))) }

// func Trim(delim String) func(S) S {
// 	return func(s S) S { return S(strings.Trim(s.String(), delim.String())) }
// }
// func TrimPrefix(prefix string, s string) string { return strings.TrimPrefix(s, prefix) }
// func TrimSuffix(suffix string, s string) string { return strings.TrimSuffix(s, suffix) }
// func TrimSpace[S ~string](s S) S                { return S(strings.TrimSpace(string(s))) }

// func TrimPrefixes(pats ...S) func(S) S {
// 	return func(inp S) S {
// 		for _, pat := range pats {
// 			if HasPrefix(pat)(inp) {
// 				return S(strings.TrimPrefix(string(inp), string(pat)))
// 			}
// 		}
// 		return inp
// 	}
// }
//
// func TrimSuffixes(pats ...S) func(S) S {
// 	return func(inp S) S {
// 		for _, pat := range pats {
// 			if HasSuffix(pat)(inp) {
// 				return S(strings.TrimSuffix(string(inp), string(pat)))
// 			}
// 		}
// 		return inp
// 	}
// }
