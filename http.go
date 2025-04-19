package blume

import "net/http"

func (s S) Request() (res Result[String]) {
	req, err := http.Get(s.String())
	if err != nil {
		return res.Fail(err)
	}
	return res.Pass(S(Buf(req.Body).Bytes()))
}

func (s S) Requests() String { return s.Request().OrExit() }
