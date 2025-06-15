package blume

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/pred/has"
)

func AllArgs(n ...int) []string { return append(Args(), Piped(os.Stdin).OrDef()...) }

func Args(n ...int) (res []string) {
	if len(os.Args) >= 1 { res = os.Args[1:] }
	if len(n) == 0       { return res }
	if len(res) < n[0]   { return res }
	return res[n[0]:]
}

func Arg(n int) Option[string] { return Index(os.Args, n+1) }

func Piped(input ...*os.File) Option[[]string] {
	var f *os.File
	if len(input) == 0 { f = os.Stdin } else { f = input[0] }

	if !has.Pipe(f) { return None[[]string]() }
	return Some(Lines(f))
}

func Lines[B any](bar B) []string {
	scanner := bufio.NewScanner(Buf(bar))
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func Entries(s string) (res Result[[]string]) {
	entries, err := fsio.ReadDir(s)
	if err != nil { return res.Fail() }
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.Path())
	}
	return res.Pass(values)
}

func Read(sar ...string) (res Result[string]) { return res.Auto(os.ReadFile(Path(sar...))) }

func Path(sar ...string) string {
	home, err := os.UserHomeDir()
	if err == nil { sar = Map[string, string](Replace("~", home))(sar) }
	fp := filepath.Join(sar...)
	absFp, err := filepath.Abs(fp)
	if err != nil { fp = string(fp) } else { fp = string(absFp) }
	if fsio.IsDir(fp) { fp = EnsureSuffix("/")(fp) }
	return fp
}

func IsSymlink(s string) (res Result[bool]) {
	stat, err := os.Stat(s)
	if err != nil { return res.Fail(err) }
	return res.Pass(stat.Mode()&fs.ModeSymlink != 0)
}

func GetPath(val string) string { return Del(Rgx(`^([A-z]*://)?[A-z|0-9|\.|-]*`))(val) }
func GetDomain(val string) string { return ReplaceRegex(`^([A-z]*://)?([A-z|0-9|\.|-]*).*`, "$2")(val) }

func Exists(s string) bool { return fsio.Exists(s) }
func Chdir(s string) (res Result[string]) { return res.Auto(os.Chdir(s)) }

func Base(s string) string { return filepath.Base(s) }
func Dir(s string) string  { return filepath.Dir(s)+"/" }

func TruePath(sar ...string) string {
	home, err := os.UserHomeDir()
	if err == nil { sar = Map[string, string](Replace("~", home))(sar) }
	fp := filepath.Join(sar...)
	if IsSymlink(string(fp)).Value {
		evaluated, err := filepath.EvalSymlinks(fp)
		if err == nil { fp = evaluated }
	}
	absFp, err := filepath.Abs(fp)
	if err != nil { fp = string(fp) } else { fp = string(absFp) }
	if fsio.IsDir(fp) { fp = EnsureSuffix("/")(fp) }
	return fp
}

func AppendTo(path string) (res *os.File, err error) { return os.OpenFile(Path(path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) }

func AppendLog[A any](f *os.File) (*os.File, func(a A) A) {
	mut := sync.Mutex{}
	return f,
		func(a A) A {
			mut.Lock()
			defer mut.Unlock()
			f.Write([]byte(fmt.Sprintf("%v\n", a)))
			return a
		}
}
