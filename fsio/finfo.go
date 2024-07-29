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

// Symlink creates a symbolic link to the destination.
func Symlink(dst string) func(string) error {
	return func(f string) (err error) {
		err = os.Symlink(f, dst)
		if err != nil {
			return
		}
		return
	}
}

// ReadDir reads the directory and returns a list of files.
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

// ArbFS is a file system that maps paths to files.
type ArbFS struct{ paths map[string]string }

// Open opens the file for reading.
func (a ArbFS) Open(path string) (file fs.File, err error) { return os.Open(a.paths[Normalize(path)]) }

// ToFS creates an [ArbFS] from the given paths, which implements the fs.FS interface.
func ToFS(paths ...string) ArbFS {
	f := ArbFS{make(map[string]string, 0)}
	for _, v := range paths {
		v = Normalize(v)
		f.paths[v] = v
	}

	return f
}

// Name returns the file name without the extension and directory.
func Name(f string) string {
	b := fp.Base(f)
	r := b[:len(b)-len(fp.Ext(b))]
	return Normalize(r)
}

// WriteAll writes the contents of the reader to the file, overwriting existing files.
func WriteAll(f string, r io.Reader) (err error) {
	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

// WriteNew writes the contents of the reader to a new file, will not overwrite existing files.
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

// AppendTo appends the contents of the reader to the file.
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

// EnsureDir creates the directory if it does not exist.
func EnsureDir(f string) error {
	if Exists(f) {
		return nil
	}
	return os.MkdirAll(f, 0o755)
}

// EnsureFile creates the file if it does not exist, along with the directory.
func EnsureFile(f string) (err error) {
	if err = EnsureDir(fp.Dir(f)); err != nil {
		return
	}

	_, err = os.Create(f)
	return
}

// This is currently incorrectly implemented.
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
