package core

import (
	"time"
)

func MustW[A, B any](fn func(A) (B, error)) func(A) B {
	return func(a A) B {
		b, err := fn(a)
		if err != nil {
			panic(err)
		}
		return b
	}
}

func IgnoreW[A, B, C any](fn func(A) (B, C)) func(A) B {
	return func(a A) B {
		b, _ := fn(a)
		return b
	}
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
func Ignore[T any](t T, err error) T { return t }

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
