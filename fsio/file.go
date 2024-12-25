package fsio

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/periaate/blume/pred"
)

func Copy[DST, SRC ~string](dst DST, src SRC, force bool) error {
	f, err := os.Open(string(src))
	if err != nil {
		return fmt.Errorf("failed to copy from [%s] to [%s] with error: [%s]", src, dst, err)
	}
	defer f.Close()

	switch force {
	case true:
		err = WriteAll(string(dst), f)
	case false:
		err = WriteNew(string(dst), f)
	}
	return err
}

// func Read[S ~string](fp S) Result[[]byte] { return AsRes(os.ReadFile(string(fp))) }

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

// ReadPipe reads from stdin and returns a slice of lines.
func ReadPipe() (res []string) {
	if HasInPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}
	}
	return
}

func HasInPipe() bool {
	a, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (a.Mode() & os.ModeCharDevice) == 0
}

func HasOutPipe() bool {
	a, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (a.Mode() & os.ModeCharDevice) == 0
}

func Args(preds ...func([]string) bool) (res []string, ok bool) {
	if len(os.Args) >= 1 {
		res = os.Args[1:]
	}
	ok = pred.And(preds...)(res)
	return
}
