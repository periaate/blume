package fsio

import (
	"path/filepath"

	. "github.com/periaate/blume/core"
)

//blume:derive String
type FilePath string

func (f FilePath) Abs() Option[FilePath] {
	res, err := filepath.Abs(string(f))
	return Either[FilePath](FilePath(res), err)
}

func (f FilePath) Base() FilePath                          { return Base(f) }
func (f FilePath) Dir() FilePath                           { return Dir(f) }
func (f FilePath) Clean() FilePath                         { return Clean(f) }
func (f FilePath) ReadsDir() Option[Array[FilePath]] { return ReadsDir(f) }
