/*
Package main is the entry point for the fsio command line tool.

Fsio is an early proof of concept for the fsio library.

Most commands accept variadic arguments.
Piped values are interpreted as normal arguments.

Commands:

	:: Get the filename without the extension.
	name		[...string, pipe]

	:: Get the directory of the strings inputted
	dir			[...string, pipe]

	:: Get the base of the strings inputted
	base		[...string, pipe]

	:: Get the absolute path of the strings inputted
	abs			[...string, pipe]

	:: Join the strings inputted
	join		[...string, pipe]

	:: Check if the file exists
	exists		[string, pipe]

	:: Ensure the directory exists
	ensure dir	[string, pipe]

	:: Ensure the file exists
	ensure file	[string, pipe]

	:: Copy the files to the destination.
	copy	dst string  [...string, pipe]

	:: Copy the files to the destination with their paths relative to current working directory.
	copyrel	dst string	[...string, pipe]

	:: Move the files to the destination.
	move	dst [...string, pipe]

	:: Create a symlink to the destination.
	sym		dst tar

	:: Read the files and print their contents.
	read	[string, pipe]
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/periaate/blume/clog"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/gen"
)

func main() {
	args := fsio.Args()
	cmd := args[0]
	args = args[1:]

	path := fsio.Normalize(args[0])
	var err error

	switch cmd {
	case "name":
		for _, arg := range args {
			fmt.Println(fsio.Normalize(fsio.Name(arg)))
		}
	case "dir":
		for _, arg := range args {
			fmt.Println(fsio.Normalize(filepath.Dir(arg)))
		}
	case "base":
		for _, arg := range args {
			fmt.Println(fsio.Normalize(filepath.Base(arg)))
		}
	case "abs":
		for _, arg := range args {
			fmt.Println(fsio.Normalize(gen.Must(filepath.Abs(arg))))
		}
	case "join":
		fmt.Println(fsio.Normalize(filepath.Join(args...)))
	case "exists":
		fmt.Println(fsio.Exists(fsio.Normalize(args[0])))
	case "ensure":
		switch args[0] {
		case "file", "f":
			path = fsio.Normalize(args[1])
			if fsio.EnsureFile(path) == nil {
				fmt.Println("Successfully ensured file", path)
			}
		case "dir", "directory", "d", "folder":
			path = fsio.Normalize(args[1])
			fallthrough
		default:
			err = fsio.EnsureDir(path)
			if err == nil {
				fmt.Println("Successfully ensured directory", path)
			}
		}
	case "copy", "cp":
		err = copyTo(path, args...)
	case "copyrel", "rel", "cr":
		err = copyRel(path, args...)
	case "move", "mv":
		err = fmt.Errorf("not implemented")
	case "sym", "symlink", "ln":
		arg := args[0]
		err = fsio.Symlink(path)(arg)
	case "read":
		for _, arg := range args {
			arg = fsio.Normalize(arg)
			if fsio.IsDir(arg) {
				continue
			}

			b, err := os.ReadFile(arg)
			if err != nil {
				clog.Error("error reading file", "file", arg, "err", err)
				continue
			}

			fmt.Println(string(b))
		}
	default:
		for _, arg := range args {
			fmt.Println(fsio.Normalize(arg))
		}
	}

	if err != nil {
		clog.Error("error running fsio", "cmd", cmd, "err", err)
		panic(err)
	}
}

func copyTo(dst string, args ...string) (err error) {
	for _, arg := range args {
		destination := filepath.Join(dst, filepath.Base(arg))
		err = fsio.Copy(destination, false)(arg)
		if err != nil {
			return
		}
	}
	return
}

func copyRel(dst string, args ...string) (err error) {
	for _, arg := range args {
		fmt.Println(filepath.Join(dst, arg))
		destination := filepath.Join(dst, arg)
		err = fsio.Copy(destination, false)(arg)
		if err != nil {
			return
		}
	}
	return
}

func move(_ string, _ ...string) (err error) { return }
