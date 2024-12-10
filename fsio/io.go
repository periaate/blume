package fsio

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/gen"
)


func B(args ...any) *bytes.Buffer {
	_, arg, _ := gen.Shifts(args)
	switch v := arg.(type) {
	case string: return bytes.NewBufferString(v)
	case []byte: return bytes.NewBuffer(v)
	default: return bytes.NewBuffer([]byte{})
	}
}

// UsePipe reads from stdin and calls the given function for each line.
func UsePipe(fn func(string)) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fn(scanner.Text())
		}
	}
}

// ReadPipe reads from stdin and returns a slice of lines.
func ReadPipe() (res []string) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}
	}
	return
}

// ReadRawPipe reads from stdin and returns a slice of bytes.
func ReadRawPipe() (res []byte) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Bytes()...)
		}
	}
	return
}

// HasPipe evaluates whether stdin is being piped in to.
func HasPipe() bool {
	a, err := os.Stdin.Stat()
	if err != nil { return false }
	return (a.Mode() & os.ModeCharDevice) == 0
}

// HasPipe evaluates whether stdin is being piped in to.
func HasOutPipe() bool {
	a, err := os.Stdout.Stat()
	if err != nil { return false }
	return (a.Mode() & os.ModeCharDevice) == 0
}

// Args returns the command-line arguments without the program name, and including any piped inputs.
func Args(opts ...gen.Condition[[]string]) Option[Array[string]] {
	args := append(os.Args[1:], ReadPipe()...)
	for _, opt := range opts {
		err := opt(args)
		if err != nil { return None[Array[string]](err) }
	}
	return Some(ToArray(args))
}

// Args returns the command-line arguments without the program name, and including any piped inputs.
func SepArgs() (res [2][]string) { return [2][]string{os.Args[1:], ReadPipe()} }

func QArgs(opts ...gen.Condition[[]string])  Option[Array[gen.String]] {
	args, err := Args(opts...).Values()
	if err != nil { return None[Array[gen.String]](err) }
	return Some[Array[gen.String]](ToArray(MapStoS[string, gen.String](args.Values()...)))
}

func StoS[A, B ~string](a A) B { return B(a) } 
func MapStoS[A, B ~string](a ...A) []B { return Map[A, B](StoS)(a) }

func Pipes() (input, output chan string) {
	input = make(chan string)
	output = make(chan string)
	go func() {
		args := os.Args[1:]
		for _, arg := range args {
			input <- arg
		}

		if HasPipe() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				input <- scanner.Text()
			}
		}
	}()

	go func() {
		for {
			a := <-output
			fmt.Println(a)
		}
	}()

	return
}
