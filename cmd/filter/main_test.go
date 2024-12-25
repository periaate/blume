package main

import (
	"fmt"
	"testing"

	. "github.com/periaate/blume"
)

func Expect(t *testing.T) func(fn func([]string) bool) func(values []string) []string {
	return func(fn func([]string) bool) func(values []string) []string {
		return func(values []string) []string {
			if !fn(values) {
				t.Fail()
			}
			return values
		}
	}
}

func For[A any](fn func(A)) func([]A) {
	return func(arr []A) {
		for _, v := range arr {
			fn(v)
		}
	}
}

// Just is composition of [All] and [Is].
func Just[C comparable](A ...C) func([]C) bool { return All(Is(A...)) }
func Ignore[A, B any](fn func(A) B) func(A)    { return func(a A) { fn(a) } }

func After[A, B, C any](fn func(A) func(B) C) func(after func(C)) func(A) func(B) C {
	return func(after func(C)) func(A) func(B) C {
		return func(a A) func(B) C {
			return func(b B) C {
				res := fn(a)(b)
				after(res)
				return res
			}
		}
	}
}

func TestFilter(t *testing.T) {
	expect := After(Expect(t))(For(func(str string) { fmt.Println(str) }))
	After(Parse)(Ignore(expect(Just("hello"))))(Join("is", "hello"))(Join("Hello", "World", "hello", "world"))
	// expect(Just("hello"))(Parse("is", "hello")("Hello", "World", "hello", "world"))
}

func Join[A any](args ...A) []A { return args }
