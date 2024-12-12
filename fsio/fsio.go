package fsio

import (
	"os"
	fp "path/filepath"
	"regexp"
	"strings"

	. "github.com/periaate/blume"
)

func Clean[S ~string](inp S) S {
	path := string(inp)
	var pre string
	var spl string
	var aft string
	aft = path
	split := Split(path, false, "://")
	if len(split) >= 2 {
		if Contains("/")(split[0]) { return S(strings.Join(split, "://")) }

		pre = split[0]
		spl = "://"
		aft = strings.Join(split[1:], "/")
		if HasPrefix("/")(aft) { aft = aft[1:] }
	}

	path = ToSlash(aft)

	regexp := regexp.MustCompile(`[/]+`)
	path = regexp.ReplaceAllString(path, "/")

	path = pre + spl + path
	return S(path)
}

func ToSlash[S ~string](path S) S { return S(strings.ReplaceAll(string(path), "\\", "/")) }

var Home = func() func() string {
	var str string
	var loaded bool
	return func() string {
		if loaded { return str }
		str, _ = os.UserHomeDir()
		loaded = true
		return str
	}
}()

// ReadDir reads the directory and returns a list of files.
func ReadDir[S ~string](inp S) Result[Array[S]] {
	f := string(inp)
	if HasPrefix("~")(f) { f = strings.Replace(f, "~", Home(), 1) }
	if !IsDir(f) { return Errf[Array[S]]("%s is not a directory", f) }
	
	entries, err := os.ReadDir(f)
	if err != nil { return Errf[Array[S]]("failed to read directory [%s] with error: [%w]", f, err) }

	res := make([]S, 0, len(entries))
	for _, entry := range entries {
		fp := entry.Name()
		if entry.IsDir() { fp += "/" }
		res = append(res, S(Join(f, fp).Unwrap())) // a panic is impossible
	}

	return Ok(ToArray(res))
}

// Name returns the file name without the extension and directory.
// TODO: create map of extensions and split by them.
func Name[S ~string](f S) S {
	b := Base(f)
	r := b[:len(b)-len(Ext(b))]
	return Clean(r)
}

func AbsPath[S ~string](f S) Result[S] {
	path, err := fp.Abs(string(f))
	return AsRes(S(path), err)
}

func Dir[S ~string](f S) S  { return S(fp.Dir(string(f))) }
func Base[S ~string](f S) S { return S(fp.Base(string(f))) }
func Ext[S ~string](f S) S   { return S(fp.Ext(string(f))) }

// IsDir checks if input is a directory.
func IsDir[S ~string](f S) bool {
	info, err := os.Stat(string(f))
	if err != nil { return false }
	return info.IsDir()
}

// Exists checks if the input exists.
func Exists[S ~string](f S) bool {
	_, err := os.Stat(string(f))
	return !os.IsNotExist(err)
}

// EnsureDir creates the directory recursively if it does not exist.
func EnsureDir[S ~string](f S) error {
	if Exists(f) { return nil }
	return os.MkdirAll(string(f), 0o755)
}

// Touch creates the file if it does not exist.
func Touch[S ~string](inp S) error {
	f := string(inp)
	if Exists(f) { return nil }
	if err := EnsureDir(Dir(f)); err != nil { return err }
	file, err := os.Create(f)
	if err != nil { return err }
	return file.Close()
}

// Join joins the path elements.
func Join[S ~string](args ...S) Result[S] {
	elems := Map(func(str S) string { return string(str) })(args)
	var res string

	elems = Filter(func(str string) bool { return str != "" })(elems)
	if len(elems) == 0 { return Errf[S]("no elements to join") }
	var isDir, isRel bool

	if len(elems) >= 1 { isDir = HasSuffix("/", `\`)(elems[len(elems)-1]) }

	isRel = HasPrefix(".", "./", `.\`)(elems[0]) && !HasPrefix("/", `\`, "..")(elems[0])

	res = Clean(strings.Join(elems, "/"))
	if isDir { res += "/" }

	res = Clean(res)
	if isRel || HasPrefix(".")(res) {
		if !HasPrefix("./", `.\`, ".")(res) {
			res = "./" + res
		} else if !HasPrefix("..")(res) {
			res = ReplacePrefix(
				"./", "./",
				`.\`, "./",
				`.`, "./",
			)(res)
		}
	}

	return Ok(S(res))
}
