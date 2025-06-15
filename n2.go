package blume

import (
	"io"
	"os"
	"reflect"
	"runtime"
)

type Buffer struct { io.ReadWriter }

func (b Buffer) CopyContents(p []byte) (res Result[int]) { return res.Auto(b.Read(p)) }
func (b Buffer) WriteContents(p any) Result[int] {
	n, err := b.Write(Buf(p).Bytes())
	if err != nil { return Err[int](err) }
	return Ok(n)
}

func FromBuf(args ...any) Buffer { return Buffer{ Buf(args...) } }

func Create(args ...string) (res Result[Buffer]) {
	path := Path(args...)
	f, err := os.Create(path)
	if err != nil { return res.Fail(err) }
	runtime.AddCleanup(f, func(s S) { f.Close() }, path)
	return res.Pass(Buffer{f})
}

type Func struct {
	Takes int
	Returns int
	Variadic bool
	Val any
	reflect.Value
	reflect.Type
	IsOption bool
	IsResult bool
	Then *Func
}

func (fn Func) Call(args ...any) (res []reflect.Value, ok bool) {
	inputs := Map[any, reflect.Value](reflect.ValueOf)(args)
	if len(inputs) < 2 {
		res = fn.Value.Call(inputs)
		if fn.Then != nil { return fn.Then.Call(Any(res)...) }
		return res, true
	}

	switch val := inputs[len(inputs)-1].Interface().(type) {
	case error: if val != nil { return }; inputs = inputs[:len(inputs)-1]
	case bool: if !val        { return }; inputs = inputs[:len(inputs)-1] }

	res = fn.Value.Call(inputs)
	if fn.Then != nil { return fn.Then.Call(Any(res)...) }
	return res, true
}

var err_type reflect.Type

func init() {
	var err error
	err_type = reflect.TypeOf(err)
}

func Function(fn any) (res Result[Func]) {
	fun := Func{}
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func { return res.Fail("function must be a function type") }
	var lastOut reflect.Type

	if no := fnType.NumOut(); no >= 2 {
		lastOut = fnType.Out(no-1)
		switch {
		case lastOut.AssignableTo(err_type): fun.IsResult = true
		case lastOut.Kind() == reflect.Bool: fun.IsOption = true
		}
	}

	fun.Takes = fnType.NumIn()
	fun.Returns = fnType.NumOut()
	fun.Variadic = fnType.IsVariadic()
	fun.Val = fn
	fun.Value = reflect.ValueOf(fn)
	fun.Type = fnType

	return res.Pass(fun)
}

