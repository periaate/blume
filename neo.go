package blume

import (
	"os"
	"time"
)

func Auto[A any](value A, handles ...any) Result[A] {
	if IsOk(value, handles...) {
		return Ok(value)
	}
	return Err[A](handles...)
}

func (s String) Stat() Result[os.FileInfo] { return Auto(os.Stat(s.String())) }
func (s String) ModTime() Result[time.Time] {
	r := s.Stat()
	if !r.IsOk() {
		return Err[time.Time](r.Other)
	}
	return Ok(r.Value.ModTime())
}

func (s String) ModMilli() Result[int64] {
	r := s.Stat()
	if !r.IsOk() {
		return Err[int64](r.Other)
	}
	return Ok(r.Value.ModTime().UnixMilli())
}
