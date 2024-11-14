package fsio

import (
	"fmt"
	"os"
	fp "path/filepath"
	"regexp"
	"strings"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/str"
)

// Normalize cleans and converts the path to use forward slashes.
func Normalize(path string) string {
	if str.HasSuffix("/")(path) {
		return fp.ToSlash(fp.Clean(path)) + "/"
	}
	return fp.ToSlash(fp.Clean(path))
}

var Home = gen.IgnoresNil(os.UserHomeDir)

// ReadDir reads the directory and returns a list of files.
func ReadDir(f string) (res []string, err error) {
	if str.HasPrefix("~")(f) {
		f = strings.Replace(f, "~", Home(), 1)
	}

	if !IsDir(f) {
		err = fmt.Errorf("%s is not a directory", f)
		return
	}

	entries, err := os.ReadDir(f)
	if err != nil {
		return
	}

	res = make([]string, 0, len(entries))

	for _, entry := range entries {
		res = append(res, fp.Join(f, entry.Name()))
	}

	return
}

// Name returns the file name without the extension and directory.
// TODO: create map of extensions and split by them.
func Name(f string) string {
	b := fp.Base(f)
	r := b[:len(b)-len(fp.Ext(b))]
	return Normalize(r)
}

func Dir(f string) string  { return fp.Dir(f) }
func Base(f string) string { return fp.Base(f) }

// Walk walks the directory and returns a list of files that pass the predicate.
func Walk(fn func(string) bool) func(string) (res []string, err error) {
	return func(root string) (res []string, err error) {
		err = fp.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !fn(path) {
				return nil
			}

			path = Normalize(path)

			res = append(res, path)
			return nil
		})

		return
	}
}

// IsDir checks if input is a directory.
func IsDir(f string) bool {
	info, err := os.Stat(f)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Exists checks if the input exists.
func Exists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

// EnsureDir creates the directory recursively if it does not exist.
func EnsureDir(f string) error {
	if Exists(f) {
		return nil
	}
	return os.MkdirAll(f, 0o755)
}

// Touch creates the file if it does not exist.
func Touch(f string) error {
	if Exists(f) {
		return nil
	}
	if err := EnsureDir(fp.Dir(f)); err != nil {
		return err
	}
	file, err := os.Create(f)
	if err != nil {
		return err
	}
	return file.Close()
}

// Join joins the path elements.
func Join(elems ...string) (res string) {
	elems = gen.Filter(func(str string) bool { return str != "" })(elems)
	if len(elems) == 0 {
		return ""
	}
	var isDir, isRel bool

	if len(elems) > 1 {
		isDir = str.HasSuffix("/", `\`)(gen.Ignore(gen.GetPop(elems)))
	}

	isRel = str.HasPrefix(".", "./", `.\`)(elems[0]) && !str.HasPrefix("/", `\`, "..")(elems[0])

	elems = gen.Map(func(str string) string {
		regexp := regexp.MustCompile(`[/\\]+`)
		return regexp.ReplaceAllString(str, "/")
	})(elems)

	res = fp.Join(elems...)
	if isDir {
		res += "/"
	}

	res = Normalize(res)
	if isRel || str.HasPrefix(".")(res) {
		if !str.HasPrefix("./", `.\`, ".")(res) {
			res = "./" + res
		} else if !str.HasPrefix("..")(res) {
			res = str.ReplacePrefix(
				"./", "./",
				`.\`, "./",
				`.`, "./",
			)(res)
		}
	}
	return res
}
