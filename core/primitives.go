package core

import (
	"time"
)

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Repeat(dur time.Duration, fn func()) (stop func()) {
	stopch := make(chan struct{})
	var stopped bool
	go func() {
		tck := time.NewTicker(dur)
		defer tck.Stop()
		defer close(stopch)
		defer func() { stopped = true }()
		for {
			select {
			case <-stopch:
				return
			case <-tck.C:
				go fn()
			}
		}
	}()

	return func() {
		if stopped {
			return
		}
		stopch <- struct{}{}
	}
}
