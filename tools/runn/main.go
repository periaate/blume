package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/yap"
)

type Command struct {
	cmd   *exec.Cmd
	label string
}

func main() {
	yap.Configure(yap.Yapfig{
		ShowLevel: true,
	})
	fmt.Println("Running command(s)...")
	opts := fsio.Args[String](func(s []string) bool { return len(s) >= 1 })
	if !opts.Ok() { yap.Fatal("runn needs to be given at least one argument") }
	args := opts.Unwrap()

	sar := [][]String{}
	cur := []String{}
	for _, arg := range args.Val {
		if arg == "??" {
			sar = append(sar, cur)
			cur = []String{}
			continue
		}
		cur = append(cur, arg)
	}
	
	if len(cur) != 0 { sar = append(sar, cur) }

	
	res := Map[[]String, Command](func(inputs []String) (r Command) {
		if len(inputs) == 0 { return }
		label := string(inputs[0])
		cmd := string(inputs[1])
		inputs = inputs[2:]
		args := []string{}
		var setDir bool
		dir := ""
		for _, input := range inputs {
			if setDir {
				setDir = !setDir
				dir = string(input)
				continue
			}
			switch input {
			case "--cd", "-c": setDir = true
			default: args = append(args, string(input))
			}
		}
		yap.Debug("mapping input to command", "label", label, "cmd", cmd, "inputs", inputs, "args", args, "dir", dir)
		ecmd := exec.Command(cmd, args...)
		if len(dir) != 0 { ecmd.Dir = dir }
		return Command{
			cmd: ecmd,
			label: label,
		}
	})(sar)

	res = Filter(func(c Command) bool { return c.cmd != nil })(res)

	err := RunCommands(res...)
	if err != nil { yap.Fatal("error running commands", "error", err) }
}

// RunCommands executes a variadic number of exec.Cmd and streams all their outputs to stdout in real-time.
func RunCommands(cmds ...Command) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(cmds))

	for _, cmd := range cmds {
		wg.Add(1)
		go func(cm Command) {
			c := cm.cmd
			defer wg.Done()

			stdoutPipe, err := c.StdoutPipe()
			if err != nil {
				errChan <- fmt.Errorf("error creating stdout pipe: %w", err)
				return
			}

			stderrPipe, err := c.StderrPipe()
			if err != nil {
				errChan <- fmt.Errorf("error creating stderr pipe: %w", err)
				return
			}

			if err := c.Start(); err != nil {
				errChan <- fmt.Errorf("error starting command: %w", err)
				return
			}

			go streamOutput(stdoutPipe, cm.label, false)
			go streamOutput(stderrPipe, cm.label, true)

			if err := c.Wait(); err != nil { errChan <- fmt.Errorf("command finished with error: %w", err) }
		}(cmd)
	}
	
	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 { return fmt.Errorf("encountered errors: %v", errs) }

	return nil
}

// streamOutput streams the given reader line by line to stdout.
func streamOutput(reader io.Reader, label string, IsErr bool) {
	scanner := bufio.NewScanner(reader)
	pad := strings.Repeat(" ", Clamp(0, 1<<32)(len("tools") - len(label)))
	for scanner.Scan() {
		switch IsErr {
		case true: yap.Error(fmt.Sprintf("%s[%s] %s", pad, label, scanner.Text()))
		default: yap.Info(fmt.Sprintf("%s[%s] %s", pad, label, scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", label, err)
	}
}
