package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/periaate/blume"
	"github.com/periaate/blume/cmd/blumefmt"
	"github.com/periaate/blume/pred/is"

	"github.com/periaate/blume/fsio"
)

func main() {
	var input []byte
	var err error
	if fsio.HasInPipe() {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("gofmt", "-s")
		inPipe, err := cmd.StdinPipe()
		if err != nil {
			os.Stdout.Write(input)
		}
		buf := blume.Buf()
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
		args := blume.Must(fsio.Args(is.NotEmpty[string]))
		cmd := exec.Command("gofmt", append([]string{"-s"}, args...)...)
		buf := blume.Buf()
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
