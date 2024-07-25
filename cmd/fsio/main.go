package main

import (
	"fmt"
	"path/filepath"

	"github.com/periaate/blume/clog"
	"github.com/periaate/blume/core"
	"github.com/periaate/blume/core/gen"
	"github.com/periaate/blume/fsio"
)

func main() {
	args := core.Args()
	cmd := args[0]
	args = args[1:]

	path := fsio.ToFinfo(args[0])
	var err error

	switch cmd {
	case "name":
		for _, arg := range args {
			fmt.Println(fsio.ToFinfo(arg).Name())
		}
	case "dir":
		for _, arg := range args {
			fmt.Println(fsio.ToFinfo(arg).Dir())
		}
	case "base":
		for _, arg := range args {
			fmt.Println(fsio.ToFinfo(arg).Base())
		}
	case "abs":
		fmt.Println(path.Abs())
	case "join":
		fmt.Println(fsio.Join(args[1:]...))
	case "exists":
		fmt.Println(path.Exists())
	case "ensure":
		switch args[0] {
		case "file", "f":
			path = fsio.ToFinfo(args[1])
			err = path.EnsureFile()
			if err == nil {
				fmt.Println("Successfully ensured file", path.String())
			}
		case "dir", "directory", "d", "folder":
			path = fsio.ToFinfo(args[1])
			fallthrough
		default:
			err = path.EnsureDir()
			if err == nil {
				fmt.Println("Successfully ensured directory", path.String())
			}
		}
	case "copy", "cp":
		err = Copy(path.String(), args...)
	case "copyrel", "rel", "cr":
		err = CopyRel(path.String(), args...)
	case "move", "mv":
		err = fmt.Errorf("not implemented")
	case "sym", "symlink", "ln":
		arg := args[0]
		err = path.Symlink(arg)
	case "read":
		for _, arg := range args {
			f := fsio.ToFinfo(arg)
			if f.IsDir() {
				continue
			}

			b, err := f.ReadAll()
			if err != nil {
				clog.Error("error reading file", "file", f.String(), "err", err)
				continue
			}

			fmt.Println(string(b))
		}
	default:
		gen.Map(args, gen.Pipe(fsio.Clean, Ln)) // lol
	}

	if err != nil {
		clog.Error("error running fsio", "cmd", cmd, "err", err)
		panic(err)
	}
}

func Ln[T any](str T) T {
	fmt.Println(str)
	return str
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
