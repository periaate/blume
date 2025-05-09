package blume

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s S) Request() (res Result[String]) {
	req, err := http.Get(s.String())
	if err != nil { return Err[S](err) }
	if req.StatusCode >= 400 { return Err[S]("bad status code: ", req.StatusCode, " ", req.Status) }
	return res.Pass(S(Buf(req.Body).Bytes()))
}

func (s S) Requests() String { return s.Request().OrExit() }

type Bytes []byte

func (b Bytes) Buf() *bytes.Buffer { return Buf(b) }
func (b Bytes) S() S               { return String(b) }
func (b Bytes) String() string     { return string(b) }


func Request[A any](method, url S, a A) (res Result[Bytes]) {
	buf := Buf()
	if err := json.NewEncoder(buf).Encode(a); err != nil { return res.Fail(err) }
	req, err := http.NewRequest(method.String(), url.String(), buf)
	if err != nil { return res.Fail(err) }
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil { return res.Fail(err) }
	defer resp.Body.Close()
	return res.Pass(Buf(resp.Body).Bytes())
}
