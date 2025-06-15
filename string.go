package blume

import (
	"fmt"
	"strings"
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

// func IsSymlink(s string) (res Result[bool]) {
// 	stat, err := os.Stat(s)
// 	if err != nil { return res.Fail(err) }
// 	return res.Pass(stat.Mode()&fs.ModeSymlink != 0)
// }

// func GetPath(val string) string { return Del(Rgx(`^([A-z]*://)?[A-z|0-9|\.|-]*`))(val) }
// func GetDomain(val string) string { return ReplaceRegex(`^([A-z]*://)?([A-z|0-9|\.|-]*).*`, "$2")(val) }

// func Exists(s string) bool { return fsio.Exists(s) }
// func Chdir(s string) Result[string] {
// 	switch err := os.Chdir(s); err {
// 	case nil: return Ok(s)
// 	default:  return Err[string](err)
// 	}
// }

// func Base(s string) string { return filepath.Base(string(s)) }
// func Dir(s string) string  { return filepath.Dir(string(s))+"/" }
// func IsDir(s string) bool  { return fsio.IsDir(s) }

// func ReadFile(s string) (res Result[string]) {
// 	bar, err := os.ReadFile(Path(s))
// 	return res.Auto(S(bar), err)
// }

// func OpenFile(s string) (res Result[*os.File]) { return res.Auto(os.Open(Path(s))) }
