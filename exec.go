package blume

import (
	"os"
	"os/exec"
)

type Command struct {
	Name String
	Args []String

	Adopt bool
}

func Exec[S ~string](name S, args ...S) *Command {
	return &Command{String(name), Map(StoS[S, String])(args), false}
}

func (c *Command) Append(args ...String) *Command {
	c.Args = append(c.Args, args...)
	return c
}

func (c *Command) Adopts() *Command {
	c.Adopt = true
	return c
}

func (c *Command) Create() *exec.Cmd {
	cmd := exec.Command(c.Name.String(), Map(StoD[String])(c.Args)...)
	cmd.Env = append(os.Environ(), "FORCE_COLOR=true")
	cmd.Env = append(os.Environ(), "CLICOLOR_FORCE=1")
	if c.Adopt {
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
	}
	return cmd
}

func (c *Command) Run() Result[String] {
	bar, err := c.Create().Output()
	if err != nil {
		return Err[String](err)
	}
	return Ok(String(bar))
}

func Run[S ~string](name S, args ...S) Result[String] {
	bar, err := Exec(name, args...).Create().Output()
	if err != nil {
		return Err[String](err)
	}
	return Ok(String(bar).TrimSpace())
}

func Runs[S ~string](name S, args ...S) String {
	bar, _ := Exec(name, args...).Create().Output()
	return String(bar).TrimSpace()
}
