package fsio

import (
	"bufio"
	"os"
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
