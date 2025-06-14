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
}

func (fn Func) Call(args ...any) []reflect.Value { return fn.Value.Call(Map[any, reflect.Value](reflect.ValueOf)(args)) }

func (f Func) Matches(t any) bool { return Function(t).AssignableTo(f.Type) }

func Function(fn any) (res Func) {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func { panic("function must be a function type") }
	return Func{
		Takes: fnType.NumIn(),
		Returns: fnType.NumOut(),
		Variadic: fnType.IsVariadic(),
		Val: fn,
		Value: reflect.ValueOf(fn),
		Type: fnType,
	}
}

