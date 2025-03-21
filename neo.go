package blume

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/periaate/blume/color"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/symbols"
)

type S = String

func Vals[K comparable, V any](m map[K]V) Array[V] {
	if m == nil {
		return Arr[V]()
	}
	arr := []V{}
	for _, v := range m {
		arr = append(arr, v)
	}
	return ToArray(arr)
}

func Keys[K comparable, V any](m map[K]V) Array[K] {
	if m == nil {
		return Arr[K]()
	}
	arr := []K{}
	for k := range m {
		arr = append(arr, k)
	}
	return ToArray(arr)
}

func (f String) Errorf(args ...any) error { return fmt.Errorf("%s", f.S(args...)) }

func IsURL[S ~string](val S) bool {
	return String(val).Contains("://") // giga scuff
}

func GetPath[S ~string](val S) S { return Del(Rgx[S](`^([A-z]*://)?[A-z|0-9|\.|-]*`))(val) }
func GetDomain[S ~string](val S) S {
	return ReplaceRegex[S](`^([A-z]*://)?([A-z|0-9|\.|-]*).*`, "$2")(val)
}

func (s String) Entries() Array[String]    { return Entries(s) }
func Entries[S ~string](s S) Array[String] { return Dir(s).Read().Must() }

func (f String) N() String { return f + "\n" }
func (f String) R() String { return f + "\r" }

func (f String) Up(lines ...int) String { return f + String(Up(ToArray(lines).Get(0).Or(1))) }
func (f String) Clean() String          { return f + String(Clean()) }
func (f String) S(args ...any) String   { return f + String(fmt.Sprint(args...)) }
func (f String) W() String              { return f + String(" ") }

func (f String) Print(args ...any) String   { fmt.Printf("%s%s", f, fmt.Sprint(args...)); return f }
func (f String) Println(args ...any) String { fmt.Printf("%s%s", f, fmt.Sprintln(args...)); return f }
func (f String) Printf(format string, args ...any) String {
	fmt.Printf("%s%s", f, fmt.Sprintf(format, args...))
	return f
}

func (f String) Color(hex string, args ...any) String {
	return f + String(ColorFg(hex)+fmt.Sprint(args...)+Reset)
}

func (f String) Spin(i int) String { return f + String(spinChars[i%len(spinChars)]) + " " }

func (f String) Info(args ...any) String { return f.Color(color.Info, symbols.Info).W().S(args...) }
func (f String) Lock(args ...any) String { return f.Color(color.Warning, symbols.Lock).W().S(args...) }
func (f String) Debug(args ...any) String {
	return f.Color(color.Pending, symbols.Debug).W().S(args...)
}
func (f String) Error(args ...any) String { return f.Color(color.Error, symbols.Error).W().S(args...) }
func (f String) Success(args ...any) String {
	return f.Color(color.Success, symbols.Success).W().S(args...)
}
func (f String) Warning(args ...any) String {
	return f.Color(color.Warning, symbols.Warning).W().S(args...)
}
func (f String) Waiting(args ...any) String {
	return f.Color(color.Waiting, symbols.Waiting).W().S(args...)
}
func (f String) Question(args ...any) String {
	return f.Color(color.Info, symbols.Question).W().S(args...)
}
func (f String) Cancelled(args ...any) String {
	return f.Color(color.Error, symbols.Cancelled).W().S(args...)
}
func (f String) InProgress(args ...any) String {
	return f.Color(color.Pending, symbols.InProgress).W().S(args...)
}

func (f String) Checkbox(done bool, args ...any) String {
	return T(done,
		f.Color(color.Success, symbols.CheckboxDone),
		f.Color(color.Warning, symbols.CheckboxEmpty),
	).S(args...)
}

func (s String) Path(args ...String) String { return Path(Prepend(s, args)...) }

func Prepend[A any](arg A, arr []A) []A { return append([]A{arg}, arr...) }

func LookupEnv(arg string) Option[String] {
	r, ok := os.LookupEnv(arg)
	if !ok {
		return None[String]()
	}
	return Some(String(r))
}

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
