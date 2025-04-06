package blume

import "compress/gzip"

func (s S) Unarchive() (res Result[String]) {
	r := s.Read()
	if !r.IsOk() {
		return r
	}
	buf := Buf([]byte(r.Value))
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return res.Fail(err)
	}
	defer reader.Close()
	return Ok(S(Buf(reader).Bytes()))
}

func (s S) Unarchives() (res String) { return s.Unarchive().Must() }

func Unarchive(s S) (res Result[String]) {
	r := s.Read()
	if !r.IsOk() {
		return r
	}
	buf := Buf([]byte(r.Value))
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return res.Fail(err)
	}
	defer reader.Close()
	return Ok(S(Buf(reader).Bytes()))
}

func Unarchives(s S) (res String) { return s.Unarchive().Must() }
