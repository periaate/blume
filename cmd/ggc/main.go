package main

import (
	"fmt"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/yap"
)

func main() {
	values := Args(func(sar []string) bool { return len(sar) >= 2 }).Must().Value
	name := values[0].Replace("/", "")
	repo := values[1].Replace("/", "")
	rest := Arr(values[2:]...)
	path := String(fmt.Sprintf("git@github.com:%s/%s", name, repo))
	cmd := Exec("git")
	if rest.Len().Ge(1) {
		cmd.Append(rest.Value[0])
	} else {
		cmd.Append("clone")
	}
	cmd.Append(path)
	if rest.Len().Ge(2) {
		cmd.Append(rest.Value[1:]...)
	}
	yap.ErrFatal(cmd.Create().Run(), "error occurred running git wrapper")
}
