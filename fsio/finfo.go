package fsio

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	fp "path/filepath"
)

// Normalize cleans and converts the path to use forward slashes.
func Normalize(path string) string { return fp.ToSlash(fp.Clean(path)) }

func Symlink(dst string) func(string) error {
	return func(f string) (err error) {
		err = os.Symlink(f, dst)
		if err != nil {
			return
		}
		return
	}
}

func ReadDir(f string) (res []string, err error) {
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

type ArbFS struct {
	paths map[string]string
}

func (a ArbFS) Open(path string) (file fs.File, err error) { return os.Open(a.paths[path]) }

func ToFS(paths ...string) ArbFS {
	f := ArbFS{make(map[string]string, 0)}
	for _, v := range paths {
		v = Normalize(v)
		f.paths[v] = v
	}

	return f
}

func Name(f string) string {
	b := fp.Base(f)
	r := b[:len(b)-len(fp.Ext(b))]
	return Normalize(r)
}

func WriteAll(f string, r io.Reader) (err error) {
	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func WriteNew(f string, r io.Reader) (err error) {
	if Exists(f) {
		return fmt.Errorf("file %s already exists", f)
	}
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func AppendTo(f string, r io.Reader) (err error) {
	// Open the file in append mode, create if not exists
	file, err := os.OpenFile(f, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

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

func IsDir(f string) bool {
	info, err := os.Stat(f)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func Exists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

func EnsureDir(f string) error {
	if Exists(f) {
		return nil
	}
	return os.MkdirAll(f, 0o755)
}

func EnsureFile(f string) (err error) {
	if err = EnsureDir(fp.Dir(f)); err != nil {
		return
	}

	_, err = os.Create(f)
	return
}

func Copy(dst string, overwrite bool) func(string) (err error) {
	fd := Normalize(dst)
	return func(src string) (err error) {
		fs := Normalize(src)
		if IsDir(fd) || IsDir(fs) {
			return
		}

		if !Exists(fs) {
			return fmt.Errorf("no such file %s exists", fs)
		}

		err = EnsureDir(fd)
		if err != nil {
			return
		}

		srcFile, err := os.Open(fs)
		if err != nil {
			return
		}
		defer srcFile.Close()

		if overwrite {
			return WriteAll(fd, srcFile)
		}

		return WriteNew(fd, srcFile)
	}
}
