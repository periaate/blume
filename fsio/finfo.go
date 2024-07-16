package fsio

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	fp "path/filepath"
)

func tf(path string) Finfo      { return Finfo(Clean(path)) }
func ToFinfo(path string) Finfo { return tf(path) }

type Finfo string

func (f Finfo) Symlink(dst string) (err error) {
	err = os.Symlink(string(f), dst)
	if err != nil {
		return
	}
	return
}

func (f Finfo) ReadDir() (res []string, err error) {
	if !f.IsDir() {
		err = fmt.Errorf("%s is not a directory", f.String())
		return
	}

	entries, err := os.ReadDir(f.String())
	if err != nil {
		return
	}

	res = make([]string, 0, len(entries))

	for _, entry := range entries {
		res = append(res, f.Join(entry.Name()))
	}

	return
}

type ArbFS struct {
	paths map[string]string
}

func (a *ArbFS) Open(path string) (file fs.File, err error) { return os.Open(a.paths[path]) }

func ToFS(paths ...string) *ArbFS {
	f := &ArbFS{make(map[string]string, 0)}
	for _, v := range paths {
		f.paths[v] = v
	}

	return f
}

func Abs(f Finfo) (string, error)    { return fp.Abs(string(f)) }
func (f Finfo) Abs() (string, error) { return fp.Abs(string(f)) }

func Open(f Finfo) (*os.File, error)    { return os.Open(string(f)) }
func (f Finfo) Open() (*os.File, error) { return os.Open(string(f)) }

func ReadAll(f Finfo) ([]byte, error)    { return os.ReadFile(string(f)) }
func (f Finfo) ReadAll() ([]byte, error) { return os.ReadFile(string(f)) }

func String(f Finfo) string    { return string(f) }
func (f Finfo) String() string { return string(f) }

func dir(f Finfo) Finfo    { return tf(fp.Dir(string(f))) }
func (f Finfo) dir() Finfo { return tf(fp.Dir(string(f))) }

func base(f Finfo) Finfo    { return tf(fp.Base(string(f))) }
func (f Finfo) base() Finfo { return tf(fp.Base(string(f))) }

func Dir(f Finfo) string    { return fp.Dir(string(f)) }
func (f Finfo) Dir() string { return fp.Dir(string(f)) }

func Base(f Finfo) string    { return fp.Base(string(f)) }
func (f Finfo) Base() string { return fp.Base(string(f)) }

func Join(a ...string) string           { return Clean(fp.Join(a...)) }
func (f Finfo) Join(a ...string) string { return Clean(fp.Join(append([]string{string(f)}, a...)...)) }

func (f Finfo) WriteAll(r io.Reader) (err error) {
	file, err := os.Create(string(f))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func (f Finfo) WriteNew(r io.Reader) (err error) {
	if f.Exists() {
		return fmt.Errorf("file %s already exists", f)
	}
	file, err := os.OpenFile(string(f), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func (f Finfo) AppendTo(r io.Reader) (err error) {
	// Open the file in append mode, create if not exists
	file, err := os.OpenFile(string(f), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func (f Finfo) Clean() string { return fp.ToSlash(fp.Clean(string(f))) }
func Clean(f string) string   { return fp.ToSlash(fp.Clean(f)) }

func (f Finfo) IsDir() bool {
	info, err := os.Stat(string(f))
	if err != nil {
		return false
	}
	return info.IsDir()
}

func (f Finfo) Exists() bool {
	_, err := os.Stat(string(f))
	return !os.IsNotExist(err)
}

func (f Finfo) EnsureDir() error {
	if f.Exists() {
		return nil
	}
	return os.MkdirAll(f.String(), 0o755)
}

func (f Finfo) EnsureFile() (err error) {
	if err = f.dir().EnsureDir(); err != nil {
		return
	}

	_, err = os.Create(string(f))
	return
}

func Copy(dst, src string, overwrite bool) (err error) {
	fd := ToFinfo(dst)
	fs := ToFinfo(src)
	if fd.IsDir() || fs.IsDir() {
		return
	}

	if !fs.Exists() {
		return fmt.Errorf("no such file %s exists", fs)
	}

	err = fd.EnsureDir()
	if err != nil {
		return
	}

	srcFile, err := fs.Open()
	if err != nil {
		return
	}
	defer srcFile.Close()

	if overwrite {
		return fd.WriteAll(srcFile)
	}

	fd.WriteNew(srcFile)
	return nil
}
