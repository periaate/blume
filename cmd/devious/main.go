package main

import (
	"os"
	"os/exec"

	"github.com/periaate/blume/cmd/devious/binary"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/yap"
)

var projType func(string, string, string)

func Go(root, entry, name string) {
	bin := binary.Binary(name)
	tar := fsio.Join("/bin", bin)
	cmd := exec.Command("go", "build", "-o", tar, entry)
	if debug {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}
	cmd.Dir = root
	cmd.Dir = root
	yap.ErrFatal(cmd.Run(), "couldn't run command", "err")
	Done(tar, name)
}

func Rust(root, entry, name string) {
	bin := binary.Binary(name)
	cargoTarget := fsio.Join(root, "target", "debug", binary.Binary(name))
	tar := fsio.Join("/bin", bin)
	cmd := exec.Command("cargo", "build")
	if debug {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}
	cmd.Dir = root
	yap.ErrFatal(cmd.Run(), "couldn't run command", "err")

	if err := fsio.Copy(tar, cargoTarget, true); err != nil { panic(err) }
	Done(tar, name)
}

func Done(tar, name string) {
	if fsio.Exists(tar) {
		yap.Info("build succeeded!", "compiled", String(name), "to", String(tar))
	} else {
		yap.Error("build failed...", "couldn't compile", name, "to", tar)
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
	args := fsio.Args[String](func(s []string) bool { return len(s) >= 1 }).Must()

	debugOpt := args.First(func(s String) bool { return Is("-d", "--debug")(s.String()) })
	if debugOpt.Ok {
		yap.SetLevel(yap.L_Debug)
		debug = true
	}

	arg := args.Shift().Must()

	found := fsio.FindFirst("C:/github.com/",
		func(s String) bool {
			return !s.Contains("/.", "node_modules", "target", "build", "data", "Modules", "mpv.net")
		},
		fsio.IsDir,
		func(f String) bool { return fsio.Base(f) == arg },
	).Must()
	entries, err := fsio.ReadDir(found)
	if err != nil { panic(err) }
	entry := entries.First(IsEntry).Or("")

	yap.Debug("looking for entry in match", entry)

	root := fsio.Ascend(found, IsProject[String]).Must()
	root = fsio.Clean(fsio.Dir(root))
	yap.Debug("found", root)

	projType(root.String(), entry.String(), arg.String())
}

func IsProject[S ~string](s S) (res bool) {
	if Is("Cargo.toml", "go.mod", "package.json")(string(fsio.Base(s))) {
		switch string(fsio.Base(s)) {
		case "Cargo.toml": projType = Rust
		case "go.mod": projType = Go
		case "package.json": projType = Go
		}
		return true
	}

	result, err := fsio.ReadDir(s)
	if err != nil { panic(err) }
	r := result.Filter(func(fp S) bool {
		return Is("Cargo.toml", "go.mod", "package.json")(string(fsio.Base(fp)))
	})
	yap.Debug("in [IsProject], ran filter", "len", r.Len())

	return r.Len() > 0
}

func With[A, B any](transform func(A) B) func(fns ...func(B) bool) func(A) bool {
	return func(fns ...func(B) bool) func(A) bool {
		pred := PredOr(fns...)
		return func(input A) bool { return pred(transform(input)) }
	}
}

func IsEntry[S ~string](s S) bool {
	match := fsio.Base(s)
	return Is("main.go", "Cargo.toml")(string(match))
}
