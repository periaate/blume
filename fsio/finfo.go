package fsio

import (
	"os"
	"path/filepath"
)

func ToFinfo(path string) Finfo { return Finfo(Finfo(path).Clean()) }

type Finfo string

func (f Finfo) String() string { return string(f) }

func (f Finfo) Dir() string  { return filepath.Dir(string(f)) }
func (f Finfo) Base() string { return filepath.Base(string(f)) }

func (f Finfo) Open() (*os.File, error) { return os.Open(string(f)) }
func (f Finfo) Read() ([]byte, error)   { return os.ReadFile(string(f)) }

func (f Finfo) Clean() string { return filepath.ToSlash(filepath.Clean(string(f))) }

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
	return os.MkdirAll(f.Dir(), 0o755)
}

func (f Finfo) EnsureFile() error {
	if f.Exists() {
		return nil
	}
	err := os.MkdirAll(f.Dir(), 0o755)
	if err != nil {
		return err
	}
	_, err = os.Create(string(f))
	return err
}
