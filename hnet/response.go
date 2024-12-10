package hnet

import (
	"fmt"
	"io"
	"net/http"

	. "github.com/periaate/blume/core"
)

type Response struct {
	*http.Response
	NetError
}

func (r Response) Assert(conds ...Predicate[any]) Response {
	if r.NetError != nil { return r }

	for _, cond := range conds {
		if !cond(r) {
			r.NetError = Bad_Request.Def("failed assertions")
			return r
		}
	}

	return r
}

func Println(s string) { fmt.Println(s) }

func (r Response) String(f func(string)) Response {
	if r.NetError != nil { return r }
	r.UseBody(func(r io.Reader) bool {
		bar, err := io.ReadAll(r)
		if err != nil { return false }
		f(string(bar))
		return true
	})
	return r
}

func UseBody(f func(io.Reader) bool) Monadic[Response, Response] {
	return func(r Response) Response {
		defer r.Body.Close()
		if r.NetError != nil { return r }

		r.NetError = Def(r.StatusCode)
		if r.NetError != nil { return r }

		if !f(r.Body) { r.NetError = Free(500, "Body assertion failed") }
		return r
	}
}

func (r Response) UseBody(f func(io.Reader) bool) Response {
	if r.NetError != nil { return r }
	return UseBody(f)(r)
}
