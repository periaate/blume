package fsio

import (
	"fmt"
	"os"
	fp "path/filepath"
	"regexp"
	"strings"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/lazy"
)

var _ = Zero[any]

func Clean[S ~string](inp S) S {
	path := string(inp)
	var pre string
	var spl string
	var aft string
	aft = path
	split := gen.SplitWithAll(path, false, "://")
	if len(split) >= 2 {
		if gen.Contains("/")(split[0]) { return S(strings.Join(split, "://")) }

		pre = split[0]
		spl = "://"
		aft = strings.Join(split[1:], "/")
		if gen.HasPrefix("/")(aft) { aft = aft[1:] }
	}

	path = ToSlash(aft)

	regexp := regexp.MustCompile(`[/]+`)
	path = regexp.ReplaceAllString(path, "/")

	path = pre + spl + path
	return S(path)
}

func ToSlash(path string) string { return strings.ReplaceAll(path, "\\", "/") }

var Home = lazy.Niladic(func() string { str, _ := os.UserHomeDir(); return str })

// ReadDir reads the directory and returns a list of files.
func ReadsDir[S ~string](fp S) Option[Array[FilePath]] {
	f := string(fp)
	var r Array[FilePath]
	res := []FilePath{}
	f = gen.ReplacePrefix("~", Home())(f)

	if !IsDir(f) { return None[Array[FilePath]]("couldn't read directory as [%s] is not a directory", f) }

	entries, err := os.ReadDir(f)
	if err != nil { return None[Array[FilePath]]("failed to read directory [%s] with error: [%w]", f, err) }

	res = make([]FilePath, 0, len(entries))

	for _, entry := range entries {
		fp := entry.Name()
		if entry.IsDir() { fp += "/" }
		res = append(res, FilePath(Join(f, fp)))
	}

	r = ToArray(res)
	return Some(r)
}

// ReadDir reads the directory and returns a list of files.
func ReadDir(f string) (res []string, err error) {
	if gen.HasPrefix("~")(f) { f = strings.Replace(f, "~", Home(), 1) }

	if !IsDir(f) {
		err = fmt.Errorf("%s is not a directory", f)
		return
	}

	entries, err := os.ReadDir(f)
	if err != nil { return }

	res = make([]string, 0, len(entries))

	for _, entry := range entries {
		fp := entry.Name()
		if entry.IsDir() { fp += "/" }
		res = append(res, Join(f, fp))
	}

	return
}

// Name returns the file name without the extension and directory.
// TODO: create map of extensions and split by them.
func Name(f string) string {
	b := Base(f)
	r := b[:len(b)-len(Ext(b))]
	return Clean(r)
}

func AbsPath[S ~string](f S) Option[S] {
	path, err := fp.Abs(string(f))
	return Either[S](S(path), err)
}

func Dir[S ~string](f S) S  { return S(fp.Dir(string(f))) }
func Base[S ~string](f S) S { return S(fp.Base(string(f))) }
func Ext(f string) string   { return fp.Ext(f) }

// Walk walks the directory and returns a list of files that pass the predicate.
func Walk(fn func(string) bool) func(string) (res []string, err error) {
	return func(root string) (res []string, err error) {
		err = fp.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil { return err }
			if !fn(path) { return nil }

			path = Clean(path)

			res = append(res, path)
			return nil
		})

		return
	}
}

// IsDir checks if input is a directory.
func IsDir[S ~string](f S) bool {
	info, err := os.Stat(string(f))
	if err != nil { return false }
	return info.IsDir()
}

// Exists checks if the input exists.
func Exists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

// EnsureDir creates the directory recursively if it does not exist.
func EnsureDir(f string) error {
	if Exists(f) { return nil }
	return os.MkdirAll(f, 0o755)
}

// Touch creates the file if it does not exist.
func Touch(f string) error {
	if Exists(f) { return nil }
	if err := EnsureDir(Dir(f)); err != nil { return err }
	file, err := os.Create(f)
	if err != nil { return err }
	return file.Close()
}

// Join joins the path elements.
func Join[S ~string](args ...S) S {
	elems := Map(func(str S) string { return string(str) })(args)
	var res string

	elems = Filter(func(str string) bool { return str != "" })(elems)
	if len(elems) == 0 { return "" }
	var isDir, isRel bool

	if len(elems) >= 1 { isDir = gen.HasSuffix("/", `\`)(elems[len(elems)-1]) }

	isRel = gen.HasPrefix(".", "./", `.\`)(elems[0]) && !gen.HasPrefix("/", `\`, "..")(elems[0])

	res = Clean(strings.Join(elems, "/"))
	if isDir { res += "/" }

	res = Clean(res)
	if isRel || gen.HasPrefix(".")(res) {
		if !gen.HasPrefix("./", `.\`, ".")(res) {
			res = "./" + res
		} else if !gen.HasPrefix("..")(res) {
			res = gen.ReplacePrefix(
				"./", "./",
				`.\`, "./",
				`.`, "./",
			)(res)
		}
	}

	return S(res)
}
