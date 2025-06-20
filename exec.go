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
	Name string
	opts CmdOption
}

type CmdOption func(*exec.Cmd) *exec.Cmd

var CmdOpt CmdOption = func(cmd *exec.Cmd) *exec.Cmd { return cmd }

func (prev CmdOption) Cd(val string) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
		cmd = prev(cmd)
		cmd.Dir = Path(val)
		return cmd
	}
}

func (prev CmdOption) Env(key, val string) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
		cmd = prev(cmd)
		cmd.Env = append(cmd.Env, string(key+"="+val))
		return cmd
	}
}

func (prev CmdOption) Sid(val bool) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
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
		if cmd == nil {return cmd}
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
		if cmd == nil { return cmd }
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
		if cmd == nil { return cmd }
		cmd = prev(cmd)
		cmd.Env = append(os.Environ(), cmd.Env...)
		return cmd
	}
}

func (prev CmdOption) UserFacing() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
		cmd = prev(cmd)
		fn := prev.Env("FORCE_COLOR", "true").
			Env("CLICOLOR_FORCE", "true")
		return fn(cmd)
	}
}

func (prev CmdOption) Args(args ...string) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
		cmd = prev(cmd)
		cmd.Args = append(cmd.Args, args...)
		return cmd
	}
}

func (prev CmdOption) Decorate(fn func(*exec.Cmd) *exec.Cmd) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
		return fn(prev(cmd))
	}
}

func (prev CmdOption) Adopt() CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
		cmd = prev(cmd)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		return cmd
	}
}

func (prev CmdOption) Signal(fn func(func(os.Signal) error)) CmdOption {
	return func(cmd *exec.Cmd) *exec.Cmd {
		if cmd == nil { return cmd }
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
						// Error.Println(err)
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

func Exec(name string, opts ...func(*exec.Cmd) *exec.Cmd) (cmd Cmd) {
	opt := func(cmd *exec.Cmd) *exec.Cmd {
		for _, opt := range opts {
			cmd = opt(cmd)
		}
		return cmd
	}
	return Cmd{ Name: name, opts: opt }
}

func Execs(name string, opts ...func(*exec.Cmd) *exec.Cmd) Result[int] { return Exec(name, opts...).Exec() }

func Run(name string, args ...string) Result[string] { return Exec(name, CmdOpt.AdoptEnv().Args(args...)).Run() }
func Runs(name string, args ...string) string { return Exec(name, CmdOpt.AdoptEnv().Args(args...)).Run().Must() }

func (c Cmd) Run() Result[string] {
	cmd := exec.Command(c.Name)
	cmd = c.opts(cmd)
	out, err := cmd.Output()
	if err != nil { return Err[string](err) }
	return Ok(string(out))
}

func (c Cmd) Exec() Result[int] {
	cmd := exec.Command(c.Name)
	cmd = c.opts(cmd)
	if err := cmd.Start(); err != nil { return Err[int](err) }
	if cmd.Wait() == nil { return Ok(0) }
	return Err[int](cmd.ProcessState.ExitCode())
}

func (c Cmd) Start() Result[*exec.Cmd] {
	cmd := exec.Command(c.Name)
	cmd = c.opts(cmd)
	err := cmd.Start()
	if err != nil { return Err[*exec.Cmd](err) }
	return Ok(cmd)
}

func Adopt(name string, args ...string) Result[int] { return Exec(name, CmdOpt.Args(args...), CmdOpt.Adopt().AdoptEnv().Foreground(true)).Exec() }
func Adopts(name string, args ...string) { Adopt(name, args...).Must() }
