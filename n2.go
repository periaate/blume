package blume

import (
	"os"
	"runtime"
	"io"
)

type IOR interface { CopyContents(bytes []byte) Result[int] }
type IOW interface { WriteContents(p any) Result[int] }
type IORW interface { IOR; IOW }

type Buffer struct { io.ReadWriter }

func (b Buffer) CopyContents(p []byte) Result[int] { return Auto(b.Read(p)) }
func (b Buffer) WriteContents(p any) Result[int] { return Auto(b.Write(Buf(p).Bytes())) }

func FromBuf(args ...any) Buffer { return Buffer{ Buf(args...) } }

func Create(args ...S) (res Result[Buffer]) {
	path := Path(args...)
	f, err := os.Create(path.String())
	if err != nil { return res.Fail(err) }
	runtime.AddCleanup(f, func(s S) { f.Close() }, path)
	return res.Pass(Buffer{f})
}
