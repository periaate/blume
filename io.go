package blume

import "io"

type Reader func(p []byte) (n int, err error)

func (r Reader) Read(p []byte) (n int, err error) { return r(p) }

type Writer func(p []byte) (n int, err error)

func (w Writer) Write(p []byte) (n int, err error) { return w(p) }

var _ io.Reader = Reader(nil)
var _ io.Writer = Writer(nil)

func FanInWriter(w io.Writer) (func(), func() io.Writer) {
	ch := make(chan []byte)
	var closed bool

	newWriter := func() io.Writer {
		return Writer(func(p []byte) (int, error) {
			if !closed {
				ch <- p
				return len(p), nil
			} else {
				return 0, io.ErrClosedPipe
			}
		})
	}

	go func() {
		for !closed {
			b := <-ch
			if closed {
				return
			}
			if _, err := w.Write(b); err != nil {
				return
			}
		}
	}()

	shutdown := func() {
		closed = true
		close(ch)
	}

	return shutdown, newWriter
}
