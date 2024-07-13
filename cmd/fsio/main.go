package main

import (
	"fmt"
	"os"
	"path/filepath"

	"blume/clog"
	"blume/core"
	"blume/core/gen"
	"blume/fsio"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		return
	}
	cmd := args[0]
	path := fsio.ToFinfo(args[1])

	if path.Exists() && !path.IsDir() {
		clog.Error("path exists and is not dir", "path", path.String())
	}

	if err := path.EnsureDir(); err != nil {
		panic(err)
	}

	var err error

	switch cmd {
	case "abs":
		fmt.Println(path.Abs())
	case "join":
		fmt.Println(fsio.Join(args[1:]...))
	case "exists":
		fmt.Println(path.Exists())
	case "ensure":
		switch args[1] {
		case "file", "f":
			path = fsio.ToFinfo(args[2])
			err = path.EnsureFile()
			if err == nil {
				fmt.Println("Successfully ensured file", path.String())
			}
		case "dir", "directory", "d", "folder":
			path = fsio.ToFinfo(args[2])
			fallthrough
		default:
			err = path.EnsureDir()
			if err == nil {
				fmt.Println("Successfully ensured directory", path.String())
			}
		}
	case "copy", "cp":
		piped := core.ReadPipe()
		piped = append(args[2:], piped...)
		err = Copy(path.String(), piped...)
	case "copyrel", "rel", "cr":
		piped := core.ReadPipe()
		piped = append(args[2:], piped...)
		err = CopyRel(path.String(), piped...)
	case "move", "mv":
		err = fmt.Errorf("not implemented")
	case "sym", "symlink", "ln":
		arg := args[2]
		err = path.Symlink(arg)
	default:
		gen.Map(args, gen.Pipe(fsio.Clean, gen.Ln)) // lol
	}

	if err != nil {
		clog.Error("error running fsio", "cmd", cmd, "err", err)
		panic(err)
	}
}

func Copy(dst string, args ...string) (err error) {
	for _, arg := range args {
		destination := filepath.Join(dst, filepath.Base(arg))
		err = fsio.Copy(destination, arg, false)
		if err != nil {
			return
		}
	}
	return
}

func CopyRel(dst string, args ...string) (err error) {
	for _, arg := range args {
		fmt.Println(filepath.Join(dst, arg))
		destination := filepath.Join(dst, arg)
		err = fsio.Copy(destination, arg, false)
		if err != nil {
			return
		}
	}
	return
}

func Move(dst string, args ...string) (err error) {
	return
}
