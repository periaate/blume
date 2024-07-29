package core

import (
	"time"
)

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func Ignore[A, B any](a A, _ B) A                      { return a }
func Ignores[A, B, C any](fn func(A) (B, C)) func(A) B { return func(a A) B { return Ignore(fn(a)) } }

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
