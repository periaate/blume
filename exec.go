package blume

import (
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

const (
	SIGTERM = syscall.SIGTERM
	SIGKILL = syscall.SIGKILL
	SIGINT  = syscall.SIGINT
)

type Cmd struct {
	Name String
	opts CmdOption
}

type CmdOption func(*exec.Cmd) *exec.Cmd

var CmdOpt CmdOption = func(cmd *exec.Cmd) *exec.Cmd { return cmd }

// Cwd [Deprecated] use [Cd] instead
func (prev CmdOption) Cwd(val String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Dir = val.String()
		return cmd
	}
}

func (prev CmdOption) Cd(val String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Dir = val.Path().String()
		return cmd
	}
}

func (prev CmdOption) Env(key, val String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Env = append(cmd.Env, string(key+"="+val))
		return cmd
	}
}

func (prev CmdOption) Sid(val bool) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{}
		}
		cmd.SysProcAttr.Setsid = val
		return cmd
	}
}

func (prev CmdOption) Pgid(val bool) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{}
		}
		cmd.SysProcAttr.Setpgid = val
		return cmd
	}
}

func (prev CmdOption) Foreground(val bool) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{}
		}
		cmd.SysProcAttr.Foreground = val
		return cmd
	}
}

func (prev CmdOption) AdoptEnv() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Env = append(os.Environ(), cmd.Env...)
		return cmd
	}
}

func (prev CmdOption) UserFacing() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		fn := prev.Env("FORCE_COLOR", "true").
			Env("CLICOLOR_FORCE", "true")
		return fn(cmd)
	}
}

func SD(s []S) []string {
	res := make([]string, 0, len(s))
	for _, el := range s { res = append(res, string(el)) }
	return res
}

func (prev CmdOption) Args(args ...String) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Args = append(cmd.Args, SD(args)...)
		return cmd
	}
}

func (prev CmdOption) Decorate(fn func(*exec.Cmd) *exec.Cmd) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		return fn(prev(cmd))
	}
}

func (prev CmdOption) Adopt() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		return cmd
	}
}

func (prev CmdOption) Signal(fn func(func(os.Signal) error)) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		fn(func(s os.Signal) error {
			err := cmd.Process.Signal(s)
			cmd.Process.Wait()
			return err
		})
		return cmd
	}
}

func Sig(signals ...os.Signal) chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signals...)
	return sigChan
}

func Signal(signals ...os.Signal) func(func(os.Signal) error) {
	fns := []func(os.Signal) error{}
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, signals...)
		sig := <-sigChan
		wg := sync.WaitGroup{}
		wg.Add(len(fns))
		for _, fn := range fns {
			go func() {
				defer wg.Done()
				if fn != nil {
					if err := fn(sig); err != nil {
						Error.Println(err)
					}
				}
			}()
		}
		wg.Wait()
		os.Exit(0)
	}()

	return func(fn func(os.Signal) error) {
		fns = append(fns, fn)
	}
}

func (prev CmdOption) Stdout(w io.Writer) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Stdout = w
		return cmd
	}
}
func (prev CmdOption) Stderr(w io.Writer) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Stderr = w
		return cmd
	}
}
func (prev CmdOption) Stdin(r io.Reader) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil {
			return cmd
		}
		cmd = prev(cmd)
		cmd.Stdin = r
		return cmd
	}
}

func Any[T any](a []T) (res []any) { for _, v := range a { res = append(res, any(v)) }; return }

func Exec(name String, opts ...func(*exec.Cmd) *exec.Cmd) (cmd Cmd) {
	if len(opts) == 0 { opts = append(opts, CmdOpt) }
	return Cmd{ Name: name, opts: Pipe[CmdOption](Any(opts)...) }
}

func Execs(name String, opts ...func(*exec.Cmd) *exec.Cmd) Result[int] { return Exec(name, opts...).Exec() }

func Run(name String, args ...String) Result[String] { return Exec(name, CmdOpt.AdoptEnv().Args(args...)).Run() }
func Runs(name String, args ...String) String { return Exec(name, CmdOpt.AdoptEnv().Args(args...)).Run().Must() }

func (c Cmd) Run() Result[String] {
	cmd := exec.Command(c.Name.String())
	cmd = c.opts(cmd)
	out, err := cmd.Output()
	if err != nil { return Err[String](err) }
	return Ok(String(out))
}

func (c Cmd) Exec() Result[int] {
	cmd := exec.Command(c.Name.String())
	cmd = c.opts(cmd)
	if err := cmd.Start(); err != nil { return Err[int](err) }
	if cmd.Wait() == nil { return Ok(0) }
	return Err[int](cmd.ProcessState.ExitCode())
}

func (c Cmd) Start() Result[*exec.Cmd] {
	cmd := exec.Command(c.Name.String())
	cmd = c.opts(cmd)
	err := cmd.Start()
	if err != nil { return Err[*exec.Cmd](err) }
	return Ok(cmd)
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
	return Exec(name, CmdOpt.Args(args...), CmdOpt.Adopt().AdoptEnv().Foreground(true)).Exec()
}
func Adopts(name String, args ...String) { Adopt(name, args...).Must() }
