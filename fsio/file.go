package fsio

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

func Copy[DST, SRC ~string](dst DST, src SRC, force bool) T.Result[any] {
	f, err := os.Open(string(src))
	if err != nil {
		return T.Results[any](nil, err)
	}
	defer f.Close()

	switch force {
	case true:
		err = WriteAll(string(dst), f)
	case false:
		err = WriteNew(string(dst), f)
	}

	return T.Results[any](nil, err)
}

func Read[S ~string](fp S) T.Result[[]byte] {
	return T.Results(os.ReadFile(string(fp)))
}

func ReadDirRecursively(fp string) (res []string) {
	dirs := []string{fp}

	for {
		fmt.Println("len(dirs):", len(dirs))
		if len(dirs) == 0 {
			break
		}

		dir := dirs[0]
		dirs = dirs[1:]
		f := gen.Is(".", ".git", ".idea", "node_modules", "./", "..", "")

		fmt.Println("reading dir:", dir)
		entries := gen.Must(os.ReadDir(dir))
		files := make([]string, 0, len(entries))
		for _, entry := range entries {
			files = append(files, filepath.Join(dir, entry.Name()))
		}

		fmt.Println("files:", files)
		for _, file := range files {
			if IsDir(file) {
				if f(file) {
					continue
				}
				dirs = append(dirs, file)
			}
			res = append(res, file)
		}
	}

	return
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

func Open(f string) (rc io.ReadCloser, err error) {
	file, err := os.Open(f)
	if err != nil {
		return
	}
	rc = file
	return
}

func Remove(f string) (err error) { return os.Remove(f) }

func ReadTo(f string, r io.Reader) (n int64, err error) {
	file, err := os.Create(f)
	if err != nil {
		return
	}
	defer file.Close()

	n, err = io.Copy(file, r)
	return
}
