package core

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func UsePipe(fn func(string)) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fn(scanner.Text())
		}
	}
}

func ReadPipe() (res []string) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}
	}
	return
}

func ReadRawPipe() (res []byte) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Bytes()...)
		}
	}
	return
}

func HasPipe() bool { return (Ignore(os.Stdin.Stat()).Mode() & os.ModeCharDevice) == 0 }

func Args() (res []string) { return append(os.Args[1:], ReadPipe()...) }

func WaitForKill() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func Prune(atLeast ...int) {
	if len(atLeast) == 0 {
		atLeast = append(atLeast, 1)
	}
	atl := atLeast[0] - 1
	os.Args = os.Args[1:]
	if len(os.Args) < atl {
		log.Fatalln("not enough arguments")
	}
}
