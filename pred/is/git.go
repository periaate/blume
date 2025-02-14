package is

import (
	"path/filepath"
)

func GitRoot[S ~string](root S) bool {
	return filepath.Base(string(root)) == ".git"
}
