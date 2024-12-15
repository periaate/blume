package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/periaate/common"
	"github.com/periaate/media/thumbnails"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: <input> |thumb <output dir>")
		os.Exit(1)
	}

	val := os.Args[1]
	val = filepath.Clean(val)
	if _, err := os.Stat(val); os.IsNotExist(err) {
		fmt.Println("output directory does not exist")
		os.Exit(1)
	}

	args := common.ReadPipe()
	for _, arg := range args {
		if _, err := os.Stat(arg); os.IsNotExist(err) {
			continue
		}
		fp := filepath.Join(val, filepath.Base(arg))
		thumbnails.QueueThumb(arg, fp)
	}

	thumbnails.Wait()
}
