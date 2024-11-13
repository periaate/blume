package fsio

import (
	"fmt"
	"io"
	"os"
)

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
