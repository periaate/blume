package blume

import (
	"os"
	"os/exec"
	"syscall"
)

type Cmd struct {
	Name String
	opts CmdOption
}

type CmdOption func(*exec.Cmd) *exec.Cmd

var CmdOpt CmdOption = func(cmd *exec.Cmd) *exec.Cmd { return cmd }

func (next CmdOption) Cwd(val String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		cmd.Dir = val.String()
		return cmd
	}
}

func (next CmdOption) Env(key, val String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		cmd.Env = append(cmd.Env, string(key+"="+val))
		return cmd
	}
}

func (next CmdOption) Pgid(val bool) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{}
		}
		cmd.SysProcAttr.Setpgid = val
		return cmd
	}
}

func (next CmdOption) Foreground(val bool) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{}
		}
		cmd.SysProcAttr.Foreground = val
		return cmd
	}
}

func (next CmdOption) AdoptEnv() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		cmd.Env = append(os.Environ(), cmd.Env...)
		return cmd
	}
}

func (next CmdOption) UserFacing() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		fn := next.Env("FORCE_COLOR", "true").
			Env("CLICOLOR_FORCE", "true")
		return fn(cmd)
	}
}

func (next CmdOption) Args(args ...String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		cmd.Args = append(cmd.Args, SD(args)...)
		return cmd
	}
}

func (next CmdOption) Adopt() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = next(cmd)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		return cmd
	}
}

func Exec(name String, opts ...func(*exec.Cmd) *exec.Cmd) (cmd Cmd) {
	if len(opts) == 0 {
		opts = append(opts, CmdOpt)
	}
	opt := Pipe(opts...)
	return Cmd{
		Name: name,
		opts: opt,
	}
}

func Execs(name String, opts ...func(*exec.Cmd) *exec.Cmd) Result[int] {
	return Exec(name, opts...).Exec()
}

func (c Cmd) Run() Result[String] {
	cmd := exec.Command(c.Name.String())
	cmd = c.opts(cmd)
	out, err := cmd.Output()
	if err != nil {
		return Err[String](err)
	}
	return Ok(String(out))
}

func (c Cmd) Exec() Result[int] {
	cmd := exec.Command(c.Name.String())
	cmd = c.opts(cmd)
	err := cmd.Start()
	if err != nil {
		return Err[int](err)
	}

	err = cmd.Wait()

	if err == nil {
		return Ok(0)
	}
	if exitError, ok := err.(*exec.ExitError); ok {
		if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
			return Ok(status.ExitStatus())
		}
	}
	return Err[int](err)
}

// func Exec[S ~string](name S, args ...S) *Command {
// 	return &Command{String(name), Map(StoS[S, String])(args), false}
// }
//
// func (c *Command) Append(args ...String) *Command {
// 	c.Args = append(c.Args, args...)
// 	return c
// }
//
// func (c *Command) Adopts() *Command {
// 	c.Adopt = true
// 	return c
// }
//
// func (c *Command) Create() *exec.Cmd {
// 	cmd := exec.Command(c.Name.String(), Map(StoD[String])(c.Args)...)
//     cmd.Env = append(os.Environ(), "FORCE_COLOR=true")
//     cmd.Env = append(cmd.Env, "CLICOLOR_FORCE=1")
// 	if c.Adopt {
// 		cmd.Stdout = os.Stdout
// 		cmd.Stdin = os.Stdin
// 		cmd.Stderr = os.Stderr
// 	}
// 	return cmd
// }
//
// func (c *Command) Run() Result[String] {
// 	bar, err := c.Create().Output()
// 	if err != nil { return Err[String](err) }
// 	return Ok(String(bar))
// }
//
// func Run[S ~string](name S, args ...S) Result[String] {
// 	bar, err := Exec(name, args...).Create().Output()
// 	if err != nil { return Err[String](err) }
// 	return Ok(String(bar).TrimSpace())
// }
//
// func Runs[S ~string](name S, args ...S) String {
// 	bar, _ := Exec(name, args...).Create().Output()
// 	return String(bar).TrimSpace()
// }

func Adopt(name String, args ...String) Result[int] {
	return Exec(name, CmdOpt.Args(args...), CmdOpt.Adopt().AdoptEnv()).Exec()
}
func Adopts(name String, args ...String) { Adopt(name, args...).Must() }
