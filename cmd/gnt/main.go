package main

import (
	"os"
	"os/exec"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/pred/has"
	"github.com/periaate/blume/pred/is"
)

func main() {
	if has.Pipe(os.Stdin) {
		Piped(os.Stdin).Value.Filter(func(s String) bool {
			return Dir(s).Read().Value.First(is.GitRoot).IsOk()
		}).Each(Logs)
		return
	}
	Dir(Args().Must().Get(0).Must()).
		Find(func(s string) bool {
			if res := Dir(s).Read(); res.IsOk() {
				return res.Value.First(is.GitRoot).IsOk()
			}
			return false
		}).
		Must().
		Each(Logs)
}

func Exec[S ~string](name string, args ...S) Command {
	return Command{name: name, args: Map(func(s S) string { return string(s) })(args)}
}

type Command struct {
	name string
	args []string
	*exec.Cmd
}

func (c Command) Args(args ...string) Command {
	c.args = append(c.args, args...)
	return c
}

// func (c Command) Err(func([]string)) Command
// func (c Command) Out(func([]string)) Command
// func (c Command) Pipe(func(string) bool) Command

// func (c Command) Start() Command
// func (c Command) Wait() Command
// func (c Command) Then() Command

func (c Command) Run(ch ...<-chan any) Result[Array[String]] {
	c.Cmd = exec.Command(c.name, c.args...)
	if res, err := c.Cmd.Output(); err == nil {
		return Ok(ToArray(Map(Stringify)(Lines(res))))
	} else {
		return Err[Array[String]](err.Error())
	}
}
