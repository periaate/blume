package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/periaate/blume/cmd/blumefmt"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
)

func main() {
	var input []byte
	var err error
	if fsio.HasPipe() {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("gofmt", "-s")
		inPipe, err := cmd.StdinPipe()
		if err != nil {
			os.Stdout.Write(input)
		}
		buf := Buf()
		cmd.Stdout = buf
		cmd.Start()

		go func() {
			defer inPipe.Close()
			inPipe.Write(input)
		}()

		if err := cmd.Wait(); err != nil {
			return
		}

		input = buf.Bytes()
	} else {
		args := fsio.IArgs[string](func(s []string) bool { return len(s) >= 1 }).Must()
		cmd := exec.Command("gofmt", append([]string{"-s"}, args.Val...)...)
		buf := Buf()
		cmd.Stdout = buf
		cmd.Run()
		input = buf.Bytes()
	}

	res, err := blumefmt.Fmt(input)
	if err != nil {
		res = input
	}
	os.Stdout.Write(res)
}
