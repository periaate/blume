package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/periaate/blume"
	"github.com/periaate/blume/cmd/devious/binary"
	"github.com/periaate/blume/pred"
	"github.com/periaate/blume/pred/filter"
	"github.com/periaate/blume/pred/is"

	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/yap"
)

var projType func(string, string, string)

func Go(root, entry, name string) {
	bin := binary.Binary(name)
	tar := filepath.Join("~/.blume/bin", bin)
	yap.Info("building for Go", tar)
	cmd := exec.Command("go", "build", "-o", tar, entry)
	if debug {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}
	cmd.Dir = root
	yap.ErrFatal(cmd.Run(), "couldn't run command", "err")
	Done(root, entry, tar, name)
}

func Rust(root, entry, name string) {
	bin := binary.Binary(name)
	cargoTarget := filepath.Join(root, "target", "debug", binary.Binary(name))
	tar := filepath.Join("~.blume/bin", bin)
	cmd := exec.Command("cargo", "build")
	if debug {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}
	cmd.Dir = root
	yap.ErrFatal(cmd.Run(), "couldn't run command", "err")

	if err := fsio.Copy(tar, cargoTarget, true); err != nil {
		panic(err)
	}
	Done(root, entry, tar, name)
}

func Done(root, entry, tar, name string) {
	if fsio.Exists(tar) {
		switch debug {
		case true:
			yap.Info("build succeeded!", "compiled", name, "to", tar, "root", root, "entry", entry)
		default:
			yap.Info("build succeeded!", "compiled", name, "to", tar)
		}
	} else {
		yap.Error("build failed...", "couldn't compile", name, "to", tar, "root", root, "entry", entry)
	}
}

var debug bool

func main() {
	yap.Configure(yap.Yapfig{
		ShowFile:  false,
		ShowLevel: false,
		ShowTime:  false,
		Level:     yap.L_Info,
	})
	args := blume.Must(fsio.Args(is.NotEmpty[string]))

	// debugOpt := args.First(func(s String) bool { return Is("-d", "--debug")(s.String()) })
	// if debugOpt.Ok {
	// 	yap.SetLevel(yap.L_Debug)
	// 	debug = true
	// }

	arg := args[0]

	found := blume.Must(fsio.First("~/github.com", pred.And(
		func(s string) bool {
			return !is.Any("/.", "node_modules", "target", "build", "data", "Modules", "mpv.net")(s)
		},
		fsio.IsDir,
		func(f string) bool { return filepath.Base(f) == arg },
	)))
	entry, _ := fsio.First(found, IsEntry)

	yap.Debug("looking for entry in match", entry)

	root := blume.Must(fsio.Ascend(found, IsProject))
	root = filepath.Clean(filepath.Dir(root))
	yap.Debug("found", root)

	projType(root, entry, arg)
}

func IsProject(s string) (res bool) {
	if is.Any("Cargo.toml", "go.mod", "package.json")(filepath.Base(s)) {
		switch filepath.Base(s) {
		case "Cargo.toml":
			projType = Rust
		case "go.mod":
			projType = Go
		case "package.json":
			projType = Go
		}
		return true
	}

	result, err := fsio.ReadDir(s)
	if err != nil {
		return false
	}
	r := filter.Filter(func(fp fsio.Entry) bool {
		return is.Any("Cargo.toml", "go.mod", "package.json")(filepath.Base(fp.Path()))
	})(result)
	yap.Debug("in [IsProject], ran filter", "len", len(r))

	return len(r) > 0
}

func With[A, B any](transform func(A) B) func(fns ...func(B) bool) func(A) bool {
	return func(fns ...func(B) bool) func(A) bool {
		pred := pred.Or(fns...)
		return func(input A) bool { return pred(transform(input)) }
	}
}

func IsEntry(s string) bool {
	match := filepath.Base(s)
	return is.Any("main.go", "Cargo.toml")(match)
}
