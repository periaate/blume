package blume

import (
	"os"
	"os/exec"
)

type Command struct {
	Name String
	Args []String
}

func Exec[S ~string](name S, args ...S) *Command {
	return &Command{String(name), Map(StoS[S, String])(args)}
}

func (c *Command) Append(args ...String) {
	c.Args = append(c.Args, args...)
}

func (c *Command) Create() *exec.Cmd {
	cmd := exec.Command(c.Name.String(), Map(StoD[String])(c.Args)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd
}

func (c *Command) Run() error {
	cmd := exec.Command(c.Name.String(), Map(StoD[String])(c.Args)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Command) Silent() error {
	cmd := exec.Command(c.Name.String(), Map(StoD[String])(c.Args)...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
