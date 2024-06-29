package core

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func UsePipe(fn func(string)) {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fn(scanner.Text())
		}
	}
}

func ReadPipe() (res []string) {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}
	}
	return
}

func ReadRawPipe() (res []byte) {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Bytes()...)
		}
	}
	return
}

func HasPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) == 0
}

func Args() (res []string) {
	return append(os.Args[1:], ReadPipe()...)
}

// Filepaths returns all valid filepaths in the input array.
// Choice integer is a fitler for either files, directories, or both.
// 0: both, 1: files, -1: directories
func Filepaths(sar []string, choice int8) (res []string, errs []error) {
	for _, s := range sar {
		s, err := filepath.Abs(s)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		fi, err := os.Stat(s)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		switch choice {
		case 1:
			if !fi.Mode().IsRegular() {
				// errs = append(errs, fmt.Errorf("%s is not a file", s))
				continue
			}
		case -1:
			if !fi.Mode().IsDir() {
				// errs = append(errs, fmt.Errorf("%s is not a directory", s))
				continue
			}
		}

		res = append(res, s)
	}

	return
}

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
