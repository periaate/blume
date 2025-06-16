package blume

import (
	"fmt"
	"os"
	"reflect"
)

func (e Either[A, B]) Is(value A) bool {
	if !e.IsOk()                           { return false }
	if !reflect.TypeOf(value).Comparable() { return false }
	return reflect.ValueOf(e.Value).Equal(reflect.ValueOf(value))
}

type Op[I, O any] interface { func(I) func(I) O | func(I, I) O }

func Pattern[T any, Fn Op[T, bool]](pred Fn, args ...T) (ok bool) {
	if len(args) == 0 { return }
	var op func(T, T) bool
	switch fn := any(pred).(type) {
	case func(T) Pred[T]: op = func(t1, t2 T) bool { return fn(t2)(t1) }
	case func(T, T) bool: op = fn }
	for i, arg := range args[1:] {
		if !op(args[i], arg) { return }
	}
	return true
}

// Exit the program with a console log
func Exit(args ...any) {
	fmt.Printf("%s\n", Join(" ")(args))
	os.Exit(1)
}

func ExitWith(n int, args ...any) { fmt.Printf("%s", Join(" ")(args)); os.Exit(n) }

func ExitsWith[A any](n int) func(arg A) A { return func(arg A) A { ExitWith(n, arg); return arg } }

func OrExit[A, B any](either Either[A, B], args ...any) (res A) {
	if !either.IsOk() { Exit(fmt.Sprintf("%s [%v]", fmt.Sprint(args...), either.Other)) }
	return either.Value
}

func OrExits[A, B any](either Either[A, B]) (res A) {
	if !either.IsOk() { Exit(fmt.Sprintf("%v", either.Other)) }
	return either.Value
}
