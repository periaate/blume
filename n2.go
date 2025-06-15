package blume

import (
	"fmt"
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

type Value struct {
	reflect.Value
	Then *Value
}

func (val *Value) IsFunc  () bool { return val.Kind() == reflect.Func   }
func (val *Value) IsBool  () bool { return val.Kind() == reflect.Bool   }
func (val *Value) IsInt   () bool { return val.Kind() == reflect.Int    }
func (val *Value) IsUint  () bool { return val.Kind() == reflect.Uint   }
func (val *Value) IsUint8 () bool { return val.Kind() == reflect.Uint8  }
func (val *Value) IsUint16() bool { return val.Kind() == reflect.Uint16 }
func (val *Value) IsUint32() bool { return val.Kind() == reflect.Uint32 }
func (val *Value) IsUint64() bool { return val.Kind() == reflect.Uint64 }
func (val *Value) IsInt8  () bool { return val.Kind() == reflect.Int8   }
func (val *Value) IsInt16 () bool { return val.Kind() == reflect.Int16  }
func (val *Value) IsInt32 () bool { return val.Kind() == reflect.Int32  }
func (val *Value) IsInt64 () bool { return val.Kind() == reflect.Int64  }

func (val *Value) Call(args ...any) (res []reflect.Value, err error) { return val.call(Map[any, reflect.Value](reflect.ValueOf)(args), false) }

func (val *Value) call(inputs []reflect.Value, fromPrevious bool) (res []reflect.Value, err error) {
	if val.Kind() != reflect.Func {
		if val.Then == nil || !val.Then.IsFunc() {
			err = fmt.Errorf("attempting to call a non-function [%s] as a function", val.String())
			return
		}
		return val.Then.call([]reflect.Value{val.Value}, false)
	}

	if len(inputs) < 2 {
		res = val.Value.Call(inputs)
		if val.Then != nil { return val.Then.call(res, true) }
		return res, nil
	}

	if fromPrevious {
		switch inp := inputs[len(inputs)-1].Interface().(type) {
		case error: if inp != nil { return nil, inp }; inputs = inputs[:len(inputs)-1]
		case bool: if !inp        { return nil, fmt.Errorf("in call to [%s] the last argument [ok != true]", val.String()) }; inputs = inputs[:len(inputs)-1] }
	}

	res = val.Value.Call(inputs)
	if val.Then != nil { return val.Then.call(res, true) }
	return res, nil
}

func Function(input any) (res *Value) { return &Value{ Value: reflect.ValueOf(input) } }

