package fsio

import (
	"path/filepath"

	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/gen/T"
)

//blume:derive String
type FilePath string

func (f FilePath) Abs() T.Result[FilePath] {
	res, err := filepath.Abs(string(f))
	return T.Results(FilePath(res), err)
}

func (f FilePath) Base() FilePath                          { return Base(f) }
func (f FilePath) Dir() FilePath                           { return Dir(f) }
func (f FilePath) Clean() FilePath                         { return Clean(f) }
func (f FilePath) ReadsDir() T.Result[gen.Array[FilePath]] { return ReadsDir(f) }
