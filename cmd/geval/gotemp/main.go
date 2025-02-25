package main

import (
	"path/filepath"

	. "github.com/periaate/blume"
)

var _ String

func main() {
	println(filepath.Join("/home/periaate", "my", "path", "file.jpg"))
}
