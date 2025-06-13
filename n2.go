package blume

import (
	"io"
	"os"
	"reflect"
	"runtime"

	// "github.com/periaate/blume/pred"
)

type IOR interface { CopyContents(bytes []byte) Result[int] }
type IOW interface { WriteContents(p any) Result[int] }
type IORW interface { IOR; IOW }

type Buffer struct { io.ReadWriter }

func (b Buffer) CopyContents(p []byte) Result[int] { return Auto(b.Read(p)) }
func (b Buffer) WriteContents(p any) Result[int] {
	n, err := b.Write(Buf(p).Bytes())
	if err != nil { return Err[int](err) }
	return Ok(n)
}

func FromBuf(args ...any) Buffer { return Buffer{ Buf(args...) } }

func Create(args ...S) (res Result[Buffer]) {
	path := Path(args...)
	f, err := os.Create(path.String())
	if err != nil { return res.Fail(err) }
	runtime.AddCleanup(f, func(s S) { f.Close() }, path)
	return res.Pass(Buffer{f})
}

// Const creates a constant function that pre-fills some arguments
func Const[WantedFunctionType any](actualFunction any, constantArgs ...any) WantedFunctionType {
	var zero WantedFunctionType
	wantedType := reflect.TypeOf(zero)

	if wantedType.Kind() != reflect.Func { panic("WantedFunctionType must be a function type") }
	actualValue := reflect.ValueOf(actualFunction)
	actualType := actualValue.Type()

	if actualType.Kind() != reflect.Func { panic("actualFunction must be a function") }
	if len(constantArgs) > actualType.NumIn() { panic("too many constant arguments") }

	// // The new function will take remaining arguments after constants
	// remainingParamCount := actualType.NumIn() - len(constantArgs)
	// if wantedType.NumIn() != remainingParamCount {
	// 	panic(P.F("parameter count mismatch: wanted %d, got %d",
	// 		wantedType.NumIn(), remainingParamCount))
	// }

	// Return types must match exactly
	// if wantedType.NumOut() != actualType.NumOut() { panic("return type count mismatch") }

	for i := range wantedType.NumOut() {
		if wantedType.Out(i) != actualType.Out(i) {
			panic(P.F("return type mismatch at position %d", i))
		}
	}

	// Convert constant args to reflect.Values
	constValues := make([]reflect.Value, len(constantArgs))
	for i, arg := range constantArgs {
		constValues[i] = reflect.ValueOf(arg)
	}

	dynamicFunc := reflect.MakeFunc(wantedType, func(_ []reflect.Value) []reflect.Value {
		return actualValue.Call(constValues)
	})

	return dynamicFunc.Interface().(WantedFunctionType)
}

type Func struct {
	Takes int
	Returns int
	Variadic bool
	Val any
	reflect.Value
	reflect.Type
}

func (fn Func) Call(args ...any) []reflect.Value { return fn.Value.Call(Map(reflect.ValueOf)(args)) }

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

