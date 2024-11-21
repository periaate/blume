package fsio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/periaate/blume/gen"
)

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
	if err != nil {
		return false
	}
	return (a.Mode() & os.ModeCharDevice) == 0
}

// Args returns the command-line arguments without the program name, and including any piped inputs.
func Args() (res []string) { return append(os.Args[1:], ReadPipe()...) }

// Args returns the command-line arguments without the program name, and including any piped inputs.
func SepArgs() (res [2][]string) { return [2][]string{os.Args[1:], ReadPipe()} }

// QArgs returns the command-line arguments without the program name, and including any piped inputs.
// Returned type is an alias of []string which includes various helper functions.
// Helper functions will panic if they fail.
func QArgs() (res []Arg) {
	args := Args()
	res = make([]Arg, len(args))
	for i, arg := range args {
		res[i] = Arg(arg)
	}
	return
}

type Arg string

func (a Arg) String() string { return string(a) }
func (a Arg) Bytes() []byte  { return []byte(a) }
func (a Arg) Int() int       { return gen.Must(strconv.Atoi(string(a))) }

// func WaitForKill() {
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
// 	<-sigChan
// }
//
// func Prune(atLeast ...int) {
// 	if len(atLeast) == 0 {
// 		atLeast = append(atLeast, 1)
// 	}
// 	atl := atLeast[0] - 1
// 	os.Args = os.Args[1:]
// 	if len(os.Args) < atl {
// 		log.Fatalln("not enough arguments")
// 	}
// }

// GetArg attemts to get the index given as argument from os.Args. Returns empty string if OOB.
func GetArg(i int) (res string) {
	if len(os.Args) == 0 {
		return
	}
	args := os.Args[1:]
	if len(args) < i || len(args) == 0 {
		return
	}

	return args[i]
}

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
